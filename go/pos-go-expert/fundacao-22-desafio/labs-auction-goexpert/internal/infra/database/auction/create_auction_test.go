package auction_test

import (
	"context"
	"fullcycle-auction_go/configuration/database/mongodb"
	"fullcycle-auction_go/internal/infra/database/auction"
	"fullcycle-auction_go/internal/infra/database/bid"
	"fullcycle-auction_go/internal/usecase/auction_usecase"
	"sync"
	"time"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson"
)

func Test_CreateAuctionAndClose(t *testing.T) {

	os.Setenv("AUCTION_INTERVAL", "10ms")
	os.Setenv("MONGODB_URL", "mongodb://admin:admin@localhost:27017/?retryWrites=false")
	os.Setenv("MONGODB_DB", "auctions_test")

	ctx := context.Background()
	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	databaseConnection.CreateCollection(ctx, "auctions")
	databaseConnection.Collection("auctions").DeleteMany(ctx, bson.D{})

	var expected int64
	expected = 0

	result, _ := databaseConnection.Collection("auctions").CountDocuments(ctx, bson.M{"status": 0})
	assert.Equal(t, expected, result)

	result, _ = databaseConnection.Collection("auctions").CountDocuments(ctx, bson.M{"status": 1})
	assert.Equal(t, expected, result)

	auctionRepository := auction.NewAuctionRepository(databaseConnection)
	bidRepository := bid.NewBidRepository(databaseConnection, auctionRepository)

	uit := auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository)

	inputDTO := auction_usecase.AuctionInputDTO{
		ProductName: "ProductName_test",
		Category:    "Category_test",
		Description: "Description_test",
	}

	println(1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		uit.CreateAuction(ctx, inputDTO)
		wg.Wait()
	}()
	wg.Done()

	// pausa para dar tempo de processar o dado
	time.Sleep(getAuctionInterval())
	time.Sleep(time.Millisecond * 10)

	expected = 0
	result, _ = databaseConnection.Collection("auctions").CountDocuments(ctx, bson.M{"status": 0})
	assert.Equal(t, expected, result)

	expected = 1
	result, _ = databaseConnection.Collection("auctions").CountDocuments(ctx, bson.M{"status": 1})
	assert.Equal(t, expected, result)
}

func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		return time.Minute * 5
	}

	return duration
}

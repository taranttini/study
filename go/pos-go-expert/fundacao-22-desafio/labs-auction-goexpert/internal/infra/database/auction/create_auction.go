package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"

	//"fullcycle-auction_go/internal/infra/database/auction"
	"fullcycle-auction_go/internal/internal_error"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *AuctionRepository
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {

	var wg sync.WaitGroup
	wg.Add(1)

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	//fmt.Println(auctionEntityMongo)
	duration := getAuctionInterval()
	ticker := time.NewTicker(duration)

	fmt.Printf("Opened auction [%v] per %v \n", auctionEntityMongo.Id, duration)

	quit := make(chan struct{})
	go func(id string) {
		for {
			select {
			case <-ticker.C:
				// realizando o update quando encerrar o tempo
				closeAuction(ctx, ar, auctionEntityMongo.Id)

				fmt.Printf("Closed auction [%v] at %v \n", id, time.Now())
				ticker.Stop()
				defer wg.Done()
			case <-quit:
				// encerrando
				defer wg.Done()
				ticker.Stop()
				return
			}
		}
	}(auctionEntityMongo.Id)

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	wg.Wait()
	return nil
}

func closeAuction(ctx context.Context, ar *AuctionRepository, auctionId string) {
	filter := bson.M{"_id": auctionId}
	update := bson.M{"$set": bson.M{"status": 1}}
	_, err := ar.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		logger.Error("Error trying to update auction", err)
	}
}

func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		return time.Minute * 5
	}

	return duration
}

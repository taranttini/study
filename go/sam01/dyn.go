main package


import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"


	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"log"
	"strconv"

)

func CreateLocalClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

func main() {

	os.Setenv("AWS_ACCESS_KEY_ID", "dummy1")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy2")
	os.Setenv("AWS_SESSION_TOKEN", "dummy3")

	svc.CreateLocalClient()

	// Create table Movies
	tableName := "Movies"

	input := &dynamodb.CreateTableInput{
	    AttributeDefinitions: []*dynamodb.AttributeDefinition{
	        {
	            AttributeName: aws.String("Year"),
	            AttributeType: aws.String("N"),
	        },
	        {
	            AttributeName: aws.String("Title"),
	            AttributeType: aws.String("S"),
	        },
	    },
	    KeySchema: []*dynamodb.KeySchemaElement{
	        {
	            AttributeName: aws.String("Year"),
	            KeyType:       aws.String("HASH"),
	        },
	        {
	            AttributeName: aws.String("Title"),
	            KeyType:       aws.String("RANGE"),
	        },
	    },
	    ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
	        ReadCapacityUnits:  aws.Int64(10),
	        WriteCapacityUnits: aws.Int64(10),
	    },
	    TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
	    log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tableName)


	// #### ADD ITEM

	item := Item{
	    Year:   2015,
	    Title:  "The Big New Movie",
	    Plot:   "Nothing happens at all.",
	    Rating: 0.0,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
	    log.Fatalf("Got error marshalling new movie item: %s", err)
	}


	input := &dynamodb.PutItemInput{
	    Item:      av,
	    TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
	    log.Fatalf("Got error calling PutItem: %s", err)
	}

	year := strconv.Itoa(item.Year)

	fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)


	//### search ITEM

	movieName := "The Big New Movie"
	movieYear := "2015"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
	    TableName: aws.String(tableName),
	    Key: map[string]*dynamodb.AttributeValue{
	        "Year": {
	            N: aws.String(movieYear),
	        },
	        "Title": {
	            S: aws.String(movieName),
	        },
	    },
	})
	if err != nil {
	    log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
	    msg := "Could not find '" + *title + "'"
	    return nil, errors.New(msg)
	}
 
	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
	    panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Year:  ", item.Year)
	fmt.Println("Title: ", item.Title)
	fmt.Println("Plot:  ", item.Plot)
	fmt.Println("Rating:", item.Rating)

}


// Create struct to hold info about new item
type Item struct {
    Year   int
    Title  string
    Plot   string
    Rating float64
}



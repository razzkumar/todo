package mongodb

import (
	//"context"
	"fmt"
	//"time"

	"github.com/razzkumar/todo/db-service/utils/constants"
	"github.com/razzkumar/todo/db-service/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {

	//db_url := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(constants.DB_URL))

	if err != nil {
		logger.FailOnError(err, "Faill to create new mongodb client")
	}
	fmt.Println("----------Mongo Connect Success---------------")
	return client
}

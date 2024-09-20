// package mongo

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func ConnectMongo() (*mongo.Database, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Hour)

// 	defer cancel()

// 	clientOPtion := options.Client().ApplyURI("mongodb://mongo:27017")

// 	client, err := mongo.Connect(ctx, clientOPtion)

// 	if err != nil {
// 		return nil, err
// 	}

// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println()
// 		fmt.Println()
// 		fmt.Println()
// 		fmt.Println()
// 		return nil, err
// 	}
// 	log.Println("salodfkemdsfmiewksdvofmiosm kmdvc")
// 	return client.Database("twit"), nil

// }
package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Database, error) {
	// 10 soniya timeout bilan context yaratish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB serveriga ulanish uchun konfiguratsiya
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017") // localhost o'rniga real IP yoki domen ishlating

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// MongoDB serverini ping qilish
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("MongoDB serveriga ulanishda xatolik:", err)
		return nil, err
	}

	return client.Database("twit"), nil
}

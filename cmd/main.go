package main

import (
	"fmt"
	"log"
	"net"
	"twit-service/config"
	"twit-service/generated/twit"
	"twit-service/logs"
	twitService "twit-service/service"
	"twit-service/storage"
	mongo "twit-service/storage/modgodb"
	"twit-service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	logger := logs.InitLogger()
	cfg := config.LoadConfig()

	db, err := postgres.Connect(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Postgres bilan bog'lanilmadi: %v", err))
		panic(err)
	}
	defer db.Close()

	mdb, err := mongo.ConnectMongo()
	if err != nil {
		logger.Error(fmt.Sprintf("Mongo DB bilan boglamilmadi: %v", err))
		panic(err)
	}

	listener, err := net.Listen("tcp", cfg.TWIT_SERVICE)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	defer listener.Close()

	storage := storage.NewTwitImpl(db, mdb)
	twitService := twitService.NewService(logger, storage)

	s := grpc.NewServer()
	twit.RegisterTwitServiceClientServer(s, twitService)

	log.Printf("Service is run: %s", cfg.TWIT_SERVICE)
	if err = s.Serve(listener); err != nil{
		log.Fatal(err)
	}
}

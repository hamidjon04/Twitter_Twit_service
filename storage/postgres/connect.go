package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"twit-service/config"
	_ "github.com/lib/pq"
)

func Connect(cfg config.Config)(*sql.DB, error){
	connector := fmt.Sprintf("host = %s port = %s user = %s dbname = %s password = %s sslmode = disable", 
							cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_NAME, cfg.DB_PASSWORD)

	log.Println(connector)
	db, err := sql.Open("postgres", connector)
	if err != nil{
		log.Println(err)
		return nil, err
	}

	if err := db.Ping(); err != nil{
		log.Println(err)
		return nil, err
	}

	return db, nil
}
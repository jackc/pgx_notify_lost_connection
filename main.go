package main

import (
	"github.com/jackc/pgx"
	"log"
	"os"
	"time"
)

var conn *pgx.Conn

func main() {
	var err error
	conn, err = pgx.Connect(extractConfig())
	if err != nil {
		log.Fatalln("Unable to connection to database:", err)
	}

	log.Println("PostgreSQL PID:", conn.Pid)

	err = conn.Listen("pgx")
	if err != nil {
		log.Fatalln("Unable to listen to channel pgx:", err)
	}
	log.Println("Listening on channel pgx")

	for {
		timeout := 5 * time.Second
		notification, err := conn.WaitForNotification(timeout)
		if err == pgx.ErrNotificationTimeout {
			log.Println("ErrNotificationTimeout")
			continue
		} else if err != nil {
			log.Println("Error waiting for notification:", err)
			os.Exit(1)
		}
		log.Println("Received notification:", notification.Payload)
	}
}

func extractConfig() pgx.ConnConfig {
	var config pgx.ConnConfig

	config.KeepAlive = 5

	config.Host = os.Getenv("DB_HOST")
	if config.Host == "" {
		config.Host = "localhost"
	}

	config.User = os.Getenv("DB_USER")
	if config.User == "" {
		config.User = os.Getenv("USER")
	}

	config.Password = os.Getenv("DB_PASSWORD")

	config.Database = os.Getenv("DB_DATABASE")
	if config.Database == "" {
		config.Database = "postgres"
	}

	return config
}

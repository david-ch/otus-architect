package db

import (
	"context"
	"fmt"
	"github.com/david-ch/otus-architect/delivery-service/delivery/model"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type deliveryDb struct{}

var Delivery = deliveryDb{}

func getConnCreds() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
}

func getConn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), getConnCreds())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return conn
}

func (db *deliveryDb) CreateDeliveryPlan(r model.DeliveryPlan) {
	conn := getConn()
	defer conn.Close(context.Background())

	conn.QueryRow(
		context.Background(),
		"INSERT INTO delivery_plan (\"order_id\", \"timeslot\") VALUES ($1, $2)",
		r.OrderId,
		r.Timeslot,
	)
}

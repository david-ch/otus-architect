package db

import (
	"context"
	"fmt"
	"github.com/david-ch/otus-architect/hw19/order-service/order/model"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type orderDb struct{}

var Orders = orderDb{}

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

func (db *orderDb) Create(o *model.Order) (int64, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	row := conn.QueryRow(
		context.Background(),
		"INSERT INTO orders (user_id, product_id, status, created_by_request_id) VALUES ($1, $2, $3, $4) RETURNING id",
		o.UserID,
		o.ProductID,
		o.Status,
		o.CreatedByRequestId,
	)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		log.Print(err)
		return 0, fmt.Errorf("cannot create order")
	}

	return id, nil
}

func (db *orderDb) Get(id int64) (model.Order, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	o := model.Order{}
	err := conn.QueryRow(
		context.Background(),
		"SELECT \"id\", \"user_id\", \"product_id\", \"status\" FROM products WHERE id = $1", id).
		Scan(&o.ID, &o.UserID, &o.ProductID, &o.Status)

	return o, err
}

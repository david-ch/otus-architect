package db

import (
	"context"
	"fmt"
	"github.com/david-ch/otus-architect/warehouse-service/warehouse/model"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type warehouseDb struct{}

var Warehouse = warehouseDb{}

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

func (db *warehouseDb) CreateReservation(r model.StockReservation) {
	conn := getConn()
	defer conn.Close(context.Background())

	conn.QueryRow(
		context.Background(),
		"INSERT INTO stock_reservation (\"product_id\", \"order_id\", \"amount\") VALUES ($1, $2, $3)",
		r.ProductId,
		r.OrderId,
		r.Amount,
	)
}

func (db *warehouseDb) GetReservation(orderId string) (model.StockReservation, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	r := model.StockReservation{}
	err := conn.QueryRow(
		context.Background(),
		"SELECT \"product_id\", \"order_id\", \"amount\" FROM stock WHERE order_id = $1", orderId).
		Scan(&r.ProductId, &r.OrderId, &r.Amount)

	return r, err
}

func (db *warehouseDb) DeleteReservation(orderId string) error {
	conn := getConn()
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"DELETE FROM stock_reservation WHERE order_id = $1",
		orderId,
	)

	return err
}

func (db *warehouseDb) GetStock(productId int64) (model.Stock, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	s := model.Stock{}
	err := conn.QueryRow(
		context.Background(),
		"SELECT \"product_id\", \"amount\" FROM stock WHERE product_id = $1", productId).
		Scan(&s.ProductId, &s.Amount)

	return s, err
}

func (db *warehouseDb) UpdateStock(s model.Stock) error {
	conn := getConn()
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"UPDATE stock set \"amount\" = $1 WHERE \"product_id\" = $2",
		s.Amount,
		s.ProductId,
	)

	return err
}

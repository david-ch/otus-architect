package db

import (
	"context"
	"fmt"
	"github.com/david-ch/otus-architect/hw15/product-service/product/model"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type productDb struct{}

var Products = productDb{}

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

func (db *productDb) Create(p *model.Product) (int64, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	row := conn.QueryRow(
		context.Background(),
		"INSERT INTO products (\"name\", \"description\") VALUES ($1, $2) RETURNING id",
		p.Name,
		p.Description,
	)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		log.Print(err)
		return 0, fmt.Errorf("cannot create product")
	}

	return id, nil
}

func (db *productDb) Get(id int64) (model.Product, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	p := model.Product{}
	err := conn.QueryRow(
		context.Background(),
		"SELECT \"id\", \"name\", \"description\" FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Description)

	return p, err
}

func (db *productDb) Update(p *model.Product) error {
	conn := getConn()
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"UPDATE products set \"name\" = $1, \"description\" = $2 WHERE \"id\" = $3",
		p.Name,
		p.Description,
		p.ID,
	)

	return err
}

func (db *productDb) Delete(id int64) error {
	conn := getConn()
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"DELETE FROM products WHERE id = $1",
		id,
	)

	return err
}

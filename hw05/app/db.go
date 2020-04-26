package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
)

type userDb struct{}

var users = userDb{}

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

func (db *userDb) create(u *user) (int64, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	row := conn.QueryRow(
		context.Background(),
		"INSERT INTO users (\"username\", \"firstname\", \"lastname\", \"email\", \"phone\") VALUES ($1, $2, $3, $4, $5) RETURNING id",
		u.Username,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Phone,
	)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		log.Print(err)
		return 0, fmt.Errorf("Cannot create user")
	}

	return id, nil
}

func (db *userDb) get(id int64) (*user, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	u := user{}
	err := conn.QueryRow(
		context.Background(),
		"SELECT \"id\", \"username\", \"firstname\", \"lastname\", \"email\", \"phone\" FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.Phone)

	return &u, err
}

func (db *userDb) update(u *user) error {
	conn := getConn()
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"UPDATE users set \"username\" = $1, \"firstname\" = $2, \"lastname\" = $3, \"email\" = $4, \"phone\" = $5",
		u.Username,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Phone,
	)

	return err
}

func (db *userDb) delete(id int64) error {
	conn := getConn()
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"DELETE FROM users WHERE id = $1",
		id,
	)

	return err
}

package user

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
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
		"INSERT INTO users (\"username\", \"firstname\", \"lastname\", \"passwordhash\", \"personalstatus\") VALUES ($1, $2, $3, $4, $5) RETURNING id",
		u.Username,
		u.FirstName,
		u.LastName,
		u.PasswordHash,
		u.PersonalStatus,
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
		"SELECT \"id\", \"username\", \"firstname\", \"lastname\", \"passwordhash\", \"personalstatus\" FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.PasswordHash, &u.PersonalStatus)

	return &u, err
}

func (db *userDb) findByUsername(username string) (user, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	u := user{}
	err := conn.QueryRow(
		context.Background(),
		"SELECT \"id\", \"username\", \"firstname\", \"lastname\", \"passwordhash\", \"personalstatus\" FROM users WHERE username = $1", username).
		Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.PasswordHash, &u.PersonalStatus)

	return u, err
}

func (db *userDb) update(u *user) error {
	conn := getConn()
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"UPDATE users set \"username\" = $1, \"firstname\" = $2, \"lastname\" = $3, \"passwordhash\" = $4, \"personalstatus\" = $5 WHERE \"id\" = $6",
		u.Username,
		u.FirstName,
		u.LastName,
		u.PasswordHash,
		u.PersonalStatus,
		u.ID,
	)

	return err
}

func (db *userDb) updateByUsername(username string, u *user) error {
	conn := getConn()
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"UPDATE users set \"username\" = $1, \"firstname\" = $2, \"lastname\" = $3, \"passwordhash\" = $4, \"personalstatus\" = $5 WHERE \"username\" = $6",
		u.Username,
		u.FirstName,
		u.LastName,
		u.PasswordHash,
		u.PersonalStatus,
		username,
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

func (db *userDb) deleteByUsername(username string) error {
	conn := getConn()
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"DELETE FROM users WHERE username = $1",
		username,
	)

	return err
}

func (db *userDb) count() (int32, error) {
	conn := getConn()
	defer conn.Close(context.Background())

	var cnt int32
	err := conn.QueryRow(
		context.Background(),
		"SELECT count(*) FROM users",
	).Scan(&cnt)

	return cnt, err
}

package db

import (
	"context"
	"fmt"
	"github.com/david-ch/otus-architect/order-service/order/model"
	"hash/fnv"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4"
)

type orderDb struct{}

var Orders = orderDb{}

func getConnCreds(shardId string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_" +shardId+ "_USER"),
		os.Getenv("POSTGRES_" +shardId+ "_PASSWORD"),
		os.Getenv("POSTGRES_" +shardId+ "_HOST"),
		os.Getenv("POSTGRES_" +shardId+ "_PORT"),
		os.Getenv("POSTGRES_" +shardId+"_DB"),
	)
}

func getConn(shardId string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), getConnCreds(shardId))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return conn
}

func (db *orderDb) Create(o model.Order) error {
	conn := getConn(calculateShardId(o.ID))
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO orders ("+
			"   id, "+
			"	user_id, "+
			"	product_id, "+
			"	status, "+
			"	delivery_timeslot, "+
			"	created_by_request_id"+
			") VALUES ($1, $2, $3, $4, $5, $6)",
		o.ID,
		o.UserID,
		o.ProductID,
		o.Status,
		o.DeliveryTimeslot,
		o.CreatedByRequestId,
	)

	if err != nil {
		log.Print(err)
		return fmt.Errorf("cannot create order")
	}

	return nil
}

func (db *orderDb) Update(o model.Order) error {
	conn := getConn(calculateShardId(o.ID))
	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		"UPDATE orders set "+
			"	\"user_id\" = $1, "+
			"	\"product_id\" = $2, "+
			"	\"status\" = $3, "+
			"	\"delivery_timeslot\" = $4 "+
			" WHERE \"id\" = $5",
		o.UserID,
		o.ProductID,
		o.Status,
		o.DeliveryTimeslot,
		o.ID,
	)

	return err
}

func (db *orderDb) Get(id string) (model.Order, error) {
	conn := getConn(calculateShardId(id))
	defer conn.Close(context.Background())

	o := model.Order{}
	err := conn.QueryRow(
		context.Background(),
		"SELECT \"id\", \"user_id\", \"product_id\", \"status\", \"delivery_timeslot\" FROM orders WHERE id = $1", id).
		Scan(&o.ID, &o.UserID, &o.ProductID, &o.Status, &o.DeliveryTimeslot)

	return o, err
}

func calculateShardId(id string) string {
	const shardsNum = 2
	return strconv.Itoa(int(hash(id) % shardsNum))
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var sqlTables = `
create extension if not exists citext;

create table if not exists address (
    address 					citext 	primary key,
    display_name 				text,
    accepting_new 				bool	not null default true,
    limit_recv_size_total 		bigint	not null default -1,
    limit_recv_size_per_msg 	bigint	not null default -1,
    limit_recv_size_per_1d 		bigint	not null default -1,
    limit_recv_count_per_1d 	bigint	not null default -1,
    limit_send_size_total 		bigint	not null default -1,
    limit_send_size_per_msg 	bigint	not null default -1,
    limit_send_size_per_1d 		bigint	not null default -1,
    limit_send_count_per_1d 	bigint	not null default -1
);

-- TODO consider time-series datastore e.g. TimescaleDB
create table if not exists address_tx (
	address	citext		not null references address (address),
	ts		timestamp	not null,
	op		varchar(5)	not null, -- send, recv
	size	int			not null,
	primary key (address, ts)
)
`

func initDb() error {
	db, err := pgx.Connect(context.Background(), "")
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	_, err = db.Exec(sqlTables)
	if err != nil {
		return err
	}
	return nil
}

func getAddressDetail(c *gin.Context) {
	ctx := context.Background()

	connString := "postgres://user:password@localhost:5432/database"

	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		fmt.Println("Unable to connect to database:", err)
		return
	}

	defer pool.Close()

	rows, err := pool.Query(ctx, "SELECT * FROM table_name")
	if err != nil {
		fmt.Println("Query failed:", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		var address string
		var displayName string
		var acceptingNew bool
		var limitRecvSizeTotal int64
		var limitRecvSizePerMsg int64
		var limitRecvSizePer1d int64
		var limitRecvCountPer1d int64
		var limitSendSizeTotal int64
		var limitSendSizePerMsg int64
		var limitSendSizePer1d int64
		var limitSendCountPer1d int64
		var recvSizeTotal int64
		var recvSizePer1d int64
		var recvCountPer1d int64
		var sendSizeTotal int64
		var sendSizePer1d int64
		var sendCountPer1d int64

		err = rows.Scan(&name, &address, &displayName, &acceptingNew, &limitRecvSizeTotal,
			&limitRecvSizePerMsg, &limitRecvSizePer1d, &limitRecvCountPer1d,
			&limitSendSizeTotal, &limitSendSizePerMsg, &limitSendSizePer1d,
			&limitSendCountPer1d, &recvSizeTotal, &recvSizePer1d,
			&recvCountPer1d, &sendSizeTotal, &sendSizePer1d,
			&sendCountPer1d)

		if err != nil {
			fmt.Println("Unable to scan row:", err)
			return
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/user/", getAddressDetail)
	r.Run(":8080")
}

package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)


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
		var limitRecvSizeTotal int
		var limitRecvSizePerMsg int
		var limitRecvSizePer1d int
		var limitRecvCountPer1d int
		var limitSendSizeTotal int
		var limitSendSizePerMsg int
		var limitSendSizePer1d int
		var limitSendCountPer1d int
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
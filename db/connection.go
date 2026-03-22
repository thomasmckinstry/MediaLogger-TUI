/*
Copyright © 2025 Thomas McKinstry thomas.g.mckinstry@protonmail.com
*/

package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"strings"
)

var configs struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

func init() {
	content, _ := os.ReadFile("config.json") // TODO: Add in the util for checking if a sql query has worked
	//c.Check(err, "ERROR: Failed to read config file.")

	_ = json.Unmarshal(content, &configs)
	//c.Check(err, "ERROR: Failed to unmarshal config.")
}

func Init_connection() *pgx.Conn {
	var conn *pgx.Conn
	var ctx = context.Background()
	var err error

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("postgres://%s:%s@%s:%d/%s", configs.User, configs.Pass, configs.Host, configs.Port, configs.Name))
	conn, err = pgx.Connect(ctx, builder.String())
	if err != nil {
		log.Fatal("Unable to connect to database:", err)

		// TODO: Failure to connect to the database should lead to initializing the database now.
		// See init.go

		os.Exit(1)
	}

	return conn
}

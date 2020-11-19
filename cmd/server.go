package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type article struct {
	Id        uint
	Title     string
	Body      string
	Votes     int
	CreatedAt time.Time
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! World!")
		log.Print("hello")
	})
	ctx := context.Background()
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		db, err := pgx.Connect(ctx, "postgres://nickbadlose:bt3brgb455646RFE8d3b53ikojorg@localhost:5431/postgres")
		if err != nil {
			panic(fmt.Sprintf("db error, %v", err))
		}

		sql, _, err := squirrel.Select("*").From("articles").Limit(1).ToSql()
		if err != nil {
			panic(fmt.Sprintf("sql building error, %v", err))
		}

		rows, err := db.Query(ctx, sql)
		if err != nil {
			panic(fmt.Sprintf("db reading error, %v", err))
		}

		rows.Scan()

		articles := make([]article, 0)

		rs, err := rows.Values()
		if err != nil {
			panic(fmt.Sprintf("db reading row values, %v", err))
		}

		for _, row := range rs {
			err = row.Scan(&a.Id, &a.Title, &a.Body, &a.Votes, &a.CreatedAt)
			if err != nil {
				panic(fmt.Sprintf("scan error, %v", err))
			}
			ma, err := json.Marshal(a)
			if err != nil {
				panic(fmt.Sprintf("marshaling error, %v", err))
			}
		}
		w.WriteHeader(200)
		w.Write(ma)
		log.Print("articles")
	})

	// Staring server
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("error starting server")
		log.Fatal(err)
	}
}

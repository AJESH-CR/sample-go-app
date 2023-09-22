package repository

import (
	"context"
	"log"
	"sample-go-app/src/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

func connection() *pgxpool.Pool {
	url := "postgres://postgres:passwd@localhost:51002/app"
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func SaveTemplate(t *dto.Template) {
	db := connection()
	db.Exec(context.Background(), "INSERT INTO template(app_id, level_count, condition) VALUES (?, ?, ?)", t.AppId, len(t.Levels))
}
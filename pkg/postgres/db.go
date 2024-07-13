package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/marcelofabianov/pejota/bootstrap"
)

var db *pgx.Conn

func ConnectDB(config bootstrap.Config) (*pgx.Conn, error) {
	host := config.DbHost
	port := config.DbPort
	user := config.DbUser
	pass := config.DbPass
	dbname := config.DbName

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pass, host, port, dbname)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	db = conn
	return db, nil
}

func GetDB() *pgx.Conn {
	return db
}

func CloseDB(ctx context.Context) {
	if db != nil {
		db.Close(ctx)
	}
}

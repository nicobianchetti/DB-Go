package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

//NewPostgresDB create a new conection at the db
func NewPostgresDB() {
	once.Do(func() {
		var err error
		usr := "postgres"
		pass := "postgres"
		host := "localhost"
		port := "5432"
		dbName := "godb"
		// dsn := "postgres://postgres:postgres@localhost:5432/godb?sslmode=disable"
		dsn := "postgres://" + usr + ":" + pass + "@" + host + ":" + port + "/" + dbName + "?sslmode=disable"
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}
		// defer db.Close() //el defer lo quitamos porque se supone que siempre debe estar disponible la instancia

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}

		fmt.Println("Conectado a Postgres")
	})
}

//Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}

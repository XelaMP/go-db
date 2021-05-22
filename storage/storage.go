package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	_"github.com/lib/pq"
)

var (
	db *sql.DB
	once sync.Once
)

func NewPostgresDB(){
	once.Do(func() {
		var err error
		db, err := sql.Open("postgres",
			"postgres://postgres:56523314@localhost:5432/gopostgresdb?sslmode=disable")
		if err != nil {
			log.Fatalf("no se pudo conectar la bd: %v")
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("no se pudo hacer ping: %v")
		}
		fmt.Println("conectado a la bd postgres")
	})

}
// pool retorna unica instancia de db
func Pool() *sql.DB{
	return db
}

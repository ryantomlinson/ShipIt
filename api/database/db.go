package database

import (
	"database/sql"
	_ "github.com/lib/pq" //import postgres
	"github.com/go-gorp/gorp"
	"fmt"
	"log"
)

type DB struct {
	*sql.DB
}


const (
	DbUser = "postgres"
	DbPassword = "postgres"
	DbName = "shipit"
)

var db *gorp.DbMap

func Init() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
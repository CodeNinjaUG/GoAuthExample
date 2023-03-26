package db

import (
	 "database/sql"
	 "log"
	 "os"
	 "testing"
	_"github.com/lib/pq"
)

var testQueries *Queries
const (
	dbDriver="postgres"
	dbSource="postgres://spsuserone:@Th1nkIT@localhost:5432/sps1?sslmode=disable"
)

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err!=nil{
		log.Fatal("cannot connect to the database", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
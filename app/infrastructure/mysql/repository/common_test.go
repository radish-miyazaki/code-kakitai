package repository

import (
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db"
	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db/db_gen"
	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db/db_test"
)

var (
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	var err error

	resource, pool := db_test.CreateContainer()
	defer db_test.CloseContainer(resource, pool)

	dbCon := db_test.ConnectDB(resource, pool)
	defer dbCon.Close()

	db_test.SetupTestDB()

	fixtures, err = testfixtures.New(
		testfixtures.Database(dbCon),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("../fixtures"),
	)
	if err != nil {
		panic(err)
	}

	q := db_gen.New(dbCon)
	db.SetQuery(q)
	db.SetDBConn(dbCon)

	m.Run()
}

func resetTestData(t *testing.T) {
	t.Helper()

	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}

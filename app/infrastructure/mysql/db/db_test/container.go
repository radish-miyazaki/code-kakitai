package db_test

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/sqldef/sqldef"
	"github.com/sqldef/sqldef/database"
	"github.com/sqldef/sqldef/database/mysql"
	"github.com/sqldef/sqldef/parser"
	"github.com/sqldef/sqldef/schema"
)

var (
	userName = "root"
	password = "secret"
	dbName   = "code_kakitai_test"
	hostName = "localhost"
	port     int
)

func CreateContainer() (*dockertest.Resource, *dockertest.Pool) {
	pool, err := dockertest.NewPool("")
	//pool.MaxWait = time.Minute * 2
	pool.MaxWait = time.Minute * 1
	if err != nil {
		log.Panic("Could not connect to docker: ", err)
	}

	runOpts := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + password,
			"MYSQL_DATABASE=" + dbName,
		},
		Mounts: []string{},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
	}

	resource, err := pool.RunWithOptions(runOpts)
	if err != nil {
		log.Panic("Could not start resource: ", err)
	}

	return resource, pool
}

func CloseContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	if err := pool.Purge(resource); err != nil {
		log.Panic("Could not purge resource: ", err)
	}
}

func ConnectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {
	var db *sql.DB
	if err := pool.Retry(func() error {
		time.Sleep(time.Second * 3)

		var err error
		port, err = strconv.Atoi(resource.GetPort("3306/tcp"))
		if err != nil {
			return err
		}

		db, err = sql.Open(
			"mysql",
			fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", userName, password, hostName, port, dbName),
		)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Panic("Could not connect to database: ", err)
	}

	return db
}

func SetupTestDB() {
	desiredDDLs, err := sqldef.ReadFile("../db/schema/schema.sql")
	if err != nil {
		log.Panic("Could not read schema file: ", err)
	}

	opts := &sqldef.Options{DesiredDDLs: desiredDDLs}
	p := database.NewParser(parser.ParserModeMysql)
	db, err := mysql.NewDatabase(database.Config{
		Host:     hostName,
		Port:     port,
		User:     userName,
		Password: password,
		DbName:   dbName,
	})
	if err != nil {
		log.Panic(err)
	}
	sqldef.Run(schema.GeneratorModeMysql, db, p, opts)
}

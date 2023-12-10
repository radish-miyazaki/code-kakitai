package schema

import (
	"strconv"

	"github.com/sqldef/sqldef"
	"github.com/sqldef/sqldef/database"
	"github.com/sqldef/sqldef/database/mysql"
	"github.com/sqldef/sqldef/parser"
	"github.com/sqldef/sqldef/schema"

	"github.com/radish-miyazaki/code-kakitai/config"
)

func Migrate(schemaFile string, dryRun bool) error {
	dbCfg := config.GetConfig().DB

	port, err := strconv.Atoi(dbCfg.Port)
	if err != nil {
		return err
	}
	db, err := mysql.NewDatabase(database.Config{
		Host:     dbCfg.Host,
		Port:     port,
		User:     dbCfg.User,
		Password: dbCfg.Password,
		DbName:   dbCfg.Name,
	})
	if err != nil {
		return err
	}

	desiredDDLs, err := sqldef.ReadFile(schemaFile)
	if err != nil {
		return err
	}

	options := &sqldef.Options{
		DesiredDDLs:     desiredDDLs,
		DryRun:          dryRun,
		EnableDropTable: true,
	}

	p := database.NewParser(parser.ParserModeMssql)
	sqldef.Run(schema.GeneratorModeMysql, db, p, options)

	return nil
}

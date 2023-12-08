package schema

import (
	"strconv"

	"github.com/radish-miyazaki/code-kakitai/config"
)

func Migrate(schemaFile string) error {
	dbCfg := config.GetConfig().DB

	port, err := strconv.Atoi(dbCfg.Port)
	if err != nil {
		return err
	}

	return nil
}

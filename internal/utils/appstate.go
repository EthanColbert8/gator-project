package utils

import (
	"github.com/EthanColbert8/gator-project/internal/config"
	"github.com/EthanColbert8/gator-project/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}

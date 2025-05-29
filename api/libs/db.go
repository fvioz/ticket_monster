package libs

import (
	"api/configs"
	"sync"

	"github.com/go-pg/pg/extra/pgotel/v10"
	"github.com/go-pg/pg/v10"
)

var db *pg.DB
var dbOnce sync.Once

func DBInstance() *pg.DB {
	dbOnce.Do(func() {
		config := configs.NewDbConfig()

		dbInstance := pg.Connect(&pg.Options{
			Addr:     config.Address(),
			User:     config.Username,
			Password: config.Password,
			Database: config.Database,
			PoolSize: config.PoolSize,
		})

		dbInstance.AddQueryHook(pgotel.NewTracingHook())

		db = dbInstance
	})

	return db
}

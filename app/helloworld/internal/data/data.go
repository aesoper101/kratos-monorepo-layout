package data

import (
	"context"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/biz"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/conf"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/data/ent/schemagen"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/data/ent/schemagen/migrate"
	log2 "github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTransaction, NewGreeterRepo)

// Data .
type Data struct {
	db *schemagen.Database
}

func NewTransaction(data *Data) biz.Transaction {
	return data.db
}

// NewData .
func NewData(c *conf.Data, logger log2.Logger) (*Data, func(), error) {
	log := log2.NewHelper(logger)
	drv, err := entsql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	client := schemagen.NewClient(schemagen.Driver(drv))
	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	)
	if err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	db := drv.DB()
	if maxIdleConn := c.Database.GetMaxIdleCount(); maxIdleConn > 0 {
		db.SetMaxIdleConns(int(maxIdleConn))
	}

	if maxOpen := c.Database.GetMaxOpen(); maxOpen > 0 {
		db.SetMaxOpenConns(int(maxOpen))
	}

	if maxLifetime := c.Database.GetMaxLifeTime(); maxLifetime != nil {
		db.SetConnMaxLifetime(maxLifetime.AsDuration())
	}

	if maxIdleTime := c.Database.GetMaxIdleTime(); maxIdleTime != nil {
		db.SetConnMaxIdleTime(maxIdleTime.AsDuration())
	}

	d := &Data{
		db: schemagen.NewDatabase(schemagen.Driver(drv)),
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := drv.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

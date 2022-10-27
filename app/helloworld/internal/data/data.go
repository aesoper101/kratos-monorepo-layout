package data

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/biz"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/conf"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/data/ent"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/data/ent/migrate"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
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
	db *ent.Client
}

func NewTransaction(data *Data) biz.Transaction {
	return data
}

// NewData .
func NewData(c *conf.Data, logger log2.Logger) (*Data, func(), error) {
	log := log2.NewHelper(logger)

	db, closeDb, err := openDB(c.Database, log)
	if err != nil {
		return nil, nil, err
	}

	d := &Data{
		db: db,
	}

	cleanup := func() {
		closeDb()
	}
	return d, cleanup, nil
}

// User is the client for interacting with the User builders.
func (data *Data) User(ctx context.Context) *ent.UserClient {
	return data.db.User
}

func openDB(cfg *confpb.Database, helper *log2.Helper) (*ent.Client, func(), error) {
	db, err := sql.Open(
		cfg.Driver,
		cfg.Source,
	)
	if err != nil {
		helper.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}

	if maxIdleConn := cfg.GetMaxIdleCount(); maxIdleConn > 0 {
		db.SetMaxIdleConns(int(maxIdleConn))
	}

	if maxOpen := cfg.GetMaxOpen(); maxOpen > 0 {
		db.SetMaxOpenConns(int(maxOpen))
	}

	if maxLifetime := cfg.GetMaxLifeTime(); maxLifetime != nil {
		db.SetConnMaxLifetime(maxLifetime.AsDuration())
	}

	if maxIdleTime := cfg.GetMaxIdleTime(); maxIdleTime != nil {
		db.SetConnMaxIdleTime(maxIdleTime.AsDuration())
	}

	drv := entsql.OpenDB(cfg.Driver, db)
	client := ent.NewClient(ent.Driver(drv))
	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	)
	if err != nil {
		helper.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	cleanup := func() {
		helper.Info("message", "closing the data resources")
		if err := drv.Close(); err != nil {
			helper.Error(err)
		}
	}

	return client, cleanup, nil
}

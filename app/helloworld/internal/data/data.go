package data

import (
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/conf"
	log2 "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	//db *ent.Database
}

// NewData .
func NewData(c *conf.Data, logger log2.Logger) (*Data, func(), error) {
	log := log2.NewHelper(logger)
	//drv, err := sql.Open(
	//	c.Database.Driver,
	//	c.Database.Source,
	//)
	//if err != nil {
	//	log.Errorf("failed opening connection to sqlite: %v", err)
	//	return nil, nil, err
	//}
	// Run the auto migration tool.
	//client := ent.NewClient(ent.Driver(drv))
	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Errorf("failed creating schema resources: %v", err)
	//	return nil, nil, err
	//}

	d := &Data{
		//db: ent.NewDatabase(ent.Driver(drv)),
	}

	cleanup := func() {
		log.Info("message", "closing the data resources")
		//if err := drv.Close(); err != nil {
		//	log.Error(err)
		//}
	}
	return d, cleanup, nil
}

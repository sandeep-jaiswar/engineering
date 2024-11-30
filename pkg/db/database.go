package db

import (
	"context"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/sandeep-jaiswar/engineering/pkg/ent"
	"github.com/sandeep-jaiswar/engineering/pkg/logger"
)

func Open(databaseURL string) (*ent.Client, error) {
    driver, err := sql.Open(dialect.Postgres, databaseURL)
    if err != nil {
        logger.Logger.Error("failed to open connection")
    }
    return ent.NewClient(ent.Driver(driver)), nil
}

func Migrate(client *ent.Client) error {
    return client.Schema.Create(context.Background())
}

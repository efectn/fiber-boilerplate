package database

import (
	"context"

	"github.com/efectn/fiber-boilerplate/pkg/database/ent"
	"github.com/efectn/fiber-boilerplate/pkg/utils/config"
	"github.com/rs/zerolog/log"

	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	Ent *ent.Client
	Cfg *config.Config
}

type Seeder interface {
	Seed(*ent.Client) error
	Count() (int, error)
}

func NewDatabase(cfg *config.Config) *Database {
	db := &Database{
		Cfg: cfg,
	}

	return db
}

func (db *Database) ConnectDatabase() {
	conn, err := sql.Open("pgx", db.Cfg.DB.Postgres.DSN)
	if err != nil {
		log.Error().Err(err).Msg("An unknown error occured when to connect the database!")
	}

	drv := entsql.OpenDB(dialect.Postgres, conn)
	db.Ent = ent.NewClient(ent.Driver(drv))
}

func (db *Database) ShutdownDatabase() {
	if err := db.Ent.Close(); err != nil {
		log.Error().Err(err).Msg("An unknown error occured when to shutdown the database!")
	}
}

func (db *Database) MigrateModels() {
	if err := db.Ent.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Error().Err(err).Msg("Failed creating schema resources!")
	}
}

func (db *Database) SeedModels(seeder ...Seeder) {
	for _, v := range seeder {

		count, err := v.Count()
		if err != nil {
			log.Panic().Err(err).Msg("")
		}

		if count == 0 {
			err = v.Seed(db.Ent)
			if err != nil {
				log.Panic().Err(err).Msg("")
			}

			log.Debug().Msg("Table has seeded successfully.")
		} else {
			log.Warn().Msg("Table has seeded already. Skipping!")
		}
	}
}

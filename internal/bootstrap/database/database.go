package database

import (
	"context"

	"github.com/efectn/fiber-boilerplate/internal/ent"
	"github.com/efectn/fiber-boilerplate/utils/config"
	"github.com/rs/zerolog"

	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	Ent *ent.Client
	Log zerolog.Logger
	Cfg *config.Config
}

type Seeder interface {
	Seed(*ent.Client) error
	Count() (int, error)
}

func NewDatabase(cfg *config.Config, log zerolog.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}

	return db
}

func (db *Database) ConnectDatabase() {
	conn, err := sql.Open("pgx", db.Cfg.DB.Postgres.DSN)
	if err != nil {
		db.Log.Error().Err(err).Msg("An unknown error occurred when to connect the database!")
	} else {
		db.Log.Info().Msg("Connected the database succesfully!")
	}

	drv := entsql.OpenDB(dialect.Postgres, conn)
	db.Ent = ent.NewClient(ent.Driver(drv))
}

func (db *Database) ShutdownDatabase() {
	if err := db.Ent.Close(); err != nil {
		db.Log.Error().Err(err).Msg("An unknown error occurred when to shutdown the database!")
	}
}

func (db *Database) MigrateModels() {
	if err := db.Ent.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		db.Log.Error().Err(err).Msg("Failed creating schema resources!")
	} else {
		db.Log.Info().Msg("Models were migrated successfully!")
	}
}

func (db *Database) SeedModels(seeder ...Seeder) {
	for _, v := range seeder {

		count, err := v.Count()
		if err != nil {
			db.Log.Panic().Err(err).Msg("")
		}

		if count == 0 {
			err = v.Seed(db.Ent)
			if err != nil {
				db.Log.Panic().Err(err).Msg("")
			}

			db.Log.Debug().Msg("Table has seeded successfully.")
		} else {
			db.Log.Warn().Msg("Table has seeded already. Skipping!")
		}
	}

	db.Log.Info().Msg("Seeding was completed!")
}

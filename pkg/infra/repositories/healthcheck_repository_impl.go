package repositories

import (
	"context"
	"database/sql"
	"log"

	abstractions "github.com/leoguilen/simple-go-api/pkg/core/abstractions/repositories"
	dbcontext "github.com/leoguilen/simple-go-api/pkg/infra/context"
)

type HealthcheckRepository struct {
	DB *sql.DB
}

func NewHealthcheckRepository() abstractions.IHealthcheckRepository {
	db, err := dbcontext.NewDbContext().GetConnection()
	if err != nil {
		panic(err)
	}
	return &HealthcheckRepository{
		DB: db,
	}
}

func (hc *HealthcheckRepository) Check(ctx context.Context) error {
	if err := hc.DB.PingContext(ctx); err != nil {
		log.Printf("Ping to the database failed: %v", err.Error())
		return err
	}

	_, err := hc.DB.QueryContext(ctx, "SELECT * FROM TODO")
	if err != nil {
		log.Printf("Execute query failed: %v", err.Error())
		return err
	}

	return nil
}

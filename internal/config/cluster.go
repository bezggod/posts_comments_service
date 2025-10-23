package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DSN() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("failed to load .env file")
	}

	pgDsn := os.Getenv("POSTGRES_PG_DSN")
	sslMode := os.Getenv("POSTGRES_SSL_MODE")

	return fmt.Sprintf("%s&sslmode=%s", pgDsn, sslMode)
}

type Cluster struct {
	Pool *pgxpool.Pool
}

func NewCluster(ctx context.Context) (*Cluster, error) {
	dsn := DSN()

	fmt.Printf("dsn: %s\n", dsn)
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}

	return &Cluster{Pool: pool}, nil
}

func (c *Cluster) Close() {
	if c.Pool != nil {
		c.Pool.Close()
	}
}

package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"

	"github.com/pajarraco93/graphql-test/pkg/library/domain"
)

const (
	username = "root"
	password = "password"
	hostname = "127.0.0.1:3306"
	dbname   = "graphql_test"
)

var _ domain.Repository = &MySQLRepository{}

type MySQLRepository struct {
	engine *sql.DB
}

func NewMySQLRepository() domain.Repository {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Fatalf("Error %s when opening DB", err)
	}

	repo := &MySQLRepository{db}

	repo.createDB()
	repo.connectDB()
	repo.runMigrations()

	return repo
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbName)
}

func (repo *MySQLRepository) createDB() error {
	_, err := repo.engine.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		log.Fatalf("Error %s when creating DB\n", err)
	}

	return err
}

func (repo *MySQLRepository) connectDB() error {
	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	db := sqldblogger.OpenDriver(
		dsn(dbname),
		repo.engine.Driver(),
		loggerAdapter,
		sqldblogger.WithMinimumLevel(sqldblogger.LevelInfo),
	)
	repo.engine = db

	return nil
}

func (repo *MySQLRepository) runMigrations() {
	driver, err := mysql.WithInstance(repo.engine, &mysql.Config{})
	if err != nil {
		log.Fatalf("Error %s when creating migration driver", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./pkg/library/shared/infra/adapters/mySQL/internal/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Error %s when loading migrations", err)
	}

	err = m.Up()
}

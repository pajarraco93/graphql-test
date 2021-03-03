package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"

	"github.com/pajarraco93/graphql-test/pkg/library/domain"
)

const (
	username = "root"
	password = "secret"
	hostname = "127.0.0.1:3307"
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
	db, err := sql.Open("mysql", dsn(dbname))

	if err != nil {
		log.Fatalf("Error %s when opening DB", err)
	}

	repo.engine = db

	return err
}

func (repo *MySQLRepository) runMigrations() {
	driver, _ := mysql.WithInstance(repo.engine, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://./pkg/library/shared/infra/adapters/mySQL/internal/migrations",
		"mysql",
		driver,
	)

	err := m.Up()
	if err != nil {
		fmt.Println("Error ", err, " when executing migrations")
	}
}

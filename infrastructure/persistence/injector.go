package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hirac1220/go-clean-architecture/domain/repository"
	"github.com/kelseyhightower/envconfig"
)

type todoPersistence struct {
	db *sql.DB
}
type Config struct {
	DatabaseEngine   string `envconfig:"DB_ENGINE"`
	DatabaseUser     string `envconfig:"DB_USER"`
	DatabasePassword string `envconfig:"DB_PASSWORD"`
	DatabaseHost     string `envconfig:"DB_HOSTNAME"`
	DatabasePort     string `envconfig:"DB_PORT"`
	DatabaseName     string `envconfig:"DB_NAME"`
}

func SetConfig() {
	os.Setenv("DB_ENGINE", "mysql")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_HOSTNAME", "0.0.0.0")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "db")
}

func NewTodoPersistence() (repository.TodoRepository, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// db
	engine := cfg.DatabaseEngine
	url := connectionString(cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)

	db, err := sql.Open(engine, url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &todoPersistence{db}, err
}

func connectionString(user, password, host, port, dbname string) string {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, dbname, "parseTime=true")

	return url
}

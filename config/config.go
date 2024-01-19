package config

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/ATtendev/share/internal/log"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

type Config struct {
	DatabaseConf DatabaseConf `yaml:"DatabaseConf"`
	ServerConf   ServerConf   `yaml:"ServerConf"`
}

type ServerConf struct {
	ID     string `yaml:"ID"`
	Host   string `yaml:"Host"`
	Port   int    `yaml:"Port"`
	Secret string `yaml:"Secret"`
	Banner string `yaml:"Banner"`
}

// DatabaseConf stores database configurations.
type DatabaseConf struct {
	Host         string `yaml:"Host"`
	Port         int    `yaml:"Port"`
	Username     string `yaml:"Username"`
	Password     string `yaml:"Password"`
	DBName       string `yaml:"DBName"`
	SSLMode      string `yaml:"SSLMode"`
	Type         string `yaml:"Type"`
	MaxOpenConn  int    `yaml:"MaxOpenConn"`
	CacheTime    int    `yaml:"CacheTime"`
	DBPath       string `yaml:"DBPath"`
	MysqlConfig  string `yaml:"MysqlConfig"`
	PGConfig     string `yaml:"PGConfig"`
	SqliteConfig string `yaml:"SqliteConfig"`
}

// NewCacheDriver returns an Ent driver with cache.

// NewNoCacheDriver returns an Ent driver without cache.
func (c *DatabaseConf) NewNoCacheDriver() *entsql.Driver {
	db, err := sql.Open(c.Type, c.GetDSN())
	if err != nil {
		log.Error(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Panicf("can't connect to db %s", err)
	}
	db.SetMaxOpenConns(c.MaxOpenConn)
	driver := entsql.OpenDB(c.Type, db)
	return driver
}

// MysqlDSN returns mysql DSN.
func (c *DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True%s", c.Username, c.Password, c.Host, c.Port, c.DBName, c.MysqlConfig)
}

// PostgresDSN returns Postgres DSN.
func (c *DatabaseConf) PostgresDSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s%s", c.Username, c.Password, c.Host, c.Port, c.DBName,
		c.SSLMode, c.PGConfig)
}

// SqliteDSN returns Sqlite DSN.
func (c *DatabaseConf) SqliteDSN() string {
	if c.DBPath == "" {
		log.Error(errors.New("the database file path cannot be empty").Error())
	}

	if _, err := os.Stat(c.DBPath); os.IsNotExist(err) {
		f, err := os.OpenFile(c.DBPath, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			log.Error(fmt.Errorf("failed to create SQLite database file %q", c.DBPath).Error())
		}
		if err := f.Close(); err != nil {
			log.Error(fmt.Errorf("failed to create SQLite database file %q", c.DBPath).Error())
		}
	} else {
		if err := os.Chmod(c.DBPath, 0660); err != nil {
			log.Error(fmt.Errorf("unable to set permission code on %s: %v", c.DBPath, err).Error())
		}
	}

	return fmt.Sprintf("file:%s?_busy_timeout=100000&_fk=1%s", c.DBPath, c.SqliteConfig)
}

// GetDSN returns DSN according to the database type.
func (c *DatabaseConf) GetDSN() string {
	switch c.Type {
	case "mysql":
		return c.MysqlDSN()
	case "postgres":
		return c.PostgresDSN()
	case "sqlite3":
		return c.SqliteDSN()
	default:
		return "mysql"
	}
}

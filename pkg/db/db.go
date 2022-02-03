package db

import (
	"fmt"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	envars "github.com/netflix/go-env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	*gorm.DB
}

type Config interface {
	postgresCfg() postgres.Config
}
type configProd struct {
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func NewCloudCfg() (*configProd, error) {
	var cfg configProd
	if _, err := envars.UnmarshalFromEnviron(&cfg); err != nil {
		return nil, fmt.Errorf("could not unmarshal database configuration envars: %s", err)
	}
	return &cfg, nil
}

func (c *configProd) postgresCfg() postgres.Config {
	return postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN: fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
			c.Host, c.User, c.Name, c.Password),
	}
}

type configDev struct {
	Host     string `env:"DB_HOST,default=db"`
	Port     int    `env:"DB_PORT,default=5432"`
	Name     string `env:"DB_NAME,default=iconophilos_db"`
	User     string `env:"DB_USER,default=user"`
	Password string `env:"DB_PASSWORD,default=password"`
}

func NewLocalCfg() (*configDev, error) {
	var cfg configDev
	if _, err := envars.UnmarshalFromEnviron(&cfg); err != nil {
		return nil, fmt.Errorf("could not unmarshal database configuration envars: %s", err)
	}
	return &cfg, nil
}

func (c *configDev) postgresCfg() postgres.Config {
	return postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s sslmode=disable TimeZone=Europe/London",
			c.Host, c.User, c.Name, c.Port, c.Password),
	}
}

func Connection(cfg Config) (*Conn, error) {
	connection, err := gorm.Open(
		postgres.New(cfg.postgresCfg()),
	)

	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %s", err)
	}

	return &Conn{connection}, nil
}

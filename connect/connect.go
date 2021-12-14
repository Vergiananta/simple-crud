package connect

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	cfg *viper.Viper
	db  *sql.DB
	err error
)

type Connect interface {
	SqlDb() *sql.DB
	Config() *viper.Viper
	ApiServer() string
}

type connect struct{}

func (i *connect) SqlDb() *sql.DB {
	dbUser := i.Config().GetString("DB_USER")
	dbPassword := i.Config().GetString("DB_PASSWORD")
	dbHost := i.Config().GetString("DB_HOST")
	dbPort := i.Config().GetString("DB_PORT")
	dbName := i.Config().GetString("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection established")
	}

	return db
}

func (c *connect) Config() *viper.Viper {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	cfg = viper.GetViper()
	return cfg
}

func (c *connect) ApiServer() string {
	//TODO implement me
	panic("implement me")
}

func NewConnect() Connect {
	return &connect{}
}

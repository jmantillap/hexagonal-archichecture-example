package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/glebarez/go-sqlite"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func NewMySQLDB() (*sql.DB,error) {

	viper.SetConfigFile("./infraestructure/database/config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
        panic(fmt.Errorf("error al leer archivo de configuración: %s", err))
    }
	// Obtén los valores de configuración
    host := viper.GetString("database.mysql.host")
    port := viper.GetInt("database.mysql.port")
    user := viper.GetString("database.mysql.user")
    password := viper.GetString("database.mysql.password")
    dbname := viper.GetString("database.mysql.dbname")	
	
	if os.Getenv("MYSQL_PASSWORD")!= "" {
		fmt.Println("variable de entorno Password: " +os.Getenv("MYSQL_PASSWORD"))
		password = os.Getenv("MYSQL_PASSWORD")
	}
	
	fmt.Println("Password: " + password)
	// Usa los valores de configuración para conectarte a MySQL
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql",dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewSQLite3() (*sql.DB,error) {
	//file := "./infraestructure/database/sqlite/hexagonal.sqlite"
	viper.SetConfigFile("./infraestructure/database/config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
        panic(fmt.Errorf("error al leer archivo de configuración: %s", err))
    }
	file := viper.GetString("database.sqlite.file")


	fmt.Println("Se conectara a sqlite  ")
	db, err := sql.Open("sqlite", file)
	if err!= nil {
        return nil, err
    }
	return db, nil
}

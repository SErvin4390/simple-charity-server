package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // needed for side effects
	"github.com/jmoiron/sqlx"
)

// Config is the global configuration object
var Config *ConfigStruct

// ConfigStruct holds the various configuration options
type ConfigStruct struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
	DbConn     *sqlx.DB
	Port       string
}

// ConfigSetup sets up the config struct with data from the environment
func ConfigSetup() *ConfigStruct {
	if Config != nil {
		return Config
	}
	c := new(ConfigStruct)
	rand.Seed(time.Now().UnixNano())
	// setup the paths

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	c.Port = fmt.Sprintf(":%s", port)

	c.dbUser = os.Getenv("DB_USER")
	c.dbPassword = os.Getenv("DB_PASSWORD")
	c.dbHost = os.Getenv("DB_HOST")
	c.dbPort = os.Getenv("DB_PORT")
	c.dbName = os.Getenv("DB_NAME")

	if c.dbUser == "" {
		c.dbUser = "root"
	}
	if c.dbPassword == "" {
		c.dbPassword = "password"
	}
	if c.dbHost == "" {
		c.dbHost = "localhost"
	}
	if c.dbPort == "" {
		c.dbPort = "3306"
	}
	if c.dbName == "" {
		c.dbName = "SimpleCharity"
	}

	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.dbUser,
		c.dbPassword, c.dbHost, c.dbPort, c.dbName)

	// setup the DB
	conn, err := sqlx.Open("mysql", dbConnectionString)
	if err != nil {
		fmt.Printf("\nERROR FOUND\n%v\n\n", err)
		panic(err)
	}
	conn.SetMaxIdleConns(100)

	// check the db
	_, err = conn.Exec("set session time_zone ='-0:00'")
	maxTries := 10
	secondsToWait := 5
	if err != nil {
		// we try again every X second for Y times, if it's still bad, we panic
		for i := 1; i <= maxTries; i++ {
			fmt.Printf("\n\tDB Error, this is attempt %d of %d. Waiting %d seconds...\n", i, maxTries, secondsToWait)
			time.Sleep(time.Duration(secondsToWait) * time.Second)
			_, err := conn.Exec("set session time_zone ='-0:00'")
			if err == nil {
				break
			}
			if i == maxTries {
				panic("Could not connect to the database, shutting down")
			}
		}
	}

	c.DbConn = conn
	Config = c
	return c
}

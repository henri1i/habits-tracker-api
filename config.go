package main

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
    db_connection string
}

func NewConfig() (*Config, error) {
    db_connection, db_connection_err := getVar("DB_CONNECTION_URL")
  
    if db_connection_err != nil {
        return nil, db_connection_err
    }

    return &Config {
        db_connection,
    }, nil
}

func getVar(variable string) (string, error) {
    value, found := os.LookupEnv(variable)
    
    if !found {
        return "", errors.New(fmt.Sprintf("ENV<%s> not found", variable)) 
    }

    return value, nil
}


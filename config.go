package main

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
    db_connection string
    db_name string
    db_username string
    db_password string
}

func NewConfig() (*Config, error) {
    db_connection, db_connection_err := getVar("DB_CONNECTION_URL")
    db_name, db_name_err := getVar("DB_CONNECTION_URL")
    db_username, db_username_err := getVar("DB_CONNECTION_URL")
    db_password, db_password_err := getVar("DB_CONNECTION_URL")
   
    if db_connection_err != nil {
        return nil, db_connection_err
    }

    if db_name_err != nil {
        return nil, db_name_err 
    }

    if db_username_err != nil {
        return nil, db_username_err 
    }

    if db_password_err != nil {
        return nil, db_password_err
    }

    return &Config {
        db_connection,
        db_name,
        db_username,
        db_password,
    }, nil
}

func getVar(variable string) (string, error) {
    value, found := os.LookupEnv(variable)
    
    if !found {
        return "", errors.New(fmt.Sprintf("ENV<%s> not found")) 
    }

    return value, nil
}


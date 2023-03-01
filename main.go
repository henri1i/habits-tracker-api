package main

import "fmt"

func main() {
    config, err := NewConfig()

    fmt.Println(config)

    if err != nil {
         panic(err)
    }

    repos, err := concreteRepositories(config)

    server := NewServer("localhost:3000", repos)
    server.Listen()
}

func concreteRepositories(config *Config) (*Repositories, error) {
    habitRepository, err := NewHabitPgRepository(config.db_connection)

    if err != nil {
        return nil, err
    }

    return &Repositories{
        habits: habitRepository,
    }, nil
}


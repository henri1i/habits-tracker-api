package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

)
type Repositories struct {
    habits HabitRepository
}

type HabitRepository interface {
    Create(name string, status string, kind string) error
    Toggle(id string) error
    FetchOne(id string) (Habit, error)
    FetchAll() ([]Habit, error)
    Delete(id string) error
}

type HabitPgRepository struct {
    db *sql.DB
}

func NewHabitPgRepository(connection string) (*HabitPgRepository, error) {
    db, err := sql.Open("postgres", connection)

    if err != nil {
        log.Println("Failed to get connection string")

        return nil, err
    }
   
    return &HabitPgRepository{ db: db }, nil
}

func (r HabitPgRepository) Create(name string, status string, kind string) error {
    return nil
}

func (r HabitPgRepository) Toggle(id string) error {
    return nil
}

func (r HabitPgRepository) FetchOne(id string) (Habit, error) {
    habit := Habit {}

    return habit, nil
}

func (r HabitPgRepository) FetchAll() ([]Habit, error) {
    habits := []Habit{}

    return habits, nil
}

func (r HabitPgRepository) Delete(id string) error {
    return nil
}


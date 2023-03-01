package main

type Repositories struct {
    habits HabitRepository
}

func NewRepositories() Repositories {
    return Repositories{
        habits: HabitPgRepository{},
    }
}

type HabitRepository interface {
    Create(name string, status string, kind string) error
    Toggle(id string) error
    FetchOne(id string) (Habit, error)
    FetchAll() ([]Habit, error)
    Delete(id string) error
}

type HabitPgRepository struct {}

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


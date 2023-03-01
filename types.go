package main

type Kind int8

const (
    Good Kind = iota
    Bad
)

type Habit struct {
    Name string `json:"name"`
    Status bool `json:"status"`
    Kind Kind `json:"kind"`
}


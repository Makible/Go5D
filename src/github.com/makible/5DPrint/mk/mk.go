package mk

import "time"

type StdMovement struct {
	Axis     string
	Distance int `json: ", string"`
	Speed    int `json: ", string"`
}

type MultiMovement struct {
	Axis     string
	Distance string
	Speed    int `json: ", string"`
}

type Position struct {
	X  int
	Y  int
	Z  int
	E1 int
}

type Tool struct {
	Name  string
	Value int `json: ", string"`
}

type Job struct {
	Name      string
	Data      string
	Status    string
	StartTime time.Time
	RunTime   time.Duration
}

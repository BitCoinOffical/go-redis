package storage

import "time"

type Storage map[string]Item

type Item struct {
	Data string
	TTL  *time.Time
}

func NewStorage() Storage {
	m := make(map[string]Item)
	return m
}

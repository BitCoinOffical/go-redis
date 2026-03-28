package storage

import "time"

type Storage map[string][]Item

//Item{
//	"A": {{data,ttl},{data,ttl},{data,ttl}}
//	"B": {{data,ttl},{data,ttl},{data,ttl}}
//}

type Item struct {
	Data string
	TTL  *time.Time
}

func NewStorage() Storage {
	m := make(map[string][]Item)
	return m
}

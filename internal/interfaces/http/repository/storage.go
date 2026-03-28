package repository

import (
	"errors"
	"log"
	"time"

	"github.com/codecrafters-io/redis-starter-go/internal/adapters/secondary/storage"
	"github.com/codecrafters-io/redis-starter-go/internal/dto"
	"github.com/codecrafters-io/redis-starter-go/internal/parser"
)

const (
	NotFound = "$-1\r\n"
)

type Repo struct {
	m storage.Storage
}

func NewRepository(m storage.Storage) *Repo {
	return &Repo{m: m}
}

func (m *Repo) SetRepo(setDTO *dto.SetDTO, items []storage.Item, TTL *time.Time) string {
	data := []storage.Item{}

	for _, item := range items {
		data = append(data, item)
	}

	for _, v := range setDTO.Values {
		data = append(data, storage.Item{
			Data: v,
			TTL:  TTL,
		})
	}

	log.Println("Incoming DATA: ", data)
	m.m[setDTO.Key] = data
	return parser.SimpleString("OK")
}

func (m *Repo) GetRepo(key string) ([]storage.Item, error) {
	v, ok := m.m[key]
	if !ok {
		return nil, errors.New(NotFound)
	}

	for idx, val := range v {
		if val.TTL != nil {
			ok = time.Now().Before(*val.TTL)
			if !ok {
				slice := v
				v = append(slice[:idx], slice[idx+1:]...)
			}
		}
	}

	if len(v) == 0 {
		return nil, errors.New(NotFound)
	}

	return v, nil
}

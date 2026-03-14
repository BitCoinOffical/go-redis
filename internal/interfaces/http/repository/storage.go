package repository

import (
	"time"

	"github.com/codecrafters-io/redis-starter-go/internal/adapters/secondary/storage"
	"github.com/codecrafters-io/redis-starter-go/internal/parser"
)

type Repo struct {
	m storage.Storage
}

func NewRepository(m storage.Storage) *Repo {
	return &Repo{m: m}
}

func (m *Repo) SetRepo(key, value string, TTL *time.Time) string {
	if TTL != nil {
		m.m[key] = storage.Item{
			Data: value,
			TTL:  TTL,
		}
		return parser.SimpleString("OK")
	}
	m.m[key] = storage.Item{
		Data: value,
	}
	return parser.SimpleString("OK")
}

func (m *Repo) GetRepo(key string) *storage.Item {
	v, ok := m.m[key]
	if !ok {
		return &storage.Item{
			Data: "$-1\r\n",
		}
	}

	ok = time.Now().Before(*v.TTL)
	if !ok {
		delete(m.m, key)
		return &storage.Item{
			Data: "$-1\r\n",
		}
	}
	return &v
}

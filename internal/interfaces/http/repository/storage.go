package repository

import (
	"log"
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
		TTL:  nil,
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
	log.Println("нашли", v)

	ok = time.Now().Before(*v.TTL)
	if !ok && v.TTL != nil {
		log.Println("нашли но время истекло", v, v.TTL)
		delete(m.m, key)
		return &storage.Item{
			Data: "$-1\r\n",
		}
	}
	log.Println("возврощаем", v)
	return &v
}

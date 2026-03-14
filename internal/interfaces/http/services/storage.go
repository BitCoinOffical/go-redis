package services

import (
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/internal/adapters/secondary/storage"
	"github.com/codecrafters-io/redis-starter-go/internal/interfaces/http/repository"
)

type Service struct {
	repo *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) SetService(key, value string, options ...string) (string, error) {
	var term time.Time
	if len(options) > 0 {
		switch strings.ToUpper(options[0]) {
		case "PX":
			millisecond, err := strconv.Atoi(options[1])
			if err != nil {
				return "", ErrorConversion
			}
			term = time.Now().Add(time.Duration(millisecond) * time.Millisecond)
		case "EX":
			second, err := strconv.Atoi(options[1])
			if err != nil {
				return "", ErrorConversion
			}
			term = time.Now().Add(time.Duration(second) * time.Second)
		default:
			return "", ErrorInvalidOptions
		}
	}
	return s.repo.SetRepo(key, value, &term), nil
}

func (s *Service) GetService(key string) *storage.Item {
	return s.repo.GetRepo(key)
}

package services

import (
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/internal/adapters/secondary/storage"
	"github.com/codecrafters-io/redis-starter-go/internal/dto"
	"github.com/codecrafters-io/redis-starter-go/internal/interfaces/http/repository"
)

const (
	NotFound = "$-1\r\n"
)

type Service struct {
	repo *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) SetService(setDTO *dto.SetDTO) (string, error) {
	if setDTO.Typettl != "" && setDTO.Valuettl != 0 {
		switch strings.ToUpper(setDTO.Typettl) {
		case "PX":
			term := time.Now().Add(time.Duration(setDTO.Valuettl) * time.Millisecond)

			items, err := s.GetService(setDTO.Key)
			if err == nil {
				return s.repo.SetRepo(setDTO, items, &term), nil
			}

			return s.repo.SetRepo(setDTO, nil, &term), nil

		case "EX":

			term := time.Now().Add(time.Duration(setDTO.Valuettl) * time.Second)

			items, err := s.GetService(setDTO.Key)
			if err != nil {
				return s.repo.SetRepo(setDTO, items, &term), nil
			}

			return s.repo.SetRepo(setDTO, nil, &term), nil
		default:
			return "", ErrorInvalidOptions
		}
	}

	items, err := s.GetService(setDTO.Key)
	if err == nil {
		return s.repo.SetRepo(setDTO, items, nil), nil
	}
	return s.repo.SetRepo(setDTO, nil, nil), nil
}

func (s *Service) GetService(key string) ([]storage.Item, error) {
	return s.repo.GetRepo(key)
}

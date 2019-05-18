package service

import (
	"github.com/tommytan/garen/internal/dao"
)

// Service struct
type Service struct {
	Dao *dao.Dao
}

// New new a service and return.
func New() (s *Service) {
	s = &Service{
		Dao: dao.New(),
	}
	return s
}

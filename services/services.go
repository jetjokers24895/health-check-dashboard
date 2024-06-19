package services

import (
	"app/constants"
	"app/repositories"

	"app/dtos"

	"gorm.io/gorm"
)

type IService interface {
	AddService(input *dtos.Service) error
	UpdateService(input *dtos.Service) error
	DeleteService(id uint) error
	GetServices(status constants.Status) ([]*dtos.ServiceResponse, error) // get all
}

type service struct {
	_repo repositories.IServiceRepository
}

func NewServices(db *gorm.DB) IService {
	return &service{
		_repo: repositories.NewServiceRepository(db),
	}
}

func (s *service) AddService(input *dtos.Service) error {
	return s._repo.AddService(input)
}

func (s *service) UpdateService(input *dtos.Service) error {
	return s._repo.UpdateService(input)
}

func (s *service) DeleteService(id uint) error {
	return s._repo.DeleteService(id)
}

func (s *service) GetServices(status constants.Status) ([]*dtos.ServiceResponse, error) {
	entities, err := s._repo.GetServices(status)
	if err != nil {
		return nil, err
	}
	var rs = []*dtos.ServiceResponse{}
	for _, entity := range entities {
		rs = append(rs, &dtos.ServiceResponse{
			ID:            entity.ID,
			Name:          entity.Name,
			Command:       entity.Command,
			LastCheckTime: entity.LastCheckTime,
			Status:        int(entity.Status),
		})
	}
	return rs, nil
}

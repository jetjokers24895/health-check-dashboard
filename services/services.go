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
	_repo          repositories.IServiceRepository
	cronJobService CronJobManagerService
}

func NewServices(db *gorm.DB, cronJobService CronJobManagerService) IService {
	return &service{
		_repo:          repositories.NewServiceRepository(db),
		cronJobService: cronJobService,
	}
}

func (s *service) AddService(input *dtos.Service) error {
	id, err := s._repo.AddService(input)
	if err != nil {
		return err
	}
	s.cronJobService.Add(&Job{
		ServiceID:   id,
		ServiceName: input.Name,
		URL:         input.URL,
		_repo:       s._repo,
	})
	return nil
}

func (s *service) UpdateService(input *dtos.Service) error {
	if err := s._repo.UpdateService(input); err != nil {
		return err
	}
	s.cronJobService.Update(&Job{
		ServiceID:   input.ID,
		ServiceName: input.Name,
		URL:         input.URL,
		_repo:       s._repo,
	})
	return nil
}

func (s *service) DeleteService(id uint) error {
	if err := s._repo.DeleteService(id); err != nil {
		return err
	}

	s.cronJobService.Remove(&Job{
		ServiceID: id,
	})
	return nil
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
			URL:           entity.URL,
			LastCheckTime: entity.LastCheckTime,
			Status:        int(entity.Status),
		})
	}
	return rs, nil
}

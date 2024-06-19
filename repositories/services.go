package repositories

import (
	"app/constants"
	"app/dtos"
	"app/models"

	"gorm.io/gorm"
)

type IServiceRepository interface {
	AddService(input *dtos.Service) error
	UpdateService(input *dtos.Service) error
	DeleteService(id uint) error
	GetServices(status constants.Status) ([]*models.Services, error) // all
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) IServiceRepository {
	return &serviceRepository{
		db: db,
	}
}

func (repo *serviceRepository) AddService(input *dtos.Service) error {
	return repo.db.Save(&models.Services{
		Name:    input.Name,
		Command: input.Command,
	}).Error
}

func (repo *serviceRepository) UpdateService(input *dtos.Service) error {
	return repo.db.Save(&models.Services{
		Model: gorm.Model{
			ID: input.ID,
		},
		Name:    input.Name,
		Command: input.Command,
	}).Error
}

func (repo *serviceRepository) DeleteService(id uint) error {
	return repo.db.Delete(&models.Services{}, id).Error
}

func (repo *serviceRepository) GetServices(status constants.Status) ([]*models.Services, error) {
	var entities = []*models.Services{}

	statement := repo.db.Limit(-1).Order("id desc") // get all services
	if status != 0 {
		statement = statement.Where("status = ?", status)
	}

	if err := statement.Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

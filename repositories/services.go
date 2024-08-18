package repositories

import (
	"app/constants"
	"app/dtos"
	"app/models"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type IServiceRepository interface {
	AddService(input *dtos.Service) (uint, error)
	UpdateService(input *dtos.Service) error
	DeleteService(id uint) error
	GetServices(status constants.Status) ([]*models.Services, error) // all
	UpdateStatus(id uint, status constants.Status, logEntity *models.LogChecked) error
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) IServiceRepository {
	return &serviceRepository{
		db: db,
	}
}

func (repo *serviceRepository) AddService(input *dtos.Service) (uint, error) {
	newEntity := models.Services{
		Name: input.Name,
		URL:  input.URL,
	}
	if err := repo.db.Save(&newEntity).Error; err != nil {
		return 0, err
	}

	return newEntity.ID, nil
}

func (repo *serviceRepository) UpdateService(input *dtos.Service) error {
	return repo.db.Save(&models.Services{
		Model: gorm.Model{
			ID: input.ID,
		},
		Name: input.Name,
		URL:  input.URL,
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

func (repo *serviceRepository) UpdateStatus(id uint, status constants.Status, logEntity *models.LogChecked) error {

	err := repo.db.Save(&models.Services{
		Model: gorm.Model{
			ID: id,
		},
		Status:        int(status),
		LastCheckTime: time.Now().Unix(),
	}).Error // get all services
	if err != nil {
		log.Error(err)
		return err
	}

	if err := repo.db.Save(&models.LogChecked{
		HttpStatus:  logEntity.HttpStatus,
		ResponseTXT: logEntity.ResponseTXT,
		Status:      int(status),
		ServicesID:  logEntity.ServicesID,
	}).Error; err != nil {
		log.Error(err)
		return err
	}

	return nil
}

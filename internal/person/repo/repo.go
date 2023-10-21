package repo

import (
	"context"
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person"
	"github.com/VolodyaLarin/rsoi-lab-01/internal/utils"
	"gorm.io/gorm"
)

type PersonRepoImpl struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) *PersonRepoImpl {
	return &PersonRepoImpl{db: db}
}

func (p PersonRepoImpl) Filter(ctx context.Context, filter person.RepoFilter) (error, []person.PersonDTO) {
	var persons []PersonModel
	err := p.db.Limit(int(filter.Limit)).Offset(int(filter.From)).Find(&persons).Error
	if err != nil {
		return err, nil
	}

	res := utils.Map(persons, func(person PersonModel) person.PersonDTO {
		return person.toDTO()
	})

	return nil, res
}

func (p PersonRepoImpl) GetById(ctx context.Context, id int32) (error, *person.PersonDTO) {
	var model PersonModel
	err := p.db.First(&model, id).Error
	if err != nil {
		return err, nil
	}

	dto := model.toDTO()
	return nil, &dto

}

func (p PersonRepoImpl) CreatePerson(ctx context.Context, person person.PersonDTO) (error, *person.PersonDTO) {
	model := PersonModel{}
	model.fromDTO(person)
	model.ID = 0

	err := p.db.Create(&model).Error

	if err != nil {
		return err, nil
	}

	dto := model.toDTO()
	return nil, &dto
}

func (p PersonRepoImpl) Update(ctx context.Context, id int32, person person.PersonDTO) (error, *person.PersonDTO) {
	model := PersonModel{}
	model.fromDTO(person)
	model.ID = id

	err := p.db.Save(&model).Error
	if err != nil {
		return err, nil
	}

	dto := model.toDTO()
	return nil, &dto
}

func (p PersonRepoImpl) DeleteById(ctx context.Context, id int32) error {
	return p.db.Delete(&PersonModel{}, id).Error
}

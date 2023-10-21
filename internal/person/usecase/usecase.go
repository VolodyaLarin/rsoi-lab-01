package usecase

import (
	"context"
	"errors"
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person"
)

type PersonUsecase struct {
	repo person.RepoInterface
}

func NewPersonUsecase(repo person.RepoInterface) *PersonUsecase {
	return &PersonUsecase{repo: repo}
}

func (uc PersonUsecase) Filter(ctx context.Context) (error, []person.PersonDTO) {
	return uc.repo.Filter(ctx, person.RepoFilter{
		Limit: 1000,
		From:  0,
	})
}

func (uc PersonUsecase) GetById(ctx context.Context, id int32) (error, *person.PersonDTO) {
	return uc.repo.GetById(ctx, id)
}

func (uc PersonUsecase) PatchPerson(ctx context.Context, id int32, dto person.PersonDTO) (error, *person.PersonDTO) {
	err, pers := uc.repo.GetById(ctx, id)

	if err != nil || pers == nil {
		return errors.New("not found"), nil
	}

	if dto.Age != nil {
		pers.Age = dto.Age
	}
	if dto.Name != "" {
		pers.Name = dto.Name
	}
	if dto.Address != nil {
		pers.Address = dto.Address
	}
	if dto.Work != nil {
		pers.Work = dto.Work
	}

	err, p := uc.repo.Update(ctx, id, *pers)
	if err != nil {
		return err, nil
	}

	return nil, p
}

func (uc PersonUsecase) DeletePerson(ctx context.Context, id int32) error {
	return uc.repo.DeleteById(ctx, id)
}

func (uc PersonUsecase) CreatePerson(ctx context.Context, dto person.PersonDTO) (error, *person.PersonDTO) {
	return uc.repo.CreatePerson(ctx, dto)
}

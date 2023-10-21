package repo

import (
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person"
)

type PersonModel struct {
	ID      int32
	Name    string
	Age     *int32
	Address *string
	Work    *string
}

func (p *PersonModel) TableName() string {
	return "person"
}

func (p *PersonModel) toDTO() person.PersonDTO {
	personModelCopy := *p
	return person.PersonDTO{
		Id:      personModelCopy.ID,
		Name:    personModelCopy.Name,
		Age:     personModelCopy.Age,
		Address: personModelCopy.Address,
		Work:    personModelCopy.Work,
	}
}

func (p *PersonModel) fromDTO(dto person.PersonDTO) *PersonModel {
	if dto.Id != 0 {
		p.ID = dto.Id
	}
	if dto.Name != "" {
		p.Name = dto.Name
	}
	if dto.Age != nil {
		p.Age = dto.Age
	}
	if dto.Address != nil {
		p.Address = dto.Address
	}
	if dto.Work != nil {
		p.Work = dto.Work
	}

	return p
}

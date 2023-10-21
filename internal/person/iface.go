package person

import (
	"context"
)

type RepoFilter struct {
	Limit int64
	From  int64
}

type RepoInterface interface {
	Filter(ctx context.Context, filter RepoFilter) (error, []PersonDTO)
	GetById(ctx context.Context, id int32) (error, *PersonDTO)

	CreatePerson(ctx context.Context, person PersonDTO) (error, *PersonDTO)
	Update(ctx context.Context, id int32, person PersonDTO) (error, *PersonDTO)
	DeleteById(ctx context.Context, id int32) error
}

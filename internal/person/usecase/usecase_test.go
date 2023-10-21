package usecase

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person"
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person/repo"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func NewMockUc() (*PersonUsecase, sqlmock.Sqlmock) {
	mockDb, mock, _ := sqlmock.New()
	db, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	}), &gorm.Config{})

	rep := repo.NewPersonRepo(db)
	return NewPersonUsecase(rep), mock
}

func TestPersonUsecase_CreatePerson(t *testing.T) {
	uc, mock := NewMockUc()
	ctx := context.Background()
	work := "Cool"

	mock.ExpectBegin()
	addRow := sqlmock.NewRows([]string{"id"}).AddRow("34")
	mock.ExpectQuery("INSERT INTO \"person\"").WithArgs("Hell", nil, nil, "Cool").WillReturnRows(addRow)
	mock.ExpectCommit()

	err, dto := uc.CreatePerson(ctx, person.PersonDTO{Name: "Hell", Work: &work})

	assert.Nil(t, err, "Should't error", err)
	assert.NotNil(t, dto, "Should't nil dto")

	assert.Equal(t, dto.Id, int32(34))
	assert.Equal(t, dto.Name, "Hell")
	assert.Nil(t, dto.Age)
	assert.Nil(t, dto.Address)
	assert.Equal(t, *dto.Work, "Cool")
}

func TestPersonUsecase_PatchPerson(t *testing.T) {
	uc, mock := NewMockUc()
	ctx := context.Background()
	work := "Cool"

	//addRow := sqlmock.NewRows([]string{"id"}).AddRow("34")
	mock.ExpectQuery("SELECT (.+) FROM \"person\" WHERE (.+)").WithArgs(34).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "age", "address", "work"}).AddRow("34", "Hell", "15", nil, nil),
	)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"person\" SET .+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err, dto := uc.PatchPerson(ctx, 34, person.PersonDTO{Work: &work})

	assert.Nil(t, err, "Should't error", err)
	assert.NotNil(t, dto, "Should't nil dto")

	assert.Equal(t, dto.Id, int32(34))
	assert.Equal(t, dto.Name, "Hell")
	assert.Equal(t, *dto.Age, int32(15))
	assert.Nil(t, dto.Address)
	assert.Equal(t, *dto.Work, "Cool")
}

func TestPersonUsecase_PatchPersonNil(t *testing.T) {
	uc, mock := NewMockUc()
	ctx := context.Background()
	work := "Cool"

	//addRow := sqlmock.NewRows([]string{"id"}).AddRow("34")
	mock.ExpectQuery("SELECT (.+) FROM \"person\" WHERE (.+)").WithArgs(34).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "age", "address", "work"}),
	)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"person\" SET .+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err, dto := uc.PatchPerson(ctx, 34, person.PersonDTO{Work: &work})

	assert.NotNil(t, err, "Should error")
	assert.Nil(t, dto, "Should't be nil dto")
}

func TestPersonUsecase_GetByIdPerson(t *testing.T) {
	uc, mock := NewMockUc()
	ctx := context.Background()

	mock.ExpectQuery("SELECT (.+) FROM \"person\" WHERE (.+)").WithArgs(34).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "age", "address", "work"}).AddRow("34", "Hell", "15", nil, nil),
	)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"person\" SET .+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err, dto := uc.GetById(ctx, 34)

	assert.Nil(t, err, "Should't error", err)
	assert.NotNil(t, dto, "Should't nil dto")

	assert.Equal(t, dto.Id, int32(34))
	assert.Equal(t, dto.Name, "Hell")
	assert.Equal(t, *dto.Age, int32(15))
	assert.Nil(t, dto.Address)
	assert.Nil(t, dto.Work)
}

func TestPersonUsecase_Filter(t *testing.T) {
	uc, mock := NewMockUc()
	ctx := context.Background()

	mock.ExpectQuery("SELECT (.+) FROM \"person\"").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "age", "address", "work"}).
			AddRow("34", "Hell", "15", nil, nil).
			AddRow("35", "Hell2", "16", nil, nil),
	)

	err, dto := uc.Filter(ctx)

	assert.Nil(t, err, "Should't error", err)
	assert.NotNil(t, dto, "Should't nil dto")
	assert.Equal(t, len(dto), 2)
}

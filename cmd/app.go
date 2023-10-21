package main

import (
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person/handler"
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person/repo"
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person/usecase"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	return r
}

func main() {

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatal("can't start db con")
	}
	err = db.AutoMigrate(repo.PersonModel{})
	if err != nil {
		log.WithError(err).Fatal("can't migrate")
		return
	}

	r := setupRouter()
	apiV1R := r.Group("/api/v1/")

	personRepo := repo.NewPersonRepo(db)
	personUc := usecase.NewPersonUsecase(personRepo)
	handler.NewPersonHandlerV1(personUc).RegisterRoutes(apiV1R)

	r.Run(":8080")
}

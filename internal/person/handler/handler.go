package handler

import (
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person"
	"github.com/VolodyaLarin/rsoi-lab-01/internal/person/usecase"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Errors  struct {
		Details string `json:"details"`
	} `json:"errors"`
}

type PersonHandlerV1 struct {
	uc *usecase.PersonUsecase
	r  gin.IRouter
}

func NewPersonHandlerV1(uc *usecase.PersonUsecase) *PersonHandlerV1 {
	return &PersonHandlerV1{uc: uc}
}

func (p *PersonHandlerV1) RegisterRoutes(router gin.IRouter) {
	p.r = router

	router.GET("/persons/", p.get)
	router.POST("/persons/", p.post)
	router.GET("/persons/:id", p.getById)
	router.PATCH("/persons/:id", p.patch)
	router.DELETE("/persons/:id", p.delete)

}

func (p *PersonHandlerV1) get(ctx *gin.Context) {
	err, persons := p.uc.Filter(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, persons)

}

func (p *PersonHandlerV1) getById(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorMessage{Message: "not found"})
		return
	}

	err, p2 := p.uc.GetById(ctx, int32(id))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, p2)

}

func (p *PersonHandlerV1) post(ctx *gin.Context) {
	pd := person.PersonDTO{}
	err := ctx.ShouldBindJSON(&pd)
	if err != nil || pd.Name == "" {
		ctx.JSON(http.StatusBadRequest, &ErrorMessage{
			Message: "Validation error",
			Errors: struct {
				Details string `json:"details"`
			}{err.Error()},
		})
		return
	}

	err, p2 := p.uc.CreatePerson(ctx, pd)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	link := ctx.FullPath() + strconv.FormatInt(int64(p2.Id), 10)
	log.Warning(link)
	ctx.Header("Location", link)
	ctx.JSON(http.StatusCreated, p2)

}

func (p *PersonHandlerV1) patch(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorMessage{Message: "not found"})
		return
	}

	pd := person.PersonDTO{}
	err = ctx.ShouldBindJSON(&pd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorMessage{
			Message: "Validation error",
			Errors: struct {
				Details string `json:"details"`
			}{err.Error()},
		})
		return
	}

	err, p2 := p.uc.PatchPerson(ctx, int32(id), person.PersonDTO{
		Name:    pd.Name,
		Age:     pd.Age,
		Address: pd.Address,
		Work:    pd.Work,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	link := ctx.FullPath() + strconv.FormatInt(int64(p2.Id), 10)
	ctx.Header("Location", link)
	ctx.JSON(http.StatusOK, p2)
}

func (p *PersonHandlerV1) delete(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &ErrorMessage{Message: "not found"})
		return
	}

	err = p.uc.DeletePerson(ctx, int32(id))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

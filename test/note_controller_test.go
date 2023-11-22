package test

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"net/http/httptest"
	"restful-api-gorm-fiber/controller"
	"restful-api-gorm-fiber/data/response"
	"restful-api-gorm-fiber/repository"
	"restful-api-gorm-fiber/router"
	"restful-api-gorm-fiber/service"
	"strings"
	"testing"
	"time"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:fadel123@tcp(localhost:3356)/golang_fiber?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}

var Db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, Db)
}

func setupRouter(Db *gorm.DB) *fiber.App {
	validate := validator.New()
	noteRepository := repository.NewNoteRepositoryImpl(Db)
	noteService := service.NewNoteServiceImpl(noteRepository, validate)
	noteController := controller.NewNoteController(noteService)
	router := router.NewRouter(noteController)

	return router
}

func TestCreateNoteSucces(t *testing.T) {
	Db := OpenConnection()
	router := setupRouter(Db)
	body := strings.NewReader(`{"content":"udin sholeh 4"}`)
	request := httptest.NewRequest("POST", "/note", body)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	respons, _ := router.Test(request)

	assert.Equal(t, 201, respons.StatusCode)
	responseBody, _ := ioutil.ReadAll(respons.Body)

	webResponse := response.Response{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "ok", webResponse.Status)
}

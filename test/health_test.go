package test

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http/httptest"
	"testing"
	"webapp/controller"
	"webapp/database"
	"webapp/logger"
)

type HealthTestSuite struct {
	suite.Suite
	App *gin.Engine
}

func TestHealthTestSuite(t *testing.T) {
	suite.Run(t, &HealthTestSuite{})
}

func (s *HealthTestSuite) SetupSuite() {

	loadEnv()
	_, err := database.Connect()
	if err != nil {
		return
	}
	logger.InitLogger()
	logger.InitMetrics()

	app := gin.New()
	app.GET("/healthz", controller.Health)
	s.App = app
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Print(".env files not found: ", err)
	} else {
		log.Print("Environment variables loaded successfully!!")
	}
}

func (s *HealthTestSuite) TestIntegrationHealth() {

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	s.App.ServeHTTP(w, req)

	res := w.Code

	s.Equal(200, res)
}

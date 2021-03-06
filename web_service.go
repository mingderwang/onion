// generated by ginger from go generate -- DO NOT EDIT

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tommy351/gin-cors"
)

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
	Token      string
	Url        string
}

type OnionService struct {
}

func (s *OnionService) getDb(cfg Config) (gorm.DB, error) {
	db, err := gorm.Open("sqlite3", cfg.DbName)
	db.LogMode(true)
	return db, err
}

func (s *OnionService) Migrate(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	db.AutoMigrate(&Onion{})
	return nil
}
func (s *OnionService) Run(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	onionResource := &OnionResource{db: db}

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.Use(cors.Middleware(cors.Options{}))

	r.GET("/onion", onionResource.GetAllOnions)
	r.GET("/onion/:id", onionResource.GetOnion)
	r.GET("/onion/:id/build", onionResource.BuildOnion)
	r.POST("/onion", onionResource.CreateOnion)
	r.PUT("/onion/:id", onionResource.UpdateOnion)
	r.PATCH("/onion/:id", onionResource.PatchOnion)
	r.DELETE("/onion/:id", onionResource.DeleteOnion)

	r.Run(cfg.SvcHost)

	return nil
}

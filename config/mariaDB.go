package config

import (
	"database/sql"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

type DatabaseInterface interface {
	Connection() *gorm.DB
}

type mariaDB struct {
	mariaDB databaseConfig `yaml:"mariaDB"`
}
type databaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Id       string `yaml:"id"`
	Pass     string `yaml:"pass"`
	Database string `yaml:"database"`
}

func (c *databaseConfig) newMariaDBconfig() *databaseConfig {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (config *databaseConfig) Connection() {

	db, err := sql.Open("mysql",
		config.Id+":"+config.Pass+"@tcp("+config.Host+":"+config.Port+")/"+config.Database)

	// 최대 대기 커넥션 , 최대 오픈 커넥션
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)

	// 커넥션 유지 최대 유지 시간을 1시간으로 설정
	db.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

package repository

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ConnectionInfo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Id       string `yaml:"id"`
	Pass     string `yaml:"pass"`
	Database string `yaml:"database"`
}

func (c *ConnectionInfo) getConnectionInfo() *ConnectionInfo {

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

func (aaa string) maria() {

	var c ConnectionInfo
	c.getConnectionInfo()
	fmt.Println(c)
	fmt.Print(aaa)
	// conn, err := sql.Open("mysql", "root:dkssud123@tcp(127.0.0.1:3306)/test")
	// if err != nil {
	// 	fmt.Println(err)
	// 	//os.Exit(1)
	// }

	// // use your own select statement
	// // this is just an example statement
	// statement, err := conn.Prepare("select level from keyword")

	// if err != nil {
	// 	fmt.Println(err)
	// 	//os.Exit(1)
	// }

	// rows, err := statement.Query() // execute our select statement

	// if err != nil {
	// 	fmt.Println(err)
	// 	//os.Exit(1)
	// }

	// for rows.Next() {
	// 	var level string
	// 	rows.Scan(&level)
	// 	fmt.Println("level is :", level)
	// }

	// defer conn.Close()

}

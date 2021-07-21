package main

import (
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"gopkg.in/yaml.v2"
)

type Yaml struct {
    Schema     string `yaml:"schema"`
    ID         string `yaml:"id"`
    Version    string `yaml:"version"`
    User []User
	Url []Url
}

type User struct {
    Name     string
    Login  	 string 
	Password string
}

type Url struct {
    Login 	string 
	List_Accounts string
}

type Path struct {
	Selenium string `yaml:"selenium"`
    Gecko  	 string `yaml:"geckoDriver"`
	Chrome   string `yaml:"chromeDriver"`
}

var y *Yaml

func load() *Yaml {
	//y := Yaml{}
	if y == nil {

		yamlFile, err := ioutil.ReadFile("config.yaml") 
		if err != nil {
			log.Printf("yamlFile.Get err  #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, &y) 

		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	return y
}

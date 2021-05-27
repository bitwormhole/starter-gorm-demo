package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"

	etc_demo "github.com/bitwormhole/starter-gorm-demo/etc"
	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
)

//go:embed src/main/resources
var resources embed.FS

func main() {
	builder := config.NewBuilderFS(&resources, "src/main/resources")

	etc_demo.Config(builder)

	err := tryMain(builder)
	if err != nil {
		panic(err)
	}
}

func tryMain(cb application.ConfigBuilder) error {

	context, err := application.Run(cb.Create(), os.Args)
	if err != nil {
		return err
	}

	err = doAccessDB(context)
	if err != nil {
		return err
	}

	err = application.Loop(context)
	if err != nil {
		return err
	}

	code, err := application.Exit(context)
	if err != nil {
		return err
	}

	fmt.Println("exit with code:", code)
	return nil
}

func doAccessDB(context application.Context) error {

	src_o, _ := context.FindComponent("#db1")
	src := src_o.(datasource.Source)
	db := src.DB()

	type ENT struct {
		ID   string
		UUID string
		Unit string
	}
	var ent ENT
	db.Table("demo_article_entity").First(&ent, 9)

	js, err := json.Marshal(&ent)
	if err == nil {
		jstr := string(js)
		fmt.Println("JSON:", jstr)
	}

	return nil
}

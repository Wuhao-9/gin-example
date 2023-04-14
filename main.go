package main

import (
	"fmt"
	_ "github/Wuhao-9/go-gin-example/models"
	"github/Wuhao-9/go-gin-example/pkg/setting"
	"github/Wuhao-9/go-gin-example/router"
)

func main() {
	router.InitRouter().Run(fmt.Sprintf(":%d", setting.HTTPPort))
}

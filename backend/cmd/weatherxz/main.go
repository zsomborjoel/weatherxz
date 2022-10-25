package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/zsomborjoel/weatherxz/pkg/common"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	common.LoadEnvVariables()

	db := common.Init()
	common.Migrate(db)

	r := gin.Default()

	r.Run()
}

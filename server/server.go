package server

import (
	"transactional/config"
	"transactional/routes"
)

func Init() {
	r := routes.NewRouter()
	r.Run(":" + config.Get().PORT)
}

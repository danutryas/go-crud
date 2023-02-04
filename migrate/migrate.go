package main

import (
	"github.com/danutryas/go-crud/initializers"
	model "github.com/danutryas/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}

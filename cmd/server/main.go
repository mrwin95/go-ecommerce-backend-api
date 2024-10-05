package main

import (
	"github.com/go-ecommerce-backend-api/internal/routes"
)

func main() {
	r := routes.NewRouter()
	r.Run(":8002")
}

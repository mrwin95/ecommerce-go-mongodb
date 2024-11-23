package main

import (
	"mrwin95/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouters()
	r.Run(":8080")
}

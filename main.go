package main

import (
	"fmt"

	"github.com/8soat-grupo35/fastfood-order-production/internal/api/server"
)

func main() {
	fmt.Println("Iniciado o servidor Rest com GO")
	server.Start()
}

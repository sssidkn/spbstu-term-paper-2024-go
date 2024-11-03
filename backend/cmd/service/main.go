package main

import (
	"fmt"
	"github.com/backend/internal/config"
)

func main() {
	cfg := config.Load("config/config.yaml")
	fmt.Println(*cfg)
}

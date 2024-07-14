package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type Order struct {
	ID    string
	Price float64
}

func GenerateOrders() Order {
	return Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100,
	}
}

func main() {
	fmt.Println(GenerateOrders())
}

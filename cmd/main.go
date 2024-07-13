package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EvyOliveira/product-fee/internal/order/infra/database"
	"github.com/EvyOliveira/product-fee/internal/order/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbDriver := "mysql"
	dbUser := "user"
	dbPassword := "root"
	dbName := "orders"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp(172.22.0.1:3306)/"+dbName)
	if err != nil {
		log.Fatal("error pinging database:", err)
	}
	fmt.Println("successfully connected to database!")
	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID:    "456",
		Price: 100,
		Tax:   10,
	}

	output, err := uc.Execute(input)
	if err != nil {
		log.Fatal("error to execute a new order:", err)
	}
	fmt.Printf("the final price is: %f\n", output.FinalPrice)
}

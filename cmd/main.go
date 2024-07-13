package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EvyOliveira/product-fee/internal/order/entity"
	"github.com/EvyOliveira/product-fee/internal/order/infra/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	order, err := entity.NewOrder("123", 10, 2)
	if err != nil {
		log.Fatal("error to create a new order:", err)
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		log.Fatal("error to calculate the final price of new order:", err)
	}

	fmt.Printf("the final price is: %f\n", order.FinalPrice)

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
	err = repository.Save(order)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/EvyOliveira/product-fee/internal/order/infra/database"
	"github.com/EvyOliveira/product-fee/internal/order/usecase"
	rabbitmq "github.com/EvyOliveira/product-fee/pkg/rabbitmq"
	_ "github.com/go-sql-driver/mysql"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	maxWorkers := 3
	waitGroup := sync.WaitGroup{}

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

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		uc := usecase.NewGetTotalUseCase(repository)
		output, err := uc.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(output)
	})
	go http.ListenAndServe(":8181", nil)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		log.Fatal("unable to open a channel with rabbitmq: ", err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)

	for i := 0; i < maxWorkers; i++ {
		waitGroup.Add(1)
		i := i
		go func() {
			fmt.Println("starting worker", i)
			defer waitGroup.Done()
			worker(out, uc, i)
		}()
	}
	waitGroup.Wait()
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {
		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			fmt.Println("error unmarshalling message", err)
		}
		input.Tax = 10.0
		_, err = uc.Execute(input)
		if err != nil {
			fmt.Println("error to execute order", err)
		}
		msg.Ack(false) //notifica o RabbitMQ que já processou a mensagem e que pode retirá-la da fila
		fmt.Println("worker", workerId, "processed order", input.ID)
		time.Sleep(1 * time.Second)
	}
}

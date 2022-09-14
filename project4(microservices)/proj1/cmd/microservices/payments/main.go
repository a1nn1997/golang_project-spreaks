package main

import (
	"fmt"
	"log"
	"os"
	"github.com/go-chi/chi"
)

func main() {
	log.Println("this is payment microservices")

	defer log.Println("Closing payments microservice")

	ctx:= cmd.Context()

	paymentsInterface:= createPaymentsMicroservice()

	if err :=server.Close(); err!=nil{
		panic(err)
	}
	}
	func createPaymentsMicroservice() amqp.paymentsInterface{
		cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
		
		paymentsService := payments_app.NewPaymentService(
			payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
		)
		paymentsInterface,err := amqp.newPaymentsInterface(
			fmt.Sprintf("amqp://%s/",os.Getenv("SHOP_RABBITMQ_ADDR")),
			os.Getenv("SHOP_RABBIT_MQ_TO_PLAY_QUEUE"),
			paymentsService,
		)
		if err!=nil{
			panic(err)
		}
		return paymentsInterface
}
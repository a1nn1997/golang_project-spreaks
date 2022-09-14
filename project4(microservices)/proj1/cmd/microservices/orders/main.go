package main

import (
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting orders microservice")

	ctx := cmd.Context()
	r,closeFn := createOrderMicroservice()
	defer closeFn()
	server :=   &http.Server{Addr: os.Getenv("SHOP_ORDER_BIND_ADDR"), Handler: r}
	go func() {
		if err := server.ListenAndServe(); err!=http.ErrServerClosed{
			panic(err)
		}
	}()
		<-ctx.Done() //looks for closing event
		log.Println("Closing order microservice")

		if err :=server.Close(); err!=nil{
			panic(err)
		}
	}

func createOrderMicroservice()(router *chi.Mux, closeFn func()){
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shopHTTPClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))
	
	r := cmd.CreateRouter()
	
	order_public_http.AddRoutes(r, ordersService, orderRepo)
	order_private_http.AddRoutes(r, ordersService, orderRepo)

	return r, func(){

	}
}

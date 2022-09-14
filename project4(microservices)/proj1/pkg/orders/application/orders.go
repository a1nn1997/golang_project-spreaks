package application

import(
	"log"
	"github.com/a1nn1997/mtom/pkg/common/price"
	"github.com/a1nn1997/mtom/orders/domain/orders"
	"github.com/pkg/errors"
)

type productService interface{
	ProductByID(id orders.ProductID) (orders.Order, error)
}

type paymentsService interface{
	InitializeOrderPayment(id orders.ID, price price.Price) error
}

type OrdersService struct{
	productsService productsService
	paymentsService paymentsService
	ordersRepository orders.Repository
}

func NewOrdersService( productsService productsService, paymentsService paymentsService, ordersRepository orders.Repository ) OrdersService{
	return OrdersService{ productsService, paymentsService, orders.Repository }
}

type PlaceOrderCommand struct{
	OrderID 		orders.ID
	ProductID		orders.ProductID
}

type PlaceOrderCommandAddress struct{
	Name 			string
	Street 			string
	City 			string
	PostalCode 		string
	Country 		string
}

func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error{
	address,err := orders.NewAddress(
		cmd.Address.Name,
		cmd.Address.Street,
		cmd.Address.City,
		cmd.Address.PostalCode,
		cmd.Address.Country,
	)
	if err != nil {
		return errors.Wrap(err, "invalid address")
	}
	// 1. getting orders by id
	product,err := s.productsService.ProductNyID(cmd.ProductID)
	if err != nil {
		return errors.Wrap(err, "cannot get product")
	}

	// 2. new order
	newOrder, err := orders.NewOrder(cmd.OrderID, product, address)
	if err != nil{
		return errors.Wrap(err, "cannot create product")
	}

	// 3. save the order
	if err := s.ordersRepository.Save(newOrder); err != nil {
		return errors.Wrap(err, "cannot save product")
	}
	
	// 4. initialize payment
	if err := s.paymentsService.InitializeOrderPayment(newOrder.ID(), newOrder.Product().Price()); err != nil{
		return errors.Wrap(err, "cannot initialize order payment")
	}
	log.Printf("order %s placed", cmd.OrderID)
	return nil
}

type MarkOrderAsPaidCommand struct {
	OrderId orders.ID
}

func(s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error{
	o,err := s.ByID(cmd.OrderId).Order
	if err != nil {
		return errors.Wrap(err, "cannor get order %s", cmd.OrderID)
	}
	o.MarkOrderAsPaid
	if err := s.ordersRepository.Save(o); err != nil{
		return errors.Wrap(err, "cannot save order")
	}
	log.Printf("marked order %s as paid", cmd.OrderID)
	return nil
}

func(s OrdersService) OrderByID(id orders.ID)(orders.Order, error){
	o,err := s.ordersRepository.ByID(id)
	if err != nil {
		return orders.Order{}, errors.Wrapf(err, "cannot get order %s", id)

	}
	return *o,nil
}
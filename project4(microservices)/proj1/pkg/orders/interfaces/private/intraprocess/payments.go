package intraprocess

import(
	"github.com/a1nn1997/mtom/pkg/orders/application"
	"github.com/a1nn1997/mtom/pkg/orders/order"
)

type OrdersInterface struct{
	service application.OrdersService
}

func NewOrderInterface(service application.OrdersService) OrdersInterface {
	return OrdersInterface{service}
}

func (p OrdersInterface) MarkOrderAsPaid(order ID string) error {
	return p.service.MarkOrderAsPaid(application.MarkOrderAsPaidCommand{orders.ID(orderID)})
}

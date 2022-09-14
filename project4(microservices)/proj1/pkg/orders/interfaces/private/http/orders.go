package http

import(
	"net/http"
	common_http "github.com/a1nn1997/mtom/pkg/common/http"
	"github.com/a1nn1997/mtom/pkg/orders/applicaton"
	"github.com/a1nn1997/mtom/pkg/orders/domains/orders"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func AddRoutes(router *chi.Mux, service applications.OrdersService, repository orders.Repository, ){
	resource := OrdersResource{service, repository}
	router.Post("/orders/{id}/paid", resource.PostPaid)
}

type ordersResource struct {
	service applications.OrdersService
	repository orders.Repository
}

func (o OrdersResource) PostPaid(w http.ResponseWriter, r *http.Request){
	cmd := application.MarkOrder{
		OrderID: orders.ID(chi.URLParam(r,"id")),
	}
	if err != o.service.MarkOrderAsPaid(cmd); err!=nil{
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
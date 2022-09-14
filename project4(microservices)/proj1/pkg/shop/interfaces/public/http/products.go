package http

import (
	"net/http"
)

func AddRoutes(router *chi.Mux, productsReadModel productsReadModel){
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

type productsReadModel interface{
	AllProducts() ([]products.Product, error)

}

type productView struct{
	ID					string				`json:"id"`
	Name				string				`json:"name"`
	Description			string				`json:"description"`
	Price				priceView			`json:"price"`
}

type priceView struct{
	Cents			uint64			`json:"cents"`
	Currency		string			`json:"currency"`
}

type ProductResource struct {
	readModel productsReadModel
}

func priceViewFromPrice(p price.Price) priceView{
	return priceView{
		p.Cents(),
		 p.Currency(),
	}
}
func (p ProductResource) GetAll(w http.ResponseWriter, r *http.Request){
	products,err := p.readModel.AllProducts()
	if err != nil {
		_ = render.Render(w,r,common_http.ErrInternal(err))
		return
	}
	
	view := []productView{}

	for _,product := range products{
		view = append(view, productView{
			string(product.ID),
			product.Name(),
			product.Description(),
			priceViewFromPrice(product.Price()),
		})
	}
	render.Respond(w, r, view)
}
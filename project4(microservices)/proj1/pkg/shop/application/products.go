package application

import(

)

type ProductReadModel interface {
	AllProducts() ([]products.Product, err)
}

type ProductService struct {
	repo products.Repository
	readModel productReadModel 

}

func NewProductsService(repo products.Repository, readModel productReadModel) ProductService{
	return ProductsService{repo,readModel}
}

func (s ProductService)AllProducts()([]products.Product, error){
	return s.readModel.AllProducts()
}

type AddProductCommand struct {
	ID 							string
	Name 						string
	Description  				string
	PriceCents 					uint64
	PriceCurrency 				string
}

func (s ProductService) AddProduct(cmd AddProductCommand) error {
	price,err := price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)

	if err != nil {
		return errors.Wrap(err,"invalid product price")
	}
	
	p,err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)
	if err != nil {
		return errors.Wrap(err,"cannot create product")
	}
	if err := s.repo.Save(p); err != nil{
		return errors.Wrap(err,"cannot save product")
	}
	return nil
}
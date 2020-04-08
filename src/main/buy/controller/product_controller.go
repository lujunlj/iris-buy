package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris/src/main/buy/service"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService service.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	products, _ := p.ProductService.GetAllProduct()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": products,
		},
	}
}

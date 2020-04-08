package service

import (
	"iris/src/main/buy/dao"
	"iris/src/main/buy/model"
)

type IProductService interface {
	GetProductById(int64) (*model.Product, error)
	GetAllProduct() ([]*model.Product, error)
	DeleteProductById(int64) bool
	InsertProduct(product *model.Product) (int64, error)
	UpdateProduct(product *model.Product) error
}

type ProductService struct {
	productDao dao.IProduct
}

func NewProductService(dao dao.IProduct) IProductService {
	return &ProductService{productDao: dao}
}

func (p *ProductService) GetProductById(prodcutId int64) (*model.Product, error) {
	return p.productDao.SelectByKey(prodcutId)
}

func (p *ProductService) GetAllProduct() ([]*model.Product, error) {
	return p.productDao.SelectAll()
}

func (p *ProductService) DeleteProductById(productId int64) bool {
	return p.productDao.Delete(productId)
}

func (p *ProductService) InsertProduct(product *model.Product) (int64, error) {
	return p.productDao.Insert(product)
}

func (p *ProductService) UpdateProduct(product *model.Product) error {
	return p.productDao.Update(product)
}

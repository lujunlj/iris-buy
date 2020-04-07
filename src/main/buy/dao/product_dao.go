package dao

import (
	"database/sql"
	"errors"
	"iris/src/main/buy/model"
)

//第一步.开发接口
type IProduct interface {
	//连接数据
	Conn() error
	Insert(*model.Product) (int64, error)
	Delete(int64) bool
	Update(*model.Product) error
	SelectByKey(int64) (*model.Product, error)
	SelectAll() ([]*model.Product, error)
}

//第二步.实现接口
type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

//非显示的实现接口
//怎么判断结构体实现了接口
//自己写初始化构造函数
func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
	}
}

func (p *ProductManager) Conn() error {
	return errors.New("123")
}

//3.

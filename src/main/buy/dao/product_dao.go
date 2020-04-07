package dao

import (
	"database/sql"
	"errors"
	"iris/src/main/buy/common"
	"iris/src/main/buy/model"
	"strconv"
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

//数据库连接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		conn, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = conn
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

//插入
func (p *ProductManager) Insert(pd *model.Product) (productId int64, err error) {

	//判断连接是否存在
	//准备sql
	//传入参数
	//执行sql获得id
	if err = p.Conn(); err != nil {
		return
	}
	sql := "Insert into product set productName = ? ,productNum=?,productImage=?,productUrl=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(pd.ProductName, pd.ProductNum, pd.ProductImage, pd.ProductUrl)
	if err != nil {
		return 0, err
	}
	productId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return productId, nil
}

//删除
func (p *ProductManager) Delete(productID int64) bool {
	//1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return false
	}
	//2.准备sql
	sql := "delete from product where id =?"
	//3.根据sql 创建statement 填入参数
	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return false
	}
	_, err = stmt.Exec(productID)
	if err != nil {
		return false
	}
	return true
}

//商品的更新
func (p *ProductManager) Update(product *model.Product) (err error) {
	//1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return err
	}
	//2.准备sql
	sql := "update product set productName =?,productNum=?,productImage=?,productUrl=? where id =" + strconv.FormatInt(product.ID, 10)
	//3.根据sql 创建statement 填入参数
	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductManager) SelectByKey(productID int64) (product *model.Product, err error) {
	//1.准备连接 判断连接是否存在
	if err = p.Conn(); err != nil {
		return &model.Product{}, err
	}
	//2.准备sql
	sql := " select * from " + p.table + " where id =" + strconv.FormatInt(productID, 10)
	row, errRow := p.mysqlConn.Query(sql)
	if errRow != nil {
		return &model.Product{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &model.Product{}, nil
	}

}

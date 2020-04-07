package test

import (
	"fmt"
	"iris/src/main/buy/common"
	"iris/src/main/buy/model"
	"testing"
)

func TestDataToMap(t *testing.T) {
	data := map[string]string{
		"id":           "123",
		"productName":  "贝拉米",
		"productNum":   "4444",
		"productImage": "beilami.png",
		"productUrl":   "http://123",
	}
	product := &model.Product{}
	common.DataToStructByTagSql(data, product)
	fmt.Printf("变量为%v\n", product)
}

// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/27 10:21 下午
// @Update: xxx 2020/10/27 10:21 下午

package test

import (
	"Third/Abstractfactory"
	"testing"
)

func getMainAndDetail(factory Abstractfactory.DAOFactory)  {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestExampleRdbFactory(t *testing.T)  {
	var factory Abstractfactory.DAOFactory
	factory = &Abstractfactory.RDBDAOFactory{}
	getMainAndDetail(factory)
}

func TestExampleXMLFactory(t *testing.T)  {
	var factory Abstractfactory.DAOFactory
	factory = &Abstractfactory.XMLDAOFactory{}
	getMainAndDetail(factory)
}
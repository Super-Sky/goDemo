// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/27 9:54 下午
// @Update: xxx 2020/10/27 9:54 下午

package Abstractfactory

import "log"

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

//DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

//RDBMainDAP 为关系型数据库的OrderMainDAO实现
type RDBMainDAP struct {}

func (*RDBMainDAP) SaveOrderMain() {
	log.Println("rdb main save")
}

//RDBDetailDAO 为关系型数据库的OrderDetailDAO实现
type RDBDetailDAO struct {}

func (*RDBDetailDAO) SaveOrderDetail() {
	log.Println("rdb detail save")
}

//RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct {}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAP{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

//XMLMainDAO XML存储
type XMLMainDAO struct {}

func (*XMLMainDAO) SaveOrderMain() {
	log.Println("xml main save")
}

//XMLDetailDAO XML存储
type XMLDetailDAO struct {}

func (*XMLDetailDAO) SaveOrderDetail() {
	log.Println("xml detail save")
}

//XMLDAOFactory 是XML 抽象工厂实现
type XMLDAOFactory struct {}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}

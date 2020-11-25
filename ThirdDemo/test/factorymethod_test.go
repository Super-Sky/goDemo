// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/22 10:18 下午
// @Update: xxx 2020/10/22 10:18 下午

package test

import (
	"Third/Factorymethod"
	"testing"
)

func compute(factory Factorymethod.OperatorFactory, a ,b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T)  {
	var factory Factorymethod.OperatorFactory
	factory = Factorymethod.PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}
	factory = Factorymethod.MinusOperatorFactory{}
	if compute(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
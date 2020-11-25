// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/28 11:50 上午
// @Update: xxx 2020/10/28 11:50 上午

package test

import (
	"Third/Builder"
	"log"
	"testing"
)

func TestBuilder1(t *testing.T)  {
	builder := &Builder.Builder1{}
	director :=Builder.NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	log.Println(res)
}

func TestBuilder2(t *testing.T)  {
	builer := &Builder.Builder2{}
	director := Builder.NewDirector(builer)
	director.Construct()
	res := builer.GetResult()
	log.Println(res)
}
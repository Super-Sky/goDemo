// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/20 2:02 下午
// @Update: xxx 2020/11/20 2:02 下午

package main

import "fmt"

type animal interface {
	Move()
}

type bird struct {
	name string
}

func (self *bird) Move() {
	fmt.Printf("beast move %s\n", self.name)
}

type beast struct {
	name string
}

func (self *beast) Move() {
	fmt.Printf("beast move %s\n", self.name)
}

func animalMove(v animal) {
	temp := &bird{}
	temp = nil
	if v == temp {
		println("nil animal")
		return
	}
	v.Move()
}

func getBirdAnimal(name string) *bird {
	if name != "" {
		return &bird{name: name}
	}
	return nil
}

func main() {
	var a animal
	var b animal
	a = getBirdAnimal("big bird")
	b = getBirdAnimal("") // return interface{data:nil}
	animalMove(a) // bird move big bird
	animalMove(b) // panic
}
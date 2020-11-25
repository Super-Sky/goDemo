// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/28 1:43 下午
// @Update: xxx 2020/10/28 1:43 下午

package Prototype

//Cloneable是原型需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

type Manager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *Manager {
	return &Manager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *Manager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *Manager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
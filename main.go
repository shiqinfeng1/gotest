package main

import (
	"fmt"
	"unsafe"
)

type Action interface {
	Jump()
}

type Somebody struct {
	Action
}

func (p *Somebody) Jump() {
	fmt.Println("Somebody Jump")
}

func sliceRise(a []int) {
	a = append(a, 0)
	for i := range a {
		a[i]++
	}
	fmt.Println("a", a)
	fmt.Println("a", len(a), cap(a))
}
func main() {
	var m map[string]int
	v, ok := m["a"]
	if !ok {
		fmt.Println("not found ")
		return
	}
	fmt.Println("v=", v)
	s1 := []int{1, 2}
	fmt.Println("s1", len(s1), cap(s1))
	s2 := s1
	s2 = append(s2, 3)
	fmt.Println("s2", len(s2), cap(s2))
	sliceRise(s1)
	sliceRise(s2)
	fmt.Println(s1, s2)

	var somebody Somebody = Somebody{}                // 和Action接口无关
	var hero Somebody = Somebody{Action: &Somebody{}} // Somebody结构实现了Action接口

	// 打印结果：somebody:size=[16],value=[{<nil>}]
	fmt.Printf("somebody:size=[%d],value=[%v]\n", unsafe.Sizeof(somebody), somebody)
	// 打印结果：hero:size=[16],value=[{0xc000010200}]
	fmt.Printf("hero:size=[%d],value=[%v]\n", unsafe.Sizeof(hero), hero)

	somebody.Jump()        // 打印 Somebody Jump，Somebody 结构实现了Jump，与 Action 接口无关
	hero.Action.Jump()     // 打印 Somebody Jump，Action 接口实例化，等价于 hero.Jump()
	somebody.Action.Jump() // runtime error: invalid memory address or nil pointer dereference
}

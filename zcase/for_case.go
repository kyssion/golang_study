package zcase

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) Print() {
	fmt.Printf(p.name)
}

//ForTest01  测试循环
func ForTest01() {
	data := []field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		go v.Print()
		//time.Sleep(3 * time.Second)
	}
	//这个和js类似，输出的是 three three three 主要是for循环运行的够快了，
	//快到了for都跑完了用go v.Print() 还没有开始跑
	time.Sleep(3 * time.Second)
}

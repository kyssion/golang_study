//Package golang0001 流程控制相关
package golang0001

func Test01() {
	//switch 语句有三种形式
	str := "123123p"
	//todo 1. 使用标记的变量来进行比较
	switch str {
	case "A":
		fallthrough // 无视下一个判断直接运行
	case "B", "C":
		println("sdfsdf")
	default:
		println("heheh")
	}
	//todo 2. 相当于多个if else
	switch {
	case str == "xxx":
	case str == "yy":
	default:
	}
	//todo ： 表达式初始化赋值 待定 暂时编译不通过
	//switch xzx:=123{
	//case xzx==1:
	//default:
	//
	//}

	// todo select 语句
	// 类似switch 语句但是应用于管道操作
	var c1,c2 chan int;
	var i1,i2 int;
	select {
	case i1=<-c1:
	case c2<-i2:

	}

	//todo for 两种情况 for 和 for-range
	//todo break和 continue 提供了标签可以直接跳转到指定的位置
	//goto 语句 可以任意跳转但是块外不能跳转到块内

	//if else if 相比较多了一个赋值操作
	if x:=1;x>12 {

	}else{

	}

	//
}

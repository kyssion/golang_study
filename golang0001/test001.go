package golang0001

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func test1() {
	var _ int     //int 8 16 32 64
	var _ uint    //int 8 16 32 64
	var _ byte    // 一字节
	var _ rune    //等同于 int32 通常用来处理的utf-8 这种字符集的
	var _ uintptr //无符号整数用来存放一个之指针

	/* 浮点数 */

	var _ float32
	var _ float64

	/* 数字的最大值和最小值  所有的数据都是使用类似的写法*/
	_, _ = math.MaxFloat32, math.MaxFloat64

	/* 数字字符串转化 */
	var _, _ = strconv.Atoi("123")
	var _, _ = strconv.ParseInt("123", 10, 64) //字符串,进制，类型
	var _, _ = strconv.ParseFloat("1.23", 64)  //字符串，类型

	var _ string = strconv.FormatBool(true)
	var _ string = strconv.FormatInt(123, 32)

	/* 复数类型 用的少 */
	var _ complex128 = 1.0 + 10i
	var _ complex64 = 2.0 + 10i
	var _ = complex(1, 1)

	/* 常量 */
	const item = 123
	fmt.Println(item)

	/* 常量块 和 iota 实现二进制优化 */

	const (
		a, a1 = 1 << iota, 1<<iota - 1 //1<<0 1<<0 -1
		b, b1                          //1<<1 1<<1 -1
		c, c1                          //1<<2 1<<2 -1
		d     = 123456                 //iota 保持不变
		_, d1 = 1 << iota, 1<<iota - 1 //1<<3 1<<3 -1
		e, e1                          //1<<4 1<<4 -1
	)

	//iota 每次使用将会+1 遇到const 将会重置

	/* 特别的逻辑运算符 */
	var test int = 1
	test = test &^ 2 // test 中的所有位和 2 的二进制进行比较，如果相同为0 不同则保留
}
func jiqiao() {
	var a = 10
	var b = 11
	/* 交换两个变量的数值 */
	a, b = b, a
}

//RuneTest 测试函数
func RuneTest() {
	var str = "hello 你好"

	//golang中string底层是通过byte数组实现的，座椅直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))

	//以下两种都可以得到str的字符串长度

	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("len(str):", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字 和上面使用的utf8 的方法结果相同

	fmt.Println("rune:", len([]rune(str)))

	//字符串拼接 go和java 有点类似吧

	// 方法1 每次运算产生一个新的字符串性能差
	var _ = "a" + "b" + "c"
	//方法2

}

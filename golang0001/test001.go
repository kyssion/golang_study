//Package golang0001 变量相关的
package golang0001

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
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
	//方法2  使用Sprint 使用byte模板，内部逻辑复杂 比方法1 好
	var _ = fmt.Sprintf("%s%s%s", "a", "b", "c")
	//方法4 strings.join() 内部构建一个byte数组然后将字符一个一个进行填充
	var _ = strings.Join([]string{"a", "b"}, ",")
	//方法4 byte buffer 和 strings.builder buffer线程安全 build不是，build使用数组切片实现的，buffer使用byte实现的
	// 都可以使用Grow() 来进行容量预估
	var buffer bytes.Buffer
	buffer.Grow(100)
	buffer.WriteString("a")
	var _ = buffer.String()

	var builder strings.Builder
	builder.Grow(100)
	builder.WriteString("a")
	var _ = builder.String()

	// 字符串处理常用包

	//1. strings - 基本功能：查询，替换，比较，截断，拆分，合并，开头结尾判断，索引查询，拼接，统计等
	//2. bytes - 针对strings包中的东西提供byte级别的操作，因为string只读所有使用bytes可以降低string的复制提高性能
	//3. strconv - 提供了类型互换和双引号转译等功能
	//4. unicode - 提供了 isDigit isLetter isUpper isLower 等类似的功能用于字符串分类
}

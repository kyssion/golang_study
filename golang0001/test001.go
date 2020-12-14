package golang0001

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis/passes/composite"
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
	var _,_ := math.MaxFloat32 ,  math.MaxFloat64;

	/* 数字字符串转化 */
	var _ = strconv.AtoI("123");
	var _ = strconv.ParseInt("123",10,64); //字符串,进制，类型
	var _ = strconv.ParseFloat("1.23",64);//字符串，类型

	var _ string = strconv.FormatBool(true);
	var _ string = strconv.FormatInt(123);


	/* 复数类型 用的少 */
	var _ complex128 = 1.0+10i;
	var _ complex64 = 2.0+10i;
	var _ composite = 3.0+10i;

}	
func jiqiao{
	var a = 10;
	var b = 11;
	/* 交换两个变量的数值 */
	a,b = b,a;
}

func Rune_test() {
	var str = "hello 你好"

	//golang中string底层是通过byte数组实现的，座椅直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))

	//以下两种都可以得到str的字符串长度

	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字 和上面使用的utf8 的方法结果相同

	fmt.Println("rune:", len([]rune(str)))

}

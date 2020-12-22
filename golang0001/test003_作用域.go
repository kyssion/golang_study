package golang0001

//全局作用区域
const iii = "123"

func test()  {
	//作用域1 局部作用域
	var _ string= "123";
	{
		//作用域2 局部作用域
		var _ string = "123";
	}
}
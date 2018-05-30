package iteracter

type Any interface{}
type f func(init Any) Any

/*
这是其中的一个方法，也可以通过函数去设定值得计算,也可以通过返回多个值来设定初始值，当结果和初始化方式不一样时
*/
func Factory(f1 f, init Any) func() Any {

	in := make(chan Any)
	var res Any
	go func() {
		for {
			res = f1(init)
			in <- res
			init = res
		}
	}()

	ret := func() Any {
		return <-in
	}
	return ret
}

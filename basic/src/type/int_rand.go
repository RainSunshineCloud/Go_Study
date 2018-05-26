package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*
seed 用于初始化一个随机种子，
根据这个种子来进行接下来的随机数运算(比如seed(1)每次返回的数都是一样的，
所以需要种子进行播种
*/
func main () {

	nows := int64(time.Now().Nanosecond());
	rand.Seed(nows);
	for i := 1; i <= 10; i++ {


		res1 := rand.Int31();//表示int31位的有符号int随机数
		res2 := rand.Uint32();//表示int32位的无符号int随机数
		res3 := rand.Intn(1);//表示从0-n 的int随机数
		res4 := rand.Float32();//表示0-1之间的32位float 随机数
		res5 := rand.Float64();//表示0-1之间的64位float 随机数

		fmt.Println(res1,res2,res3,res4,res5);

	}
}
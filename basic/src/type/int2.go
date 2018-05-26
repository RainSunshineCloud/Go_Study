package main

import "fmt"
/*
总结：
记住：不同类型之间不能做运算，不管什么运算

int 各项基本运算
float浮点类型不能取余
float 也可以自增
++ 自增自减是一个独立的运算，无法带入表达式和函数。
逻辑运算符（&&，|| ，！）只能用在布尔类型上
二进制运算符（& ,|, ^，>>，<<）只能用在整型上
complex 可以做加减乘除运算（复数运算）
*/


func main () {

	basic();//基本数值运算
	compare();//比较运算
	logic();//逻辑运算
	binaryOperation(); //二进制运算
	complexs();//复数
}


func compare() {
	int1 := 1;
	var float1 float32 = 1;
	var int2 byte = 1;
	bool1 := true;

	//fmt.Println(int1 == float1,int1 == bool1,int1 == int2);//不同类型不能比较

	fmt.Println(int1 == int1,int2 == int2,float1 == float1,bool1 == bool1);


}

func basic () {
	int1 := 1;
	fmt.Println(int1);//1
	int2 := 2;
	fmt.Println(int2);//2
	int1++; //这是一个独立的运算，无法带入表达式和函数。
	int1--; //这是一个独立的运算，无法带入表达式和函数。
	fmt.Println(int1);//1

	int3 := int1 + int2;//加
	float3 := float32(int1) + float32(int2);
	fmt.Println(float3);//3
	fmt.Println(int3);//3

	int4 := int1 - int2;//减
	fmt.Println(int4);//-1

	int5 := int1 * int2;//乘
	fmt.Println(int5);//2
	float5 := float32(int1) * float32(int2);//浮点运算
	fmt.Println(float5);//2

	//注意，与c不一样的是go没有隐性转换，故整数与浮点数的计算是错误的
	int6 := int1 / int2;//除(整数和整数只会得到整数）
	fmt.Println(int6);//0
	float6 := float32(int1) / float32(int2); //0.5
	fmt.Println(float6);

	//求余
	int7 := int1 % int2;
	fmt.Println(int7);//1

	//float7 := float32(int1) % float32(int2);//浮点类型不能取余
	//fmt.Println(float7);
	var int8 byte = 1;
	var int9 byte = 2;
	fmt.Println(int8 % int9);

	float10 := 1.2;
	float10++;
	fmt.Println(float10);
}

func logic() {
	//int1 := 1;
	//int2 := 2;
	bool1 := false;
	bool2 := true;
	//fmt.Println(!int1,int1 && int2,int1 || int2);
	fmt.Println(!bool1);
	fmt.Println(bool1 && bool2);
	fmt.Println(bool1 || bool2);
	fmt.Println(returnFalse()||returnTrue());

}

func returnFalse () bool {
	return false;
}

func returnTrue () bool {
	return true;
}

func binaryOperation() {
	int1 := 1;//11
	int2 := 2;//10
	var int3 uint = 2;
	//bool1 := true;
	//bool2 := false;
	//float1 := 1.0;
	//float2 := 1.0;
	fmt.Println(int1 & int2);//10
	fmt.Println(int1 | int2);//11
	fmt.Println(int1 ^ int2);//01
	fmt.Println(int1 << int3);//110
	fmt.Println(int1 >> int3);//1
	fmt.Println(^int1);//这个暂时不懂
	fmt.Println(int2&^int1);//11 //将int2的二进制 中的第 int1 位转化为0
	fmt.Println(^int3);
	//fmt.Println(bool1 & bool2,bool1 | bool2,bool1 ^ bool2);
	//fmt.Println(float1 & float2,float1 | float2,float1 ^float2);


}

func complexs() {
	cpx1 := complex(1,2);
	cpx2 := complex(1,-2);
	fmt.Println(cpx1);
	fmt.Println(real(cpx1));
	fmt.Println(imag(cpx1));
	cpx3 := cpx1 * cpx2;
	fmt.Println(cpx3);
	cpx4 := cpx1 -cpx2;
	fmt.Println(cpx4);
	cpx5 := cpx1 + cpx2;
	fmt.Println(cpx5);
	cpx6 := cpx1 / cpx2;
	fmt.Println(cpx6);
}
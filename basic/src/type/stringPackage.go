package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "sdes 0df中 国"
	bool1 := strings.HasPrefix(str, "sd")
	bool2 := strings.HasSuffix(str, "f")

	int := strings.Count(str, "")     //在子串出现的次数(如果为空串则为unicode（中英文各1）字节个数 + 1
	int1 := strings.Compare(str, "t") //比较字符串 0 相等 -1 小于 1 大于

	index := strings.Index(str, "se") //第一次出现字串的位置找不到为-1
	map1 := strings.Map(func(rune1 rune) rune {
		if rune1 == '中' { //返回负号就是删除
			return 1<<31 - 1
		}
		return rune1
	}, str)

	repeat := strings.Repeat("s", 6)
	upper := strings.ToUpper(str)
	lower := strings.ToLower(str)
	title := strings.ToTitle(str)
	fold := strings.EqualFold(str, "sDEs0Df中国")
	field := strings.Fields(str) //按空格分开
	split := strings.Split(str, "d")
	explode := strings.SplitN(str, "", 4)
	reader := strings.NewReader(str)
	byte1, _ := reader.ReadByte()
	byte2, _ := reader.ReadByte()
	fmt.Println(bool1, bool2, int, int1, index, map1, repeat, upper, lower, title, fold, field[1], split[2])
	fmt.Println(explode[3], byte1, byte2)

}

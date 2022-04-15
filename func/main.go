/*
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-03 10:27:01
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-04-07 13:35:17
 */
package main

import (
	"regexp"
	"strings"
)

func isbool(x, y interface{}) bool {
	if x == y {
		return true
	} else {
		return false
	}
}

// 计算字符串的个数
func countString(s string) (count map[string]int) {
	sArr := strings.Split(s, " ")
	count = make(map[string]int, 4)
	for _, v := range sArr {
		if _, ok := count[v]; !ok {
			count[v] = 1
		} else {
			count[v]++
		}
	}
	return
}

/**
*分金币
*50枚金币分给某些人
*名字中包含e 或 E 分1个金币
*名字中包含i 或 I 分2个金币
*名字中包含o 或 O 分3个金币
*名字中包含u 或 U 分4个金币
*每个用户分到多少金币，还剩多少金币
 */
var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func dispathcCoin() (num int) {
	for _, v := range users {
		// 用字符串包含方法实现
		// oke := strings.Contains(v, "e")
		// okE := strings.Contains(v, "E")
		// oki := strings.Contains(v, "i")
		// okI := strings.Contains(v, "I")
		// oko := strings.Contains(v, "o")
		// okO := strings.Contains(v, "O")
		// oku := strings.Contains(v, "u")
		// okU := strings.Contains(v, "U")
		// if oke || okE {
		// 	distribution[v] += 1
		// 	coins -= 1
		// }
		// if oki || okI {
		// 	distribution[v] += 2
		// 	coins -= 2
		// }
		// if oko || okO {
		// 	distribution[v] += 3
		// 	coins -= 3
		// }
		// if oku || okU {
		// 	distribution[v] += 4
		// 	coins -= 4
		// }

		// 用正则实现
		// oke, _ := regexp.Match("e|E", []byte(v))
		// oki, _ := regexp.Match("i|I", []byte(v))
		// oko, _ := regexp.Match("o|O", []byte(v))
		// oku, _ := regexp.Match("u|U", []byte(v))
		// if oke {
		// 	distribution[v] += 1
		// 	coins -= 1
		// }
		// if oki {
		// 	distribution[v] += 2
		// 	coins -= 2
		// }
		// if oko {
		// 	distribution[v] += 3
		// 	coins -= 3
		// }
		// if oku {
		// 	distribution[v] += 4
		// 	coins -= 4
		// }
		// 以上都是只要匹配上一次就进行计算，下面进行只要匹配就计算
		// 解析正则有效性，并在字符串中查找符合的元素返回{{起始位置, 结束位置}, {起始位置, 结束位置}, ...}
		oke := regexp.MustCompile("e|E").FindAllStringIndex(v, -1)
		oki := regexp.MustCompile("i|I").FindAllStringIndex(v, -1)
		oko := regexp.MustCompile("o|O").FindAllStringIndex(v, -1)
		oku := regexp.MustCompile("u|U").FindAllStringIndex(v, -1)
		if len(oke) > 0 {
			distribution[v] += len(oke)
			coins -= len(oke)
		}
		if len(oki) > 0 {
			distribution[v] += len(oki) * 2
			coins -= len(oki) * 2
		}
		if len(oko) > 0 {
			distribution[v] += len(oko) * 3
			coins -= len(oko) * 3
		}
		if len(oku) > 0 {
			distribution[v] += len(oku) * 4
			coins -= len(oku) * 4
		}
	}

	num = coins
	return
}

// 递归 1到100累加值
func getrecursion(num uint64) uint64 {
	if num <= 1 {
		return 1
	}
	return num + getrecursion(num-1)
}

// 走台阶可以走一步 可以走两步 for 循环
func step(n int) int {
	var x int = 1
	var y int = 2
	var sum int
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	for i := 3; i <= n; i++ {
		sum = x + y
		x = y
		y = sum
	}
	return sum
}

// 走台阶可以走一步 可以走两步 递归操作
func step2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return step2(n-1) + step2(n-2)
}

// 打扫房间
// func clearRoom(n int) {
// 	var timeNum = [7]int{2, 3, 9, 6, 1, 3, 4}
// 	for i := 0; i < n; i++ {

// 	}
// }
// 罗马数字
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	last := 0
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]
		sign := 1
		if temp < last {
			sign = -1
		}
		res += temp * sign
		last = temp
	}
	return res
}

// 单链条
// func isPalindrome(head *ListNode) bool {
// 	h := head
// 	fmt.Printf("%T,%v", h, h.Val)
// 	for head != nil {
// 		h = head.Next
// 	}
// 	return true
// }

func main() {
	// 罗马数字
	// num := romanToInt("X")
	// fmt.Println(num)
	// fmt.Println("for 循环操作：", step(10))
	// fmt.Println("递归操作：", step2(10))
	// a := [...]int{1, 2, 3, 4}
	// b := [4]int{1, 2, 3, 4}

	// why := isbool(a, b)
	// fmt.Println(why) //true

	// // 计算重复的字符串个数
	// s := "How do you do"
	// countNum := countString(s)
	// for k, v := range countNum {
	// 	fmt.Println(k, v)
	// }

	// // 分金币
	// left := dispathcCoin()
	// fmt.Println("剩下：", left)
	// fmt.Println("每人分得明细为：", distribution)

	// // 递归recursion
	// ret := getrecursion(100)
	// fmt.Println("累加：", ret)
}

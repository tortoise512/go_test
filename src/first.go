package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
	"unicode"
)

func main() {
	var (
		a     int = 100
		count int
		b     float32 = 3.1415
		c     bool    = true
		d     string  = "hello world!我是奥特曼"
	)
	// fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
	fmt.Println(d, reflect.TypeOf(d))
	fmt.Printf("a:%T\n", a)
	fmt.Printf("a:%b\n", a)
	fmt.Printf("a:%t\n", c)
	for _, c := range d {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
	arr := [...]int{1, 3, 5, 7, 8}
	count = 0
	for k, i := range arr {
		count += i
		print(k, "\n")
	}
	fmt.Println(count)
	count = 0
	for i := 0; i < len(arr); i++ {
		count += arr[i]
	}
	fmt.Println(count)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
	arr1 := []int{}
	for i := 0; i < 100; i++ {
		arr1 = append(arr1, i)
	}
	fmt.Println(arr1)
	fmt.Println(cap(arr1))
	fmt.Println(len(arr1))
	s := arr1[:120]
	fmt.Println(s)
	a1 := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := a1[:5]
	fmt.Println(s1)
	fmt.Println(cap(s1))
	s2 := make([]int, 100, 100)
	copy(s2, s1)
	fmt.Println(s2)
	s3 := make([]string, 5, 10)
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	for i := 0; i < 10; i++ {
		s3 = append(s3, fmt.Sprintf("%v", i))
	}
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	fmt.Printf("%T", s3[1])
	s4 := [...]int{3, 7, 8, 9, 1}
	sort.Ints(s4[:])
	fmt.Println(s4)
	s5 := [...]int{3, 7, 8, 9, 1}
	s6 := s5[:]
	s6[1] = 1000
	fmt.Println(s5, s6)
	s7 := []string{"hello", "aworld", "hella"}
	// sort.Strings(s7)
	sort.Sort(sort.StringSlice(s7))
	fmt.Println(s7)
	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo)
	value, ok := userInfo["password"]
	if ok {
		fmt.Println(value, ok)
	} else {
		fmt.Println("123")
	}
	for k, v := range userInfo {
		fmt.Println(k, v)
	}
	delete(userInfo, "password")
	fmt.Println(userInfo)

	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	// 元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("-----------------------")
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value1, ok := sliceMap[key]
	if !ok {
		value1 = make([]string, 0, 2)
	}
	value1 = append(value1, "北京", "上海")
	sliceMap[key] = value1
	fmt.Println(sliceMap)
	fmt.Println(strings.Repeat("=", 50))

	str := "how do you do"
	// ar :=[]string
	ab := strings.Fields(str)
	fmt.Println(ab)
	m1 := make(map[string]int, 10)
	for _, w := range ab {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	for kw, vw := range m1 {
		fmt.Println(kw, vw)
	}
	kk := make(map[string]int, 10)
	if _, ok := kk["aa"]; !ok {
		print("aaaaaaaaaaaaaaaa\n")
	}
	type Map map[string][]int
	m := make(Map)
	sa := []int{1, 2, 3}
	// sa = append(sa, 3)
	// fmt.Printf("%+v\n", sa)
	m["q1mi"] = sa // [1 2 3]
	fmt.Println(m["q1mi"])
	sa = append(sa[:1], sa[2:]...) // [1 3]
	fmt.Printf("%+v\n", sa)
	fmt.Printf("%+v\n", m["q1mi"])
	fmt.Printf("%p,%p\n", sa, m["q1mi"])
	sa = sa[:3]
	fmt.Println(sa)
	s12 := sum1(10, 20, 30, 40, 50)
	fmt.Println(s12)
	ret2 := calc(10, 20, add)
	fmt.Println(ret2)
}

func sum1(x ...int) (sum int) {
	fmt.Println(x)
	sum = 0
	for _, v := range x {
		sum += v
	}
	return
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func add(x, y int) int {
	return x + y

}

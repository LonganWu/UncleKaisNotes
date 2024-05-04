package main

//fuzz 模糊测试
//模糊测试就是指半自动或自动地为程序提供非法的、非预期、随机的数据，
//并监控程序在这些输入数据下是否会出现崩溃、内置断言失败、内存泄露、安全漏洞等情况。
//go-fuzz的工作流程如下：
//1）生成随机数据；
//2）将上述数据作为输入传递给被测程序；
//3）观察是否有崩溃记录（crash），如果发现崩溃记录，则说明找到了潜在的bug。

//使用方法
//1. 安装go-fuzz
//$ go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
//$ git clone https://github.com/dvyukov/go-fuzz-corpus.git
//$ cd go-fuzz-corpus
//$ cd png
//$ go-fuzz-build
//2.写一个方法 func Fuzz(data []byte) int
// 不能写在main包

func Reverse(s string) string {
	bs := []byte(s)
	length := len(bs)
	for i := 0; i < length/2; i++ {
		bs[i], bs[length-i-1] = bs[length-i-1], bs[i]
	}
	return string(bs)
}

func ReverseNew(s string) string {
	rs := []rune(s) // 将字符串转换为 Unicode 代码点的切片
	length := len(rs)
	if length == 1 {
		return s
	}
	for i := 0; i < length/2; i++ {
		rs[i], rs[length-i-1] = rs[length-i-1], rs[i]
	}
	return string(rs) // 将 Unicode 代码点切片转换回字符串
}

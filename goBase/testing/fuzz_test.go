package main

import (
	"testing"
	"unicode/utf8"
)

//可以在testdata/fuzz/FuzzXxx目录下以文件的形式放置其初始语料
// testdata/fuzz/FuzzXxx/corpus1
//go test fuzz v1
//[]byte("ABC\xa8\x8c\xb3G\xfc")

// 构造出来的随机数据被称为语料(corpus)。
func FuzzName(f *testing.F) {
	// 设置种子语料(可选)

	// 执行Fuzzing
	f.Fuzz(func(t *testing.T) {

	})
}
func FuzzReverse(f *testing.F) {
	str_slice := []string{"abc", "bb"}
	for _, v := range str_slice {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T, str string) {
		rev_str1 := Reverse(str)
		rev_str2 := Reverse(rev_str1)
		if str != rev_str2 {
			t.Errorf("fuzz test failed. str:%s, rev_str1:%s, rev_str2:%s", str, rev_str1, rev_str2)
		}
		if utf8.ValidString(str) && !utf8.ValidString(rev_str1) {
			t.Errorf("reverse result is not utf8. str:%s, len: %d, rev_str1:%s", str, len(str), rev_str1)
		}
	})
}

func FuzzReverseNew(f *testing.F) {
	str_slice := []string{"abc", "bb"}
	for _, v := range str_slice {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T, str string) {
		if !utf8.ValidString(str) {
			t.Skipf("skipping invalid UTF-8 string: %q", str)
		}
		if len(str) < 2 {
			t.Skipf("skipping string with length less than 2: %q", str)
		}
		rev_str1 := ReverseNew(str)
		rev_str2 := ReverseNew(rev_str1)
		if str != rev_str2 {
			t.Errorf("fuzz test failed. str:%s, rev_str1:%s, rev_str2:%s", str, rev_str1, rev_str2)
			t.Errorf("fuzz test failed. str:%v, rev_str1:%v, rev_str2:%v", []rune(str), []rune(rev_str1), []rune(rev_str2))
		}
		if utf8.ValidString(str) && !utf8.ValidString(rev_str1) {
			t.Errorf("reverse result is not utf8. str:%s, len: %d, rev_str1:%s", str, len(str), rev_str1)
		}
	})
}

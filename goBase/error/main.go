package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

//参考资料：《Go语言精进之路1》

// 1. 错误就是值
// 2. 错误处理就是基于值比较后的决策
// 3. Go语言需要显式处理错误
// 4. 错误检查是必需的
// 5. 错误处理应该在最接近问题的地方进行
// 6. 错误信息应该是可读的和有意义的
// 7. 按惯例，函数或方法通常将error类型返回值放在返回值列表的末尾

// 四个错误处理策略
// 1. 所有错误一致处理（尽量使用、最常见）优点：错误构造方和处理方耦合低
// 2. 按指定错误值处理
// 3. 按指定错误类型处理（获取上下文）（其次使用）
// 4. 按错误类型分类处理

//代码所在栈帧越低（越接近于main函数栈帧），if err != nil就越不常见；反之，代码在栈中的位置越高（更接近于网络I/O操作或操作系统API调用），if err != nil就越常见

//todo 补充更多错误处理经验。比如错误日志在底层记录还是应用接口层记录。

func main() {
	//一. 错误的两种构造方式
	// 1.返回一个新的错误
	err1 := errors.New("错误1")
	fmt.Printf("err1错误值：[%v]\n", err1)
	// 2.包装一个存在的错误,形成错误树
	err2 := fmt.Errorf("错误2，包含：%w", err1)
	fmt.Printf("err2错误值：[%v]\n", err2)
	//二. errors.Unwrap()，从错误数中取出被包装的error
	err1 = errors.Unwrap(err2)
	fmt.Printf("Unwrap err2错误值：[%v]\n", err1)
	//三. errors.Join():将多个error封装为一个error, Join后的error调用Unwrap返回<nil>
	err3 := errors.Join(err1, err2)
	fmt.Printf("err3错误值：[%v]\n", err3)
	//四. errors.Is()bool：从错误树查找是否存在指定错误值。
	fmt.Printf(" err3 Is err1 ? %v\n", errors.Is(err3, err1))
	fmt.Printf(" err3 Is err2 ? %v\n", errors.Is(err3, err2))
	//五. errors.As()bool: 从错误树取出第一个匹配到的错误
	fmt.Printf(" err3 As err1 ? %v\n", errors.As(err3, &err1))
	fmt.Printf(" err2 As err1 ? %v\n", errors.As(err2, &err1))
}

func doSomething() error { return errors.New("new error") }

// 错误处理策略一   最常见的错误处理策略
// ================================================================
// <所有错误一致处理>策略: 不关心错误上下文（发生任何错误都将进入唯一错误处理路径）。
// ================================================================
func strategy1() error {
	err := doSomething()
	if err != nil {
		// 不关心err变量底层错误值所携带的具体上下文信息
		// 执行简单错误处理逻辑并返回
		return err
	}
	return nil
}

// 错误处理策略二
// ================================================================
// <特定错误特定处理>策略： 通过Is方法从错误链中搜索是否存在该错误值
// 一般错误值变量以ErrXXX格式命名
// ================================================================
var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

func strategy2() {
	// 错误处理代码
	err := doSomething()
	if err != nil {
		if errors.Is(err, ErrNegativeCount) {

		} else if errors.Is(err, ErrInvalidUnreadByte) {

		} else if errors.Is(err, ErrInvalidUnreadRune) {

		} else if errors.Is(err, ErrBufferFull) {

		} else {

		}
	}
}

//
// 错误处理策略三    一般自定义导出的错误类型以XXXError的形式命名
// ================================================================
// <按错误类型(class)处理>策略： 通过As方法错误类型比较（断言），
// 调用As方法后会提取错误链中类型符合的错误，这样可以获取错误的上下文信息。
// ================================================================
//

type MyError struct {
	s string
	i int //假设i是其他上下文信息
}

func (e MyError) Error() string {
	return e.s
}
func strategy3() {
	// 错误处理代码
	err := doSomething()
	var e *MyError
	//从错误链中提取一个MyError类型的错误赋值到e
	if errors.As(err, e) {
		fmt.Println("上下文i=", e.i)
	}
}

//
//  错误处理策略四
// ================================================================
//  错误分批处理策略：类似策略三
// ================================================================
//

// Error 自定义错误接口
type Error interface {
	error
	Timeout() bool   // 是超时类错误吗？
	Temporary() bool // 是临时性错误吗？
}

type OpError struct {
	//...
	// Err is the error that occurred during the operation.
	Err error
}

// 自定义实现
type temporary interface {
	Temporary() bool
}

func (e *OpError) Temporary() bool {
	if ne, ok := e.Err.(*os.SyscallError); ok {
		t, ok := ne.Err.(temporary)
		return ok && t.Temporary()
	}
	t, ok := e.Err.(temporary)
	return ok && t.Temporary()
}

// 检测方式
func strategy4() {
	e := doSomething()
	if e != nil {
		//
		if ne, ok := e.(net.Error); ok && ne.Temporary() {
			// 这里对临时性错误进行处理
			time.Sleep(1)
		}
		return
	}
}

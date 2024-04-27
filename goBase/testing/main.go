package main

// 1. 包内测试
// 白盒测试、测试代码和实现代码耦合、需要维护。
// 有包循环了引用可能
// 2. 包外测试
// 黑盒测试、面向API，与实现代码没有耦合，不会产生包循环引用；但覆盖率不如包内测试。
// 3.export_test.go用于导出符号供包外测试使用，提升测试覆盖率。

func main() {

}

// Error/Errorf、Fatal/Fatalf
//一旦进入这些分支，即代表该测试失败。不同的是Error/Errorf并不会立刻终止当前goroutine的执行，还会继续执行该goroutine后续的测试，而Fatal/Fatalf则会立刻停止当前goroutine的测试执行。

//使用testdata管理测试依赖的外部数据文件，Go工具链将忽略名为testdata的目录。

//golden文件模式：实现了testdata目录下测试依赖的预期结果数据文件的数据采集与测试代码的融合。

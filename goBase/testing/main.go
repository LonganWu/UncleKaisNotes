package main

// 1. 包内测试
// 白盒测试、测试代码和实现代码耦合、需要维护。
// 有包循环了引用可能
// 2. 包外测试
// 黑盒测试、面向API，与实现代码没有耦合，不会产生包循环引用；但覆盖率不如包内测试。
// 3.export_test.go用于导出符号供包外测试使用，提升测试覆盖率。

//  4. fake->使用简单的接口实现替换实际依赖的外部组件，如（服务、客户端、数据库等）
//     缺点：并不具备在测试前对返回结果进行预设置的能力。
//  5. stub：对返回结果有一定预设控制能力的替身。增强了对替身返回结果的间接控制能力，这种控制可以通过测试前对调用结果预设置来实现。
//     不过，stub替身通常仅针对计划之内的结果进行设置，对计划之外的请求也无能为力。
//     gostub（https://github.com/prashantv/gostub）的第三方包可以用于简化stub替身的管理和编写。
//  6. mock 只用于实现某接口的实现类型的替身。gomock（https://github.com/golang/mock），该框架通过代码生成的方式生成实现某接口的替身类型。
//	gomock（https://github.com/golang/mock），该框架通过代码生成的方式生成实现某接口的替身类型。
//	这个框架分两部分：一部分是用于生成mock替身的mockgen二进制程序，另一部分则是生成的代码所要使用的gomock包。
//	go install github.com/golang/mock/mockgen@v1.6.0
//生成mock代码
//go:generate mockgen -source=./fake.go -destination=./mock.go -package=main Mailer

// Error/Errorf、Fatal/Fatalf
//一旦进入这些分支，即代表该测试失败。不同的是Error/Errorf并不会立刻终止当前goroutine的执行，还会继续执行该goroutine后续的测试，而Fatal/Fatalf则会立刻停止当前goroutine的测试执行。

//使用testdata管理测试依赖的外部数据文件，Go工具链将忽略名为testdata的目录。

//golden文件模式：实现了testdata目录下测试依赖的预期结果数据文件的数据采集与测试代码的融合。

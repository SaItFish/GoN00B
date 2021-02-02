# 简单工厂模式

go 语言没有构造函数一说，所以一般会定义 NewXXX 函数来初始化相关类。
NewXXX 函数返回接口时就是简单工厂模式，也就是说 Golang 的一般推荐做法就是简单工厂。

在这个 simplefactory 包中只有 API 接口和 NewAPI 函数为包外可见，封装了实现细节。
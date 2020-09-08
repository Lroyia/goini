# goini

go语言下简单的ini配置文件读取实现

## 安装方法

```bash
go get github.com/lroyia/goini
```

## 使用方法

1、加载文件
```go
// 加载文件
conf, err := goini.Read("test.ini")
```

2、读取section下某个键值
```go
value := conf.GetValueBySection("ina", "mat")
```

3、读取section下所有键值
```go
items := conf.GetAllItemInSection("ina")
```

4、读取第一个成功匹配key的值
```go
conf.GetValueByItem("mat")
```
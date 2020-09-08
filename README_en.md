# goini

A simple implementation of ini reader with golang.

## Install

```bash
go get github.com/Lroyia/goini
```

## How to use

1、load file
```go
// load file
conf, err := goini.Read("test.ini")
```

2、get value of specify key in section
```go
value := conf.GetValueBySection("ina", "mat")
```

3、get all item in section
```go
items := conf.GetAllItemInSection("ina")
```

4、get value of key which first match
```go
conf.GetValueByItem("mat")
```
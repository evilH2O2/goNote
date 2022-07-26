VSCode 快捷键

导入 fmt 包后：

`变量名.print!`: 自动补全，前提变量要存在
```go
var i int = 0
fmt.Printf("i: %v\n", i)
```

`f.Name().print!`: 函数调用
```go
fmt.Printf("f.Name(): %v\n", f.Name())
```

<br/>

`语句/表达式.var!`: 自动补全变量名

```go
f, err := os.Create("test")
```

<br/>

`iferr`
```go
if err != nil {
    return nil, err
}
```

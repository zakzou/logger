logger
======

一个模仿python logging的简单日志类，简单包装了golang里面的log.Logger类


# 例子

使用标准错误输出
```go
var log = logger.New("", logger.Lerror, logger.LstdFlags | logger.Lrelativefile)
```

使用具体的文件
```go
var log = logger.New("logger.log", logger.Lerror, logger.LstdFlags | logger.Lrelativefile)
log.Errorf("this is error. error info %s", err)
```

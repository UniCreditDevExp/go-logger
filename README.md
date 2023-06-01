# go-logger
Is a module to be imported in order to centrally manage logs for all the applications

## How to use
``` go
err := log.Init()
if err != nil {
    panic(err)
}
log.Info("message")

pw := "myPW"
log.Log.Filter(pwd)
log.Info(pw) // Will produce [REDACTED]
```

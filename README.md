# go-logger
A module based on zerolog in order to centrally manage logs for all the microservices

## How to use
``` go
err := log.Init()

if err != nil {
    panic(err)
}

// Hide sensible data
pw := "myPW"
log.Filter(pwd)
log.Info(pw) // Will produce [REDACTED]

// INFO
log.Info("message")
log.Infof("%s","message")

// DEBUG
log.Debug("message")
log.Debugf("%s","message")

// WARN
log.Warn("message")
log.Warnf("%s","message")

// ERROR
log.Error("message")
log.Errorf("%s","message")

// FATAL
log.Fatal("message")
log.Fatalf("%s","message")

// PANIC
log.Panic("message")
log.Panicf("%s","message")
```

# LOG Filters - Hide sensitive data

``` go
pw := "myPW"
log.Filter(pwd)
log.Info(pw) // Will produce [REDACTED]
```

# Share log filters 

You can share log filters between different instances or microservices by
configuring the logger to store data in a database

### Redis
Set following environment variables  

`REDIS_ADDRESS`

`REDIS_USERNAME`

`REDIS_PASSWORD`


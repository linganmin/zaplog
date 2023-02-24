# zaplog

> go version 1.19.3

## Install

```bash
 go get github.com/linganmin/zaplog@v0.0.2
```

## Usage

```golang


    ctx := context.WithValue(context.Background(), "request_id", uuid.NewString())
    
    logger := zaplog.FromContext(ctx)
    
    logger.Errorf("this is a test error msg error: %+v", errors.New("name filed is required"))
    logger.Infof("this is a test info msg name: %+v", "saboran")
    
    type user struct {
        Name string
        Age  int
        Addr string
    }
    logger.Debugf("this is a test debug msg user: %+v", user{
        Name: "saboran",
        Age:  18,
        Addr: "The People's Republic of China",
    })

```

### Output 

It will Print content on console and write in log file  `$pwd/logs/YYYY-mm-DD.log`

Such as

```json
{"level":"error","time":"2023-02-24T10:35:42.979+0800","path_line":"china_regions/main.go:16","msg":"this is a test error msg error: name filed is required","request_id":"c98e9c81-9d25-4868-9d65-ac980e103057","stack":"main.main\n\t/Users/linganmin/dev/company/lingan/china_regions/main.go:16\nruntime.main\n\t/usr/local/Cellar/go/1.19.3/libexec/src/runtime/proc.go:250"}
{"level":"info","time":"2023-02-24T10:35:42.979+0800","path_line":"china_regions/main.go:17","msg":"this is a test info msg name: saboran","request_id":"c98e9c81-9d25-4868-9d65-ac980e103057"}
{"level":"debug","time":"2023-02-24T10:35:42.979+0800","path_line":"china_regions/main.go:24","msg":"this is a test debug msg user: {Name:saboran Age:18 Addr:The People's Republic of China}","request_id":"c98e9c81-9d25-4868-9d65-ac980e103057"}

```
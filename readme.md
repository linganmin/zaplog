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
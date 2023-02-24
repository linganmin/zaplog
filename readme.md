# zaplog

> go version 1.19.3

## Install

## Usage

```golang

    ctx := context.Background()
    ctx = context.WithValue(ctx, metadatax.Keys[0], uuid.NewString())
    ctx = context.WithValue(ctx, metadatax.Keys[1], uuid.NewString())
    
    logger := FromContext(ctx)
    
    logger.Errorf("this is a test error msg error: %+v", errors.New("name filed is required"))
    logger.Infof("this is a test info msg name: %+v", "saboran")
    logger.Debugf("this is a test debug msg user: %+v", user{
        Name: "saboran",
        Age:  18,
        Addr: "The People's Republic of China",
    })

```
# go-bqworker

<p align="left"> 
   <a href="https://hits.seeyoufarm.com"/><img src="https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fgjbae1212%2Fgo-bqworker"/></a>
   <a href="https://goreportcard.com/report/github.com/gjbae1212/go-bqworker"><img src="https://goreportcard.com/badge/github.com/gjbae1212/go-bqworker"/></a>
   <a href="https://godoc.org/github.com/gjbae1212/go-bqworker"><img src="https://godoc.org/github.com/gjbae1212/go-bqworker?status.svg"/></a>            
   <a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-GREEN.svg" alt="license"/></a>
</p>

**go-esworker** is an async worker that data can bulk insert, update to the BigQuery.  
It's a library for **golang**.


## Getting Started
### Install
```bash
$ go get -u github.com/gjbae1212/go-bqworker
```

### Usage
```go
// pseudo code
package main
import (
    "context"
    "github.com/gjbae1212/go-bqworker"
)

func main() {
   schema := []*bqworker.TableSchema{your-tables-schema}
   cfg := bqworker.NewConfig("gcp proejct id", "gcp JWT bytes", schema , queue-size, worker-size, worker-queue, delay-time)
   
   streamer, _ := bqworker.NewStreamer(cfg, func(err error){ })
   
   ctx := context.Background()
   
   // row is an object implemented bqworker.Row. 
    
   // async
   streamer.AddRow(ctx, row)     
   // sync 
   streamer.AddRowSync(ctx, row)
}  
```

## LICENSE
This project is following The MIT.

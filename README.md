errors [![godoc badge](http://godoc.org/github.com/plimble/errors?status.png)](http://godoc.org/github.com/plimble/errors)
========


## Installation

```go
$ go get -u github.com/plimble/errors
```

## Usage

```go

func main(){

  err1 = errors.New("error message")
  err2 = errors.Newf("%s", "error message")
  err3 = errors.NewMsg(
      errors.HTTP(404),
      errors.Msg("error message"),
      errors.Type("Invalid Arguments"),
      errors.Code("123"),
      errors.Dev("dev err message"),
  )
}


```
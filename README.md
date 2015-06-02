errors [![godoc badge](http://godoc.org/github.com/plimble/errors?status.png)](http://godoc.org/github.com/plimble/errors)
========


## Installation

```go
$ go get -u github.com/plimble/errors
```

## Usage

```go

func main(){

  err1 = errors.New("error message").Http(404).Code("1000").Type("Result not found")

  err2 := errors.New("error message").Http(errors.BadReq).Code("1001").Type("Invalid Arguments")

  errors.EnableDevMsg = true
  err2.DevMsg("Field is empty")
}


```
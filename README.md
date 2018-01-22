```go

package main  //always need

import (
  //"fmt", etc.
)

const (
  //declare w/ global scope,
  ```
  _name_ **type**
  ```go
)
var (
  //
)

func main() {
  func syntax() {
    fmt.Println("...")  //fmt.Printf("%g:%g", 1, 10)  //=> 1:10

    a := 5  //declare variable w/in function; type set implicitly

    for i := 0; i < 10; i++ {
      //
    }

    if /* variables; condition */ {
      //
    } else {
      //
    }

    switch /* ... */ { //w/o condition, defaults to `true`--can use like if/elsif/else
      case /* value */ :
        //runs first true case; break is implicit
      default:
        //
    }

    return  //explicit
  }
  
  defer func child() {
    returns := "after parent"
    return  //=> returns  --implicitly
  }
  
  type Vertex struct {
    property1 string
    property2 int
  }
}

```

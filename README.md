# learngo

A simple tutorial for myself as I would like to refer to key concepts ( specially in concurrency ). 
Each directory has a single file with package main so that the tutorial is somple and one can easily build it 
by running go build
or
go build filename

One can experiment firther and I have added comments to explain what is being done.

```go
package main

import (
	"fmt"
)

type Outer struct {
	name string
	Inner
}

type Inner struct {
	age uint
}

//type InnerInterface interface {
//	GetAge() uint
//}

func (inner Inner) GetAge() uint {
	return inner.age
}

func main() {
	fmt.Println(" Hello behaviour")
	inner := Inner{age: 23}
	outer := Outer{name: "Beaut", Inner: inner}
	fmt.Println(outer)
	fmt.Println(outer.age)
	outer.GetAge()

}
```

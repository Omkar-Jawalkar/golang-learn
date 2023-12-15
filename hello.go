package main

import (
	"fmt"
	"reflect"
	"strconv"
)


var pl = fmt.Println

func main() {

	cV1 := "15"
	cV2, err := strconv.Atoi(cV1)
	pl(cV2, err, reflect.TypeOf(cV2))
	
	
	
}
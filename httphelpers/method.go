package httphelpers

import "fmt"

func Method(method string) Processor {
	return Condition(func(r BarebonesRequest) bool {
		fmt.Println("Method", r.Method, method)
		return r.Method == method
	})
}

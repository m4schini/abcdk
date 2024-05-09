package main

import (
	"context"
	"fmt"
	"github.com/m4schini/abcdk/kv"
)

func main() {
	kv, err := kv.OpenKV(context.TODO(), "valkey://127.0.0.1/0")
	if err != nil {
		panic(err)
	}

	err = kv.Set(context.TODO(), "example", []byte("example_value"))
	if err != nil {
		panic(err)
	}

	v, err := kv.Get(context.TODO(), "example")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(v))
}

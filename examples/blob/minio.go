package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/m4schini/abcdk/blob"
	"io"
	"os"
	"time"
)

func main() {
	key := "example.txt"
	bucket, err := blob.OpenBucket("s3:///abcdk?key=uDPXpUjzId29dXM2t2lR&secret=vgIE5aqj9n7psjReHNJwigdpc8K0Z9q8HZ79kdQl")
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader([]byte("this is a example text"))
	url, err := bucket.Upload(context.TODO(), key, r)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)

	fmt.Println("sleeping 1 minute")
	time.Sleep(time.Minute)

	rf, err := bucket.Download(context.TODO(), key)
	if err != nil {
		panic(err)
	}
	t, err := os.CreateTemp("", "abcdk")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(t, rf)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Name())

	fmt.Println("sleeping 1 minute")
	time.Sleep(time.Minute)

	err = bucket.Delete(context.TODO(), "example.txt")
	if err != nil {
		panic(err)
	}
}

package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	p := new([10]SyncedBuffer) // type *SyncedBuffer
	var v SyncedBuffer         // type  SyncedBuffer
	fmt.Println(p)
	fmt.Println(v)

}

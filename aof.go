package main

import (
	"bufio"
	"os"
	"sync"
	"time"
)

type Aof struct {
	file *os.File
	rd   *bufio.Reader
	mu   sync.Mutex
}

func newAof(path string) (*Aof,error){
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666 );
	if err!=nil {
		return nil, err
	}
	Aof := &Aof{
		file:f, 
		rd: bufio.NewReader(f)
	}

	go func() {
		for{
			Aof.mu.Lock()
			Aof.file.Sync()
			Aof.mu.Unlock()
			time.Sleep(time.Second)
		}
	}()

	return Aof, nil
}

package main

import (
	"fmt"
	"errors"
	"sync"
	"time"
	"math/rand"
)

const (
	numberOfGoroutines = 2
)

//Wait group for the two go-routines of producer and consumer.
var wg sync.WaitGroup

type RingBuffer struct {
	start int
	end int
	writeCount int
	readCount int
	data [10]string
}

func (r *RingBuffer) write(item string) {
	if (r.end + 1) % len(r.data) != r.start {
		r.data[r.end] = item
		r.end = (r.end + 1) % len(r.data)
		r.writeCount ++
	}
}

func (r *RingBuffer) read() (string, error) {
	if r.start != r.end {
		r.readCount++
		item := r.data[r.start]
		r.start = (r.start + 1) % len(r.data)
		return item, nil
	}

	return "", errors.New("Empty buffer")
}

func (r RingBuffer) isEmpty() bool {
	return r.start == r.end
}

func (r RingBuffer) isFull() bool {
	return (r.end + 1) % len(r.data) == r.start
}

func producer(r *RingBuffer,
	stop chan bool,
	spaceAvailable *sync.Cond,
	contentAvailable *sync.Cond) {

	rand.Seed(86)
	
	defer wg.Done()

	for {
		is_stop := false
		select {
		case <- stop:
			is_stop = true
		case <- time.After(10 * time.Millisecond):
			is_stop = false
		}

		if is_stop {
			fmt.Println("Producer stopping now.")
			return
		}

		spaceAvailable.L.Lock()

		for r.isFull() {
			spaceAvailable.Wait();
		}

		item := fmt.Sprintf("%d", rand.Intn(10000))
		r.write(item)
		fmt.Println("Producer wrote: ", item)
		contentAvailable.Signal()
		spaceAvailable.L.Unlock()
	}
}

func consumer(r *RingBuffer,
	stop chan bool,
	spaceAvailable *sync.Cond,
	contentAvailable *sync.Cond) {
	
	defer wg.Done()

	for {
		is_stop := false
		select {
		case <- stop:
			is_stop = true
		case <- time.After(10 * time.Millisecond):
			is_stop = false
		}

		if is_stop {
			fmt.Println("Stopping Consumer")
			return
		}

		contentAvailable.L.Lock()
		
		for r.isEmpty() {
			contentAvailable.Wait();
		}
			
		item, _ := r.read()
		fmt.Println("Consumer Read: ", item)
		spaceAvailable.Signal()
		contentAvailable.L.Unlock()
	}
}

func main() {
	wg.Add(numberOfGoroutines)
	fmt.Println("Starting")

	//Create a channel for both producer and consumer to listen and react so they can stop.
	stopButton := make(chan bool)
	
	lock := sync.Mutex{}
	spaceAvailable := sync.NewCond(&lock)
	contentAvailable := sync.NewCond(&lock)
	
	
	r := RingBuffer{}
	go producer(&r, stopButton, spaceAvailable, contentAvailable)
	go consumer(&r, stopButton, spaceAvailable, contentAvailable)

	time.Sleep(10 * time.Second)
	
	close(stopButton)
	
	wg.Wait()
}

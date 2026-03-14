package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var ops int32

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Need at least one phrase and file parameters!")
		return
	}

	phrase := os.Args[1]
	A := make(chan int)
	jobNumbers := len(os.Args) - 2	
	atomic.AddInt32(&ops, int32(jobNumbers))

	for i := 2; i < len(os.Args); i++ {
		fn := os.Args[i]
		//reads the file and fins the number of occurences of a given phrase
		go scanFile(fn, phrase, A)

	}

	//calculate totoal number of occurence
	total := sum(A)
	fmt.Println("total occurences: ", total)
}

func sum(in <- chan int) int {
	total := 0
	for x := range in {
		total += x
	}

	return total
}

func scanFile(fn, phrase string, out chan<- int) {
	defer func() {
		atomic.AddInt32(&ops, -1)
		if atomic.LoadInt32(&ops) == 0 {
			close(out)
			return
		}
	}()


	f, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		if phrase == scanner.Text() {
			count++
		}
	}
	fmt.Println("file: ", f.Name(), ", occurenecce: ", count)
	//send to the channel
	out <- count 
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

}
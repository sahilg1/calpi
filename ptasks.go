/*This application is used to apply load on the CPU and is meant for testing on containers.
The program is written in golang.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	for {
		if checkval() {
			threads := numthreads()
			writeout()
			for i := 0; i < threads; i++ { //creates 10 million threads. this can be altered to put different load on the CPU
				wg.Add(1)
				go calc() //calls thread to calculate the value of pi
			}
			wg.Wait()
		}
	}

}

func calc() { //function to calculate the value of pi
	var N = 10000
	var sum float64
	var term float64
	var sign = 1.0
	for k := 0; k < N; k++ {
		term = float64((1.0) / (float64(2.0)*float64(k) + float64(1.0)))
		time.Sleep(5 * time.Millisecond)
		sum = sum + float64(sign)*term
		sign = -sign
	}
	fmt.Println("Pi=", float64(sum*4.0))
	wg.Done()
}

//This function checks the value of the S3 bucket to check to stop or continue running
func checkval() bool {
	res, err := http.Get("https://s3.amazonaws.com/sahgupta-cpu-testing/s3file.txt")
	checkerr(err)
	scanner := bufio.NewScanner(res.Body)
	scanner.Scan()
	fmt.Println(scanner.Text())
	if scanner.Text() == "stop" || scanner.Text() == "Stop" {
		res.Body.Close()
		return false
	}
	res.Body.Close()
	return true
}

func numthreads() int {
	res, err := http.Get("https://s3.amazonaws.com/sahgupta-cpu-testing/NumThreads.txt")
	checkerr(err)
	scanner := bufio.NewScanner(res.Body)
	scanner.Scan()
	nthreads, _ := strconv.Atoi(scanner.Text())
	return nthreads
}

func writeout() {
	num := numthreads()
	num = num / 2
	nthreads := strconv.Itoa(num)

	usr, er1 := user.Current()
	checkerr(er1)
	fmt.Println(usr.HomeDir)
	f, err := os.Create("NumThreads.txt")
	checkerr(err)
	defer f.Close()
	_, e := f.WriteString(nthreads)
	checkerr(e)
	fmt.Printf("Wrote number of threads for next container")

}

func checkerr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

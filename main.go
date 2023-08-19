package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func checkandSaveBody(){
	
}

func f1(wg *sync.WaitGroup){
	fmt.Println("f1 (goroutine) execution started")

	for i:=0 ; i<3 ; i++{
		fmt.Println("f1 , i=", i)
		time.Sleep(time.Second)
	}

	fmt.Println("f1 execution finished")
	wg.Done()
}

func f2(){
	fmt.Println("f2 execution started")
	for i:= 0 ; i<9;i++{
		fmt.Println("f2 ,  i=i" , i)
	}

	fmt.Println("f2 execution finished")
}



func main(){

	var wg sync.WaitGroup

	 wg.Add(1)
	fmt.Println("main execution started")
	fmt.Println("no of cpus:" , runtime.NumCPU()) 
	fmt.Println("no of goroutines" , runtime.NumGoroutine())
	

	fmt.Println("os" , runtime.GOOS)
	fmt.Println("arch"  , runtime.GOARCH)

	fmt.Println("gomaxprocs" , runtime.GOMAXPROCS(0))

     go f1(&wg) 

 fmt.Println("no of goroutines after go f1()",runtime.NumGoroutine())


  f2()


  wg.Wait()

 fmt.Println("main execution ended")


}
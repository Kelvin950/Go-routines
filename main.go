package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	

	"strings"
	"sync"
	"time"
)

func checkandSaveBody(url string , wg *sync.WaitGroup){
 

 res ,err := http.Get(url) 
  if err!=nil{

	fmt.Printf("%s is down\n" , url)
  }else {

	defer res.Body.Close() 

  
	
	 if res.StatusCode ==200 {
		  
	bodyBytes ,err := ioutil.ReadAll(res.Body) 
	 

    file := strings.Split(url , "//")[1] 
     file+=  ".txt" 
   
	  fmt.Printf("writing to file %s" ,file)
	  err= ioutil.WriteFile(file ,bodyBytes , 0664) 

   if err!=nil{

	log.Fatal("failed")
   }


	 }
	 
  }

  wg.Done()
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

	var wg *sync.WaitGroup


// 	fmt.Println("main execution started")
// 	fmt.Println("no of cpus:" , runtime.NumCPU()) 
// 	fmt.Println("no of goroutines" , runtime.NumGoroutine())
	

// 	fmt.Println("os" , runtime.GOOS)
// 	fmt.Println("arch"  , runtime.GOARCH)

// 	fmt.Println("gomaxprocs" , runtime.GOMAXPROCS(0))

//      go f1(&wg) 

//  fmt.Println("no of goroutines after go f1()",runtime.NumGoroutine())


//   f2()


//   wg.Wait()

//  fmt.Println("main execution ended")

 
var url = []string{"https://www.golang1.org", "https://www.google.com", "https://www.medium.com"}
	 wg.Add(len(url))

for _,a :=range url{

	checkandSaveBody(a ,wg)
}


wg.Wait()
}
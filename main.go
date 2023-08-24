package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"time"

	"strings"
)



func f1(n int ,ch chan int)   {

 ch <- n 


}

func checkandSaveBody(url string , ch chan string){

 res ,err := http.Get(url)
  if err!=nil{

	s := fmt.Sprintf("%s is down\n" , url)
	s +=fmt.Sprintf("  value %v" , err)
	ch <-  s
  }else {time.Sleep(time.Second *2) 


	
	defer res.Body.Close()

	var s string
	 if res.StatusCode ==200 {

	bodyBytes ,err := ioutil.ReadAll(res.Body)

    file := strings.Split(url , "//")[1]
     file+=  ".txt"

	 s = fmt.Sprintf("writing to file %s" ,file)
	  err= ioutil.WriteFile(file ,bodyBytes , 0664)

   if err!=nil{

	s +=fmt.Sprintf(" failed %v" , err)
   }
   

	 }

	  ch <- s
  }

 
}

// func f1(wg *sync.WaitGroup){
// 	fmt.Println("f1 (goroutine) execution started")

// 	for i:=0 ; i<3 ; i++{
// 		fmt.Println("f1 , i=", i)
// 		time.Sleep(time.Second)
// 	}

// 	fmt.Println("f1 execution finished")
// 	wg.Done()
// }

// func f2(){
// 	fmt.Println("f2 execution started")
// 	for i:= 0 ; i<9;i++{
// 		fmt.Println("f2 ,  i=i" , i)
// 	}

// 	fmt.Println("f2 execution finished")
// }


func factorial(n int  , c chan int ){

	 f := 1 

	 for i := 2 ;  i<=n ; i++{
		f *=  i 
	 }
	
	 c <-f
}


func main(){





	ch := make( chan float64)

 power :=  func(n float64 , ch chan float64) {
 
  
 ch <- math.Pow(n ,2)	 
 }
// var wg sync.WaitGroup 

// wg.Add(3) 
 
 var  i = 0.0
for   i < 50.0  {
 go power(i ,ch)


 i++ 
}

for i:=0 ; i<50;i++{

fmt.Println( <-ch)
}

// go func(wg *sync.WaitGroup){
//  fmt.Println("hello world")

//  wg.Done()
// }(&wg)


// wg.Wait()

// sum := func( wg * sync.WaitGroup ,n ...float64 ){
  

//   fmt.Println(fmt.Sprintf("%.2f" , n[0]+n[1]))
 
  
//   wg.Done()
// }


// go sum(&wg , 1.0 , 2.0)
// go sum(&wg , 1.0 , 2.0)
// go sum(&wg , 1.0 , 2.0)

// wg.Wait()
	// var wg sync.WaitGroup


// 	fmt.Println("main execution started")
// 	fmt.Println("no of cpus:" , runtime.NumCPU()) 
// 	fmt.Println("no of goroutines" , runtime.NumGoroutine())
	

// 	fmt.Println("os" , runtime.GOOS)
// 	fmt.Println("arch"  , runtime.GOARCH)

// 	fmt.Println("gomaxprocs" , runtime.GOMAXPROCS(0))

//      go f1(&wg) 

//  fmt.Println("no of goroutines after go f1()",runtime.NumGoroutine())


//   "strconv"
	// "strings"f2()


//   wg.Wait()

//  fmt.Println("main execution ended")

 
// var url = []string{"https://www.golang1.org", "https://www.google.com", "https://www.medium.com"}
// 	 wg.Add(len(url))


//  ch :=make(chan  string) //unbuffered channel 
 

// c1 :=make(chan  int) 
//  _=c1
// c2 :=make(chan int, 3) //buffered channel


// // go func(c chan int) {
// //  fmt.Println("func go routine start sending data into the channel")
// //  c <-10
// //  fmt.Println("func go routine after sending data into the channel")
// // }(c1)
 

// // fmt.Println("main goroutine sleeps for 2 seconds")
// // time.Sleep(time.Second *2) 


// // fmt.Println("main go routine start receiving data")
// // d := <-c1 

// // fmt.Println(d) 
 
// go func(c chan int) {

// 	for  i :=1 ; i<=5 ;i++{
 
// 		fmt.Printf("func goroutine %d  starts sending data\n" , i)
// 		c <- i 
// 		fmt.Printf("func goroutine %d after sending data\n" ,i)
	
// 	}
// 	close(c)
// }(c2)

// fmt.Println("main sleeps for 2 secons")
// time.Sleep(time.Second*2) 





// for v := range c2 {

// 	fmt.Println("main go received value" , v)
// }
	  
// for _,a :=range url{

//   go checkandSaveBody(a  , ch) 
 
// }

// fmt.Println("number of channels " ,runtime.NumGoroutine())

// for  i:=0 ; i<len(url) ;i++{

// 	fmt.Printf(" %s\n" , <-ch )
}
// fmt.Printf("number of goroutines %d", runtime.NumGoroutine() )
// wg.Wait()

// var ch chan int
 


// ch =  make(chan int)

// fmt.Println(ch)

//  c := make(chan int) 
  
// //  fmt.Println(c)
// // <- channel operator 

// c <- 10 


// num := <- c 
 
// _ = num 
// fmt.Println(<-c)


// close(c)

// c :=  make(chan int) 

// //only receiving  
// c1:=make(<-chan string) 

// //only send 

// c2 := make(chan<- string )

// fmt.Printf("%T , %T , %T" , c1 ,c, c2)
 
// go f1(10 ,c)
 

// n := <-c 

// fmt.Printf("%d\n" ,n)

// ch := make(chan int) 

// defer close(ch) 

// go factorial(5, ch) 


// nums := <-ch 

// fmt.Printf("%d\n" , nums)

// for  i:=1 ;  i<20 ; i++{

// 	go factorial(i , ch) 
// 	f := <-ch 
// 	fmt.Printf("%d\n" , f)
// } 



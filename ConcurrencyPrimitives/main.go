package main

import "fmt"
import "time"

func someFuncRoutine(msg string) {
	//
	// channel <- msg
	fmt.Println(msg)
}

// func someFuncChannel(msg string, channel chan) bool{
// 	channel <- msg
// }


func goRoutines(){
	/*     GO Routines */

	go someFuncRoutine("1")
	go someFuncRoutine("2")

	time.Sleep(time.Second * 2)
	//Others wont run 
	fmt.Println("Hi")

}

func goChannels(){
	/* go Channels */

	// mychannel := make(chan string)

	// go func(){
	// 	mychannel <- "data"
	// }()

	// msg := <-mychannel
	// fmt.Println(msg)

	//declare channel
	ch1 := make(chan string)
	ch2 := make(chan string)
	// Routines
	go func(){
	ch1 <- "data"
	}()
	
	go func(){
		ch2 <- "data"
	}()

	msg1 := <-ch1;
	msg2 := <-ch2;
	//Others wont run 
	fmt.Println(msg1)
	fmt.Println(msg2)
}
func main(){

	goRoutines()
	goChannels()

}
// package main

// import (
// 	"fmt"
// 	"os/exec"
// 	"time"
// )

// func main(){
// 	// making a command
// 	cmd := exec.Command("ping", "google.com")

// 	//using non blocking execution and storing output in a var

// 	err := cmd.Start()
// 	if err != nil{
// 		panic(err)
// 	}

// 	// lets get the pid of the process

// 	fmt.Println("a process has started with pid %d", cmd.Process.Pid)

// 	// making a go routine that monitors the process

// 	go func(){
// 		for{
// 			fmt.Println("process is being monitored")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	e := cmd.Wait()

// 	fmt.Println("program exited due to error-->%v", e)


// }
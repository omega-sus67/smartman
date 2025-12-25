package main

import(
	"fmt"
	"bufio"
	"os/exec"
	"time"
	"os"
)

func main(){
	// defining a process
	cmd := exec.Command("ping", "google.com")

	// making a pipe for an output
	stdoutpipe, err := cmd.StdoutPipe()
	if err != nil{
		panic(err)
	}

	//starting the process
	for i := range 3{
		fmt.Println("starting the process in ", i + 1)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Process execution started")

	if err := cmd.Start(); err != nil{
		panic(err)
	}
	
	// anoimaly detector working here
	brain := addData(10, 2.0)
	scanner := bufio.NewScanner(stdoutpipe)
	
	go func(){
		for scanner.Scan() {
		line := scanner.Text()

		latency, err := parseLatency(line)

		if err != nil{
			continue
		}

		isAnomaly, status := brain.addDataAndCheck(latency)

		if isAnomaly{
			fmt.Printf("a Anomaly is detected : %s latency : %.2fms\n", status, latency)
		} else{
			fmt.Printf("Normal : %.2fms\n", latency)
		}
	}
	}()

	//killing the process in 5 secs

	go func(){
		fmt.Println("Killing the process in ", 20)
		time.Sleep(20 * time.Second)
		//killing the process now
		if err := cmd.Process.Signal(os.Kill); err != nil{
			fmt.Println("Failed to kill the process!! %v", err)
		}
	}()

	err1 := cmd.Wait()

	if err1 != nil{
		fmt.Printf("worker has died unaturally : %v\n", err)
	} else{
		fmt.Println("the worker has died naturally")
	}
}
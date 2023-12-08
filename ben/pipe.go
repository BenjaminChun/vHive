package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"bufio"
	"io"
	"log"
)

func readStdout(pipe io.Reader) {
	scanner := bufio.NewScanner(pipe)

	for scanner.Scan() {
		// Process each line of stdout
		line := scanner.Text()
		fmt.Println("Received from stdout:", line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func runTestOut(wg *sync.WaitGroup, output chan<- []byte) {
	defer wg.Done()

	cmd := exec.Command("./test.out")

	// Capture the output of ./test.out
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for ./test.out:", err)
		return
	}
	fmt.Println(stdout)

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting ./test.out:", err)
		return
	}

	go readStdout(stdout)

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for ./test.out:", err)
	}
}

func runVirtToPhysOut(input <-chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	// Receive the output of ./test.out
	output1 := <-input

	// Parse the output to extract vaddr and pid
	fields := strings.Fields(string(output1))
	if len(fields) < 4 {
		fmt.Println("Error parsing output of ./test.out. Expected at least 4 fields.")
		return
	}

	vaddr := fields[1]
	pid := fields[3]

	// Run ./virt_to_phys_user.out with extracted arguments
	cmd := exec.Command("./virt_to_phys_user.out", pid, vaddr)

	// Pipe the output of ./test.out as input to ./virt_to_phys_user.out
	cmd.Stdin = strings.NewReader(string(output1))

	// Capture the output of ./virt_to_phys_user.out
	output2, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running ./virt_to_phys_user.out:", err)
		return
	}

	// Print the result of ./virt_to_phys_user.out
	fmt.Println("Result of ./virt_to_phys_user.out:", string(output2))
}

func main() {
	// Use a WaitGroup to wait for the completion of goroutines
	var wg sync.WaitGroup

	// Use a channel to pass the output of ./test.out to ./virt_to_phys_user.out
	outputChan := make(chan []byte)

	// Run ./test.out concurrently in a goroutine
	wg.Add(1)
	go runTestOut(&wg, outputChan)

	// Run ./virt_to_phys_user.out concurrently in a goroutine
	wg.Add(1)
	go runVirtToPhysOut(outputChan, &wg)

	// Wait for all goroutines to complete
	wg.Wait()

	// Close the channel
	close(outputChan)
}

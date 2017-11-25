package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var rightLine = regexp.MustCompile(os.Args[1])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	handleErr(cmd.Start())

	// pipe all stderr to our stderr without processing
	stderr, err := cmd.StderrPipe()
	handleErr(err)
	go func() {
		_, err = io.Copy(os.Stderr, stderr)
		handleErr(err)
	}()

	stdout, err := cmd.StdoutPipe()
	handleErr(err)
	scanner := bufio.NewScanner(stdout)
	foundLine := false
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if rightLine.MatchString(line) {
			foundLine = true
			break
		}
	}
	if !foundLine {
		log.Fatal("Didn't find line")
	}
}

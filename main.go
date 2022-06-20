package main

import (
	"github.com/vadimkononenko/event-loop/engine"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	//if input, err := os.Open(inputFile); err == nil {
	//	defer input.Close()
	//	scanner := bufio.NewScanner(input)
	//	for scanner.Scan() {
	//		commandLine := scanner.Text()
	//		cmd := parse(commandLine) // parse the line to get a Command
	//		eventLoop.Post(cmd)
	//	}
	//}

	eventLoop.AwaitFinish()
}

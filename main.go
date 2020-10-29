package main

import (
	"a/cmd"
	"log"
)

func init(){

	//setting flags for log level
	log.SetFlags(3)
}

func main() {
	cmd.Execute()
}

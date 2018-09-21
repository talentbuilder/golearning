package main

import (
	"log"
	"os"

	_ "github/golearning/chapter2/matchers"
	"github/golearning/chapter2/search"
)

// init is called prior to main.
func init() {  //所有模块都会先运行init
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout) //将标准输出打到log中
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")  //搜索president这个term
}

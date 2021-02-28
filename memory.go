package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type smallStruct struct {
	a, b int64
	c, d float64
}

func main() {
	memPrint()
	smallAllocation()
	memPrint()

	for i := 0; i < 1024; i++ { //1024
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}
	}

	memPrint()

}

//go:noinline
func smallAllocation() *smallStruct {
	// stackPrint()
	return &smallStruct{}
}

// func stackPrint() {
// 	stackSlice := make([]byte, 512)
// 	s := runtime.Stack(stackSlice, false)
// 	fmt.Printf("%s", stackSlice[0:s])
// }

func memPrint() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	currentTime := time.Now()
	fmt.Println(currentTime.Format("--2006-01-02 15:04:05.000000000"))

	p := message.NewPrinter(language.English)
	p.Printf("Alloc %+v\n", mem.Alloc)
	p.Printf("TotalAlloc %+v\n", mem.TotalAlloc)
	p.Printf("HeapAlloc %+v\n", mem.HeapAlloc)
	p.Printf("HeapSys %+v\n", mem.HeapSys)
	p.Printf("Sys %+v\n", mem.Sys)
}

// bigBytes allocates 100 megabytes
func bigBytes() *[]byte {
	s := make([]byte, 1024*1024*1024)
	return &s
}

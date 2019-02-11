package main

import "fmt"

var i int
var array [30]string
var flags [3]chan int
var end chan int

type contexto struct {
	flag    int
	next    int
	max     int
	toWrite string
}

func trabalho(ctx *contexto) {
	fmt.Printf("Thread %+v started\n", ctx)
	for i <= ctx.max {
		<-flags[ctx.flag]
		fmt.Printf("Thread %d to write [%s] on %d\n", ctx.flag, ctx.toWrite, i)
		array[i] = ctx.toWrite
		i++
		if i >= ctx.max {
			end <- 1
		}
		flags[ctx.next] <- 1
	}
	fmt.Printf("Thread %d exiting\n", ctx.next)
	end <- 1
}

func printArray() {
	for i := 0; i < 30; i++ {
		fmt.Printf("%s,", array[i])
	}
	fmt.Printf("\n")
}

func main() {
	i = 0
	for j := 0; j < 30; j++ {
		array[j] = "?"
	}
	flags[0] = make(chan int)
	flags[1] = make(chan int)
	flags[2] = make(chan int)
	end = make(chan int)

	ctx1 := contexto{flag: 0, next: 1, max: 27, toWrite: "a"}
	ctx2 := contexto{flag: 1, next: 2, max: 28, toWrite: "b"}
	ctx3 := contexto{flag: 2, next: 0, max: 29, toWrite: "c"}

	go trabalho(&ctx1)
	go trabalho(&ctx2)
	go trabalho(&ctx3)

	// start everything
	flags[0] <- 1

	<-end
	<-end
	<-end

	fmt.Printf("Resultado:\n")
	printArray()
}

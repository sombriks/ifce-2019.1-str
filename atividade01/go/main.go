package main

import "fmt"

var i int
var array [30]string
var flags [3]int

type contexto struct {
	flag    int
	next    int
	toWrite string
}

func trabalho(ctx contexto) {
	fmt.Printf("Thread %+v started\n", ctx)
	for i < 30 {
		// busy-wait
		for flags[ctx.flag] == 1 {

		}
		fmt.Printf("Thread %d to write on %d\n", ctx.next, i)

		array[i] = ctx.toWrite
		i++
		flags[ctx.flag] = 1
		flags[ctx.next] = 0
	}
	fmt.Printf("Thread %d exiting\n", ctx.next)
}

func printArray() {
	for i := 0; i < 30; i++ {
		fmt.Printf("%s ", array[i])
	}
	fmt.Printf("\n")
}

func main() {
	i = 0
	for j := 0; j < 30; j++ {
		array[j] = " "
	}
	flags[0] = 0
	flags[1] = 1
	flags[2] = 1

	ctx1 := contexto{flag: 0, next: 1, toWrite: "a"}
	ctx2 := contexto{flag: 1, next: 2, toWrite: "b"}
	ctx3 := contexto{flag: 2, next: 0, toWrite: "c"}

	go trabalho(ctx1)
	go trabalho(ctx2)
	go trabalho(ctx3)

	// busy-wait
	for i < 30 {
	}

	fmt.Printf("Resultado:\n")
	printArray()
}

package main

import (
	"fmt"
	"strconv"
	"apps/fibonacci/fibonacciclientproxy"
	"executionenvironment/executionenvironment"
	"shared/net"
	"shared/factories"
	"shared/parameters"
)

func main(){

	// start configuration
	executionenvironment.ExecutionEnvironment{}.Deploy("MiddlewareFibonacciServer.conf")

	// proxy to naming service
	namingClientProxy := factories.LocateNaming()

	// register
	fiboProxy := fibonacciclientproxy.FibonacciClientProxy{Host:netshared.ResolveHostIp(),Port:parameters.FIBONACCI_PORT}
	namingClientProxy.Register("Fibonacci", fiboProxy)

	fmt.Println("Fibonacci Server ready at port "+strconv.Itoa(fiboProxy.Port))

	fmt.Scanln()
	fmt.Println("done")
}

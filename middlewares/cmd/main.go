package main

import (
	"middlewares_examples/functions"
)

func execute(name string, f functions.MyFunction) {
	f(name)
}

func main() {
	name := "EDTeam Community"
	execute(name, functions.MiddlewareLog(functions.Greeting))
	execute(name, functions.MiddlewareLog(functions.SayBye))

}

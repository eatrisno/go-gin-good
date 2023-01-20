package main

import "github.com/eatrisno/go-gin-good/routers"

func main() {
	routersInit := routers.InitRouter()
	routersInit.Run()
}

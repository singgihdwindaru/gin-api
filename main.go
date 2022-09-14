package main

import "gin-api/src/routers"

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}

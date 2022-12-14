package main

import "log"

const PORT = ":7788"

func main() {
	router := SetUpRouter()

	err := router.Run(PORT)
	if err != nil {
		log.Fatal("server start ERROR", err)
	} else {
		log.Fatal("server is running")
	}

}

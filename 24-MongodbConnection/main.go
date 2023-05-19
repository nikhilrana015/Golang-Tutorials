package main

import (
	"fmt"
	"log"
	"net/http"

	routers "github.com/nikhilrana/Golang-Tutorials/24-MongodbConnection/routers"
)

func main() {

	fmt.Println("Server getting starts ...")
	router := routers.GetRouter()
	fmt.Println("Server starts ...")

	log.Fatal(http.ListenAndServe(":8000", router))

}

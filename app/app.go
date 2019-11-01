package app

import (
	"fmt"
	"github.com/x777/mvc_pattern/controllers"
	"net/http"
)
var port string = ":8081"

// Starting point of server
func StartApp() {
	fmt.Printf("Server started at port %v\n", port)
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(port, nil); err !=nil{
		panic(err)
	}
}

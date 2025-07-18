package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/elements", Elements)
	http.HandleFunc("/generic", Generic)

	//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
	// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world")
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println(fmt.Sprintf("Numbers of bytes written: %d",
			n))
	})
	//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
	// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
	_ = http.ListenAndServe(":8080", nil)
}

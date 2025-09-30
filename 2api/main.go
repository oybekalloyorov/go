package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", mux)
	
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	}else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}



// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	const serverAddr = "127.0.0.1:8081"

// 	fmt.Println("Server ishlamoqda:", serverAddr)

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodGet {
// 			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 			return
// 		}
// 		w.Write([]byte("Hola Caracola"))
// 	})

// 	log.Fatal(http.ListenAndServe(serverAddr, nil))
// }

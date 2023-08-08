package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	http.HandleFunc("/", HomeHandler)

	// Use `PORT` provided in environment or default to 3000
	var port = envPortOr("3000")
	fmt.Println(port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Server Stopped. Error is %v", err)
		os.Exit(1)
	}
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	file, _ := os.Open("./src/home/home.html")
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	fmt.Fprint(w, string(byteValue))
}

// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":3000"`.
func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}

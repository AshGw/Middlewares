package main

import (
	"fmt"
	"net/http"
)

const PORT = ":6969"
const SECURE = ""

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello mfs!"))
	})

	http.Handle("/", STSMiddleware(XCTOMiddleware(CSPMiddleware(CORSMiddleware(handler))))) // u get the point

	fmt.Printf("Server is listening on http%s://localhost%s\n", SECURE, PORT)
	http.ListenAndServe(PORT, nil)
}

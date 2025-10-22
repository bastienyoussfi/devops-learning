package main
import "fmt"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tiny go container!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

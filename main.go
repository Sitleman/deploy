package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func MagicNumberHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	magicNumber := rand.Int()

	w.Header().Set("X-Version", os.Getenv("VERSION"))
	w.Write([]byte(strconv.Itoa(magicNumber)))
}

func main() {
	http.HandleFunc("/", MagicNumberHandler)

	fmt.Println("Start service")

	http.ListenAndServe(":4000", nil)
}

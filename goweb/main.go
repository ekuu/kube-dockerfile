package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("hello world, my name: %s\n", os.Getenv("MY_NAME"))))
	})
	log.Println("listening...")
	for i := 0; i < math.MaxInt64; i++ {
		log.Printf("count: %d\n", i)
		time.Sleep(time.Second)
	}
	log.Println("listening...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

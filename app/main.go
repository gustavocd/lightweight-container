package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	APPID = os.Getenv("APPID")
	PORT  = os.Getenv("PORT")
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "appid: "+APPID+" home page: says hello!")
}

func app1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "appid: "+APPID+" app1 page: says hello!")
}

func app2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "appid: "+APPID+" app2 page: says hello!")
}

func admin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "appid: "+APPID+" ADMIN page: very few people should see this")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/app1", app1)
	http.HandleFunc("/app2", app2)
	http.HandleFunc("/admin", admin)
	log.Println("Goapp with ID: " + APPID + " is listening on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

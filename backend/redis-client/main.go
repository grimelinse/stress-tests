package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/instrumentisto/go-rtmp-bot/model"
	"log"
	"net/http"
	"os"

	"github.com/instrumentisto/go-rtmp-bot/controller"
	"github.com/instrumentisto/go-rtmp-bot/redis"
)

var (
	redis_url = flag.String("redis", "localhost:6379", "redis url")
	api_addrs = flag.String("api.addrs", ":8083",
		"Address to listen http requests for API")
	listener *redis.RedisListener
)

func main() {
	flag.Parse()
	defer os.Exit(1)
	log.Printf("START REDIS CLIENT ON %v REDIS URL: %v", *api_addrs, *redis_url)
	listener = redis.NewRedisListener(
		*redis_url, "", 0, controller.AppHandler{make(chan *model.Signal)})
	go listener.Listen()
	defer listener.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/start_test", startTest)
	router.HandleFunc("/stop_test", stopTest)
	router.HandleFunc("/status", getStatus)
	log.Fatal(http.ListenAndServe(*api_addrs, router))
}

// Start test request handler.
// Runs stress test with request parameters.
// Before run test stops tests if the test is running.
func startTest(w http.ResponseWriter, r *http.Request) {
	writeAccessHeaders(w)
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		fmt.Fprintln(w, model.GetResponse(2))
		return
	}
	decoder := schema.NewDecoder()
	start_request := new(model.StartRequest)
	err = decoder.Decode(start_request, r.PostForm)

	if err != nil {
		log.Fatal(err)
		fmt.Fprintln(w, model.GetResponse(2))
		return
	}
	if start_request.ClientCount == 0 || start_request.ModelCount == 0 {
		fmt.Fprintln(w, model.GetResponse(2))
		return
	}
	err = listener.Write("stress-test:model_count", start_request.ModelCount)
	if err != nil {
		log.Printf("can not write model count: %v", err)
		fmt.Fprintln(w, model.GetResponse(2))
		return
	}
	err = listener.Write("stress-test:client_count", start_request.ClientCount)
	if err != nil {
		log.Printf("can not write client count: %v", err)
		fmt.Fprintln(w, model.GetResponse(2))
		return
	}
	err = listener.Call(redis.START_COMMAND)
	if err != nil {
		log.Printf("can not write client count: %v", err)
		fmt.Fprintln(w, model.GetResponse(2))
		return
	}
	fmt.Fprintln(w, model.GetResponse(1))
}

// Stop test request handler.
// Stops all tests.
// Clean report.
func stopTest(w http.ResponseWriter, r *http.Request) {
	writeAccessHeaders(w)
	log.Println("STOP TEST")
	listener.Call(redis.STOP_COMMAND)
	fmt.Fprintln(w, model.GetResponse(0))
}

// Returns status of current test.
func getStatus(w http.ResponseWriter, r *http.Request) {
	writeAccessHeaders(w)
	stress_test_map, err := listener.GetMap("stress_test:status")
	if err != nil {
		fmt.Fprintln(w, model.GetResponse(2))
		log.Printf("error getting status map: %v", err)
		return
	}
	if len(stress_test_map) == 0 {
		fmt.Fprintln(w, model.GetResponse(2))
		log.Println("status map clear!!!")
		return
	}
	for key, value := range stress_test_map {
		log.Printf("server: %s status: %s", key, value)
		if value != "ready" {
			fmt.Fprintln(w, model.GetResponse(1))
			return
		}
	}
	fmt.Fprintln(w, model.GetResponse(0))
}

// Writes CORS headers to HTTP response.
func writeAccessHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods",
		"POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	fmt.Println("Sever runnig on port", os.Getenv("PORT"))
	http.ListenAndServe("127.0.0.1:"+os.Getenv("PORT"), mux)
}

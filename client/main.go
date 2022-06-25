package main

import (
	"encoding/json"
	"fmt"
	pb "github.com/CodeDSA/hello/pb"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"time"
)

func main() {
	//	Connect to Add service
	conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	codeClient := pb.NewComputeServiceClient(conn)

	routes := mux.NewRouter()
	routes.HandleFunc("/", indexHandler).Methods("GET")
	routes.HandleFunc("/add/{a}/{b}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()

		req := &pb.CodeRequest{Problem: string(vars["a"]), Code: string(vars["b"])}
		if resp, err := codeClient.ComputeCode(ctx, req); err == nil {
			msg := resp.Result
			fmt.Println(msg)
		} else {
			msg := fmt.Sprintf("Internal server error: %s", err.Error())
			fmt.Println(msg)
		}
	}).Methods("GET")

	fmt.Println("Application is running on : 8080 .....")
	http.ListenAndServe(":8080", routes)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode("Server is running")
}

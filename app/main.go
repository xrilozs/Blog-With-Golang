package main

import (
	"app/controller"
	"app/helper"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	helper.ConnectDb()
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/register", controller.AddUser).Methods("POST")
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")
	r.HandleFunc("/refresh", controller.RefreshUser).Methods("POST")
	r.HandleFunc("/posts", controller.AddBlog).Methods("POST")
	r.HandleFunc("/posts/{id}", controller.UpdateBlog).Methods("PUT")
	r.HandleFunc("/posts", controller.GetAllBlog).Methods("GET")
	r.HandleFunc("/posts/{id}", controller.GetBlogByID).Methods("GET")
	r.HandleFunc("/posts/{id}", controller.DeleteBlog).Methods("DELETE")
	r.HandleFunc("/posts/{id}/comments", controller.GetCommentsByPostId).Methods("GET")
	r.HandleFunc("/posts/{id}/comments", controller.AddComment).Methods("POST")

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

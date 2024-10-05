package controller

import (
	"app/helper"
	"app/models"
	"app/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetAllBlog(w http.ResponseWriter, r *http.Request) {
	blogs, err := service.GetBlogs()
	if err != nil {
		helper.SendResponse(w, 400, "Get All Blog is failed", nil)
		return
	}

	helper.SendResponse(w, 200, "Get All Blog is success", blogs)
	return
}

func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		helper.SendResponse(w, 200, "Missing parameter", nil)
		return
	}
	idNum, _ := strconv.Atoi(idStr)

	blog, err := service.GetBlogById(idNum)
	if err != nil {
		helper.SendResponse(w, 400, "Get Blog by ID is failed", nil)
		return
	}

	helper.SendResponse(w, 200, "Get Blog by ID is success", blog)
	return
}

func AddBlog(w http.ResponseWriter, r *http.Request) {
	token, err := helper.GetAuthHeader(r)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}
	registeredClaim, err := helper.ValidateToken(token)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	user, err := service.GetUserByEmail(registeredClaim.Subject)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	var reqBlog models.Blog
	errParse := json.NewDecoder(r.Body).Decode(&reqBlog)
	if errParse != nil {
		helper.SendResponse(w, 400, errParse.Error(), nil)
		return
	}

	if reqBlog.Content == "" || reqBlog.Title == "" {
		helper.SendResponse(w, 400, "Invalid param", nil)
		return
	}

	reqBlog.CreatedAt = time.Now()
	reqBlog.AuthorID = int(user.ID)
	blog, err := service.AddBlog(reqBlog)
	if err != nil {
		helper.SendResponse(w, 400, "Add Blog is failed", nil)
		return
	}

	helper.SendResponse(w, 200, "Add Blog is success", blog)
	return
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	token, err := helper.GetAuthHeader(r)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}
	registeredClaim, err := helper.ValidateToken(token)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	user, err := service.GetUserByEmail(registeredClaim.Subject)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	var reqBlog models.Blog
	errParse := json.NewDecoder(r.Body).Decode(&reqBlog)
	if errParse != nil {
		helper.SendResponse(w, 400, errParse.Error(), nil)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		helper.SendResponse(w, 400, "Missing parameter", nil)
		return
	}
	idNum, _ := strconv.Atoi(idStr)

	if reqBlog.Content == "" || reqBlog.Title == "" {
		helper.SendResponse(w, 400, "Invalid param", nil)
		return
	}

	existingBlog, _ := service.GetBlogById(idNum)
	if existingBlog == nil {
		helper.SendResponse(w, 400, "Blog is not found", nil)
		return
	}

	if existingBlog.AuthorID != int(user.ID) {
		helper.SendResponse(w, 400, "Invalid access", nil)
		return
	}

	existingBlog.Title = reqBlog.Title
	existingBlog.Content = reqBlog.Content
	existingBlog.UpdatedAt = time.Now()
	blog, err := service.UpdateBlog(existingBlog)
	if err != nil {
		helper.SendResponse(w, 400, "Update Blog is failed", nil)
		return
	}

	helper.SendResponse(w, 200, "Update Blog is success", blog)
	return
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	token, err := helper.GetAuthHeader(r)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}
	registeredClaim, err := helper.ValidateToken(token)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	user, err := service.GetUserByEmail(registeredClaim.Subject)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		helper.SendResponse(w, 400, "Missing parameter", nil)
		return
	}
	idNum, _ := strconv.Atoi(idStr)

	existingBlog, err := service.GetBlogById(idNum)
	if err != nil {
		helper.SendResponse(w, 400, "Blog is not found", nil)
		return
	}

	if existingBlog.AuthorID != int(user.ID) {
		helper.SendResponse(w, 400, "Invalid access", nil)
		return
	}

	blog, err := service.DeleteBlog(existingBlog)
	if err != nil {
		helper.SendResponse(w, 400, "Delete Blog is failed", nil)
		return
	}

	helper.SendResponse(w, 200, "Delete Blog is success", blog)
	return
}

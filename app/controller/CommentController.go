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

func GetCommentsByPostId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		helper.SendResponse(w, 200, "Missing parameter", nil)
		return
	}
	idNum, _ := strconv.Atoi(idStr)

	comments, err := service.GetCommentsByPostId(idNum)
	if err != nil {
		helper.SendResponse(w, 400, "Get Comment by Post ID is failed", nil)
		return
	}

	helper.SendResponse(w, 200, "Get Comment by Post ID is success", comments)
	return
}

func AddComment(w http.ResponseWriter, r *http.Request) {
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
		helper.SendResponse(w, 200, "Missing parameter", nil)
		return
	}
	idNum, _ := strconv.Atoi(idStr)

	blog, err := service.GetBlogById(idNum)
	if err != nil {
		helper.SendResponse(w, 400, "Blog is not found", nil)
		return
	}

	var reqComment models.Comment
	errParse := json.NewDecoder(r.Body).Decode(&reqComment)
	if errParse != nil {
		helper.SendResponse(w, 400, errParse.Error(), nil)
		return
	}

	if reqComment.Content == "" {
		helper.SendResponse(w, 400, "Invalid param", nil)
		return
	}

	reqComment.CreatedAt = time.Now()
	reqComment.PostID = blog.ID
	reqComment.AuthorName = user.Name
	comment, err := service.AddComment(reqComment)
	if err != nil {
		helper.SendResponse(w, 400, "Add Comment is failed", nil)
		return
	}

	helper.SendResponse(w, 200, "Add Comment is success", comment)
	return
}

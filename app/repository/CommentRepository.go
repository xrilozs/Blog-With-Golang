package repository

import (
	"app/helper"
	"app/models"
)

func GetCommentsByPostId(postId int) ([]models.Comment, error) {
	var comments []models.Comment

	result := helper.DBConn.Where("post_id = ?", postId).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}

func GetCommentById(id int) (*models.Comment, error) {
	var comment models.Comment

	result := helper.DBConn.First(&comment, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil
}

func AddComment(comment models.Comment) (*models.Comment, error) {
	resultCreate := helper.DBConn.Create(&comment)
	if resultCreate.Error == nil {
		return nil, resultCreate.Error
	}

	return &comment, nil
}

func UpdateComment(comment models.Comment) (*models.Comment, error) {
	resultUpdate := helper.DBConn.Save(&comment)
	if resultUpdate.Error == nil {
		return nil, resultUpdate.Error
	}

	return &comment, nil
}

func DeleteComment(comment models.Comment) (*models.Comment, error) {
	result := helper.DBConn.Delete(comment)
	if result.Error == nil {
		return nil, result.Error
	}

	return &comment, nil
}

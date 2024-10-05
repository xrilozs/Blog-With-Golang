package service

import (
	"app/models"
	"app/repository"
)

func GetCommentsByPostId(postId int) ([]models.Comment, error) {
	comments, err := repository.GetCommentsByPostId(postId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func GetCommentById(id int) (*models.Comment, error) {
	comment, err := repository.GetCommentById(id)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func AddComment(newComment models.Comment) (*models.Comment, error) {
	comment, err := repository.AddComment(newComment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func UpdateComment(updateComment models.Comment) (*models.Comment, error) {
	comment, err := repository.UpdateComment(updateComment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func DeleteComment(deleteComment models.Comment) (*models.Comment, error) {
	comment, err := repository.DeleteComment(deleteComment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

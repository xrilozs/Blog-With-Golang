package repository

import (
	"app/helper"
	"app/models"
)

func GetBlogs() ([]models.Blog, error) {
	var blogs []models.Blog

	result := helper.DBConn.Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}

	return blogs, nil
}

func GetBlogById(id int) (*models.Blog, error) {
	var blog models.Blog

	result := helper.DBConn.First(&blog, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blog, nil
}

func AddBlog(blog models.Blog) (*models.Blog, error) {
	resultCreate := helper.DBConn.Create(&blog)
	if resultCreate.Error == nil {
		return nil, resultCreate.Error
	}

	return &blog, nil
}

func UpdateBlog(blog *models.Blog) (*models.Blog, error) {
	resultUpdate := helper.DBConn.Save(&blog)
	if resultUpdate.Error == nil {
		return nil, resultUpdate.Error
	}

	return blog, nil
}

func DeleteBlog(blog *models.Blog) (*models.Blog, error) {
	result := helper.DBConn.Delete(blog)
	if result.Error == nil {
		return nil, result.Error
	}

	return blog, nil
}

package service

import (
	"app/models"
	"app/repository"
)

func GetBlogs() ([]models.Blog, error) {
	blogs, err := repository.GetBlogs()
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func GetBlogById(id int) (*models.Blog, error) {
	blog, err := repository.GetBlogById(id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func AddBlog(newBlog models.Blog) (*models.Blog, error) {
	blog, err := repository.AddBlog(newBlog)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func UpdateBlog(updateBlog *models.Blog) (*models.Blog, error) {
	blog, err := repository.UpdateBlog(updateBlog)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func DeleteBlog(deleteBlog *models.Blog) (*models.Blog, error) {
	blog, err := repository.DeleteBlog(deleteBlog)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

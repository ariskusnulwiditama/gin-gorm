package models

import (
	_ "github.com/jinzhu/gorm"
)

type Book struct {
	ID     uint   `json:"id" gorm: "primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateBookInput struct {
	//ID     uint   `json:"id" gorm: "primary_key"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBook struct {
	Title string `json:"title"`
	Author string `json:"author"`
}
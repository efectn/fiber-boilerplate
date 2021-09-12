package models

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Articles map[int]Article
var IDs int

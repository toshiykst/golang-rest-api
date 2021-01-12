package main

import (
	"github.com/toshiykst/golang-rest-api/app/infrastructure"
	"gorm.io/gorm"
)

// Post is a structure of the post.
type Post struct { // TODO: Move to domain.
	gorm.Model
	Title   string
	Content string
}

func main() {
	infrastructure.Setup()
}

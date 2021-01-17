package main

import (
	i "github.com/toshiykst/golang-rest-api/app/infrastructure"
)

func main() {
	db := i.NewDB()
	i.RunRouter(db)
}

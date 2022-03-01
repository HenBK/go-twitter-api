package main

import (
	"fmt"

	"github.com/henbk/go-twitter-api/db"
	"github.com/henbk/go-twitter-api/handlers"
)

func main() {
	fmt.Println(db.GetConnectionStatus())
	handlers.Handlers()
}

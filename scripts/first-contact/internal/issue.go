package internal

import "fmt"

type TimeContainer struct {
	Time  float64
	Error error
}

type Issue struct {
	CreatedAt   string `json:"created_at"`
	NumComments int    `json:"comments"`
	CommentsURL string `json:"comments_url"`
}

type Comment struct {
	User struct {
		Login string `json:"login"`
	} `json:"user"`
	CreatedAt string `json:"created_at"`
}

func main() {
	fmt.Println("vim-go")
}

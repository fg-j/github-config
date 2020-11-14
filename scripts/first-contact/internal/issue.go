package internal

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

//go:generate faux --interface CommentGetter --output fakes/comment_getter.go
type CommentGetter interface {
	GetFirstReply() (Comment, error)
	GetCreatedAt() string
}

func (i *Issue) GetComments() ([]Comment, error) {
	return nil, nil
}

func (i *Issue) GetCreatedAt() string {
	return ""
}

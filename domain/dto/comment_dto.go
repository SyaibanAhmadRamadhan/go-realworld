package dto

type RequestCreateComment struct {
	ArticleID string `json:"article_id"`
	AuthorID  string `json:"-"`
	Body      string `json:"body"`
}

type RequestUpdateComment struct {
	ArticleID string `json:"article_id"`
	AuthorID  string `json:"-"`
	Body      string `json:"body"`
}

type ResponseComment struct {
	ID        string       `json:"id"`
	ArticleID string       `json:"article_id"`
	Author    ResponseUser `json:"author"`
	Body      string       `json:"body"`
	CreatedAt string       `json:"created_at"`
	UpdateAt  string       `json:"update_at"`
}

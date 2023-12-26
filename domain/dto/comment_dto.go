package dto

type RequestCreateComment struct {
	ArticleId string `json:"article_id"  validate:"required,ulid"`
	AuthorId  string `json:"-"`
	Body      string `json:"body"        validate:"required"`
}

type RequestUpdateComment struct {
	CommentId string `json:"comment_id"  validate:"required"`
	ArticleId string `json:"article_id"  validate:"required"`
	AuthorId  string `json:"-"`
	Body      string `json:"body"        validate:"required"`
}

type RequestDeleteComment struct {
	CommentId string `json:"comment_id"  validate:"required"`
	ArticleId string `json:"article_id"  validate:"required"`
	AuthorId  string `json:"-"`
}

type RequestFindAllComment struct {
	ArticleId     string
	LastCommentId string
}

type ResponseComments struct {
	Comments  []ResponseComment `json:"comments"`
	LastId    string            `json:"last_id"`
	ArticleId string            `json:"article_id"`
}

type ResponseComment struct {
	Id        string       `json:"id"`
	ArticleId string       `json:"article_id"`
	Author    ResponseUser `json:"author"`
	Body      string       `json:"body"`
	CreatedAt string       `json:"created_at"`
	UpdateAt  string       `json:"update_at"`
}

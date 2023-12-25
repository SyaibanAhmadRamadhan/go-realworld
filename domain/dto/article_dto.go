package dto

type RequestCreateArticle struct {
	TagNames    []string `json:"tag_names"    validate:"required,dive,ulid,max=50,min=5"`
	AuthorID    string   `json:"author_id"    validate:"required,ulid"`
	Slug        string   `json:"slug"         validate:"required,max=80,min=10"`
	Title       string   `json:"title"        validate:"required,max=100,min=15"`
	Description string   `json:"description"  validate:"required,max=255,min=25"`
	Body        string   `json:"body"         validate:"required,min=50"`
}

type RequestUpdateArticle struct {
	ID          string   `json:"-"            validate:"required,ulid"`
	TagNames    []string `json:"tag_names"    validate:"required,dive,ulid,max=50,min=5"`
	AuthorID    string   `json:"author_id"    validate:"required,ulid"`
	Slug        string   `json:"slug"         validate:"required,max=80,min=10"`
	Title       string   `json:"title"        validate:"required,max=100,min=15"`
	Description string   `json:"description"  validate:"required,max=255,min=25"`
	Body        string   `json:"body"         validate:"required,min=50"`
}

type RequestFindOneArticle struct {
	ArticleID     string
	LastCommentID string
}

type RequestFindAllArticle struct {
	Pagination RequestPaginate
}

type DataCommentsArticle struct {
	Comments []ResponseComment `json:"comments"`
	LastID   string            `json:"last_id"`
}
type ResponseArticle struct {
	ID           string              `json:"id"`
	Tags         []ResponseTag       `json:"tags"`
	Author       ResponseUser        `json:"author"`
	DataComments DataCommentsArticle `json:"data_comments,omitempty"`
	Slug         string              `json:"slug"`
	Title        string              `json:"title"`
	Description  string              `json:"description"`
	Body         string              `json:"body"`
	CreatedAt    string              `json:"created_at"`
	UpdatedAt    string              `json:"updated_at"`
}

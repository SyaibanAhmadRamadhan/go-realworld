package dto

type RequestCreateArticle struct {
	TagNames    []string `json:"tag_names" validate:"required,dive,max=25"`
	AuthorId    string   `json:"author_id" validate:"required,ulid"`
	Title       string   `json:"title" validate:"required,max=100,min=15"`
	Description string   `json:"description" validate:"required,max=255,min=25"`
	Body        string   `json:"body" validate:"required,min=50"`
}

type RequestUpdateArticle struct {
	Id          string   `json:"-"            validate:"required,ulid"`
	TagNames    []string `json:"tag_names"    validate:"required,dive,max=25"`
	AuthorId    string   `json:"author_id"    validate:"required,ulid"`
	Title       string   `json:"title"        validate:"required,max=100,min=15"`
	Description string   `json:"description"  validate:"required,max=255,min=25"`
	Body        string   `json:"body"         validate:"required,min=50"`
}

type RequestFindAllArticle struct {
	Pagination RequestPaginate
	TagName    string
}

type ResponseArticle struct {
	Id            string        `json:"id"`
	Tags          []ResponseTag `json:"tags"`
	Author        ResponseUser  `json:"author"`
	TotalFavorite int64         `json:"total_favorite,omitempty"`
	Slug          string        `json:"slug"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Body          string        `json:"body"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}

type ResponseArticles struct {
	Articles []ResponseArticle `json:"articles"`
	Total    int64             `json:"total"`
}

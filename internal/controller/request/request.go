package request

type BookRequest struct {
	Author      string `form:"author"`
	Title       string `form:"title"`
	Description string `form:"description"`
}

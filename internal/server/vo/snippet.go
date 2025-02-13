package vo

type SnippetCreateForm struct {
	Title       string
	Content     string
	Expires     int
	FieldErrors map[string]error
}

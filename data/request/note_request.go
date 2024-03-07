package request

type CreateNoteRequest struct {
	Content string `validate:"required,min=2,max=100" json:"content"`
}

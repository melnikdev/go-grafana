package request

type CreateMovieRequest struct {
	Title string `validate:"required,min=1,max=200" json:"title"`
}

type UpdateMovieRequest struct {
	Id    int    `validate:"required"`
	Title string `validate:"required,max=200,min=1" json:"title"`
}

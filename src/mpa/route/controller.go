package route

type GetController interface {
	ServeGet(*Context) error
}

type PostController interface {
	ServePost(*Context) error
}

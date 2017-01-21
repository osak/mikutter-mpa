package route

type GetController interface {
	ServeGet(*Context) (View, error)
}

type PostController interface {
	ServePost(*Context) (View, error)
}

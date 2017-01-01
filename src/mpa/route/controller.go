package route

type Controller interface {
	ServeGet(*Context) error
	ServePost(*Context) error
}

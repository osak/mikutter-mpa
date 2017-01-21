package route

type View interface {
	Render(ctx *Context) error
}

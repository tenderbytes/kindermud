package mud

type CommandHandler func(arg string) error

type Controller struct {
	Activator func()
	Handler   CommandHandler
}

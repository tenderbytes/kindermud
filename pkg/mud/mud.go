package mud

type CommandHandler func(arg string, conn *Connection) error

type Controller struct {
	OnEnter func(conn *Connection)
	OnLeave func(conn *Connection)
	Handler CommandHandler
}

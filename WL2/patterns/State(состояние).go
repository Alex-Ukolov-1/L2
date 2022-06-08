package main

import (
	"fmt"
)

type state interface {
	open(c *Connection)
	close(c *Connection)
}

type CloseState struct{}

func (cs CloseState) open(c *Connection) {
	fmt.Println("open the connection")
	c.setState(OpenState{})
}

func (cs CloseState) close(c *Connection) {
	fmt.Println("connection is already closed")
}

type OpenState struct{}

func (os OpenState) open(c *Connection) {
	fmt.Println("connection is already open")
}

func (os OpenState) close(c *Connection) {
	fmt.Println("close the connection")
	c.setState(CloseState{})
}

type Connection struct {
	_state state
}

func (c *Connection) Open() {
	c._state.open(c)
}

func (c *Connection) Close() {
	c._state.close(c)
}

func (c *Connection) setState(state state) {
	c._state = state
}

func main() {

	con := Connection{CloseState{}}
	con.Open()
	con.Open()
	con.Close()
	con.Close()
}

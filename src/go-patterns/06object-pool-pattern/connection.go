package main

type connection struct {
	id string
}

func (c *connection) getId() string {
	return c.id
}

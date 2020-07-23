package main

import (
	"fmt"
)

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

func test() {

	c := &ConfigOne{}
	c.String()
}

func main() {
	test()
}

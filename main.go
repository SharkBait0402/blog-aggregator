package main

import (
	"github.com/sharkbait0402/blog-aggregator/internal/config"
	"fmt"
)

func main() {

	cfg, err := config.Read()
	if err!=nil {
		fmt.Errorf("read unsuccessful")
	}

	st := state{}
	st.cfg = &cfg

	cmds := commands {
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)


	cfg, err = config.Read()
		if err!=nil {
			fmt.Errorf("read unsuccessful")
		}

	fmt.Printf("%+v\n", cfg)

}

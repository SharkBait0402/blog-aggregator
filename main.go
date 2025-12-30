package main

import (
	"github.com/sharkbait0402/blog-aggregator/internal/config"
	"fmt"
	"os"
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

	args := os.Args

	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	cmd:= command {
		name: os.Args[1],
		args: os.Args[2:],
	}

	if cmd.name == "login" {
		err=handlerLogin(&st, cmd)

		if err!=nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	cfg, err = config.Read()
		if err!=nil {
			fmt.Errorf("read unsuccessful")
		}

	fmt.Printf("%+v\n", cfg)

}

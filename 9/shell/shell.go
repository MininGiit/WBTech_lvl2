package shell

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"errors"
)

type Command struct {
	Command string
	Arg     []string
}

func (cfg *Command) ParseCommand(str string) error{
	splittedStr := strings.Split(str, " ")

	cfg.Command = splittedStr[0]

	if cfg.Command == "pwd" || cfg.Command == "ps" {
		return nil
	}
	if len(splittedStr) < 2 {
		return errors.New("command not found")
	} else {
		cfg.Arg = []string{}
		cfg.Arg = append(cfg.Arg, splittedStr...)
	}
	return nil
}

func Shell() {
	root := "/home/slavik/"
	var com Command
	for {
		if root != "/home/slavik/" {
			fmt.Print("slavik@slavikBook:~", root, "$ ")
		} else {
			fmt.Print("slavik@slavikBook:~$ ")
		}
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		pipe := strings.Split(input, " | ")
		for _, command := range pipe {
			err := com.ParseCommand(command)
			if err != nil {
				continue
			}
			switch com.Command {
			case "pwd":
				fmt.Println(root)
			case "ps":
				err := Ps()
				if err != nil {
					fmt.Println(err)
				}
			case "echo":
				for i := 1; i < len(com.Arg); i++ {
					fmt.Print(com.Arg[i], " ")
				}
				fmt.Println()
			case "cd":
				err := Cd(&root, com.Arg[1])
				if err != nil {
					fmt.Println(err)
					continue
				}
			case "kill":
				for i := 1; i < len(com.Arg); i++ {
					pid, err := strconv.Atoi(com.Arg[i])

					if err != nil {
						fmt.Println(err)
						continue
					}
					if pid == 0 {
						continue
					}
					err = Kill(pid)
					if err != nil {
						fmt.Println(err)
					}
				}
			default:
				continue
			}
			
		}
	}
}
package commands

import (
	"RedisClone/store"
	"fmt"
)

type CommandError struct {
	message string
}

type ActionFunc func([]string) any 

func set(words []string) any {
	if len(words) < 2 {
		return &CommandError{message: "SET must have key and value"}
	}

	key, value := words[0], words[1]
	store.Set(key, value)

	return nil
}

func get(words []string) any {
	if len(words) < 1 {
		return &CommandError{message: "GET must have key"}
	}

	key := words[0]
	val, ok := store.Get(key)

	if !ok {
		return &CommandError{message: "GET operation failed"} 
	}

	return val 
}

var commandMap = map[string]ActionFunc{
	"SET": set,
	"GET": get,
}

func ExecuteCommand(words []string) (any) {
	command := words[0]

	if commandFunc, ok := commandMap[command]; ok {
		return commandFunc(words[1:])
	}

	return fmt.Sprintf("%s command not defined", command)
}

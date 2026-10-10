package commands

import (
	"RedisClone/store"
	"fmt"
)

type CommandError struct {
	message string
}

func (e *CommandError) Error() string {
	return e.message
}

type CommandFunc func([]string) (any, *CommandError)

func set(words []string) (any, *CommandError) {
	if len(words) < 2 {
		return nil, &CommandError{message: "SET must have key and value"}
	}

	key, value := words[0], words[1]
	store.Set(key, value)

	return nil, nil
}

func get(words []string) (any, *CommandError) {
	if len(words) < 1 {
		return nil, &CommandError{message: "GET must have key"}
	}

	key := words[0]
	val, ok := store.Get(key)

	if !ok {
		return nil, nil 
	}

	return val, nil
}

var commandMap = map[string]CommandFunc{
	"SET": set,
	"GET": get,
}

func ExecuteCommand(words []string) (any, *CommandError) {
	if len(words) < 2 {
		return nil, &CommandError{"All commands must have at least 2 arguments!"}
	}

	command := words[0]

	if commandFunc, ok := commandMap[command]; ok {
		return commandFunc(words[1:])
	}

	return nil, &CommandError{fmt.Sprintf("%s is not defined as a command", command)}
}

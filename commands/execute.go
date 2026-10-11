package commands

import (
	"RedisClone/store"
	"fmt"
	"strconv"
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
	conv, err := strconv.Atoi(value)

	if err == nil {
		store.Set(key, conv)
	} else {
		store.Set(key, value)
	}

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

func delete(words []string) (any, *CommandError) {
	if len(words) < 1 {
		return nil, &CommandError{message: "DELETE must have at least 1 key"}
	}

	for _, key := range words {
		store.Delete(key)
	}

	return nil, nil 
}

func exists(words []string) (any, *CommandError) {
	if len(words) < 1 {
		return nil, &CommandError{message: "EXISTS must have a key provided"}
	}

	val, err := get(words)

	if val != nil {
		return true, nil 
	}

	return false, err 
}

func increment(words []string) (any, *CommandError) {
	if len(words) < 1 {
		return nil, &CommandError{message: "INCR must have a key provided"}
	}

	newVal, err := store.Increment(words[0], 1)

	if err != nil {
		return nil, &CommandError{message: err.Error()}
	}

	return newVal, nil 
}

func decrement(words []string) (any, *CommandError) {
	if len(words) < 1 {
		return nil, &CommandError{message: "DECR must have a key provided"}
	}

	newVal, err := store.Increment(words[0], -1)

	if err != nil {
		return nil, &CommandError{message: err.Error()}
	}

	return newVal, nil  
}

var commandMap = map[string]CommandFunc{
	"SET": set,
	"GET": get,
	"DELETE": delete,
	"EXISTS": exists,
	"INCR": increment,
	"DECR": decrement,
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

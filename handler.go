package main

import (
	"sync"
)

var Handlers = map[string]func([]Value) Value{
	"PING":    ping,
	"GET":     get,
	"SET":     set,
	"HGET":    hget,
	"HSET":    hset,
	"HGETALL": hgetall,
}

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}
var HSETs = map[string]map[string]string{}
var HSEtsMu = sync.RWMutex{}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	return Value{typ: "string", str: args[0].bulk}

}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}
	key := args[0].bulk
	value := args[1].bulk

	SETsMu.Lock()
	SETs[key] = value

	SETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "Err wrong number of arguments"}
	}

	key := args[0].bulk

	SETsMu.RLock()
	value, ok := SETs[key]

	SETsMu.Unlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: "error", str: "ERR wrong number of argument"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	value := args[2].bulk

	HSEtsMu.Lock()

	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}

	HSETs[hash][key] = value

	HSEtsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}
	hash := args[0].bulk
	key := args[1].bulk

	HSEtsMu.Lock()
	value, ok := HSETs[hash][key]
	HSEtsMu.Unlock()
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}

func hgetall(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	hash := args[0].bulk

	HSEtsMu.Lock()
	value, ok := HSETs[hash]

	if !ok {
		return Value{typ: "null"}
	}
	values := []Value{}
	for i, v := range value {
		values = append(values, Value{typ: "bulk", bulk: i})
		values = append(values, Value{typ: "bulk", bulk: v})
	}
	return Value{typ: "array", array: values}
}

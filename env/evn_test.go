package env_test

import (
	"testing"

	"github.com/pro200/go-modules/env"
)

func TestEnv(t *testing.T) {
	env.Set(map[string]string{
		"STRING": "hello world",
		"INT":    "123456",
		"FLOAT":  "1234.56",
	})

	strVal := env.Get("STRING")
	intVal := env.GetInt("INT")
	floatVal := env.GetFloat("FLOAT")

	//fmt.Println("STRING:", strVal)
	//fmt.Println("INT:", intVal)
	//fmt.Println("FLOAT:", floatVal)

	if strVal != "hello world" || intVal != 123456 || floatVal != 1234.56 {
		t.Error("Wrong result")
	}
}

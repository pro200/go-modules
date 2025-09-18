package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var isLoadedEnv = false
var ENV = map[string]string{}

func loadEnv() {
	// 로딩 순위
	// 1. 현재 디렉토리의 .파일명.env
	// 2. 현재 디렉토리의 .config.env
	// 3. 상위 디렉토리의 .config.env
	if isLoadedEnv {
		return
	}

	exPaths := strings.Split(os.Args[0], "/")
	fileName := exPaths[len(exPaths)-1]

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	paths := strings.Split(path, "/")

	envFiles := []string{
		strings.Join(paths, "/") + "/." + fileName + ".env",
		strings.Join(paths, "/") + "/.config.env",
	}

	if len(paths) > 1 {
		envFiles = append(envFiles, strings.Join(paths[:len(paths)-1], "/")+"/.config.env")
	}

	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			isLoadedEnv = true
			return
		}
	}

	panic("not found ." + fileName + ".env or .config.env")
}

func Set(env map[string]string) {
	isLoadedEnv = true
	ENV = env
}

func Get(key string) string {
	if !isLoadedEnv {
		loadEnv()
	}

	if result, ok := ENV[key]; ok {
		return result
	}

	result := os.Getenv(key)
	if result == "" {
		panic("no env " + key)
	}

	return result
}

func GetInt(key string) int {
	r := Get(key)
	intVal, _ := strconv.Atoi(r)
	return intVal
}

func GetFloat(key string) float64 {
	r := GetInt(key)
	return float64(r)
}

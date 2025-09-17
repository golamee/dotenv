package core

import (
	"os"
	"strings"
)

func Load(filenames ...string) (err error) {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}

	for _, filename := range filenames {
		err = load(filename, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func load(filename string, overload bool) error {
	envMap, err := read(filename)
	if err != nil {
		return err
	}

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] || overload {
			_ = os.Setenv(key, value)
		}
	}
	return nil
}

func read(filename string) (envMap map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return parse(file)
}

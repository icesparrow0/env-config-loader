package env_config_loader

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Environment struct {
	Paths []string
	Name string
}

func NewEnv(paths []string, name string) Environment{
	return Environment{Paths: paths, Name: name}
}

func (Env Environment) Load(fileName string, v interface{}) error{
	pathLength := len(Env.Paths)
	var name string = Env.Name
	if len(name) == 0 {
		name = "local"
	}
	var error error = nil
	for i := 0; i < pathLength; i++ {
		path := Env.Paths[i]
		configPath := fmt.Sprintf("%s/%s/%s", path, name, fileName)
		_, err := os.Stat(configPath)
		if err != nil && os.IsNotExist(err) {
			continue
		}
		file, err := ioutil.ReadFile(configPath)
		if err != nil {
			error = err
			break
		}
		err = yaml.Unmarshal(file, v)
		if err != nil {
			error = err
			break
		}
	}
	if error != nil {
		panic(fmt.Sprintf("load config: %s failed, error: %v", fileName, error))
	}
	return error
}

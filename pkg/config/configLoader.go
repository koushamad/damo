package config

import (
	"github.com/lara-store/damo/pkg/env"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

var configs map[string]interface{}

func init() {
	configs = make(map[string]interface{})
}

func GetConfigs[T interface{}](name string, _ T) (map[string]T, error) {

	if configs[name] != nil {
		return configs[name].(map[string]T), nil
	}

	configs := make(map[string]T)

	yamlFile, err := ioutil.ReadFile(getConfigFilePath(name))

	if err != nil {
		return configs, err
	}

	err = yaml.Unmarshal(yamlFile, &configs)

	if err != nil {
		return configs, err
	}

	SetConfigs(name, configs)

	return configs, nil
}

func SetConfigs[T interface{}](name string, data map[string]T) {
	configs[name] = data
}

func Get[T interface{}](key string, defaultValue T) T {
	keys := strings.Split(key, ".")

	if len(keys) != 2 {
		return defaultValue
	}

	configs, err := GetConfigs(keys[0], defaultValue)

	if err != nil {
		return defaultValue
	}

	value := configs[keys[1]]
	check(&value, defaultValue)

	return value
}

func check(value interface{}, defaultValue interface{}) {
	v := reflect.ValueOf(value)

	if v.Elem().Kind() == reflect.String {
		if strings.HasPrefix(v.Elem().String(), "env:") {
			key := strings.TrimPrefix(v.Elem().String(), "env:")
			envValue := env.Get(key, "")

			if envValue == "" {
				v.Elem().SetString(defaultValue.(string))
			} else {
				v.Elem().SetString(envValue)
			}
		}
	}
}

func getConfigFilePath(name string) string {
	filename := name + ".yaml"
	configPath := "./configs/"
	environment := env.Get("SERVICE_NAME", "")
	environmentConfigFile := configPath + environment + "/" + filename
	configFile := configPath + filename

	if environment != "" {
		var _, err = os.Stat(environmentConfigFile)

		if os.IsNotExist(err) {
			return configFile
		}
	}

	return environmentConfigFile
}

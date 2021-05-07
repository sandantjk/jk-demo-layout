package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestConfig(t *testing.T) {
	config := Config{
		Server0Port: 8080,
		Server1Port: 8081,
	}
	out, err := yaml.Marshal(config)

	t.Log(err)
	fmt.Println(string(out))
}

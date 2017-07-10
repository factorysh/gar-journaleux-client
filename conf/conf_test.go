package conf

import (
	"testing"
)

func TestReadClientConfig(t *testing.T) {
	var conf ClientConfig
	err := ReadClientConf(&conf, "../../contrib/client.config.example.toml")
	if err != nil {
		t.Errorf("Error reading client config")
	}

	if conf.Project.Name != "example_project" {
		t.Errorf("Project is not equal to example_project")
	}

	if conf.Servers["prod"].Address != "prod.example.com:50051" {
		t.Errorf("Prod server address is not prod.example.com:50051")
	}

	if conf.Servers["preprod"].Address != "preprod.example.com" {
		t.Errorf("Preprod server address is not preprod.example.com")
	}

}

package conf

import (
	"testing"
)

func TestReadClientConfig(t *testing.T) {
	var conf ClientConfig
	err := ReadClientConf(&conf, "/home/papey/repos/bt/gitlab/factory/journaleux/contrib/client.config.example.toml")
	if err != nil {
		t.Errorf("Error reading client config")
	}

	if conf.Project.Name != "example_project" {
		t.Errorf("Project is not equal to example_project")
	}

	if conf.Servers["prod"].Domain != "prod.example.com" {
		t.Errorf("Prod server domain is not domain.example.com")
	}

	if conf.Servers["prod"].Port != "50051" {
		t.Errorf("Prod server port is not 50051")
	}

	if conf.Servers["preprod"].Port != "" {
		t.Errorf("Preprod port should be undefined")
	}

}

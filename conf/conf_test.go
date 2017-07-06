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

}

package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DriftCfg struct {
	AtlantisUrl        string
	AtlantisToken      string
	AtlantisRepoPath   string
	AtlantisConfigPath string
}
type Repo struct {
	Ref  string
	Name string
}

type ServerCfg struct {
	ApiEndpoint string `yaml:"apiEndpoint"`
	Repos       []Repo `yaml:"repos"`
}

type VcsServers struct {
	GithubServer *ServerCfg `yaml:"github"`
	GitlabServer *ServerCfg `yaml:"gitlab"`
}

func GetDriftCfg(atlantisUrl, atlantisToken, atlantisRepoPath, atlantisConfigPath string) (DriftCfg, error) {
	var d DriftCfg

	if atlantisUrl == "" {
		return d, fmt.Errorf("ATLANTIS_URL environment variable is required but not set")
	}
	d.AtlantisUrl = atlantisUrl

	if atlantisToken == "" {
		return d, fmt.Errorf("ATLANTIS_TOKEN environment variable is required but not set")
	}
	d.AtlantisToken = atlantisToken

	if atlantisRepoPath == "" {
		return d, fmt.Errorf("ATLANTIS_REPO_PATH environment variable is required but not set")
	}
	d.AtlantisRepoPath = atlantisRepoPath

	if atlantisConfigPath == "" {
		return d, fmt.Errorf("ATLANTIS_CONFIG_PATH environment variable is required but not set")
	}
	d.AtlantisConfigPath = atlantisConfigPath

	return d, nil
}

func LoadVcsConfig(repoCfgPath string) (*VcsServers, error) {
	var cfg VcsServers
	if fileExists(repoCfgPath) {
		f, err := os.ReadFile(repoCfgPath)
		if err != nil {
			return &cfg, err
		}
		err = yaml.Unmarshal(f, &cfg)
		if err != nil {
			return &cfg, err
		}
		return &cfg, nil
	}
	return &cfg, fmt.Errorf("could not find config file")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

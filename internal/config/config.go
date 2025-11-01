package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type DriftCfg struct {
	GithubRepoRef string

	AtlantisUrl        string
	AtlantisToken      string
	AtlantisRepo       string
	AtlantisConfigPath string

	AtlantisRepoSetting []string
}

type AtlanstisSetting struct {
	Projects []struct {
		Name     string `json:"name"`
		Dir      string `json:"dir"`
		Workflow string `json:"terraform"`
	} `json:"projects"`
}

type Repo struct {
	Ref  string
	Name string
}

type ServerCfg struct {
	ApiEndpoint string `yaml:"apiEndpoint"`
	Repos       []Repo `yaml:"repos"`
}

func GetDriftCfg(githubRepoToken, atlantisUrl, atlantisToken, atlantisRepoPath, atlantisConfigPath string) (DriftCfg, error) {
	var d DriftCfg

	if githubRepoToken == "" {
		return d, fmt.Errorf("GITHUB_REPO_REF environment variable is required but not set")
	}
	d.GithubRepoRef = githubRepoToken

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
	d.AtlantisRepo = atlantisRepoPath

	if atlantisConfigPath == "" {
		return d, fmt.Errorf("ATLANTIS_CONFIG_PATH environment variable is required but not set")
	}
	d.AtlantisConfigPath = atlantisConfigPath

	return d, nil
}

func LoadAtlantisConfig(repoCfgPath string) (*AtlanstisSetting, error) {

	var ats AtlanstisSetting

	if fileExists(repoCfgPath) {
		f, err := os.ReadFile(repoCfgPath)
		if err != nil {
			return &ats, err
		}
		err = yaml.Unmarshal(f, &ats)

		if err != nil {
			return &ats, err
		}

		return &ats, nil
	}

	return &ats, errors.New("could not find config file")

}

func fileExists(filename string) bool {
	home, _ := os.Getwd()
	file := filepath.Join(home, filename)

	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

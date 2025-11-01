package config_test

import (
	"os"
	"testing"

	"github.com/jukie/atlantis-drift-detection/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGetDriftCfg(t *testing.T) {
	os.Setenv("GITHUB_REPO_REF", "main")
	os.Setenv("ATLANTIS_URL", "https://api.github.com")
	os.Setenv("ATLANTIS_TOKEN", "token")
	os.Setenv("ATLANTIS_REPO_PATH", "zkfmapf123/atlantis-fargate")
	os.Setenv("ATLANTIS_CONFIG_PATH", "/path/to/atlantis.yaml")
	defer os.Clearenv()

	expectedCfg := config.DriftCfg{
		GithubRepoRef:      "main",
		AtlantisUrl:        "https://api.github.com",
		AtlantisToken:      "token",
		AtlantisRepo:       "zkfmapf123/atlantis-fargate",
		AtlantisConfigPath: "/path/to/atlantis.yaml",
	}

	cfg, err := config.GetDriftCfg("main", "https://api.github.com", "token", "zkfmapf123/atlantis-fargate", "/path/to/atlantis.yaml")
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg, cfg)
}

func TestGetDriftCfgMissingEnvVar(t *testing.T) {
	os.Unsetenv("GITHUB_REPO_REF")
	// os.Unsetenv("ATLANTIS_URL")
	// os.Unsetenv("ATLANTIS_TOKEN")
	// os.Unsetenv("ATLANTIS_REPO")
	// os.Unsetenv("ATLANTIS_CONFIG_PATH")
	defer os.Clearenv()

	_, err := config.GetDriftCfg("", "https://api.github.com", "token", "zkfmapf123/atlantis-fargate", "/path/to/atlantis.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "GITHUB_REPO_REF environment variable is required but not set")
}

// func TestLoadVcsConfig(t *testing.T) {
// 	cfgYAML := `version: 3
// projects:
//   - name: 1
//     dir: examples/ec2
//     workflow: terraform
//   - name: 2
//     dir: examples/ec2
//     workflow: terraform
// `
// 	tmpfile, err := os.CreateTemp("", "atlantis_temp.yaml")
// 	assert.NoError(t, err)
// 	defer os.Remove(tmpfile.Name())

// 	err = os.WriteFile(tmpfile.Name(), []byte(cfgYAML), 0644)
// 	assert.NoError(t, err)

// 	cfg, err := config.LoadAtlantisConfig(tmpfile.Name())
// 	assert.NoError(t, err)
// 	assert.Equal(t, len(cfg.Projects), 2)
// }

func TestLoadVcsConfigMissingFile(t *testing.T) {
	cfgPath := "/path/to/nonexistent.yaml"

	_, err := config.LoadAtlantisConfig(cfgPath)
	assert.Error(t, err)
}

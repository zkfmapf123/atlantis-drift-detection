package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/jukie/atlantis-drift-detection/internal/config"
	"github.com/jukie/atlantis-drift-detection/internal/drift"
	"github.com/jukie/atlantis-drift-detection/internal/vcs"
)

var (
	GH_URL = "https://api.github.com"
)

func main() {
	var githubToken, githubRepoToken, atlantisUrl, atlantisToken, atlantisRepo, atlantisConfigPath string

	// Define flags
	flag.StringVar(&githubToken, "GITHUB_TOKEN", os.Getenv("GITHUB_TOKEN"), "API token for Github")
	flag.StringVar(&githubRepoToken, "GITHUB_REPO_REF", os.Getenv("GITHUB_REPO_REF"), "Github Repo Ref") // main or master

	flag.StringVar(&atlantisUrl, "ATLANTIS_URL", os.Getenv("ATLANTIS_URL"), "Atlantis URL")
	flag.StringVar(&atlantisToken, "ATLANTIS_TOKEN", os.Getenv("ATLANTIS_TOKEN"), "Atlantis Token")
	flag.StringVar(&atlantisRepo, "ATLNATIS_REPO", os.Getenv("ATLNATIS_REPO"), "Atlantis Repo") // [user]/[repo]
	flag.StringVar(&atlantisConfigPath, "ATLANTIS_CONFIG_PATH", os.Getenv("ATLANTIS_CONFIG_PATH"), "Atlantis Config Path")
	flag.Parse()

	validateTokens(githubToken)

	driftCfg, err := config.GetDriftCfg(
		githubRepoToken,
		atlantisUrl,
		atlantisToken,
		atlantisRepo,
		atlantisConfigPath,
	)

	if err != nil {
		log.Fatalln(err)
	}
	servers, err := config.LoadAtlantisConfig(driftCfg.AtlantisConfigPath)

	if err != nil {
		log.Fatalln(err)
	}
	executeDriftCheck(servers, githubToken, driftCfg)
}

func validateTokens(githubToken string) {
	if githubToken == "" {
		log.Fatalln("Error: Both GitLab and GitHub tokens are not provided but at least one is required. Set GITLAB_TOKEN or GITHUB_TOKEN environment variables, or pass them using the --gitlab-token and/or --github-token flags.")
	}
}

/*
github : https://api.github.com
*/
func executeDriftCheck(servers *config.AtlanstisSetting, githubToken string, driftCfg config.DriftCfg) {

	ghClient, err := vcs.NewGithubClient(GH_URL, githubToken)
	if err != nil {
		log.Fatalln("failed to setup github client")
	}

	driftRunner(ghClient, []config.Repo{
		{
			Ref:  driftCfg.GithubRepoRef,
			Name: strings.Split(driftCfg.AtlantisRepo, "/")[1],
		},
	}, driftCfg)
}

func driftRunner(client vcs.Client, repos []config.Repo, driftCfg config.DriftCfg) {
	for _, r := range repos {
		err := drift.Run(client, r, driftCfg)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

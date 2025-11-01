test:
	go test -v ./...

clear:
	clear

.PHONY: build
build: clear
	go build -o atlantis-drift-detection . && sudo mv atlantis-drift-detection /usr/local/bin/

run: build
	atlantis-drift-detection \
		--GITHUB_TOKEN test \
		--ATLANTIS_URL https://atlantis.dev.leedonggyu.com \
		--ATLANTIS_TOKEN test \
		--ATLNATIS_REPO zkfmapf123/atlantis-fargate \
		--ATLANTIS_CONFIG_PATH atlantis.yaml



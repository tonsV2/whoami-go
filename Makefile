tag ?= latest

version ?= $(shell yq e '.version' helm/Chart.yaml)

app:
	go build -o whoami-go -ldflags "-s -w" .

image:
	KO_DOCKER_REPO=tons ko publish --base-import-paths --tags=$(tag) .

alpine-image:
	docker-compose build prod && docker-compose push prod

helm-package:
	@helm package --sign --key helm --keyring ~/.gnupg/pubring.gpg helm/

publish-helm:
	@curl --user "$(CHARTMUSEUM_AUTH_USER):$(CHARTMUSEUM_AUTH_PASS)" \
        -F "chart=@whoami-go-$(version).tgz" \
        -F "prov=@whoami-go-$(version).tgz.prov" \
        https://helm-charts.fitfit.dk/api/charts

.PHONY: image helm publish-helm

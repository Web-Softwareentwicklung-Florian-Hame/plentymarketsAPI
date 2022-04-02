#
# Makefile include for sqlboiler
# desc: generate sqlboiler structs
# link: https://github.com/golang/mock/gomock
#

.PHONY: generate-plentymarkets-api
generate-plentymarkets-api:
	rm -rf api
	openapi-generator-cli generate -g go -i openapi.yml -o api --package-name api
	rm api/go.mod
	rm api/go.sum

package gen

//go:generate curl -H "Accept: application/vnd.github.v3.raw" -L "https://api.github.com/repos/whkelvin/stamp-api-spec/contents/openapi.yaml" -o "./pkg/api/generated/spec/openapi.yaml"

//go:generate oapi-codegen -config ./oapi-gen-config.yaml ./pkg/api/generated/spec/openapi.yaml

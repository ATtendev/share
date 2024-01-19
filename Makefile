.PHONY: gen-ent
gen-ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./store/db/ent/template/*.tmpl" ./store/db/ent/schema --feature sql/execquery,intercept

.PHONY: migrate
migrate:
	go run ./store/db/migrations/migrate.go --config ./config/config.yaml

.PHONY: dev
dev:
	air --build.cmd "go build -o build/share  bin/share/main.go" --build.bin "./build/share --config ./config/config.yaml"

.PHONY: gen-swagger
gen-swagger:
	swag init --generalInfo /api/v1/v1.go --outputTypes go,yaml --output ./api/v1
	swag fmt --dir ./api/v1/

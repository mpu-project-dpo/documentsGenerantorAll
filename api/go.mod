module dpo-document-api

go 1.23.1

replace (
	app-kit => ./../pkg/app-kit
	nats => ./../pkg/nats
	nats-contracts => ./../pkg/nats-contracts
)

require (
	app-kit v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.27.0
	nats v0.0.0-00010101000000-000000000000
	nats-contracts v0.0.0-00010101000000-000000000000
)

require (
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/nats-io/nats.go v1.37.0 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/sync v0.9.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

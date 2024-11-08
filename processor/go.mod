module dpo-document-processor

go 1.23.1

replace (
	github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit => ./../pkg/app-kit
	github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats => ./../pkg/nats
	github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts => ./../pkg/nats-contracts
)

require (
	github.com/lukasjarosch/go-docx v0.5.0
	github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit v0.0.0-00010101000000-000000000000
	github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats v0.0.0-00010101000000-000000000000
	github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.27.0
)

require (
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/nats-io/nats.go v1.37.0 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

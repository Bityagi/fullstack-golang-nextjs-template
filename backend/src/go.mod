module bityagi

go 1.18

replace build/code/spec => ./../build/generated

require (
	build/code/spec v0.0.0-00010101000000-000000000000
	github.com/go-chi/cors v1.2.1
	github.com/stretchr/testify v1.8.2
	go.uber.org/zap v1.24.0
)

require github.com/go-chi/chi/v5 v5.0.8

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect

)

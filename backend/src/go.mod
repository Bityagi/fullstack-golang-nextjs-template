module bityagi

go 1.18

replace build/code/spec => ./../build/generated

require (
	build/code/spec v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.24.0
	github.com/stretchr/testify v1.8.2
)

require github.com/go-chi/chi/v5 v5.0.3 // indirect

require (
	github.com/stretchr/testify v1.8.2 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect

)

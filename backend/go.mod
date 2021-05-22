module github.com/rohaquinlop/Trainin-exercise/backend

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/go-chi/chi/v5 v5.0.3
	github.com/machinebox/graphql v0.2.2 // indirect
	github.com/matryer/is v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rohaquinlop/Trainin-exercise/backend/loaddata v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/vektah/gqlparser/v2 v2.1.0
)

replace github.com/rohaquinlop/Trainin-exercise/backend/loaddata => ./loaddata

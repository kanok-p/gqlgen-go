//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"graphql-gen/graph/repository"
)

type Resolver struct {
	Repo repository.Repository
}

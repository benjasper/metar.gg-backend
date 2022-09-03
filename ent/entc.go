//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yaml"),
		entgql.WithRelaySpec(false),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
		entc.FeatureNames("sql/upsert"),
	}

	if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}

package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature privacy,entql,schema/snapshot,sql/modifier,sql/schemaconfig,namedges,sql/lock --template ./template ./schema --target ./dbx

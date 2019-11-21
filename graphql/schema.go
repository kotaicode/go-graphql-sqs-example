//go:generate go-bindata -pkg main -nometadata -ignore ".go" -ignore ".un~" -ignore ".swp" .
package main

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func parseSchema() (*graphql.Schema, error) {
	b, err := schemaGqlBytes()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return graphql.ParseSchema(string(b), &resolver{})
}

type greeting struct {
	Msg *string
}

type resolver struct {
}

func (r *resolver) Hello(args greeting) (string, error) {
	if args.Msg != nil {
		return *args.Msg, nil
	}
	return "world", nil
}

func (r *resolver) ReceiveMessage() (string, error) {
	res, err := receiveSqsMessage()
	return res, err
}

func (r *resolver) SendMessage(args greeting) (graphql.ID, error) {
	sqsId, err := sendSqsMessage(*args.Msg)
	if err != nil {
		log.Err(err)
		return graphql.ID(""), err
	}
	return graphql.ID(sqsId), nil
}

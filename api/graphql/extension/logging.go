package extensions

import (
	"context"
	"log/slog"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type LoggingExtension struct{}

func (LoggingExtension) ExtensionName() string {
	return "LoggingExtension"
}

func (LoggingExtension) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (LoggingExtension) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	start := time.Now()
	opCtx := graphql.GetOperationContext(ctx)

	respHandler := next(ctx)

	return func(ctx context.Context) *graphql.Response {
		resp := respHandler(ctx)
		duration := time.Since(start)

		slog.Info("[GraphQL] new operation called",
			"operation", opCtx.Operation.Operation,
			"operation", opCtx.OperationName,
			"name", duration,
			"errors", resp.Errors,
		)

		return resp
	}
}

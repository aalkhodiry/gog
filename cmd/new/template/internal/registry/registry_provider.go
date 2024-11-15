package registry

import (
	"github.com/PROJECT_NAME/internal/config"
	"github.com/PROJECT_NAME/internal/db"
	"github.com/PROJECT_NAME/internal/domains/post"
	"github.com/PROJECT_NAME/internal/domains/user"
	"github.com/PROJECT_NAME/internal/errors"
	"github.com/PROJECT_NAME/internal/logger"
	"github.com/jmoiron/sqlx"
)

type RegistryProvider interface {
	db.DBProvider
	config.ConfigProvider
	logger.LoggerProvider

	// errors
	errors.ErrorProvider
	errors.ErrorHandlerProvider

	// domains
	// user
	user.RepositoryProvider
	user.ServiceProvider
	user.HandlerProvider

	// post
	post.RepositoryProvider
	post.ServiceProvider
	post.HandlerProvider
}

func (r *Registry) DB() *sqlx.DB {
	return r.db
}

func (r *Registry) Config() *config.Config {
	return r.config
}

func (r *Registry) Logger() logger.Logger {
	if r.logger == nil {
		r.logger = logger.NewLogger(r)
	}

	return r.logger
}

func (r *Registry) NewError(c errors.ErrorCode, m string) *errors.AppError {
	return &errors.AppError{
		Code:    c,
		Message: m,
	}
}

func (r *Registry) ErrorHandler() *errors.Handler {
	if r.errorHandler == nil {
		r.errorHandler = errors.NewErrorHandler(r)
	}

	return r.errorHandler
}

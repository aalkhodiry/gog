package post

import (
	"context"

	"github.com/PROJECT_NAME/internal/config"
	"github.com/PROJECT_NAME/internal/domains/user"
	"github.com/PROJECT_NAME/internal/errors"
	"github.com/PROJECT_NAME/internal/logger"
	"github.com/google/uuid"
)

type (
	ServiceProvider interface {
		PostService() *Service
	}

	serviceDependencies interface {
		logger.LoggerProvider
		config.ConfigProvider
		errors.ErrorProvider
		RepositoryProvider
		user.ServiceProvider
	}

	Service struct {
		d serviceDependencies
	}
)

func NewService(d serviceDependencies) *Service {
	return &Service{d: d}
}

func (s *Service) CreatePost(ctx context.Context, dto *CreatePostDTO) error {
	user, err := s.d.UserService().GetUserByID(ctx, string(dto.AuthorID))
	if err != nil {
		return err
	}

	if err := s.d.PostRepository().createPost(ctx, &Post{
		Title:    dto.Title,
		Content:  dto.Content,
		AuthorID: user.ID,
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetPostByID(ctx context.Context, id string) (*Post, error) {
	post := &Post{ID: uuid.MustParse(id)}

	if err := s.d.PostRepository().getPostByID(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

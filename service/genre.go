package service

import (
	"context"
	"int-gateway/client"
	"int-gateway/response"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type genre struct {
	logger *logrus.Logger
}

type GenreServicer interface {
	CreateGenre(ctx context.Context, name string, description string) (*response.Genre, error)
	GetGenre(ctx context.Context, ID string) (*response.Genre, error)
	GetGenreByName(ctx context.Context, name string) (*response.Genre, error)
	UpdateGenre(ctx context.Context, ID string, name string, description string) (*response.Genre, error)
	ListGenres(ctx context.Context) (response.Genres, error)
}

func InitiateGenreService(logger *logrus.Logger) GenreServicer {
	return &genre{logger: logger}
}

func (g *genre) CreateGenre(ctx context.Context, name string, description string) (*response.Genre, error) {
	c := client.NewGenreClient()
	resp, err := c.CreateGenre(ctx, name, description)
	if err != nil {
		g.logger.Error("Error while creating genre, ", err)
		return nil, errors.Wrap(err, "Error while creating genre!")
	}
	return resp.ToResponse(), nil
}

func (g *genre) GetGenre(ctx context.Context, ID string) (*response.Genre, error) {
	c := client.NewGenreClient()
	resp, err := c.GetGenre(ctx, ID)
	if err != nil {
		g.logger.Error("Error while getting genre by id, ", err)
		return nil, errors.Wrap(err, "Error while getting genre by id!")
	}
	return resp.ToResponse(), nil
}

func (g *genre) GetGenreByName(ctx context.Context, name string) (*response.Genre, error) {
	c := client.NewGenreClient()
	resp, err := c.GetGenreByName(ctx, name)
	if err != nil {
		g.logger.Error("Error while getting genre by name, ", err)
		return nil, errors.Wrap(err, "Error while getting genre by name!")
	}
	return resp.ToResponse(), nil
}

func (g *genre) UpdateGenre(ctx context.Context, ID string, name string, description string) (*response.Genre, error) {
	c := client.NewGenreClient()
	resp, err := c.UpdateGenre(ctx, ID, name, description)
	if err != nil {
		g.logger.Error("Error while updating genre, ", err)
		return nil, errors.Wrap(err, "Error while updating genre!")
	}
	return resp.ToResponse(), nil
}

func (g *genre) ListGenres(ctx context.Context) (response.Genres, error) {
	c := client.NewGenreClient()
	resp, err := c.ListGenres(ctx)
	if err != nil {
		g.logger.Error("Error while listing genres, ", err)
		return nil, errors.Wrap(err, "Error while listing genres!")
	}
	return resp.ToResponse(), nil
}

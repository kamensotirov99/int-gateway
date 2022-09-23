
package client

import (
	"context"
	pb "int-gateway/_proto"
	"int-gateway/models"
	"time"

	"github.com/pkg/errors"
)

type genre struct{}

func NewGenreClient() GenreClienter {
	return &genre{}
}

type GenreClienter interface {
	CreateGenre(ctx context.Context, name string, description string) (*models.Genre, error)
	GetGenre(ctx context.Context, ID string) (*models.Genre, error)
	GetGenreByName(ctx context.Context, name string) (*models.Genre, error)
	UpdateGenre(ctx context.Context, ID string, name string, description string) (*models.Genre, error)
	ListGenres(ctx context.Context) (models.Genres, error)
}

func (g *genre) CreateGenre(ctx context.Context, name string, description string) (*models.Genre, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewGenreSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.CreateGenreRequest{
		Name: name,
		Description: description,
	}
	resp, err := c.CreateGenre(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Create genre request")
	}

	return resp.ToModel().(*models.Genre), nil
}

func (g *genre) GetGenre(ctx context.Context, ID string) (*models.Genre, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewGenreSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.GetGenre(ctx, &pb.GetByIDRequest{Id: ID})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Get genre by id request")
	}

	return resp.ToModel().(*models.Genre), nil
}

func (g *genre) GetGenreByName(ctx context.Context, name string) (*models.Genre, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewGenreSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.GetGenreByName(ctx, &pb.GetByNameRequest{Name: name})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Get genre by name request")
	}

	return resp.ToModel().(*models.Genre), nil
}

func (g *genre) UpdateGenre(ctx context.Context, ID string, name string, description string) (*models.Genre, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewGenreSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.Genre{
		Id: ID,
		Name: name,
		Description: description,
	}
	resp, err := c.UpdateGenre(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Update genre request")
	}

	return resp.ToModel().(*models.Genre), nil
}

func (g *genre) ListGenres(ctx context.Context) (models.Genres, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewGenreSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.ListGenres(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during List genres request")
	}
	return resp.ToModel(), nil
}

func toShortGenresPb(genresModel models.ShortGenres) *pb.ShortGenres {
	genres := pb.ShortGenres{}
	for _, genre := range genresModel {
		genres.Genres = append(genres.Genres, &pb.ShortGenre{
			Id: genre.ID,
			Name: genre.Name,
		})
	}
	return &genres
}
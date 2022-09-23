package service

import (
	"context"
	"int-gateway/client"
	"int-gateway/models"
	"int-gateway/response"
	"mime/multipart"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type show struct {
	logger       *logrus.Logger
	client       client.ShowClienter
	uploadSeriesClient client.UploadClienter
	uploadMovieClient client.UploadClienter
}

type ShowServicer interface {
	CreateShow(ctx context.Context, title string, sType string, postersPath []string, releaseDate time.Time, endDate time.Time, rating float64, length *models.ShowLength, trailerURL string, genres models.ShortGenres, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, starring models.ShortCelebrities, description string, seasons models.ShortSeasons) (*response.Show, error)
	GetShow(ctx context.Context, ID string) (*response.Show, error)
	UpdateShow(ctx context.Context, ID string, title string, sType string, postersPath []string, releaseDate time.Time, endDate time.Time, rating float64, length *models.ShowLength, trailerURL string, genres models.ShortGenres, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, starring models.ShortCelebrities, description string, seasons models.ShortSeasons) (*response.Show, error)
	ListShows(ctx context.Context) (response.Shows, error)
	UploadSeriesPosters(ctx context.Context, seriesID string, images []*multipart.FileHeader) (*response.Show, error)
	DeleteSeriesPoster(ctx context.Context, seriesID string, image string) error

	UploadMoviePosters(ctx context.Context, movieID string, images []*multipart.FileHeader) (*response.Show, error)
	DeleteMoviePoster(ctx context.Context, movieID string, image string) error
}

func InitiateShowService(logger *logrus.Logger, client client.ShowClienter, uploadSeriesClient client.UploadClienter,uploadMovieClient client.UploadClienter) ShowServicer {
	return &show{logger: logger,
		client:       client,
		uploadSeriesClient: uploadSeriesClient,
		uploadMovieClient:uploadMovieClient,

	}
}

func (s *show) CreateShow(ctx context.Context, title string, sType string, postersPath []string, releaseDate time.Time, endDate time.Time, rating float64, length *models.ShowLength, trailerURL string, genres models.ShortGenres, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, starring models.ShortCelebrities, description string, seasons models.ShortSeasons) (*response.Show, error) {
	resp, err := s.client.CreateShow(ctx, title, sType, postersPath, releaseDate, endDate, rating, length, trailerURL, genres, directedBy, producedBy, writtenBy, starring, description, seasons)
	if err != nil {
		s.logger.Error("Error while creating show, ", err)
		return nil, errors.Wrap(err, "Error while creating show!")
	}
	return resp.ToResponse(), nil
}

func (s *show) GetShow(ctx context.Context, ID string) (*response.Show, error) {
	resp, err := s.client.GetShow(ctx, ID)
	if err != nil {
		s.logger.Error("Error while getting show, ", err)
		return nil, errors.Wrap(err, "Error while getting show!")
	}
	return resp.ToResponse(), nil
}

func (s *show) UpdateShow(ctx context.Context, ID string, title string, sType string, postersPath []string, releaseDate time.Time, endDate time.Time, rating float64, length *models.ShowLength, trailerURL string, genres models.ShortGenres, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, starring models.ShortCelebrities, description string, seasons models.ShortSeasons) (*response.Show, error) {
	resp, err := s.client.UpdateShow(ctx, ID, title, sType, postersPath, releaseDate, endDate, rating, length, trailerURL, genres, directedBy, producedBy, writtenBy, starring, description, seasons)
	if err != nil {
		s.logger.Error("Error while updating show, ", err)
		return nil, errors.Wrap(err, "Error while updating show!")
	}
	return resp.ToResponse(), nil
}

func (s *show) ListShows(ctx context.Context) (response.Shows, error) {
	resp, err := s.client.ListShows(ctx)
	if err != nil {
		s.logger.Error("Error while listing shows, ", err)
		return nil, errors.Wrap(err, "Error while listing shows!")
	}
	return resp.ToResponse(), nil
}

func (s *show) UploadSeriesPosters(ctx context.Context, ID string, images []*multipart.FileHeader) (*response.Show, error) {
	response, err := s.uploadSeriesClient.UploadPostersFS(ctx, images, ID)
	if err != nil {
		s.logger.Error("Error while uploading series posters in file server, ", err)
		return nil, errors.Wrap(err, "Error while uploading series posters in file server!")
	}

	resp, err := s.client.UploadSeriesPostersService(ctx, ID, response.PostersPath)
	if err != nil {
		s.logger.Error("Error while uploading series posters in database, ", err)
		return nil, errors.Wrap(err, "Error while uploading series posters in database!")
	}
	return resp.ToResponse(), nil
}

func (s *show) DeleteSeriesPoster(ctx context.Context, ID string, image string) error {
	err := s.client.DeleteSeriesPosterFS(ctx, ID, image)
	if err != nil {
		s.logger.Error("Error while deleting season poster in file server, ", err)
		return errors.Wrap(err, "Error while deleting season poster in file server")
	}

	err = s.client.DeleteSeriesPosterService(ctx, ID, image)
	if err != nil {
		s.logger.Error("Error while deleting series poster in database, ", err)
		return errors.Wrap(err, "Error while deleting series poster in database!")
	}
	return nil
}

func(s *show)UploadMoviePosters(ctx context.Context, movieID string, images []*multipart.FileHeader) (*response.Show, error){
	response, err := s.uploadMovieClient.UploadPostersFS(ctx, images, movieID)
	if err != nil {
		s.logger.Error("Error while uploading movie posters in file server, ", err)
		return nil, errors.Wrap(err, "Error while uploading movie posters in file server!")
	}

	resp, err := s.client.UploadMoviePostersService(ctx, movieID, response.PostersPath)
	if err != nil {
		s.logger.Error("Error while uploading movie posters in database, ", err)
		return nil, errors.Wrap(err, "Error while uploading movie posters in database!")
	}
	return resp.ToResponse(), nil
}

func (s *show) DeleteMoviePoster(ctx context.Context, ID string, image string) error {
	err := s.client.DeleteMoviePosterFS(ctx, ID, image)
	if err != nil {
		s.logger.Error("Error while deleting movie poster in file server, ", err)
		return errors.Wrap(err, "Error while deleting movie poster in file server")
	}

	err = s.client.DeleteMoviePosterService(ctx, ID, image)
	if err != nil {
		s.logger.Error("Error while deleting movie poster in database, ", err)
		return errors.Wrap(err, "Error while deleting movie poster in database!")
	}
	return nil
}
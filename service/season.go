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

type season struct {
	logger       *logrus.Logger
	client       client.SeasonClienter
	uploadClient client.UploadClienter
}

type SeasonServicer interface {
	CreateSeason(ctx context.Context, showID string, title string, trailerURL string, postersPath []string, releaseDate time.Time, rating float64, resume string, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, episodes models.ShortEpisodes) (*response.Season, error)
	GetSeason(ctx context.Context, ID string) (*response.Season, error)
	UpdateSeason(ctx context.Context, ID string, showID string, title string, trailerURL string, postersPath []string, releaseDate time.Time, rating float64, resume string, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, episodes models.ShortEpisodes) (*response.Season, error)
	UploadSeasonPosters(ctx context.Context, seriesID string, seasonID string, images []*multipart.FileHeader) (*response.Season, error)
	DeleteSeasonPoster(ctx context.Context, seriesID string, seasonID string, image string) error
	ListShowSeasons(ctx context.Context, ShowID string) (response.Seasons, error)
	ListSeasonsCollection(ctx context.Context) (response.Seasons, error)
}

func InitiateSeasonService(logger *logrus.Logger, client client.SeasonClienter, uploadClient client.UploadClienter) SeasonServicer {
	return &season{
		logger:       logger,
		client:       client,
		uploadClient: uploadClient,
	}
}

func (s *season) CreateSeason(ctx context.Context, showID string, title string, trailerURL string, postersPath []string, releaseDate time.Time, rating float64, resume string, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, episodes models.ShortEpisodes) (*response.Season, error) {
	resp, err := s.client.CreateSeason(ctx, showID, title, trailerURL, postersPath, releaseDate, rating, resume, directedBy, producedBy, writtenBy, episodes)
	if err != nil {
		s.logger.Error("Error while creating season, ", err)
		return nil, errors.Wrap(err, "Error while creating season!")
	}
	return resp.ToResponse(), nil
}

func (s *season) GetSeason(ctx context.Context, ID string) (*response.Season, error) {
	resp, err := s.client.GetSeason(ctx, ID)
	if err != nil {
		s.logger.Error("Error while getting season, ", err)
		return nil, errors.Wrap(err, "Error while getting season!")
	}
	return resp.ToResponse(), nil
}

func (s *season) UpdateSeason(ctx context.Context, ID string, showID string, title string, trailerURL string, postersPath []string, releaseDate time.Time, rating float64, resume string, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, episodes models.ShortEpisodes) (*response.Season, error) {
	resp, err := s.client.UpdateSeason(ctx, ID, showID, title, trailerURL, postersPath, releaseDate, rating, resume, directedBy, producedBy, writtenBy, episodes)
	if err != nil {
		s.logger.Error("Error while updating season, ", err)
		return nil, errors.Wrap(err, "Error while updating season!")
	}
	return resp.ToResponse(), nil
}

func (s *season) UploadSeasonPosters(ctx context.Context, seriesID string, seasonID string, images []*multipart.FileHeader) (*response.Season, error) {
	response, err := s.uploadClient.UploadPostersFS(ctx, images, seriesID, seasonID)
	if err != nil {
		s.logger.Error("Error while uploading season posters in file server, ", err)
		return nil, errors.Wrap(err, "Error while uploading season posters in file server!")
	}

	resp, err := s.client.UploadSeasonPostersService(ctx, seasonID, response.PostersPath)
	if err != nil {
		s.logger.Error("Error while uploading season posters in database, ", err)
		return nil, errors.Wrap(err, "Error while uploading season posters in database!")
	}
	return resp.ToResponse(), nil
}

func (s *season) DeleteSeasonPoster(ctx context.Context, seriesID string, seasonID string, image string) error {
	err := s.client.DeleteSeasonPosterFS(ctx, seriesID, seasonID, image)
	if err != nil {
		s.logger.Error("Error while deleting season poster in file server, ", err)
		return errors.Wrap(err, "Error while deleting season poster in file server")
	}

	err = s.client.DeleteSeasonPosterService(ctx, seriesID, seasonID, image)
	if err != nil {
		s.logger.Error("Error while deleting season poster in database, ", err)
		return errors.Wrap(err, "Error while deleting season poster in database!")
	}
	return nil
}

func (s *season) ListShowSeasons(ctx context.Context, ShowID string) (response.Seasons, error) {
	resp, err := s.client.ListShowSeasons(ctx, ShowID)
	if err != nil {
		s.logger.Error("Error while listing show seasons, ", err)
		return nil, errors.Wrap(err, "Error while listing show seasons!")
	}
	return resp.ToResponse(), nil
}

func (s *season) ListSeasonsCollection(ctx context.Context) (response.Seasons, error) {
	resp, err := s.client.ListSeasonsCollection(ctx)
	if err != nil {
		s.logger.Error("Error while listing seasons, ", err)
		return nil, errors.Wrap(err, "Error while listing seasons!")
	}
	return resp.ToResponse(), nil
}

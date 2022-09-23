package service

import (
	"context"
	"int-gateway/client"
	"int-gateway/models"
	"int-gateway/response"
	"mime/multipart"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type episode struct {
	logger       *logrus.Logger
	client       client.EpisodeClienter
	uploadClient client.UploadClienter
}

type EpisodeServicer interface {
	CreateEpisode(ctx context.Context, seasonID string, title string, postersPath []string, trailerURL string, length *models.ShowLength, rating float64, resume string, writtenBy models.FilmCrews, producedBy models.FilmCrews, directedBy models.FilmCrews, starring models.ShortCelebrities) (*response.Episode, error)
	GetEpisode(ctx context.Context, ID string) (*response.Episode, error)
	UpdateEpisode(ctx context.Context, ID string, seasonID string, title string, postersPath []string, trailerURL string, length *models.ShowLength, rating float64, resume string, writtenBy models.FilmCrews, producedBy models.FilmCrews, directedBy models.FilmCrews, starring models.ShortCelebrities) (*response.Episode, error)
	UploadEpisodePosters(ctx context.Context, seriesID string, seasonID string, episodeID string, images []*multipart.FileHeader) (*response.Episode, error)
	DeleteEpisodePoster(ctx context.Context, seriesID string, seasonID string, episodeID string, image string) error
	ListSeasonEpisodes(ctx context.Context, seasonID string) (response.Episodes, error)
	ListCollectionEpisodes(ctx context.Context) (response.Episodes, error)
}

func InitiateEpisodeService(logger *logrus.Logger, client client.EpisodeClienter, uploadClient client.UploadClienter) EpisodeServicer {
	return &episode{
		logger:       logger,
		client:       client,
		uploadClient: uploadClient,
	}
}

func (e *episode) CreateEpisode(ctx context.Context, seasonID string, title string, postersPath []string, trailerURL string, length *models.ShowLength, rating float64, resume string, writtenBy models.FilmCrews, producedBy models.FilmCrews, directedBy models.FilmCrews, starring models.ShortCelebrities) (*response.Episode, error) {
	resp, err := e.client.CreateEpisode(ctx, seasonID, title, postersPath, trailerURL, length, rating, resume, writtenBy, producedBy, directedBy, starring)
	if err != nil {
		e.logger.Error("Error while creating episode, ", err)
		return nil, errors.Wrap(err, "Error while creating episode!")
	}
	return resp.ToResponse(), nil
}

func (e *episode) GetEpisode(ctx context.Context, ID string) (*response.Episode, error) {
	resp, err := e.client.GetEpisode(ctx, ID)
	if err != nil {
		e.logger.Error("Error while getting episode, ", err)
		return nil, errors.Wrap(err, "Error while getting episode!")
	}
	return resp.ToResponse(), nil
}

func (e *episode) UpdateEpisode(ctx context.Context, ID string, seasonID string, title string, postersPath []string, trailerURL string, length *models.ShowLength, rating float64, resume string, writtenBy models.FilmCrews, producedBy models.FilmCrews, directedBy models.FilmCrews, starring models.ShortCelebrities) (*response.Episode, error) {
	resp, err := e.client.UpdateEpisode(ctx, ID, seasonID, title, postersPath, trailerURL, length, rating, resume, writtenBy, producedBy, directedBy, starring)
	if err != nil {
		e.logger.Error("Error while updating episode, ", err)
		return nil, errors.Wrap(err, "Error while updating episode!")
	}
	return resp.ToResponse(), nil
}

func (e *episode) UploadEpisodePosters(ctx context.Context, seriesID string, seasonID string, episodeID string, images []*multipart.FileHeader) (*response.Episode, error) {
	response, err := e.uploadClient.UploadPostersFS(ctx, images, seriesID, seasonID, episodeID)
	if err != nil {
		e.logger.Error("Error while uploading episode posters in file server, ", err)
		return nil, errors.Wrap(err, "Error while uploading episode posters in file server!")
	}

	resp, err := e.client.UploadEpisodePostersService(ctx, episodeID, response.PostersPath)
	if err != nil {
		e.logger.Error("Error while uploading episode posters in database, ", err)
		return nil, errors.Wrap(err, "Error while uploading episode posters in database!")
	}
	return resp.ToResponse(), nil
}

func (e *episode) DeleteEpisodePoster(ctx context.Context, seriesID string, seasonID string, episodeID string, image string) error {
	err := e.client.DeleteEpisodePosterFS(ctx, seriesID, seasonID, episodeID, image)
	if err != nil {
		e.logger.Error("Error while deleting episode poster in file server, ", err)
		return errors.Wrap(err, "Error while deleting episode poster in file server")
	}

	err = e.client.DeleteEpisodePosterService(ctx, seriesID, seasonID, episodeID, image)
	if err != nil {
		e.logger.Error("Error while deleting episode poster in database, ", err)
		return errors.Wrap(err, "Error while deleting episode poster in database!")
	}
	return nil
}

func (e *episode) ListSeasonEpisodes(ctx context.Context, seasonID string) (response.Episodes, error) {
	resp, err := e.client.ListSeasonEpisodes(ctx, seasonID)
	if err != nil {
		e.logger.Error("Error while listing season episodes, ", err)
		return nil, errors.Wrap(err, "Error while listing season episodes!")
	}
	return resp.ToResponse(), nil
}

func (e *episode) ListCollectionEpisodes(ctx context.Context) (response.Episodes, error) {
	resp, err := e.client.ListCollectionEpisodes(ctx)
	if err != nil {
		e.logger.Error("Error while listing episodes, ", err)
		return nil, errors.Wrap(err, "Error while listing episodes!")
	}
	return resp.ToResponse(), nil
}

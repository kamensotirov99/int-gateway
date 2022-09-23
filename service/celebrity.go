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

type celebrity struct {
	logger       *logrus.Logger
	client       client.CelebrityClienter
	uploadClient client.UploadClienter
}

type CelebrityServicer interface {
	CreateCelebrity(ctx context.Context, name string, occupation []string, postersPath []string, dateOfBirth time.Time, dateOfDeath time.Time, placeOfBirth string, gender models.Gender, bio string) (*response.Celebrity, error)
	GetCelebrity(ctx context.Context, ID string) (*response.Celebrity, error)
	UpdateCelebrity(ctx context.Context, ID string, name string, occupation []string, postersPath []string, dateOfBirth time.Time, dateOfDeath time.Time, placeOfBirth string, gender models.Gender, bio string) (*response.Celebrity, error)
	UploadCelebrityPosters(ctx context.Context, ID string, images []*multipart.FileHeader) (*response.Celebrity, error)
	DeleteCelebrityPoster(ctx context.Context, ID string, image string) error
	ListCelebrities(ctx context.Context) (response.Celebrities, error)
}

func InitiateCelebrityService(logger *logrus.Logger, client client.CelebrityClienter, uploadClient client.UploadClienter) CelebrityServicer {
	return &celebrity{
		logger:       logger,
		client:       client,
		uploadClient: uploadClient,
	}
}

func (c *celebrity) CreateCelebrity(ctx context.Context, name string, occupation []string, postersPath []string, dateOfBirth time.Time, dateOfDeath time.Time, placeOfBirth string, gender models.Gender, bio string) (*response.Celebrity, error) {
	resp, err := c.client.CreateCelebrity(ctx, name, occupation, postersPath, dateOfBirth, dateOfDeath, placeOfBirth, gender, bio)
	if err != nil {
		c.logger.Error("Error while creating celebrity, ", err)
		return nil, errors.Wrap(err, "Error while creating celebrity!")
	}
	return resp.ToResponse(), nil
}

func (c *celebrity) GetCelebrity(ctx context.Context, ID string) (*response.Celebrity, error) {
	resp, err := c.client.GetCelebrity(ctx, ID)
	if err != nil {
		c.logger.Error("Error while getting celebrity, ", err)
		return nil, errors.Wrap(err, "Error while getting celebrity!")
	}
	return resp.ToResponse(), nil
}

func (c *celebrity) UpdateCelebrity(ctx context.Context, ID string, name string, occupation []string, postersPath []string, dateOfBirth time.Time, dateOfDeath time.Time, placeOfBirth string, gender models.Gender, bio string) (*response.Celebrity, error) {
	resp, err := c.client.UpdateCelebrity(ctx, ID, name, occupation, postersPath, dateOfBirth, dateOfDeath, placeOfBirth, gender, bio)
	if err != nil {
		c.logger.Error("Error while updating celebrity, ", err)
		return nil, errors.Wrap(err, "Error while updating celebrity!")
	}
	return resp.ToResponse(), nil
}

func (c *celebrity) UploadCelebrityPosters(ctx context.Context, ID string, images []*multipart.FileHeader) (*response.Celebrity, error) {
	response, err := c.uploadClient.UploadPostersFS(ctx, images, ID)
	if err != nil {
		c.logger.Error("Error while uploading celebrity posters in file server, ", err)
		return nil, errors.Wrap(err, "Error while uploading celebrity posters in file server!")
	}

	resp, err := c.client.UploadCelebrityPostersService(ctx, ID, response.PostersPath)
	if err != nil {
		c.logger.Error("Error while uploading celebrity posters in database, ", err)
		return nil, errors.Wrap(err, "Error while uploading celebrity posters in database!")
	}
	return resp.ToResponse(), nil
}

func (c *celebrity) DeleteCelebrityPoster(ctx context.Context, ID string, image string) error {
	err := c.client.DeleteCelebrityPosterFS(ctx, ID, image)
	if err != nil {
		c.logger.Error("Error while deleting celebrity poster in file server, ", err)
		return errors.Wrap(err, "Error while deleting season poster in file server")
	}

	err = c.client.DeleteCelebrityPosterService(ctx, ID, image)
	if err != nil {
		c.logger.Error("Error while deleting celebrity poster in database, ", err)
		return errors.Wrap(err, "Error while deleting celebrity poster in database!")
	}
	return nil
}

func (c *celebrity) ListCelebrities(ctx context.Context) (response.Celebrities, error) {
	resp, err := c.client.ListCelebrities(ctx)
	if err != nil {
		c.logger.Error("Error while listing celebrities, ", err)
		return nil, errors.Wrap(err, "Error while listing celebrities!")
	}
	return resp.ToResponse(), nil
}

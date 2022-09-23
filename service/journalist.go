package service

import (
	"context"
	"int-gateway/client"
	"int-gateway/response"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type journalist struct {
	logger *logrus.Logger
}

type JournalistServicer interface {
	CreateJournalist(ctx context.Context, name string) (*response.Journalist, error)
	GetJournalist(ctx context.Context, ID string) (*response.Journalist, error)
	GetJournalistByName(ctx context.Context, name string) (*response.Journalist, error)
	ListJournalists(ctx context.Context) (response.Journalists, error)
	UpdateJournalist(ctx context.Context, ID string, name string) (*response.Journalist, error)
}

func InitiateJournalistService(logger *logrus.Logger) JournalistServicer {
	return &journalist{logger: logger}
}

func (j *journalist) CreateJournalist(ctx context.Context, name string) (*response.Journalist, error) {
	c := client.NewJournalistClient()
	resp, err := c.CreateJournalist(ctx, name)
	if err != nil {
		j.logger.Error("Error while creating journalist, ", err)
		return nil, errors.Wrap(err, "Error while creating the journalist")
	}

	return resp.ToResponse(), nil
}

func (j *journalist) GetJournalist(ctx context.Context, ID string) (*response.Journalist, error) {
	c := client.NewJournalistClient()
	resp, err := c.GetJournalist(ctx, ID)
	if err != nil {
		j.logger.Error("Error while getting journalist, ", err)
		return nil, errors.Wrap(err, "Error while getting journalist")
	}

	return resp.ToResponse(), nil
}

func (j *journalist) GetJournalistByName(ctx context.Context, name string) (*response.Journalist, error) {
	c := client.NewJournalistClient()
	resp, err := c.GetJournalistByName(ctx, name)
	if err != nil {
		j.logger.Error("Error while getting journalist by name, ", err)
		return nil, errors.Wrap(err, "Error while getting journalist by name")
	}

	return resp.ToResponse(), nil
}

func (j *journalist) ListJournalists(ctx context.Context) (response.Journalists, error) {
	c := client.NewJournalistClient()
	resp, err := c.ListJournalists(ctx)
	if err != nil {
		j.logger.Error("Error while getting all journalists, ", err)
		return nil, errors.Wrap(err, "Error while getting all journalists")
	}

	return resp.ToResponse(), nil
}

func (j *journalist) UpdateJournalist(ctx context.Context, ID string, name string) (*response.Journalist, error) {
	c := client.NewJournalistClient()
	resp, err := c.UpdateJournalist(ctx, ID, name)
	if err != nil {
		j.logger.Error("Error while updating journalist, ", err)
		return nil, errors.Wrap(err, "Error while updating the journalist")
	}

	return resp.ToResponse(), nil
}

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

type article struct {
	logger       *logrus.Logger
	client       client.ArticleClienter
	uploadClient client.UploadClienter
}

type ArticleServicer interface {
	CreateArticle(ctx context.Context, title string, releaseDate time.Time, postersPath []string, description string, journalistName string) (*response.Article, error)
	GetArticle(ctx context.Context, ID string) (*response.Article, error)
	UpdateArticle(ctx context.Context, ID string, title string, releaseDate time.Time, postersPath []string, description string, journalistModel *models.Journalist) (*response.Article, error)
	ListArticles(ctx context.Context,elementCount int) (response.Articles, error)
	ListArticlesByJournalist(ctx context.Context, journalistID string) (response.Articles, error)
	UploadArticlePosters(ctx context.Context, articleID string, images []*multipart.FileHeader) (*response.Article, error)
	DeleteArticlePoster(ctx context.Context, articleID string, image string) error
}

func InitiateArticleService(logger *logrus.Logger, client client.ArticleClienter, uploadClient client.UploadClienter) ArticleServicer {
	return &article{logger: logger,
		client:       client,
		uploadClient: uploadClient}
}

func (a *article) CreateArticle(ctx context.Context, title string, releaseDate time.Time, postersPath []string, description string, journalistName string) (*response.Article, error) {
	resp, err := a.client.CreateArticle(ctx, title, releaseDate, postersPath, description, journalistName)
	if err != nil {
		a.logger.Error("Error while creating article ", err)
		return nil, errors.Wrap(err, "Error while creating article!")
	}
	return resp.ToResponse(), nil
}

func (a *article) GetArticle(ctx context.Context, ID string) (*response.Article, error) {
	c := client.NewArticleClient()
	resp, err := c.GetArticle(ctx, ID)
	if err != nil {
		a.logger.Error("Error while getting article ", err)
		return nil, errors.Wrap(err, "Error while getting article!")
	}
	return resp.ToResponse(), nil
}

func (a *article) UpdateArticle(ctx context.Context, ID string, title string, releaseDate time.Time, postersPath []string, description string, journalistModel *models.Journalist) (*response.Article, error) {
	resp, err := a.client.UpdateArticle(ctx, ID, title, releaseDate, postersPath, description, journalistModel)
	if err != nil {
		a.logger.Error("Error while updating article ", err)
		return nil, errors.Wrap(err, "Error while updating article!")
	}
	return resp.ToResponse(), nil
}

func (a *article) ListArticles(ctx context.Context,elementCount int) (response.Articles, error) {
	resp, err := a.client.ListArticles(ctx,elementCount)
	if err != nil {
		a.logger.Error("Error while listing articles ", err)
		return nil, errors.Wrap(err, "Error while listing articles!")
	}
	return resp.ToResponse(), nil
}

func (a *article) ListArticlesByJournalist(ctx context.Context, journalistID string) (response.Articles, error) {
	resp, err := a.client.ListArticlesByJournalist(ctx, journalistID)
	if err != nil {
		a.logger.Error("Error while listing articles by journalist ", err)
		return nil, errors.Wrap(err, "Error while listing articles by journalist!")
	}
	return resp.ToResponse(), nil
}

func (a *article) UploadArticlePosters(ctx context.Context, articleID string, images []*multipart.FileHeader) (*response.Article, error) {
	response, err := a.uploadClient.UploadPostersFS(ctx, images, articleID)
	if err != nil {
		a.logger.Error("Error while uploading article posters in file server, ", err)
		return nil, errors.Wrap(err, "Error while uploading article posters in file server!")
	}

	resp, err := a.client.UploadArticlePostersService(ctx, articleID, response.PostersPath)
	if err != nil {
		a.logger.Error("Error while uploading article posters in database, ", err)
		return nil, errors.Wrap(err, "Error while uploading article posters in database!")
	}
	return resp.ToResponse(), nil
}

func (a *article) DeleteArticlePoster(ctx context.Context, articleID string, image string) error {
	err := a.client.DeleteArticlePosterFS(ctx, articleID, image)
	if err != nil {
		a.logger.Error("Error while deleting article poster in file server, ", err)
		return errors.Wrap(err, "Error while deleting article poster in file server")
	}

	err = a.client.DeleteArticlePosterService(ctx, articleID, image)
	if err != nil {
		a.logger.Error("Error while deleting article poster in database, ", err)
		return errors.Wrap(err, "Error while deleting article poster in database!")
	}
	return nil
}

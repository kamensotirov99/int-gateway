package client

import (
	"context"
	pb "int-gateway/_proto"
	"int-gateway/models"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type article struct{}

func NewArticleClient() ArticleClienter {
	return &article{}
}

type uploadArticle struct {
	client pb.FileServerArticleSvcClient
}

func NewUploadArticleClient() Clienter {
	return &uploadArticle{}
}

func (a *uploadArticle) getClient(ctx context.Context, conn *grpc.ClientConn) Clienter {
	return &uploadArticle{
		client: pb.NewFileServerArticleSvcClient(conn),
	}
}

func (a *uploadArticle) getStream(ctx context.Context) (grpc.ClientStream, error) {
	stream, err := a.client.UploadArticlePosters(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error while uploading article images in file server!")
	}
	return stream, nil
}

func (a *uploadArticle) getRequestMessage(imageExtension string, IDs []string) proto.Message {
	return &pb.UploadArticlePostersFSRequest{
		ArticleId:      IDs[0],
		ImageExtention: imageExtension,
	}
}

func (a *uploadArticle) getChunkData(buffer []byte, n int) proto.Message {
	return &pb.UploadArticlePostersFSRequest{
		ChunkData: buffer[:n],
	}
}

type ArticleClienter interface {
	CreateArticle(ctx context.Context, title string, releaseDate time.Time, postersPath []string, description string, journalistName string) (*models.Article, error)
	GetArticle(ctx context.Context, ID string) (*models.Article, error)
	UpdateArticle(ctx context.Context, ID string, title string, releaseDate time.Time, postersPath []string, description string, journalistModel *models.Journalist) (*models.Article, error)
	ListArticles(ctx context.Context,elementCount int) (models.Articles, error)
	ListArticlesByJournalist(ctx context.Context, journalistID string) (models.Articles, error)
	UploadArticlePostersService(ctx context.Context, articleID string, postersPath []string) (*models.Article, error)
	DeleteArticlePosterService(ctx context.Context, articleID string, image string) error
	DeleteArticlePosterFS(ctx context.Context, articleID string, image string) error
}

func (a *article) CreateArticle(ctx context.Context, title string, releaseDate time.Time, postersPath []string, description string, journalistName string) (*models.Article, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewArticleSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.CreateArticleRequest{
		Title:       title,
		ReleaseDate: timestamppb.New(releaseDate),
		PostersPath: postersPath,
		Description: description,
		Journalist:  &pb.CreateJournalistRequest{Name: journalistName},
	}
	resp, err := c.CreateArticle(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Create article request")
	}

	return resp.ToModel().(*models.Article), nil
}

func (a *article) GetArticle(ctx context.Context, ID string) (*models.Article, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewArticleSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.GetArticle(ctx, &pb.GetByIDRequest{Id: ID})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Get article request")
	}

	return resp.ToModel().(*models.Article), nil
}

func (a *article) UpdateArticle(ctx context.Context, ID string, title string, releaseDate time.Time, postersPath []string, description string, journalistModel *models.Journalist) (*models.Article, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewArticleSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.Article{
		Id:          ID,
		Title:       title,
		ReleaseDate: timestamppb.New(releaseDate),
		PostersPath: postersPath,
		Description: description,
		Journalist:  &pb.ShortJournalist{Id: journalistModel.ID},
	}
	resp, err := c.UpdateArticle(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Update article request")
	}

	return resp.ToModel().(*models.Article), nil
}

func (a *article) ListArticles(ctx context.Context,elementCount int) (models.Articles, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewArticleSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.ListArticles(ctx, &pb.ListArticlesRequest{ElementCount:int32(elementCount)})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during List articles request")
	}
	return resp.ToModel(), nil
}

func (a *article) ListArticlesByJournalist(ctx context.Context, journalistID string) (models.Articles, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewArticleSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.ListArticlesByJournalist(ctx, &pb.GetByIDRequest{Id: journalistID})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during List articles by journalist request")
	}
	return resp.ToModel(), nil
}

func (a *article) UploadArticlePostersService(ctx context.Context, articleID string, postersPath []string) (*models.Article, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewArticleSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := client.UploadArticlePosters(ctx, &pb.UploadArticlePostersServiceRequest{
		ArticleId:   articleID,
		PostersPath: postersPath,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Updating article posters request in database")
	}
	return resp.ToModel().(*models.Article), nil
}

func (a *article) DeleteArticlePosterService(ctx context.Context, articleID string, image string) error {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewArticleSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = client.DeleteArticlePoster(ctx, &pb.DeleteArticlePosterRequest{
		ArticleId: articleID,
		Image:     image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during delete article poster request in database")
	}
	return nil
}

func (a *article) DeleteArticlePosterFS(ctx context.Context, articleID string, image string) error {
	conn, err := dialGrpcFileServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc file server")
	}
	defer conn.Close()
	client := pb.NewFileServerArticleSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	_, err = client.DeleteArticlePoster(ctx, &pb.DeleteArticlePosterRequest{
		ArticleId: articleID,
		Image:     image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during delete article poster request in file server")
	}
	return nil
}

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

type show struct{}

func NewShowClient() ShowClienter {
	return &show{}
}

type uploadSeries struct {
	client pb.FileServerSeriesSvcClient
}

func NewUploadSeriesClient() Clienter {
	return &uploadSeries{}
}

func (s *uploadSeries) getClient(ctx context.Context, conn *grpc.ClientConn) Clienter {
	return &uploadSeries{
		client: pb.NewFileServerSeriesSvcClient(conn),
	}
}

func (s *uploadSeries) getStream(ctx context.Context) (grpc.ClientStream, error) {
	stream, err := s.client.UploadSeriesPosters(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error while uploading series images in file server!")
	}
	return stream, nil
}

func (s *uploadSeries) getRequestMessage(imageExtension string, IDs []string) proto.Message {
	return &pb.UploadSeriesPostersFSRequest{
		SeriesId:       IDs[0],
		ImageExtention: imageExtension,
	}
}

func (s *uploadSeries) getChunkData(buffer []byte, n int) proto.Message {
	return &pb.UploadSeriesPostersFSRequest{
		ChunkData: buffer[:n],
	}
}


type uploadMovie struct {
	client pb.FileServerMovieSvcClient
}

func NewUploadMovieClient() Clienter {
	return &uploadMovie{}
}

func (s *uploadMovie) getClient(ctx context.Context, conn *grpc.ClientConn) Clienter {
	return &uploadMovie{
		client: pb.NewFileServerMovieSvcClient(conn),
	}
}

func (s *uploadMovie) getStream(ctx context.Context) (grpc.ClientStream, error) {
	stream, err := s.client.UploadMoviePosters(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error while uploading movie images in file server!")
	}
	return stream, nil
}

func (s *uploadMovie) getRequestMessage(imageExtension string, IDs []string) proto.Message {
	return &pb.UploadMoviePostersFSRequest{
		MovieId:       IDs[0],
		ImageExtention: imageExtension,
	}
}

func (s *uploadMovie) getChunkData(buffer []byte, n int) proto.Message {
	return &pb.UploadMoviePostersFSRequest{
		ChunkData: buffer[:n],
	}
}


type ShowClienter interface {
	CreateShow(ctx context.Context, title string, sType string, postersPath []string, releaseDate time.Time, endDate time.Time, rating float64, length *models.ShowLength, trailerURL string, genres models.ShortGenres, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, starring models.ShortCelebrities, description string, seasons models.ShortSeasons) (*models.Show, error)
	GetShow(ctx context.Context, ID string) (*models.Show, error)
	UpdateShow(ctx context.Context, ID string, title string, sType string, postersPath []string, releaseDate time.Time, endDate time.Time, rating float64, length *models.ShowLength, trailerURL string, genres models.ShortGenres, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, starring models.ShortCelebrities, description string, seasons models.ShortSeasons) (*models.Show, error)
	ListShows(ctx context.Context) (models.Shows, error)
	UploadSeriesPostersService(ctx context.Context, seriesID string, postersPath []string) (*models.Show, error)
	DeleteSeriesPosterService(ctx context.Context, seriesID string, image string) error
	DeleteSeriesPosterFS(ctx context.Context, seriesID string, image string) error
	UploadMoviePostersService(ctx context.Context, movieID string, postersPath []string) (*models.Show, error)
	DeleteMoviePosterService(ctx context.Context, movieID string, image string) error
	DeleteMoviePosterFS(ctx context.Context, movieID string, image string) error
}

func (s *show) CreateShow(ctx context.Context, title string, sType string, postersPath []string, releaseDate time.Time, endDate time.Time, rating float64, length *models.ShowLength, trailerURL string, genres models.ShortGenres, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, starring models.ShortCelebrities, description string, seasons models.ShortSeasons) (*models.Show, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewShowSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	showLength := &pb.ShowLength{
		Hours:   int32(length.Hours),
		Minutes: int32(length.Minutes),
	}
	request := pb.CreateShowRequest{
		Title:       title,
		Type:        sType,
		PostersPath: postersPath,
		ReleaseDate: timestamppb.New(releaseDate),
		EndDate:     timestamppb.New(endDate),
		Rating:      rating,
		Length:      showLength,
		TrailerUrl:  trailerURL,
		Genres:      toShortGenresPb(genres),
		DirectedBy:  toFilmCrewPb(directedBy),
		ProducedBy:  toFilmCrewPb(producedBy),
		WrittenBy:   toFilmCrewPb(writtenBy),
		Starring:    toShortCelebritiesPb(starring),
		Description: description,
		Seasons:     toShortSeasonsPb(seasons),
	}
	resp, err := c.CreateShow(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Create show request")
	}

	return resp.ToModel().(*models.Show), nil
}

func (s *show) GetShow(ctx context.Context, ID string) (*models.Show, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewShowSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.GetShow(ctx, &pb.GetByIDRequest{Id: ID})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Get show request")
	}

	return resp.ToModel().(*models.Show), nil
}

func (s *show) UpdateShow(ctx context.Context, ID string, title string, sType string, postersPath []string, releaseDate time.Time, endDate time.Time, rating float64, length *models.ShowLength, trailerURL string, genres models.ShortGenres, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, starring models.ShortCelebrities, description string, seasons models.ShortSeasons) (*models.Show, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewShowSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	showLength := &pb.ShowLength{
		Hours:   int32(length.Hours),
		Minutes: int32(length.Minutes),
	}
	request := pb.Show{
		Id:          ID,
		Title:       title,
		Type:        sType,
		PostersPath: postersPath,
		ReleaseDate: timestamppb.New(releaseDate),
		EndDate:     timestamppb.New(endDate),
		Rating:      rating,
		Length:      showLength,
		TrailerUrl:  trailerURL,
		Genres:      toShortGenresPb(genres),
		DirectedBy:  toFilmCrewPb(directedBy),
		ProducedBy:  toFilmCrewPb(producedBy),
		WrittenBy:   toFilmCrewPb(writtenBy),
		Starring:    toShortCelebritiesPb(starring),
		Description: description,
		Seasons:     toShortSeasonsPb(seasons),
	}
	resp, err := c.UpdateShow(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Update show request")
	}

	return resp.ToModel().(*models.Show), nil
}

func (s *show) ListShows(ctx context.Context) (models.Shows, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewShowSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.ListShows(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during List shows request")
	}
	return resp.ToModel(), nil
}

func (s *show) UploadSeriesPostersService(ctx context.Context, seriesID string, postersPath []string) (*models.Show, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewShowSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := client.UploadSeriesPosters(ctx, &pb.UploadSeriesPostersServiceRequest{
		SeriesId:    seriesID,
		PostersPath: postersPath,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Updating series posters request in database")
	}
	return resp.ToModel().(*models.Show), nil
}

func (s *show) DeleteSeriesPosterService(ctx context.Context, seriesID string, image string) error {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewShowSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = client.DeleteSeriesPoster(ctx, &pb.DeleteSeriesPosterRequest{
		SeriesId: seriesID,
		Image:    image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Updating deleted series poster request in database")
	}
	return nil
}

func (s *show) DeleteSeriesPosterFS(ctx context.Context, seriesID string, image string) error {
	conn, err := dialGrpcFileServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc file server")
	}
	defer conn.Close()
	client := pb.NewFileServerSeriesSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	_, err = client.DeleteSeriesPoster(ctx, &pb.DeleteSeriesPosterRequest{
		SeriesId: seriesID,
		Image:    image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Deleting series poster request in file server")
	}
	return nil
}

func (s *show) UploadMoviePostersService(ctx context.Context, movieID string, postersPath []string) (*models.Show, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewShowSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := client.UploadMoviePosters(ctx, &pb.UploadMoviePostersServiceRequest{
		MovieId:    movieID,
		PostersPath: postersPath,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Updating movie posters request in database")
	}
	return resp.ToModel().(*models.Show), nil
}

func (s *show) DeleteMoviePosterService(ctx context.Context, movieID string, image string) error {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewShowSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = client.DeleteMoviePoster(ctx, &pb.DeleteMoviePosterRequest{
		MovieId: movieID,
		Image:    image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Updating deleted movie poster request in database")
	}
	return nil
}

func (s *show) DeleteMoviePosterFS(ctx context.Context, movieID string, image string) error {
	conn, err := dialGrpcFileServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc file server")
	}
	defer conn.Close()
	client := pb.NewFileServerMovieSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	_, err = client.DeleteMoviePoster(ctx, &pb.DeleteMoviePosterRequest{
		MovieId: movieID,
		Image:    image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Deleting movie poster request in file server")
	}
	return nil
}
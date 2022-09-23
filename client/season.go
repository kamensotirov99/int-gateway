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

type season struct{}

type uploadSeason struct {
	client pb.FileServerSeasonSvcClient
}

func NewUploadSeasonClient() Clienter {
	return &uploadSeason{}
}

func NewSeasonClient() SeasonClienter {
	return &season{}
}

func (s *uploadSeason) getClient(ctx context.Context, conn *grpc.ClientConn) Clienter {
	return &uploadSeason{
		client: pb.NewFileServerSeasonSvcClient(conn),
	}
}

func (s *uploadSeason) getStream(ctx context.Context) (grpc.ClientStream, error) {
	stream, err := s.client.UploadSeasonPosters(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error while uploading season images in file server!")
	}
	return stream, nil
}

func (s *uploadSeason) getRequestMessage(imageExtension string, IDs []string) proto.Message {
	return &pb.UploadSeasonPostersFSRequest{
		SeriesId:       IDs[0],
		SeasonId:       IDs[1],
		ImageExtention: imageExtension,
	}
}

func (s *uploadSeason) getChunkData(buffer []byte, n int) proto.Message {
	return &pb.UploadSeasonPostersFSRequest{
		ChunkData: buffer[:n],
	}
}

type SeasonClienter interface {
	CreateSeason(ctx context.Context, showID string, title string, trailerURL string, postersPath []string, releaseDate time.Time, rating float64, resume string, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, episodes models.ShortEpisodes) (*models.Season, error)
	GetSeason(ctx context.Context, ID string) (*models.Season, error)
	UpdateSeason(ctx context.Context, ID string, showID string, title string, trailerURL string, postersPath []string, releaseDate time.Time, rating float64, resume string, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, episodes models.ShortEpisodes) (*models.Season, error)
	UploadSeasonPostersService(ctx context.Context, seasonID string, postersPath []string) (*models.Season, error)
	DeleteSeasonPosterService(ctx context.Context, seriesID string, seasonID string, image string) error
	DeleteSeasonPosterFS(ctx context.Context, seriesID string, seasonID string, image string) error
	ListShowSeasons(ctx context.Context, ShowID string) (models.Seasons, error)
	ListSeasonsCollection(ctx context.Context) (models.Seasons, error)
}

func (s *season) CreateSeason(ctx context.Context, showID string, title string, trailerURL string, postersPath []string, releaseDate time.Time, rating float64, resume string, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, episodes models.ShortEpisodes) (*models.Season, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()

	c := pb.NewSeasonSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	request := pb.CreateSeasonRequest{
		Title:       title,
		TrailerUrl:  trailerURL,
		Resume:      resume,
		Rating:      rating,
		ReleaseDate: timestamppb.New(releaseDate),
		WrittenBy:   toFilmCrewPb(writtenBy),
		ProducedBy:  toFilmCrewPb(producedBy),
		DirectedBy:  toFilmCrewPb(directedBy),
		Episodes:    toShortEpisodesGrpc(episodes),
		PostersPath: postersPath,
		ShowId:      showID,
	}
	resp, err := c.CreateSeason(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Create request")
	}

	return resp.ToModel().(*models.Season), nil
}

func (s *season) GetSeason(ctx context.Context, ID string) (*models.Season, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewSeasonSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := c.GetSeason(ctx, &pb.GetByIDRequest{Id: ID})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Get season request")
	}

	return resp.ToModel().(*models.Season), nil
}

func (s *season) UpdateSeason(ctx context.Context, ID string, showID string, title string, trailerURL string, postersPath []string, releaseDate time.Time, rating float64, resume string, directedBy models.FilmCrews, producedBy models.FilmCrews, writtenBy models.FilmCrews, episodes models.ShortEpisodes) (*models.Season, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()

	c := pb.NewSeasonSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	request := pb.Season{
		Id:          ID,
		Title:       title,
		TrailerUrl:  trailerURL,
		Resume:      resume,
		Rating:      rating,
		ReleaseDate: timestamppb.New(releaseDate),
		WrittenBy:   toFilmCrewPb(writtenBy),
		ProducedBy:  toFilmCrewPb(producedBy),
		DirectedBy:  toFilmCrewPb(directedBy),
		Episodes:    toShortEpisodesGrpc(episodes),
		PostersPath: postersPath,
		ShowId:      showID,
	}
	resp, err := c.UpdateSeason(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during UpdateSeason request")
	}

	return resp.ToModel().(*models.Season), nil
}

func (s *season) UploadSeasonPostersService(ctx context.Context, seasonID string, postersPath []string) (*models.Season, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewSeasonSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := client.UploadSeasonPosters(ctx, &pb.UploadSeasonPostersServiceRequest{
		SeasonId:    seasonID,
		PostersPath: postersPath,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Updating season posters request in database")
	}
	return resp.ToModel().(*models.Season), nil
}

func (s *season) DeleteSeasonPosterService(ctx context.Context, seriesID string, seasonID string, image string) error {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewSeasonSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = client.DeleteSeasonPoster(ctx, &pb.DeleteSeasonPosterRequest{
		SeriesId: seriesID,
		SeasonId: seasonID,
		Image:    image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Updating deleted season poster request in database")
	}
	return nil
}

func (s *season) DeleteSeasonPosterFS(ctx context.Context, seriesID string, seasonID string, image string) error {
	conn, err := dialGrpcFileServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc file server")
	}
	defer conn.Close()
	client := pb.NewFileServerSeasonSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	_, err = client.DeleteSeasonPoster(ctx, &pb.DeleteSeasonPosterRequest{
		SeriesId: seriesID,
		SeasonId: seasonID,
		Image:    image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Deleting season poster request in file server")
	}
	return nil
}

func (s *season) ListShowSeasons(ctx context.Context, ShowID string) (models.Seasons, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewSeasonSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := c.ListShowSeasons(ctx, &pb.GetByIDRequest{Id: ShowID})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during ListShowSeasons request")
	}
	return resp.ToModel(), nil
}

func (s *season) ListSeasonsCollection(ctx context.Context) (models.Seasons, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewSeasonSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := c.ListSeasonsCollection(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during ListSeasonsCollection request")
	}
	return resp.ToModel(), nil
}

func toShortSeasonsPb(shortSeasons models.ShortSeasons) *pb.ShortSeasons {
	seasons := pb.ShortSeasons{}
	for _, season := range shortSeasons {
		seasons.Seasons = append(seasons.Seasons, &pb.ShortSeason{
			Id:          season.ID,
			Title:       season.Title,
			PostersPath: season.PostersPath,
			Rating:      season.Rating,
		})
	}
	return &seasons
}

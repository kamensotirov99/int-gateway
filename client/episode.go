package client

import (
	"context"
	pb "int-gateway/_proto"
	"int-gateway/models"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type episode struct{}

type uploadEpisode struct {
	client pb.FileServerEpisodeSvcClient
}

func NewEpisodeClient() EpisodeClienter {
	return &episode{}
}

func NewUploadEpisodeClient() Clienter {
	return &uploadEpisode{}
}

func (e *uploadEpisode) getClient(ctx context.Context, conn *grpc.ClientConn) Clienter {
	return &uploadEpisode{
		client: pb.NewFileServerEpisodeSvcClient(conn),
	}
}

func (e *uploadEpisode) getStream(ctx context.Context) (grpc.ClientStream, error) {
	stream, err := e.client.UploadEpisodePosters(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error while uploading episode images in file server!")
	}
	return stream, nil
}

func (e *uploadEpisode) getRequestMessage(imageExtension string, IDs []string) proto.Message {
	return &pb.UploadEpisodePostersFSRequest{
		SeriesId:       IDs[0],
		SeasonId:       IDs[1],
		EpisodeId:      IDs[2],
		ImageExtention: imageExtension,
	}
}

func (e *uploadEpisode) getChunkData(buffer []byte, n int) proto.Message {
	return &pb.UploadEpisodePostersFSRequest{
		ChunkData: buffer[:n],
	}
}

type EpisodeClienter interface {
	CreateEpisode(ctx context.Context, seasonID string, title string, postersPath []string, trailerURL string, length *models.ShowLength, rating float64, resume string, writtenBy models.FilmCrews, producedBy models.FilmCrews, directedBy models.FilmCrews, starring models.ShortCelebrities) (*models.Episode, error)
	GetEpisode(ctx context.Context, ID string) (*models.Episode, error)
	UpdateEpisode(ctx context.Context, ID string, seasonID string, title string, postersPath []string, trailerURL string, length *models.ShowLength, rating float64, resume string, writtenBy models.FilmCrews, producedBy models.FilmCrews, directedBy models.FilmCrews, starring models.ShortCelebrities) (*models.Episode, error)
	UploadEpisodePostersService(ctx context.Context, episodeID string, postersPath []string) (*models.Episode, error)
	DeleteEpisodePosterService(ctx context.Context, seriesID string, seasonID string, episodeID string, image string) error
	DeleteEpisodePosterFS(ctx context.Context, seriesID string, seasonID string, episodeID string, image string) error
	ListSeasonEpisodes(ctx context.Context, seasonID string) (models.Episodes, error)
	ListCollectionEpisodes(ctx context.Context) (models.Episodes, error)
}

func (e *episode) CreateEpisode(ctx context.Context, seasonID string, title string, postersPath []string, trailerURL string, length *models.ShowLength, rating float64, resume string, writtenBy models.FilmCrews, producedBy models.FilmCrews, directedBy models.FilmCrews, starring models.ShortCelebrities) (*models.Episode, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewEpisodeSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	showLength := &pb.ShowLength{
		Hours:   int32(length.Hours),
		Minutes: int32(length.Minutes),
	}
	request := pb.CreateEpisodeRequest{
		SeasonId:    seasonID,
		Title:       title,
		PostersPath: postersPath,
		TrailerUrl:  trailerURL,
		ShowLength:  showLength,
		Rating:      rating,
		Resume:      resume,
		WrittenBy:   toFilmCrewPb(writtenBy),
		ProducedBy:  toFilmCrewPb(producedBy),
		DirectedBy:  toFilmCrewPb(directedBy),
		Starring:    toShortCelebritiesPb(starring),
	}
	resp, err := c.CreateEpisode(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Create episode request")
	}

	return resp.ToModel().(*models.Episode), nil
}

func (e *episode) GetEpisode(ctx context.Context, ID string) (*models.Episode, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewEpisodeSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.GetEpisode(ctx, &pb.GetByIDRequest{Id: ID})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Get episode request")
	}

	return resp.ToModel().(*models.Episode), nil
}

func (e *episode) UpdateEpisode(ctx context.Context, ID string, seasonID string, title string, postersPath []string, trailerURL string, length *models.ShowLength, rating float64, resume string, writtenBy models.FilmCrews, producedBy models.FilmCrews, directedBy models.FilmCrews, starring models.ShortCelebrities) (*models.Episode, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewEpisodeSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	showLength := &pb.ShowLength{
		Hours:   int32(length.Hours),
		Minutes: int32(length.Minutes),
	}
	request := pb.Episode{
		Id:          ID,
		SeasonId:    seasonID,
		Title:       title,
		PostersPath: postersPath,
		TrailerUrl:  trailerURL,
		ShowLength:  showLength,
		Rating:      rating,
		Resume:      resume,
		WrittenBy:   toFilmCrewPb(writtenBy),
		ProducedBy:  toFilmCrewPb(producedBy),
		DirectedBy:  toFilmCrewPb(directedBy),
		Starring:    toShortCelebritiesPb(starring),
	}
	resp, err := c.UpdateEpisode(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Update episode request")
	}

	return resp.ToModel().(*models.Episode), nil
}

func (e *episode) UploadEpisodePostersService(ctx context.Context, episodeID string, postersPath []string) (*models.Episode, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewEpisodeSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := client.UploadEpisodePosters(ctx, &pb.UploadEpisodePostersServiceRequest{
		EpisodeId:   episodeID,
		PostersPath: postersPath,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Updating episode posters request in database")
	}
	return resp.ToModel().(*models.Episode), nil
}

func (e *episode) DeleteEpisodePosterService(ctx context.Context, seriesID string, seasonID string, episodeID string, image string) error {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewEpisodeSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = client.DeleteEpisodePoster(ctx, &pb.DeleteEpisodePosterRequest{
		SeriesId:  seriesID,
		SeasonId:  seasonID,
		EpisodeId: episodeID,
		Image:     image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Deleting episode poster request in database")
	}
	return nil
}

func (e *episode) DeleteSeasonPosterService(ctx context.Context, seriesID string, seasonID string, episodeID string, image string) error {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewEpisodeSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = client.DeleteEpisodePoster(ctx, &pb.DeleteEpisodePosterRequest{
		SeriesId:  seriesID,
		SeasonId:  seasonID,
		EpisodeId: episodeID,
		Image:     image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Deleting episode poster request in database")
	}
	return nil
}

func (e *episode) DeleteEpisodePosterFS(ctx context.Context, seriesID string, seasonID string, episodeID string, image string) error {
	conn, err := dialGrpcFileServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc file server")
	}
	defer conn.Close()
	client := pb.NewFileServerEpisodeSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	_, err = client.DeleteEpisodePoster(ctx, &pb.DeleteEpisodePosterRequest{
		SeriesId:  seriesID,
		SeasonId:  seasonID,
		EpisodeId: episodeID,
		Image:     image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Deleting episode poster request in file server")
	}
	return nil
}

func (e *episode) ListSeasonEpisodes(ctx context.Context, seasonID string) (models.Episodes, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewEpisodeSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.ListSeasonEpisodes(ctx, &pb.GetByIDRequest{Id: seasonID})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during List season episodes request")
	}
	return resp.ToModel(), nil
}

func (e *episode) ListCollectionEpisodes(ctx context.Context) (models.Episodes, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	c := pb.NewEpisodeSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.ListCollectionEpisodes(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during List episodes request")
	}
	return resp.ToModel(), nil
}

func toShortEpisodesGrpc(episodesModel models.ShortEpisodes) *pb.ShortEpisodeList {
	pbEpisodes := pb.ShortEpisodeList{}
	for _, shortEp := range episodesModel {
		pbEpisodes.ShortEpisodes = append(pbEpisodes.ShortEpisodes, &pb.ShortEpisode{
			Id:          shortEp.ID,
			Title:       shortEp.Title,
			PostersPath: shortEp.PostersPath,
			Rating:      shortEp.Rating,
			Resume:      shortEp.Resume,
		})
	}
	return &pbEpisodes
}

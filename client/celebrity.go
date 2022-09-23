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

type celebrity struct{}

type uploadCelebrity struct {
	client pb.FileServerCelebritySvcClient
}

func NewCelebrityClient() CelebrityClienter {
	return &celebrity{}
}

func NewUploadCelebrityClient() Clienter {
	return &uploadCelebrity{}
}

func (c *uploadCelebrity) getClient(ctx context.Context, conn *grpc.ClientConn) Clienter {
	return &uploadCelebrity{
		client: pb.NewFileServerCelebritySvcClient(conn),
	}
}

func (c *uploadCelebrity) getStream(ctx context.Context) (grpc.ClientStream, error) {
	stream, err := c.client.UploadCelebrityPosters(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error while uploading celebrity images in file server!")
	}
	return stream, nil
}

func (c *uploadCelebrity) getRequestMessage(imageExtension string, IDs []string) proto.Message {
	return &pb.UploadCelebrityPostersFSRequest{
		CelebrityId:    IDs[0],
		ImageExtention: imageExtension,
	}
}

func (c *uploadCelebrity) getChunkData(buffer []byte, n int) proto.Message {
	return &pb.UploadCelebrityPostersFSRequest{
		ChunkData: buffer[:n],
	}
}

type CelebrityClienter interface {
	CreateCelebrity(ctx context.Context, name string, occupation []string, postersPath []string, dateOfBirth time.Time, dateOfDeath time.Time, placeOfBirth string, gender models.Gender, bio string) (*models.Celebrity, error)
	GetCelebrity(ctx context.Context, ID string) (*models.Celebrity, error)
	UpdateCelebrity(ctx context.Context, ID string, name string, occupation []string, postersPath []string, dateOfBirth time.Time, dateOfDeath time.Time, placeOfBirth string, gender models.Gender, bio string) (*models.Celebrity, error)
	UploadCelebrityPostersService(ctx context.Context, ID string, postersPath []string) (*models.Celebrity, error)
	DeleteCelebrityPosterService(ctx context.Context, ID string, image string) error
	DeleteCelebrityPosterFS(ctx context.Context, ID string, image string) error
	ListCelebrities(ctx context.Context) (models.Celebrities, error)
}

func (c *celebrity) CreateCelebrity(ctx context.Context, name string, occupation []string, postersPath []string, dateOfBirth time.Time, dateOfDeath time.Time, placeOfBirth string, gender models.Gender, bio string) (*models.Celebrity, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewCelebritySvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.CreateCelebrityRequest{
		Name:         name,
		PostersPath:  postersPath,
		DateOfBirth:  timestamppb.New(dateOfBirth),
		DateOfDeath:  timestamppb.New(dateOfDeath),
		PlaceOfBirth: placeOfBirth,
		Gender:       string(gender),
		Bio:          bio,
		Occupation:   occupation,
	}
	resp, err := client.CreateCelebrity(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Create celebrity request")
	}

	return resp.ToModel().(*models.Celebrity), nil
}

func (c *celebrity) GetCelebrity(ctx context.Context, ID string) (*models.Celebrity, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewCelebritySvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := client.GetCelebrity(ctx, &pb.GetByIDRequest{Id: ID})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Get celebrity request")
	}

	return resp.ToModel().(*models.Celebrity), nil
}

func (c *celebrity) UpdateCelebrity(ctx context.Context, ID string, name string, occupation []string, postersPath []string, dateOfBirth time.Time, dateOfDeath time.Time, placeOfBirth string, gender models.Gender, bio string) (*models.Celebrity, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewCelebritySvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.Celebrity{
		Id:           ID,
		Name:         name,
		PostersPath:  postersPath,
		DateOfBirth:  timestamppb.New(dateOfBirth),
		DateOfDeath:  timestamppb.New(dateOfDeath),
		PlaceOfBirth: placeOfBirth,
		Gender:       string(gender),
		Bio:          bio,
		Occupation:   occupation,
	}
	resp, err := client.UpdateCelebrity(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Update celebrity request")
	}

	return resp.ToModel().(*models.Celebrity), nil
}

func (c *celebrity) UploadCelebrityPostersService(ctx context.Context, ID string, postersPath []string) (*models.Celebrity, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewCelebritySvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := client.UploadCelebrityPosters(ctx, &pb.UploadCelebrityPostersServiceRequest{
		CelebrityId: ID,
		PostersPath: postersPath,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Updating celebrity posters request in database")
	}
	return resp.ToModel().(*models.Celebrity), nil
}

func (c *celebrity) DeleteCelebrityPosterService(ctx context.Context, ID string, image string) error {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewCelebritySvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = client.DeleteCelebrityPoster(ctx, &pb.DeleteCelebrityPosterRequest{
		CelebrityId: ID,
		Image:       image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Deleting celebrity poster request in database")
	}
	return nil
}

func (c *celebrity) DeleteCelebrityPosterFS(ctx context.Context, ID string, image string) error {
	conn, err := dialGrpcFileServer(ctx)
	if err != nil {
		return errors.Wrap(err, "Error on dialing Grpc file server")
	}
	defer conn.Close()
	client := pb.NewFileServerCelebritySvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	_, err = client.DeleteCelebrityPoster(ctx, &pb.DeleteCelebrityPosterRequest{
		CelebrityId: ID,
		Image:       image,
	})
	if err != nil {
		return errors.Wrap(err, "Grpc service returned an error during Deleting celebrity poster request in file server")
	}
	return nil
}

func (c *celebrity) ListCelebrities(ctx context.Context) (models.Celebrities, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()
	client := pb.NewCelebritySvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := client.ListCelebrities(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during List celebrities request")
	}
	return resp.ToModel(), nil
}

func toFilmCrewPb(filmCrews models.FilmCrews) *pb.FilmCrew {
	crews := pb.FilmCrew{}
	for _, crew := range filmCrews {
		crews.FilmCrew = append(crews.FilmCrew, &pb.FilmStaff{
			Id:          crew.ID,
			Name:        crew.Name,
			PostersPath: crew.PostersPath,
		})
	}
	return &crews
}

func toShortCelebritiesPb(shortCelebs models.ShortCelebrities) *pb.ShortCelebrities {
	celebs := pb.ShortCelebrities{}
	for _, celeb := range shortCelebs {
		celebs.ShortCelebs = append(celebs.ShortCelebs, &pb.ShortCelebrity{
			Id:          celeb.ID,
			Name:        celeb.Name,
			RoleName:    celeb.RoleName,
			PostersPath: celeb.PostersPath,
		})
	}
	return &celebs
}

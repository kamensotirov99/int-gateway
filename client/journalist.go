package client

import (
	"context"
	"time"

	pb "int-gateway/_proto"
	"int-gateway/models"

	"github.com/pkg/errors"
)

type journalist struct{}

func NewJournalistClient() JournalistClienter {
	return &journalist{}
}

type JournalistClienter interface {
	CreateJournalist(ctx context.Context, name string) (*models.Journalist, error)
	GetJournalist(ctx context.Context, ID string) (*models.Journalist, error)
	GetJournalistByName(ctx context.Context, name string) (*models.Journalist, error)
	ListJournalists(ctx context.Context) (models.Journalists, error)
	UpdateJournalist(ctx context.Context, ID string, name string) (*models.Journalist, error)
}

func (j *journalist) CreateJournalist(ctx context.Context, name string) (*models.Journalist, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()

	c := pb.NewJournalistSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.CreateJournalistRequest{
		Name: name,
	}
	resp, err := c.CreateJournalist(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during Create request")
	}

	return &models.Journalist{
		ID:   resp.Id,
		Name: resp.Name,
	}, nil
}

func (j *journalist) GetJournalist(ctx context.Context, ID string) (*models.Journalist, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()

	c := pb.NewJournalistSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.GetByIDRequest{
		Id: ID,
	}
	resp, err := c.GetJournalist(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during GetJournalist request")
	}

	return &models.Journalist{
		ID:   resp.Id,
		Name: resp.Name,
	}, nil
}

func (j *journalist) GetJournalistByName(ctx context.Context, name string) (*models.Journalist, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()

	c := pb.NewJournalistSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.GetByNameRequest{
		Name: name,
	}
	resp, err := c.GetJournalistByName(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during GetJournalistByName request")
	}

	return &models.Journalist{
		ID:   resp.Id,
		Name: resp.Name,
	}, nil
}

func (j *journalist) ListJournalists(ctx context.Context) (models.Journalists, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing  Grpc service")
	}
	defer conn.Close()

	c := pb.NewJournalistSvcClient(conn)

	// creating context with 15 seconds wait time for response
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.ListJournalists(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during ListJournalists request")
	}
	journalists := []*models.Journalist{}
	for _, c := range resp.Journalists {
		journalists = append(journalists, &models.Journalist{ID: c.Id,
			Name: c.Name})
	}
	return journalists, nil

}

func (j *journalist) UpdateJournalist(ctx context.Context, ID string, name string) (*models.Journalist, error) {
	conn, err := dialGrpcServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc service")
	}
	defer conn.Close()

	c := pb.NewJournalistSvcClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	request := pb.Journalist{
		Id:   ID,
		Name: name,
	}
	resp, err := c.UpdateJournalist(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Grpc service returned an error during UpdateJournalist request")
	}

	return &models.Journalist{
		ID:   resp.Id,
		Name: resp.Name,
	}, nil
}

package client

import (
	"bufio"
	"context"
	pb "int-gateway/_proto"
	"int-gateway/models"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type uploadClient struct {
	client Clienter
}

func NewUploadClient(c Clienter) UploadClienter {
	return &uploadClient{
		client: c,
	}
}

type Clienter interface {
	getClient(ctx context.Context, conn *grpc.ClientConn) Clienter
	getStream(ctx context.Context) (grpc.ClientStream, error)
	getRequestMessage(imageExtension string, IDs []string) proto.Message
	getChunkData(buffer []byte, n int) proto.Message
}

type UploadClienter interface {
	UploadPostersFS(ctx context.Context, images []*multipart.FileHeader, IDs ...string) (*models.UploadResponse, error)
}

func (c *uploadClient) UploadPostersFS(ctx context.Context, images []*multipart.FileHeader, IDs ...string) (*models.UploadResponse, error) {
	conn, err := dialGrpcFileServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error on dialing Grpc file server")
	}
	defer conn.Close()
	client := c.client.getClient(ctx, conn)

	postersPath := pb.UploadPostersResponse{}
	for i := range images {
		image, err := images[i].Open()
		if err != nil {
			return nil, errors.Wrap(err, "Error while opening file headers!")
		}
		defer image.Close()

		ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()

		stream, err := client.getStream(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "Error while uploading images in file server!")
		}

		req := c.client.getRequestMessage(filepath.Ext(images[i].Filename), IDs)
		err = stream.SendMsg(req)
		if err != nil {
			return nil, errors.Wrap(err, "Error while sending image to the file server!")
		}

		reader := bufio.NewReader(image)
		buffer := make([]byte, 1024)

		for {
			n, err := reader.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, errors.Wrap(err, "Error reading chunk!")
			}

			req := c.client.getChunkData(buffer, n)
			err = stream.SendMsg(req)
			if err != nil {
				return nil, errors.Wrap(err, "Error while sending chunk to the file server!")
			}
		}

		err = stream.CloseSend()
		if err != nil {
			return nil, errors.Wrap(err, "Error while closing stream!")
		}

		res := pb.UploadPosterResponse{}
		err = stream.RecvMsg(&res)
		if err != nil {
			return nil, errors.Wrap(err, "Error while receiving response!")
		}
		postersPath.PostersPath = append(postersPath.PostersPath, res.PosterPath)
	}
	return postersPath.ToModel(), nil
}

func dialGrpcServer(ctx context.Context) (*grpc.ClientConn, error) {
	serviceHost := "localhost"
	servicePort := "2002"
	serverAddress := serviceHost + ":" + servicePort

	return grpc.DialContext(ctx, serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func dialGrpcFileServer(ctx context.Context) (*grpc.ClientConn, error) {
	serviceHost := "ArgoXInterns"
	servicePort := "2004"
	serverAddress := serviceHost + ":" + servicePort

	return grpc.DialContext(ctx, serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

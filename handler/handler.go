package handler

import (
	"encoding/json"

	"int-gateway/client"
	"int-gateway/service"

	"net/http"

	"github.com/sirupsen/logrus"
)

const MAX_UPLOAD_SIZE = 32 << 20

type Handler struct {
	Article    service.ArticleServicer
	Celebrity  service.CelebrityServicer
	Episode    service.EpisodeServicer
	Genre      service.GenreServicer
	Journalist service.JournalistServicer
	Season     service.SeasonServicer
	Show       service.ShowServicer
}

func InitiateHandler(logger *logrus.Logger) *Handler {
	return &Handler{
		Article:    service.InitiateArticleService(logger,client.NewArticleClient(), client.NewUploadClient(client.NewUploadArticleClient())),
		Episode:    service.InitiateEpisodeService(logger,client.NewEpisodeClient(), client.NewUploadClient(client.NewUploadEpisodeClient())),
		Celebrity:  service.InitiateCelebrityService(logger, client.NewCelebrityClient(), client.NewUploadClient(client.NewUploadCelebrityClient())),
		Genre:      service.InitiateGenreService(logger),
		Journalist: service.InitiateJournalistService(logger),
		Season:     service.InitiateSeasonService(logger, client.NewSeasonClient(), client.NewUploadClient(client.NewUploadSeasonClient())),
		Show:       service.InitiateShowService(logger, client.NewShowClient(), client.NewUploadClient(client.NewUploadSeriesClient()), client.NewUploadClient(client.NewUploadMovieClient())),
	}
}

func createResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

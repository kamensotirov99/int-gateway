package handler

import (
	"encoding/json"
	"int-gateway/request"
	"int-gateway/response"
	"net/http"
)

func (h *Handler) CreateEpisode(w http.ResponseWriter, r *http.Request) {
	episode := request.CreateEpisode{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&episode); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if episode.SeasonID == "" {
		http.Error(w, "Missing seasonId parameter", http.StatusBadRequest)
		return
	}
	if episode.Title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	if episode.TrailerURL == "" {
		http.Error(w, "Missing trailerUrl parameter", http.StatusBadRequest)
		return
	}
	if episode.Starring == nil || len(episode.Starring) == 0 {
		http.Error(w, "Missing starring parameter", http.StatusBadRequest)
		return
	}
	if episode.ProducedBy == nil || len(episode.ProducedBy) == 0 {
		http.Error(w, "Missing producedBy parameter", http.StatusBadRequest)
		return
	}
	if episode.DirectedBy == nil || len(episode.DirectedBy) == 0 {
		http.Error(w, "Missing directedBy parameter", http.StatusBadRequest)
		return
	}
	if episode.WrittenBy == nil || len(episode.WrittenBy) == 0 {
		http.Error(w, "Missing writtenBy parameter", http.StatusBadRequest)
		return
	}
	if episode.Resume == "" {
		http.Error(w, "Missing resume parameter", http.StatusBadRequest)
		return
	}
	if episode.Length.Hours == 0 && episode.Length.Minutes == 0 {
		http.Error(w, "Missing length parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Episode.CreateEpisode(r.Context(), episode.SeasonID, episode.Title, episode.PostersPath, episode.TrailerURL, episode.Length.ToModel(), episode.Rating, episode.Resume, episode.WrittenBy.ToModel(), episode.ProducedBy.ToModel(), episode.DirectedBy.ToModel(), episode.Starring.ToModel())
	if err != nil {
		http.Error(w, "Error while creating episode", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetEpisode(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idString := params["id"]
	if len(idString) == 0 || idString[0] == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	resp, err := h.Episode.GetEpisode(r.Context(), idString[0])
	if err != nil {
		http.Error(w, "Error while getting episode", http.StatusInternalServerError)
		return
	}
	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateEpisode(w http.ResponseWriter, r *http.Request) {
	episode := request.UpdateEpisode{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&episode); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if episode.ID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	if episode.SeasonID == "" {
		http.Error(w, "Missing seasonId parameter", http.StatusBadRequest)
		return
	}
	if episode.Title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	if episode.TrailerURL == "" {
		http.Error(w, "Missing trailerUrl parameter", http.StatusBadRequest)
		return
	}
	if episode.Starring == nil || len(episode.Starring) == 0 {
		http.Error(w, "Missing starring parameter", http.StatusBadRequest)
		return
	}
	if episode.ProducedBy == nil || len(episode.ProducedBy) == 0 {
		http.Error(w, "Missing producedBy parameter", http.StatusBadRequest)
		return
	}
	if episode.DirectedBy == nil || len(episode.DirectedBy) == 0 {
		http.Error(w, "Missing directedBy parameter", http.StatusBadRequest)
		return
	}
	if episode.WrittenBy == nil || len(episode.WrittenBy) == 0 {
		http.Error(w, "Missing writtenBy parameter", http.StatusBadRequest)
		return
	}
	if episode.Resume == "" {
		http.Error(w, "Missing resume parameter", http.StatusBadRequest)
		return
	}
	if episode.Length.Hours == 0 && episode.Length.Minutes == 0 {
		http.Error(w, "Missing length parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Episode.UpdateEpisode(r.Context(), episode.ID, episode.SeasonID, episode.Title, episode.PostersPath, episode.TrailerURL, episode.Length.ToModel(), episode.Rating, episode.Resume, episode.WrittenBy.ToModel(), episode.ProducedBy.ToModel(), episode.DirectedBy.ToModel(), episode.Starring.ToModel())
	if err != nil {
		http.Error(w, "Error while updating episode", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UploadEpisodePosters(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "Files are too big!", http.StatusRequestEntityTooLarge)
		return
	}
	receivedImages, correct := r.MultipartForm.File["images"]
	if !correct {
		http.Error(w, "Missing images key or the value is empty", http.StatusBadRequest)
		return
	}
	receivedSeriesID, correct := r.Form["seriesId"]
	if !correct || receivedSeriesID[0] == "" {
		http.Error(w, "Missing seriesId key or the value is empty", http.StatusBadRequest)
		return
	}
	receivedSeasonID, correct := r.Form["seasonId"]
	if !correct || receivedSeasonID[0] == "" {
		http.Error(w, "Missing seasonId key or the value is empty", http.StatusBadRequest)
		return
	}
	receivedEpisodeID, correct := r.Form["episodeId"]
	if !correct || receivedEpisodeID[0] == "" {
		http.Error(w, "Missing episodeId key or the value is empty", http.StatusBadRequest)
		return
	}

	resp, err := h.Episode.UploadEpisodePosters(r.Context(), receivedSeriesID[0], receivedSeasonID[0], receivedEpisodeID[0], receivedImages)
	if err != nil {
		http.Error(w, "Error while uploading episode posters", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteEpisodePoster(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	seriesID := params["seriesId"]
	seasonID := params["seasonId"]
	episodeID := params["episodeId"]
	image := params["image"]
	if len(seriesID) == 0 || seriesID[0] == "" {
		http.Error(w, "Missing seriesId parameter", http.StatusBadRequest)
		return
	}
	if len(seasonID) == 0 || seasonID[0] == "" {
		http.Error(w, "Missing seasonId parameter", http.StatusBadRequest)
		return
	}
	if len(episodeID) == 0 || episodeID[0] == "" {
		http.Error(w, "Missing episodeId parameter", http.StatusBadRequest)
		return
	}
	if len(image) == 0 || image[0] == "" {
		http.Error(w, "Missing image parameter", http.StatusBadRequest)
		return
	}

	err := h.Episode.DeleteEpisodePoster(r.Context(), seriesID[0], seasonID[0], episodeID[0], image[0])
	if err != nil {
		http.Error(w, "Error while deleting episode poster", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListEpisodes(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var episodes response.Episodes
	var err error
	if params["seasonId"] != nil {
		seasonIdString := params["seasonId"]
		if len(seasonIdString) == 0 || seasonIdString[0] == "" {
			http.Error(w, "Missing seasonId parameter", http.StatusBadRequest)
			return
		}
		episodes, err = h.Episode.ListSeasonEpisodes(r.Context(), seasonIdString[0])
		if err != nil {
			http.Error(w, "Error while getting episodes by seasonId", http.StatusInternalServerError)
			return
		}
	} else {
		episodes, err = h.Episode.ListCollectionEpisodes(r.Context())
		if err != nil {
			http.Error(w, "Error while listing all episodes!", http.StatusInternalServerError)
			return
		}
	}

	err = createResponse(w, episodes)
	if err != nil {
		http.Error(w, "Error while encoding response body!", http.StatusBadRequest)
		return
	}
}

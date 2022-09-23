package handler

import (
	"encoding/json"
	"int-gateway/request"
	"int-gateway/response"
	"net/http"
)

func (h *Handler) CreateSeason(w http.ResponseWriter, r *http.Request) {
	season := request.CreateSeason{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&season); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if season.ShowID == "" {
		http.Error(w, "Missing show id parameter", http.StatusBadRequest)
		return
	}
	if season.Title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	if season.TrailerURL == "" {
		http.Error(w, "Missing trailerUrl parameter", http.StatusBadRequest)
		return
	}
	if season.Resume == "" {
		http.Error(w, "Missing resume parameter", http.StatusBadRequest)
		return
	}
	if season.ReleaseDate.IsZero() {
		http.Error(w, "Missing release date parameter", http.StatusBadRequest)
		return
	}
	if season.WrittenBy == nil || len(season.WrittenBy) == 0 {
		http.Error(w, "Missing writtenBy parameter", http.StatusBadRequest)
		return
	}
	if season.ProducedBy == nil || len(season.ProducedBy) == 0 {
		http.Error(w, "Missing producedBy parameter", http.StatusBadRequest)
		return
	}
	if season.DirectedBy == nil || len(season.DirectedBy) == 0 {
		http.Error(w, "Missing directedBy parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Season.CreateSeason(r.Context(), season.ShowID, season.Title, season.TrailerURL, season.PostersPath, season.ReleaseDate,
		season.Rating, season.Resume, season.DirectedBy.ToModel(), season.ProducedBy.ToModel(), season.WrittenBy.ToModel(), season.Episodes.ToModel())
	if err != nil {
		http.Error(w, "Error while creating season", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetSeason(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idString := params["id"]
	if len(idString) == 0 || idString[0] == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	resp, err := h.Season.GetSeason(r.Context(), idString[0])
	if err != nil {
		http.Error(w, "Error while getting season", http.StatusInternalServerError)
		return
	}
	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateSeason(w http.ResponseWriter, r *http.Request) {
	season := request.UpdateSeason{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&season); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if season.ID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	if season.ShowID == "" {
		http.Error(w, "Missing showId parameter", http.StatusBadRequest)
		return
	}
	if season.Title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	if season.TrailerURL == "" {
		http.Error(w, "Missing trailerUrl parameter", http.StatusBadRequest)
		return
	}
	if season.Resume == "" {
		http.Error(w, "Missing resume parameter", http.StatusBadRequest)
		return
	}
	if season.Rating <= 0 {
		http.Error(w, "Missing rating parameter", http.StatusBadRequest)
		return
	}
	if season.ReleaseDate.IsZero() {
		http.Error(w, "Missing release date parameter", http.StatusBadRequest)
		return
	}
	if season.WrittenBy == nil || len(season.WrittenBy) == 0 {
		http.Error(w, "Missing writtenBy parameter", http.StatusBadRequest)
		return
	}
	if season.ProducedBy == nil || len(season.ProducedBy) == 0 {
		http.Error(w, "Missing producedBy parameter", http.StatusBadRequest)
		return
	}
	if season.DirectedBy == nil || len(season.DirectedBy) == 0 {
		http.Error(w, "Missing directedBy parameter", http.StatusBadRequest)
		return
	}
	if season.Episodes == nil || len(season.Episodes) == 0 {
		http.Error(w, "Missing episodes parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Season.UpdateSeason(r.Context(), season.ID, season.ShowID, season.Title, season.TrailerURL, season.PostersPath, season.ReleaseDate,
		season.Rating, season.Resume, season.DirectedBy.ToModel(), season.ProducedBy.ToModel(), season.WrittenBy.ToModel(), season.Episodes.ToModel())
	if err != nil {
		http.Error(w, "Error while updating season", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) 	UploadSeasonPosters(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.Season.UploadSeasonPosters(r.Context(), receivedSeriesID[0], receivedSeasonID[0], receivedImages)
	if err != nil {
		http.Error(w, "Error while uploading season posters", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteSeasonPoster(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	seriesID := params["seriesId"]
	seasonID := params["seasonId"]
	image := params["image"]
	if len(seriesID) == 0 || seriesID[0] == "" {
		http.Error(w, "Missing seriesId parameter", http.StatusBadRequest)
		return
	}
	if len(seasonID) == 0 || seasonID[0] == "" {
		http.Error(w, "Missing seasonId parameter", http.StatusBadRequest)
		return
	}
	if len(image) == 0 || image[0] == "" {
		http.Error(w, "Missing image parameter", http.StatusBadRequest)
		return
	}

	err := h.Season.DeleteSeasonPoster(r.Context(), seriesID[0], seasonID[0], image[0])
	if err != nil {
		http.Error(w, "Error while deleting season poster", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListSeasons(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var seasons response.Seasons
	var err error
	if params["showId"] != nil {
		showIdString := params["showId"]
		if len(showIdString) == 0 || showIdString[0] == "" {
			http.Error(w, "Missing showId parameter", http.StatusBadRequest)
			return
		}
		seasons, err = h.Season.ListShowSeasons(r.Context(), showIdString[0])
		if err != nil {
			http.Error(w, "Error while getting seasons by showId", http.StatusInternalServerError)
			return
		}

	} else {
		seasons, err = h.Season.ListSeasonsCollection(r.Context())
		if err != nil {
			http.Error(w, "Error while listing all seasons!", http.StatusInternalServerError)
			return
		}
	}

	err = createResponse(w, seasons)
	if err != nil {
		http.Error(w, "Error while encoding response body!", http.StatusBadRequest)
		return
	}
}

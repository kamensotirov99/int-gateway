package handler

import (
	"encoding/json"
	"int-gateway/request"
	"net/http"
)

func (h *Handler) 	CreateShow(w http.ResponseWriter, r *http.Request) {
	show := request.CreateShow{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&show); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if show.Title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	if show.Genres == nil || len(show.Genres) == 0 {
		http.Error(w, "Missing genres parameter", http.StatusBadRequest)
		return
	}
	if show.Type == "" {
		http.Error(w, "Missing type parameter", http.StatusBadRequest)
		return
	}
	if show.ReleaseDate.IsZero() {
		http.Error(w, "Missing releaseDate parameter", http.StatusBadRequest)
		return
	}
	if show.TrailerURL == "" {
		http.Error(w, "Missing trailerUrl parameter", http.StatusBadRequest)
		return
	}
	if show.Description == "" {
		http.Error(w, "Missing description parameter", http.StatusBadRequest)
		return
	}
	if show.Starring == nil || len(show.Starring) == 0 {
		http.Error(w, "Missing starring parameter", http.StatusBadRequest)
		return
	}
	if show.ProducedBy == nil || len(show.ProducedBy) == 0 {
		http.Error(w, "Missing producedBy parameter", http.StatusBadRequest)
		return
	}
	if show.DirectedBy == nil || len(show.DirectedBy) == 0 {
		http.Error(w, "Missing directedBy parameter", http.StatusBadRequest)
		return
	}
	if show.WrittenBy == nil || len(show.WrittenBy) == 0 {
		http.Error(w, "Missing writtenBy parameter", http.StatusBadRequest)
		return
	}
	if show.Length.Hours == 0 && show.Length.Minutes == 0 {
		http.Error(w, "Missing length parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Show.CreateShow(r.Context(), show.Title, show.Type, show.PostersPath, show.ReleaseDate, show.EndDate, show.Rating, show.Length.ToModel(), show.TrailerURL, show.Genres.ToModel(), show.DirectedBy.ToModel(), show.ProducedBy.ToModel(), show.WrittenBy.ToModel(), show.Starring.ToModel(), show.Description, show.Seasons.ToModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetShow(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idString := params["id"]
	if len(idString) == 0 || idString[0] == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	resp, err := h.Show.GetShow(r.Context(), idString[0])
	if err != nil {
		http.Error(w, "Error while getting show", http.StatusInternalServerError)
		return
	}
	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateShow(w http.ResponseWriter, r *http.Request) {
	show := request.UpdateShow{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&show); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if show.ID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	if show.Title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	if show.Type == "" {
		http.Error(w, "Missing type parameter", http.StatusBadRequest)
		return
	}
	if show.ReleaseDate.IsZero() {
		http.Error(w, "Missing releaseDate parameter", http.StatusBadRequest)
		return
	}
	if show.TrailerURL == "" {
		http.Error(w, "Missing trailerUrl parameter", http.StatusBadRequest)
		return
	}
	if show.Description == "" {
		http.Error(w, "Missing description parameter", http.StatusBadRequest)
		return
	}
	if show.Starring == nil || len(show.Starring) == 0 {
		http.Error(w, "Missing starring parameter", http.StatusBadRequest)
		return
	}
	if show.ProducedBy == nil || len(show.ProducedBy) == 0 {
		http.Error(w, "Missing producedBy parameter", http.StatusBadRequest)
		return
	}
	if show.DirectedBy == nil || len(show.DirectedBy) == 0 {
		http.Error(w, "Missing directedBy parameter", http.StatusBadRequest)
		return
	}
	if show.WrittenBy == nil || len(show.WrittenBy) == 0 {
		http.Error(w, "Missing writtenBy parameter", http.StatusBadRequest)
		return
	}
	resp, err := h.Show.UpdateShow(r.Context(), show.ID, show.Title, show.Type, show.PostersPath, show.ReleaseDate, show.EndDate, show.Rating, show.Length.ToModel(), show.TrailerURL, show.Genres.ToModel(), show.DirectedBy.ToModel(), show.ProducedBy.ToModel(), show.WrittenBy.ToModel(), show.Starring.ToModel(), show.Description, show.Seasons.ToModel())
	if err != nil {
		http.Error(w, "Error while updating show", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ListShows(w http.ResponseWriter, r *http.Request) {
	response, err := h.Show.ListShows(r.Context())
	if err != nil {
		http.Error(w, "Error while listing all shows!", http.StatusInternalServerError)
		return
	}

	err = createResponse(w, response)
	if err != nil {
		http.Error(w, "Error while encoding response body!", http.StatusBadRequest)
		return
	}
}

func (h *Handler) UploadSeriesPosters(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "Files are too big!", http.StatusRequestEntityTooLarge)
		return
	}
	receivedImages, correct := r.MultipartForm.File["images"]
	if !correct {
		http.Error(w, "Missing images key or the value is empty ", http.StatusBadRequest)
		return
	}
	receivedSeriesID, correct := r.Form["seriesId"]
	if !correct || receivedSeriesID[0] == "" {
		http.Error(w, "Missing seriesId key or the value is empty", http.StatusBadRequest)
		return
	}

	resp, err := h.Show.UploadSeriesPosters(r.Context(), receivedSeriesID[0], receivedImages)
	if err != nil {
		http.Error(w, "Error while uploading series posters", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteSeriesPoster(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	seriesID := params["seriesId"]
	image := params["image"]
	if len(seriesID) == 0 || seriesID[0] == "" {
		http.Error(w, "Missing seriesID parameter", http.StatusBadRequest)
		return
	}
	if len(image) == 0 || image[0] == "" {
		http.Error(w, "Missing image parameter", http.StatusBadRequest)
		return
	}

	err := h.Show.DeleteSeriesPoster(r.Context(), seriesID[0], image[0])
	if err != nil {
		http.Error(w, "Error while deleting series posters ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UploadMoviePosters(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "Files are too big!", http.StatusRequestEntityTooLarge)
		return
	}
	receivedImages, correct := r.MultipartForm.File["images"]
	if !correct {
		http.Error(w, "Missing images key or the value is empty ", http.StatusBadRequest)
		return
	}
	receivedMovieID, correct := r.Form["movieId"]
	if !correct || receivedMovieID[0] == "" {
		http.Error(w, "Missing movieId key or the value is empty", http.StatusBadRequest)
		return
	}

	resp, err := h.Show.UploadMoviePosters(r.Context(), receivedMovieID[0], receivedImages)
	if err != nil {
		http.Error(w, "Error while uploading movie posters", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteMoviePoster(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	movieID := params["movieId"]
	image := params["image"]
	if len(movieID) == 0 || movieID[0] == "" {
		http.Error(w, "Missing movieID parameter", http.StatusBadRequest)
		return
	}
	if len(image) == 0 || image[0] == "" {
		http.Error(w, "Missing image parameter", http.StatusBadRequest)
		return
	}

	err := h.Show.DeleteMoviePoster(r.Context(), movieID[0], image[0])
	if err != nil {
		http.Error(w, "Error while deleting movie posters ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

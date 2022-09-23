package handler

import (
	"encoding/json"
	"int-gateway/models"
	"int-gateway/request"
	"net/http"
)

func (h *Handler) CreateCelebrity(w http.ResponseWriter, r *http.Request) {
	celebrity := request.CreateCelebrity{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&celebrity); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if celebrity.Name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}
	if celebrity.Occupation == nil || len(celebrity.Occupation) == 0 {
		http.Error(w, "Missing occupation parameter", http.StatusBadRequest)
		return
	}
	if celebrity.DateOfBirth.IsZero() {
		http.Error(w, "Missing dateOfBirth parameter", http.StatusBadRequest)
		return
	}
	if celebrity.PlaceOfBirth == "" {
		http.Error(w, "Missing placeOfBirth parameter", http.StatusBadRequest)
		return
	}
	if celebrity.Gender == "" {
		http.Error(w, "Missing gender parameter", http.StatusBadRequest)
		return
	}
	if celebrity.Bio == "" {
		http.Error(w, "Missing bio parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Celebrity.CreateCelebrity(r.Context(), celebrity.Name, celebrity.Occupation, celebrity.PostersPath, celebrity.DateOfBirth, celebrity.DateOfDeath, celebrity.PlaceOfBirth, models.Gender(celebrity.Gender), celebrity.Bio)
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

func (h *Handler) GetCelebrity(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idString := params["id"]
	if len(idString) == 0 || idString[0] == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	resp, err := h.Celebrity.GetCelebrity(r.Context(), idString[0])
	if err != nil {
		http.Error(w, "Error while getting celebrity", http.StatusInternalServerError)
		return
	}
	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateCelebrity(w http.ResponseWriter, r *http.Request) {
	celebrity := request.UpdateCelebrity{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&celebrity); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if celebrity.ID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	if celebrity.Name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}
	if celebrity.Occupation == nil || len(celebrity.Occupation) == 0 {
		http.Error(w, "Missing occupation parameter", http.StatusBadRequest)
		return
	}
	if celebrity.DateOfBirth.IsZero() {
		http.Error(w, "Missing dateOfBirth parameter", http.StatusBadRequest)
		return
	}
	if celebrity.PlaceOfBirth == "" {
		http.Error(w, "Missing placeOfBirth parameter", http.StatusBadRequest)
		return
	}
	if celebrity.Gender == "" {
		http.Error(w, "Missing gender parameter", http.StatusBadRequest)
		return
	}
	if celebrity.Bio == "" {
		http.Error(w, "Missing bio parameter", http.StatusBadRequest)
		return
	}
	resp, err := h.Celebrity.UpdateCelebrity(r.Context(), celebrity.ID, celebrity.Name, celebrity.Occupation, celebrity.PostersPath, celebrity.DateOfBirth, celebrity.DateOfDeath, celebrity.PlaceOfBirth, models.Gender(celebrity.Gender), celebrity.Bio)
	if err != nil {
		http.Error(w, "Error while updating celebrity", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UploadCelebrityPosters(w http.ResponseWriter, r *http.Request) {
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
	receivedCelebrityID, correct := r.Form["celebrityId"]
	if !correct || receivedCelebrityID[0] == "" {
		http.Error(w, "Missing celebrityId key or the value is empty", http.StatusBadRequest)
		return
	}

	resp, err := h.Celebrity.UploadCelebrityPosters(r.Context(), receivedCelebrityID[0], receivedImages)
	if err != nil {
		http.Error(w, "Error while uploading celebrity posters", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteCelebrityPoster(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	celebrityID := params["celebrityId"]
	image := params["image"]
	if len(celebrityID) == 0 || celebrityID[0] == "" {
		http.Error(w, "Missing celebrityId parameter", http.StatusBadRequest)
		return
	}
	if len(image) == 0 || image[0] == "" {
		http.Error(w, "Missing image parameter", http.StatusBadRequest)
		return
	}

	err := h.Celebrity.DeleteCelebrityPoster(r.Context(), celebrityID[0], image[0])
	if err != nil {
		http.Error(w, "Error while deleting celebrity poster", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListCelebrities(w http.ResponseWriter, r *http.Request) {
	response, err := h.Celebrity.ListCelebrities(r.Context())
	if err != nil {
		http.Error(w, "Error while listing all celebrities!", http.StatusInternalServerError)
		return
	}

	err = createResponse(w, response)
	if err != nil {
		http.Error(w, "Error while encoding response body!", http.StatusBadRequest)
		return
	}
}

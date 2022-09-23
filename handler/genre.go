package handler

import (
	"encoding/json"
	"int-gateway/request"
	"int-gateway/response"
	"net/http"
)

func (h *Handler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	genre := request.CreateGenre{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&genre); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if genre.Name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}
	if genre.Description == "" {
		http.Error(w, "Missing description parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Genre.CreateGenre(r.Context(), genre.Name, genre.Description)
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

func (h *Handler) GetGenre(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	genre := &response.Genre{}
	var err error
	if params["id"] != nil {
		idString := params["id"]
		if len(idString) == 0 || idString[0] == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}
		genre, err = h.Genre.GetGenre(r.Context(), idString[0])
		if err != nil {
			http.Error(w, "Error while getting genre by id", http.StatusInternalServerError)
			return
		}

	} else if params["name"] != nil {
		nameString := params["name"]
		if len(nameString) == 0 || nameString[0] == "" {
			http.Error(w, "Missing name parameter", http.StatusBadRequest)
			return
		}
		genre, err = h.Genre.GetGenreByName(r.Context(), nameString[0])
		if err != nil {
			http.Error(w, "Error while getting genre by name", http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, "Invalid parameter", http.StatusInternalServerError)
		return
	}
	err = createResponse(w, &genre)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	genre := request.UpdateGenre{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&genre); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if genre.ID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	if genre.Name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}
	if genre.Description == "" {
		http.Error(w, "Missing description parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Genre.UpdateGenre(r.Context(), genre.ID, genre.Name, genre.Description)
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

func (h *Handler) ListGenres(w http.ResponseWriter, r *http.Request) {
	response, err := h.Genre.ListGenres(r.Context())
	if err != nil {
		http.Error(w, "Error while listing all genres!", http.StatusInternalServerError)
		return
	}

	err = createResponse(w, response)
	if err != nil {
		http.Error(w, "Error while encoding response body!", http.StatusBadRequest)
		return
	}
}

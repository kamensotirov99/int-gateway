package handler

import (
	"encoding/json"
	"int-gateway/request"
	"int-gateway/response"

	"net/http"
)

func (h *Handler) CreateJournalist(w http.ResponseWriter, r *http.Request) {
	journalist := request.CreateJournalist{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&journalist); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if journalist.Name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Journalist.CreateJournalist(r.Context(), journalist.Name)
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

func (h *Handler) GetJournalist(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var journalist *response.Journalist
	var err error
	if params["id"] != nil {
		idString := params["id"]
		if len(idString) == 0 || idString[0] == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}
		journalist, err = h.Journalist.GetJournalist(r.Context(), idString[0])
		if err != nil {
			http.Error(w, "Error while getting journalist", http.StatusInternalServerError)
			return
		}

	} else if params["name"] != nil {
		nameString := params["name"]
		if len(nameString) == 0 || nameString[0] == "" {
			http.Error(w, "Missing name parameter", http.StatusBadRequest)
			return
		}
		journalist, err = h.Journalist.GetJournalistByName(r.Context(), nameString[0])
		if err != nil {
			http.Error(w, "Error while getting journalist by name", http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, "Invalid parameter", http.StatusInternalServerError)
		return
	}
	err = createResponse(w, journalist)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) ListJournalists(w http.ResponseWriter, r *http.Request) {
	journalists, err := h.Journalist.ListJournalists(r.Context())
	if err != nil {
		http.Error(w, "Error while reading the list of journalists", http.StatusInternalServerError)
		return
	}
	err = createResponse(w, journalists)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateJournalist(w http.ResponseWriter, r *http.Request) {
	journalist := request.UpdateJournalist{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&journalist); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if journalist.ID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	if journalist.Name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Journalist.UpdateJournalist(r.Context(), journalist.ID, journalist.Name)
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

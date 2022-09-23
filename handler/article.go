package handler

import (
	"encoding/json"
	"int-gateway/models"
	"int-gateway/request"
	"int-gateway/response"
	"net/http"
	"strconv"
)

func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	article := request.CreateArticle{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&article); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if article.Title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	if article.ReleaseDate.IsZero() {
		http.Error(w, "Missing releaseDate parameter", http.StatusBadRequest)
		return
	}
	if article.Description == "" {
		http.Error(w, "Missing description parameter", http.StatusBadRequest)
		return
	}
	if article.JournalistName == "" {
		http.Error(w, "Missing journalist name parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Article.CreateArticle(r.Context(), article.Title, article.ReleaseDate, article.PostersPath, article.Description, article.JournalistName)
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

func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idString := params["id"]
	if len(idString) == 0 || idString[0] == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	resp, err := h.Article.GetArticle(r.Context(), idString[0])
	if err != nil {
		http.Error(w, "Error while getting article", http.StatusInternalServerError)
		return
	}
	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := request.UpdateArticle{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&article); err != nil {
		http.Error(w, "Error while decoding request body!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if article.ID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	if article.Title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	if article.ReleaseDate.IsZero() {
		http.Error(w, "Missing releaseDate parameter", http.StatusBadRequest)
		return
	}
	if article.Description == "" {
		http.Error(w, "Missing description parameter", http.StatusBadRequest)
		return
	}
	if article.Journalist.ID == "" {
		http.Error(w, "Missing journalist id parameter", http.StatusBadRequest)
		return
	}

	resp, err := h.Article.UpdateArticle(r.Context(), article.ID, article.Title, article.ReleaseDate, article.PostersPath, article.Description, &models.Journalist{ID: article.Journalist.ID})
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

func (h *Handler) ListArticles(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var articles response.Articles
	var err error
	if params["journalistId"] != nil {
		journalistIdString := params["journalistId"]
		if len(journalistIdString) == 0 || journalistIdString[0] == "" {
			http.Error(w, "Missing journalist id parameter", http.StatusBadRequest)
			return
		}
		articles, err = h.Article.ListArticlesByJournalist(r.Context(), journalistIdString[0])
		if err != nil {
			http.Error(w, "Error while getting articles by journalist id", http.StatusInternalServerError)
			return
		}

	} else {

		if params["elementCount"] == nil {
			articles, err = h.Article.ListArticles(r.Context(), 0)
			if err != nil {
				http.Error(w, "Error while listing all articles!", http.StatusInternalServerError)
				return
			}
		} else {
			elementCount := params["elementCount"]
			var elementInt = 0
			elementInt, err = strconv.Atoi(elementCount[0])
			if err != nil {
				http.Error(w, "Error while converting elementCount!", http.StatusInternalServerError)
				return
			}
			articles, err = h.Article.ListArticles(r.Context(), elementInt)
			if err != nil {
				http.Error(w, "Error while listing all articles!", http.StatusInternalServerError)
				return
			}
		}

	}

	err = createResponse(w, articles)
	if err != nil {
		http.Error(w, "Error while encoding response body!", http.StatusBadRequest)
		return
	}
}

func (h *Handler) UploadArticlePosters(w http.ResponseWriter, r *http.Request) {
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
	receivedArticleID, correct := r.Form["articleId"]
	if !correct || receivedArticleID[0] == "" {
		http.Error(w, "Missing articleId key or the value is empty", http.StatusBadRequest)
		return
	}

	resp, err := h.Article.UploadArticlePosters(r.Context(), receivedArticleID[0], receivedImages)
	if err != nil {
		http.Error(w, "Error while uploading article posters", http.StatusBadRequest)
		return
	}

	err = createResponse(w, resp)
	if err != nil {
		http.Error(w, "Error while encoding the data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteArticlePoster(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	articleID := params["articleId"]
	image := params["image"]
	if len(articleID) == 0 || articleID[0] == "" {
		http.Error(w, "Missing articleID parameter", http.StatusBadRequest)
		return
	}
	if len(image) == 0 || image[0] == "" {
		http.Error(w, "Missing image parameter", http.StatusBadRequest)
		return
	}

	err := h.Article.DeleteArticlePoster(r.Context(), articleID[0], image[0])
	if err != nil {
		http.Error(w, "Error while deleting article posters ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

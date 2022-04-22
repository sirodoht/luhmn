package document

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type API struct {
	store  Store
	logger *zap.Logger
}

func NewAPI(store Store) *API {
	return &API{
		store: store,
	}
}

func (api *API) InsertHandler(w http.ResponseWriter, r *http.Request) {
	type ReqBody struct {
		Title string
		Body  string
	}
	decoder := json.NewDecoder(r.Body)
	var rb ReqBody
	err := decoder.Decode(&rb)
	if err != nil {
		panic(err)
	}

	if rb.Title == "" || rb.Body == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now()
	d := &Document{
		Title:     rb.Title,
		Body:      rb.Body,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err = api.store.Insert(r.Context(), d)
	if err != nil {
		panic(err)
	}
}

func (api *API) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	docs, err := api.store.GetAll(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		api.logger.With(
			zap.Error(err),
		).Error("failed to get all documents")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
	}
	w.Write(res)
}

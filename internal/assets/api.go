package assets

import (
	"assets/internal/api"
	serviceErrors "assets/internal/errors"
	"assets/pkg/pagination"
	"encoding/json"
	"io"
	"net/http"
)

type APIAssets interface {
	Upload(http.ResponseWriter, *http.Request)
	GetAsset(http.ResponseWriter, *http.Request)
	DeleteAsset(http.ResponseWriter, *http.Request)
	GetList(http.ResponseWriter, *http.Request)
}

type API struct {
	service ServiceAssets
}

func NewAPI(service ServiceAssets) APIAssets {
	return &API{service: service}
}

func (a *API) Upload(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		uid := r.Context().Value("uid").(uint32)

		assetName := r.PathValue("asset_name")

		assetData, err := io.ReadAll(r.Body)
		if err != nil {
			api.JSONResponseError(w, err)
			return
		}

		_, err = a.service.CreateAsset(r.Context(), uid, assetName, assetData)

		if err != nil {
			api.JSONResponseError(w, err)
			return
		}

		//Тут по-хорошему возвращать бы *entity.Asset из a.service.CreateAsset
		//но он в ответе не требуется
		data, err := json.Marshal(map[string]string{"status": "ok"})
		if err != nil {
			api.JSONResponseError(w, err)
			return
		}
		api.JSONResponse(w, data)

	default:
		api.JSONResponseError(w, serviceErrors.ErrorMethodNotAllowed)
		return
	}
}

func (a *API) GetAsset(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		assetName := r.PathValue("asset_name")
		uid := r.Context().Value("uid").(uint32)

		data, err := a.service.GetAsset(r.Context(), uid, assetName)

		if err != nil {
			api.JSONResponseError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)

	default:
		api.JSONResponseError(w, serviceErrors.ErrorMethodNotAllowed)
		return
	}
}

func (a *API) DeleteAsset(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		assetName := r.PathValue("asset_name")
		uid := r.Context().Value("uid").(uint32)

		deleted, err := a.service.DeleteAsset(r.Context(), uid, assetName)

		if err != nil {
			api.JSONResponseError(w, err)
			return
		}
		//
		data, err := json.Marshal(map[string]interface{}{"deleted": deleted})
		if err != nil {
			api.JSONResponseError(w, err)
			return
		}
		api.JSONResponse(w, data)

	default:
		api.JSONResponseError(w, serviceErrors.ErrorMethodNotAllowed)
		return
	}
}

func (a *API) GetList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		pages, err := a.service.GetList(r.Context(), pagination.NewFromRequest(r, -1))
		if err != nil {
			api.JSONResponseError(w, err)
			return
		}

		data, err := json.Marshal(pages)
		if err != nil {
			api.JSONResponseError(w, err)
			return
		}
		api.JSONResponse(w, data)
	default:
		api.JSONResponseError(w, serviceErrors.ErrorMethodNotAllowed)
		return
	}

}

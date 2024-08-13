package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simple-microservice/internal/model"
	"simple-microservice/internal/services"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type status string

const (
	statusSuccess status = "success"
	statusError   status = "error"
)

type structuredResponse struct {
	Data   interface{}       `json:"data"`
	Errors []structuredError `json:"errors"`
}

type structuredError struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg"`
}

type ApiResponse struct {
	Status    status          `json:"status"`
	Data      interface{}     `json:"data,omitempty"`
	ErrorType model.ErrorType `json:"errorType,omitempty"`
	Error     string          `json:"error,omitempty"`
}

type APIHandler struct {
	opts    APIHandlerOptions
	baseAPI BaseAPI
}

type BaseAPI struct {
	contextTimeOut time.Duration
}

type APIHandlerOptions struct {
	timeout time.Duration
}

func NewAPIHandler(opts APIHandlerOptions, baseAPI BaseAPI) (*APIHandler, error) {
	if opts.timeout == 0 || baseAPI.contextTimeOut == 0 {
		return nil, fmt.Errorf("incorrect api options")
	}
	aH := &APIHandler{opts: APIHandlerOptions{timeout: opts.timeout}, baseAPI: BaseAPI{baseAPI.contextTimeOut}}
	return aH, nil
}

func RespondError(w http.ResponseWriter, apiErr model.BaseApiError, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&ApiResponse{
		Status:    statusError,
		Data:      data,
		ErrorType: apiErr.Type(),
		Error:     apiErr.Error(),
	})
	if err != nil {
		http.Error(w, "Error Marshalling the json response", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		log.Fatalf("Error writing response:%s\n", b)
	}

	var code int
	switch apiErr.Type() {
	case model.ErrorBadData:
		code = http.StatusBadRequest
	case model.ErrorInternal:
		code = http.StatusInternalServerError
	case model.ErrorNone:
		code = http.StatusNotFound
	case model.ErrorTimeout:
		code = http.StatusRequestTimeout
	default:
		code = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func writeHttpResponse(w http.ResponseWriter, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&ApiResponse{
		Status: statusSuccess,
		Data:   data,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		log.Fatalf("Error writing response:%s\n", b)
	}

}

func (aH *APIHandler) HandleError(w http.ResponseWriter, err error, statusCode int) bool {
	if err == nil {
		return false
	}
	if statusCode == http.StatusInternalServerError {
		log.Fatalf("Error writing response")
	}
	structuredResp := structuredResponse{
		Errors: []structuredError{
			{
				Code: statusCode,
				Msg:  err.Error(),
			},
		},
	}
	resp, _ := json.Marshal(&structuredResp)
	http.Error(w, string(resp), statusCode)
	return true
}

func (aH *APIHandler) WriteJSON(w http.ResponseWriter, r *http.Request, response interface{}) {
	marshall := json.Marshal
	if prettyPrint := r.FormValue("pretty"); prettyPrint != "" && prettyPrint != "false" {
		marshall = func(v interface{}) ([]byte, error) {
			return json.MarshalIndent(v, "", "    ")
		}
	}
	resp, _ := marshall(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (aH *APIHandler) Respond(w http.ResponseWriter, data interface{}) {
	writeHttpResponse(w, data)
}

func (aH *APIHandler) AccByHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), aH.baseAPI.contextTimeOut)
	defer cancel()
	select {
	case <-time.After(aH.opts.timeout):
		RespondError(w, &model.ApiError{Typ: model.ErrorTimeout}, "Request Timed Out")
	case <-ctx.Done():
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal}, "Internal Server Error")
	default:
		healthStr := r.URL.Query().Get("health")
		health, err := strconv.Atoi(healthStr)
		if err != nil {
			RespondError(w, &model.ApiError{Typ: model.ErrorBadData}, "Invalid Health")
			return
		}
		var filteredAccounts []model.AccountInfo
		for _, account := range services.Accounts {
			if (account.AccHealth) == model.AccountHealth(health) {
				filteredAccounts = append(filteredAccounts, account)
			}
		}
		if len(filteredAccounts) > 0 {
			writeHttpResponse(w, filteredAccounts)
			return
		}

		RespondError(w, &model.ApiError{Typ: model.ErrorNotFound}, "Account Not found")

	}

}

func (aH *APIHandler) AccByType(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), aH.baseAPI.contextTimeOut)
	defer cancel()
	select {
	case <-time.After(aH.opts.timeout):
		RespondError(w, &model.ApiError{Typ: model.ErrorTimeout}, "Request Timed Out")
	case <-ctx.Done():
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal}, "Internal Server Error")
	default:
		typeStr := r.URL.Query().Get("type")
		accType, err := strconv.Atoi(typeStr)
		if err != nil {
			RespondError(w, &model.ApiError{Typ: model.ErrorBadData}, "Invalid Type")
			return
		}
		var filteredAccounts []model.AccountInfo
		for _, account := range services.Accounts {
			if (account.AccType) == model.AccountType(accType) {
				filteredAccounts = append(filteredAccounts, account)
			}
		}
		if len(filteredAccounts) > 0 {
			writeHttpResponse(w, filteredAccounts)
			return
		}

		RespondError(w, &model.ApiError{Typ: model.ErrorNotFound}, "Account Not found")

	}

}

func (aH *APIHandler) getAllAcc(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), aH.baseAPI.contextTimeOut)
	defer cancel()
	select {
	case <-time.After(aH.opts.timeout):
		RespondError(w, &model.ApiError{Typ: model.ErrorTimeout}, "Request Timed Out")
	case <-ctx.Done():
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal}, "Internal Server Error")
	default:
		var allAccounts = services.Accounts
		if len(allAccounts) > 0 {
			writeHttpResponse(w, services.Accounts)
			return
		}

		RespondError(w, &model.ApiError{Typ: model.ErrorNotFound}, "No Accounts Found")

	}

}

func (aH *APIHandler) Home(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), aH.baseAPI.contextTimeOut)
	defer cancel()
	select {
	case <-time.After(aH.opts.timeout):
		RespondError(w, &model.ApiError{Typ: model.ErrorTimeout}, "Request Timed Out")
	case <-ctx.Done():
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal}, "Internal Server Error")
	default:

		writeHttpResponse(w, "Welcome to Home page")

	}

}

func setupRoutes() *http.ServeMux {
	opts := APIHandlerOptions{
		timeout: 60 * time.Second,
	}
	baseAPI := BaseAPI{contextTimeOut: 60 * time.Second}
	aH, err := NewAPIHandler(opts, baseAPI)
	if err != nil {
		log.Printf("Error setting up routing: %v", err)
		return http.NewServeMux()
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", aH.Home)
	mux.HandleFunc("/accountsByHealth", aH.AccByHealth)
	mux.HandleFunc("/accountsByType", aH.AccByType)
	mux.HandleFunc("/getAll", aH.getAllAcc)
	return mux
}

func SetupHTTPHandlers() {
	setupRoutes()
}

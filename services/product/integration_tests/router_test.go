package integrationtests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/S6-Wallmarkt/Wallmarkt/services/product/api"
)

func TestGetAllRoute_Success(t *testing.T) {
	router := api.SetupRouter()

	// Create a new request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getall", nil)
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)

	// Check if the right content is returned
	expectedBody := `[{"id":1,"name":"Candle","description":"Light up rooms with this candle","price":9.95,"color":"White","types":["Lighting","Decorative"]},{"id":2,"name":"Wall-torch","description":"Light up rooms with this wall-torch","price":25.5,"color":"Steel","types":["Lighting","Decorative"]},{"id":3,"name":"Basket","description":"Store items in a basket","price":5,"color":"Wood","types":["Decorative","Storage"]},{"id":4,"name":"Star-lamp","description":"Light up rooms with this wall mounted star lamp","price":14.95,"color":"White","types":["Lighting","Decorative"]},{"id":5,"name":"Deer-mount","description":"Creepy deer head mount for your wall","price":999.99,"color":"Brown","types":["Creepy","Decorative"]}]`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetByIDRoute_Success(t *testing.T) {
	router := api.SetupRouter()

	// Create a new request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getbyid/1", nil)
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)

	// Check if the right content is returned
	expectedBody := `{"id":1,"name":"Candle","description":"Light up rooms with this candle","price":9.95,"color":"White","types":["Lighting","Decorative"]}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetByIDRoute_InvalidID(t *testing.T) {
	router := api.SetupRouter()

	// Create a new request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getbyid/notanID", nil)
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, w.Body)

	// Check if the right content is returned
	expectedBody := `{"error":"Invalid ID"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetByIDRoute_NonExistingID(t *testing.T) {
	router := api.SetupRouter()

	// Create a new request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getbyid/6", nil)
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.NotEmpty(t, w.Body)

	// Check if the right content is returned
	expectedBody := `{"error":"Product not found"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetByTypeRoute_Success(t *testing.T) {
	router := api.SetupRouter()

	// Create a new request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getbytype/Lighting", nil)
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)

	// Check if the right content is returned
	expectedBody := `[{"id":1,"name":"Candle","description":"Light up rooms with this candle","price":9.95,"color":"White","types":["Lighting","Decorative"]},{"id":2,"name":"Wall-torch","description":"Light up rooms with this wall-torch","price":25.5,"color":"Steel","types":["Lighting","Decorative"]},{"id":4,"name":"Star-lamp","description":"Light up rooms with this wall mounted star lamp","price":14.95,"color":"White","types":["Lighting","Decorative"]}]`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetByTypeRoute_NonExistingType(t *testing.T) {
	router := api.SetupRouter()

	// Create a new request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getbytype/notatype", nil)
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.NotEmpty(t, w.Body)

	// Check if the right content is returned
	expectedBody := `{"error":"Products not found"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestAddRoute_Success(t *testing.T) {
	router := api.SetupRouter()

	// Create a new request with a JSON body
	jsonBody := `{"id":6,"name":"TestItem","description":"This is a test item added by the add endpoint","price":12.34,"color":"blank","types":["Test"]}`
	reqBody := strings.NewReader(jsonBody)

	// Create a new request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/add", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, w.Body)

	// Check if the right content is returned
	expectedBody := `{"id":6,"name":"TestItem","description":"This is a test item added by the add endpoint","price":12.34,"color":"blank","types":["Test"]}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestAddRoute_NotValidBody(t *testing.T) {
	router := api.SetupRouter()

	// Create a new request with a JSON body
	jsonBody := `{"notvalid":"body"}`
	reqBody := strings.NewReader(jsonBody)

	// Create a new request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/add", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, w.Body)

	// Check if the right content is returned
	expectedBody := `{"error":"Invalid JSON data"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestAddRoute_AlreadyExistingID(t *testing.T) {
	router := api.SetupRouter()

	// Create a new request with a JSON body
	jsonBody := `{"id":3,"name":"TestItem","description":"This is a test item added by the add endpoint","price":12.34,"color":"blank","types":["Test"]}`
	reqBody := strings.NewReader(jsonBody)

	// Create a new request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/add", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, w.Body)

	// Check if the right content is returned
	expectedBody := `{"error":"product already exists"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

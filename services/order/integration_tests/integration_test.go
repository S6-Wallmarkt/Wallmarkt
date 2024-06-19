package integrationtests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/S6-Wallmarkt/Wallmarkt/services/order/api"
	"github.com/S6-Wallmarkt/Wallmarkt/services/order/configs"
	"github.com/stretchr/testify/assert"
)

func TestGetAllRoute_Success(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/getall", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)
}

func TestAddRoute_InvalidJSON(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"notvalid":"body"}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/add", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := `{"error":"Invalid JSON data"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestAddRoute_NoProducts(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"customer_id": "facebook12345", "payed": true, "products": []}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/add", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := `{"error":"No products provided"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestAddRoute_Success(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"customer_id": "integrationtest", "payed": true, "products": ["6651d1aa02fb7f8a4414e17d"]}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/add", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, w.Body)
}

func TestGetByIDRoute_NonExisting(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/getbyid/6656f34ace0c27ef71d06999", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
	expectedBody := `{"error": "order not found"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetByIDRoute_NoID(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/getbyid/", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetByIDRoute_Success(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/getbyid/6656f380ce0c27ef71d06981", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"customer_id":"auth0|66538aa12019c08e25cb17b5", "id":"6656f380ce0c27ef71d06981", "payed":true, "products":["6651d1aa02fb7f8a4414e17d", "6651d1ec02fb7f8a4414e17f", "6651d22c02fb7f8a4414e183"], "shipped":false}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetByCustomerRoute_NonExisting(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/getbycustomer/some-customer", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
	expectedBody := `{"error":"No orders found"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetByCustomerRoute_Success(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/getbycustomer/auth0|66538aa12019c08e25cb17b5", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)
}

func TestDeleteRoute_Success(t *testing.T) {
	// Arrange
	// Set up database connection
	mongodbUri := "mongodb://localhost:27017"
	configs.InitMongoDB(mongodbUri)
	configs.InitCollections()
	router := api.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("DELETE", "/delete/6656f49ace0c27ef71d06984", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `"Deleted order: 6656f49ace0c27ef71d06984"`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

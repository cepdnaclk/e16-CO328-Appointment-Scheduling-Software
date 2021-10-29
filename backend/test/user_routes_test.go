package test

import (
	"appoiment-backend/controllers"
	"appoiment-backend/database"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert" // add Testify package
)


func TestHelloRoute(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
	  description  string // description of the test case
	  route        string // route path to test
	  expectedCode int    // expected HTTP status code
	}{
	  // First test case
	  {
		description:  "get HTTP status 200",
		route:        "/hello",
		expectedCode: 200,
	  },
	  // Second test case
	  {
		description:  "get HTTP status 404, when route is not exists",
		route:        "/not-found",
		expectedCode: 404,
	  },
	}
	
	// Define Fiber app.
	app := fiber.New()
	// Create route with GET method for test
	app.Get("/hello", func(c *fiber.Ctx) error {
	  // Return simple string as response
	  return c.SendString("Hello, World!")
	})
  
	// Iterate through test single test cases
	for _, test := range tests {
	  // Create a new http request with the route from the test case
	  req := httptest.NewRequest("GET", test.route, nil)
  
	  // Perform the request plain with the app,
	  // the second argument is a request latency
	  // (set to -1 for no latency)
	  resp, _ := app.Test(req, 1)
  
	  // Verify, if the status code is as expected
	  assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}


func TestLoginRoute(t *testing.T){
	tests := []struct {
		description  string 
		route        string 
		expectedCode int 
		body []byte
	  }{
		// First test case
		{
		  description:  "get HTTP status 200",
		  route:        "/login",
		  expectedCode: 200,
		  body :  []byte(`{"email": "randikavirajmax@gmail.com", "password": "anjalee"}`),
		},
		// Second test case
		{
		  description:  "get HTTP status 409, when body is not passes",
		  route:        "/login",
		  expectedCode: 409,
		  body: nil,
		},
		{
			description:  "get HTTP status 409, when password is wrong",
			route:        "/login",
			expectedCode: 409,
			body : []byte(`{"email": "randikavirajmax@gmail.com", "password": "wrong"}`),
		},
	  }

	database.Connect("mongodb://localhost:27017")
	defer database.Disconnect()
	
	app := fiber.New()

	app.Post("/login",controllers.Login)

	for _, test := range tests {
		form := url.Values{}
		form.Set("email", "randikavirajmax@gmail.com") 
		form.Set("password", "anjalee") 
		req := httptest.NewRequest("POST", test.route, strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		resp, _ := app.Test(req, 1)
	
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		
	}
}
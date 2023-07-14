package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jarcoal/httpmock"
	"github.com/trunov/erply-assignement-task/user-service/internal/domain"
	"github.com/trunov/erply-assignement-task/user-service/internal/repository/erply"
)

var customerData = domain.CustomerInput{
	FirstName: "John",
	LastName:  "Doe",
	Email:     "johndoe@example.com",
	Phone:     "+372555555",
	TwitterID: "john_doe",
}

var customerDBData = erply.Customer{
	ID:                   7,
	CustomerID:           7,
	FullName:             "Doe, John",
	CompanyTypeID:        0,
	FirstName:            "John",
	LastName:             "Doe",
	PersonTitleID:        0,
	EmailEnabled:         1,
	GroupID:              14,
	CountryID:            "0",
	Phone:                "+372555555",
	Email:                "johndoe@example.com",
	Birthday:             "0000-00-00",
	GroupName:            "Vaikimisi grupp",
	CustomerType:         "PERSON",
	LastModifierUsername: "kirill.trunov",
	TwitterID:            "john_doe",
}

var erplySampleCustomerData = erply.Customer{
	ID:            123,
	CustomerID:    123,
	FullName:      "Test",
	CompanyName:   "Test",
	CompanyTypeID: 15,
	FirstName:     "Test",
	LastName:      "Test",
	PersonTitleID: 123,
}

func mockPostRequestHandler(req *http.Request) (*http.Response, error) {
	sessionKey := "123123"

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	// Define the expected payload for verifyUser request
	authPayload := url.Values{
		"username":        {Username},
		"password":        {Password},
		"request":         {"verifyUser"},
		"sendContentType": {"1"},
		"clientCode":      {ClientCode},
	}

	addCustomerPayload := url.Values{
		"sessionKey":      {sessionKey},
		"request":         {"saveCustomer"},
		"clientCode":      {ClientCode},
		"firstName":       {customerData.FirstName},
		"lastName":        {customerData.LastName},
		"email":           {customerData.Email},
		"phone":           {customerData.Phone},
		"twitterID":       {customerData.TwitterID},
		"sendContentType": {"1"},
	}

	if string(body) == authPayload.Encode() {
		responseBody := erply.GetVerifyUserResponse{
			Status: erply.Status{
				Request:           "verifyUser",
				RequestUnixTime:   1689260537,
				ResponseStatus:    "ok",
				ErrorCode:         0,
				GenerationTime:    0.4205901622772217,
				RecordsTotal:      1,
				RecordsInResponse: 1,
			},
			Records: []erply.UserInfo{{
				UserID:       "5",
				EmployeeName: "Kirill",
				SessionKey:   sessionKey,
			}},
		}
		responseJSON, _ := json.Marshal(responseBody)

		return httpmock.NewStringResponse(http.StatusOK, string(responseJSON)), nil
	}

	if string(body) == addCustomerPayload.Encode() {
		responseBody := erply.SaveCustomerResponse{
			Status: erply.Status{
				ResponseStatus: "ok",
			},
			Records: []erply.SaveCustomerRecord{{}},
		}
		responseJSON, _ := json.Marshal(responseBody)

		return httpmock.NewStringResponse(http.StatusOK, string(responseJSON)), nil
	}

	responseBody := erply.GetCustomerResponse{
		Status: erply.Status{
			Request:           "getCusomers",
			RequestUnixTime:   1689260537,
			ResponseStatus:    "ok",
			ErrorCode:         0,
			GenerationTime:    0.4205901622772217,
			RecordsTotal:      1,
			RecordsInResponse: 1,
		},
		Records: []erply.Customer{erplySampleCustomerData},
	}
	responseJSON, _ := json.Marshal(responseBody)

	return httpmock.NewStringResponse(http.StatusOK, string(responseJSON)), nil
}

func (s *TestSuite) TestSaveCustomer() {
	request := domain.CustomerInput{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Phone:     "+372555555",
		TwitterID: "john_doe",
	}

	url := fmt.Sprintf("https://%s.erply.com/api/", ClientCode)

	buf := bytes.NewBufferString("")
	err := json.NewEncoder(buf).Encode(request)
	s.Require().NoError(err)

	httpmock.RegisterResponder(
		http.MethodPost,
		url,
		mockPostRequestHandler,
	)

	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/customer", buf)
	s.Require().NoError(err)

	// Set the Authorization header with the bearer token
	token := Auth
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := s.server.Client().Do(req)
	s.Require().NoError(err)

	defer res.Body.Close()

	s.Require().Equal(http.StatusCreated, res.StatusCode)
}

func (s *TestSuite) TestGettingCustomerFromErply() {
	url := fmt.Sprintf("https://%s.erply.com/api/", ClientCode)

	httpmock.RegisterResponder(
		http.MethodPost,
		url,
		mockPostRequestHandler,
	)

	req, err := http.NewRequest(http.MethodGet, s.server.URL+"/customer/1", nil)
	s.Require().NoError(err)

	// Set the Authorization header with the bearer token
	token := Auth
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := s.server.Client().Do(req)
	s.Require().NoError(err)

	defer res.Body.Close()

	var respCustomer erply.Customer
	err = json.NewDecoder(res.Body).Decode(&respCustomer)
	s.Require().NoError(err)

	s.Require().Equal(http.StatusOK, res.StatusCode)
	s.Require().Equal(erplySampleCustomerData, respCustomer)
}

func (s *TestSuite) TestGettingCustomerFromCache() {
	url := fmt.Sprintf("https://%s.erply.com/api/", ClientCode)

	httpmock.RegisterResponder(
		http.MethodPost,
		url,
		mockPostRequestHandler,
	)

	// id number 7 is defined in fixtures
	req, err := http.NewRequest(http.MethodGet, s.server.URL+"/customer/7", nil)
	s.Require().NoError(err)

	token := Auth
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := s.server.Client().Do(req)
	s.Require().NoError(err)

	defer res.Body.Close()

	var respCustomer erply.Customer
	err = json.NewDecoder(res.Body).Decode(&respCustomer)
	s.Require().NoError(err)

	s.Require().Equal(http.StatusOK, res.StatusCode)

	s.Require().Equal(customerDBData.ID, respCustomer.ID)
	s.Require().Equal(customerDBData.CustomerID, respCustomer.CustomerID)
	s.Require().Equal(customerDBData.FullName, respCustomer.FullName)
}

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

	return httpmock.NewStringResponse(http.StatusOK, ""), nil
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

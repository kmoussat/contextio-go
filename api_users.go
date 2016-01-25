package ciolite

// Api functions that support: https://context.io/docs/lite/users

// Imports
import (
	"fmt"
)

// GetUsersResponse ...
type GetUsersResponse struct {
	ID             string   `json:"id,omitempty"`
	Username       string   `json:"username,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	FirstName      string   `json:"first_name,omitempty"`
	LastName       string   `json:"last_name,omitempty"`
	ResourceURL    string   `json:"resource_url,omitempty"`

	EmailAccounts []GetUsersEmailAccountsResponse `json:"email_accounts,omitempty"`

	Created         int `json:"created,omitempty"`
	Suspended       int `json:"suspended,omitempty"`
	PasswordExpired int `json:"password_expired,omitempty"`
}

// CreateUserResponse ...
type CreateUserResponse struct {
	Success bool   `json:"success,omitempty"`
	ID      string `json:"id,omitempty"`

	EmailAccount CreateEmailAccountResponse `json:"email_account,omitempty"`

	ResourceURL       string `json:"resource_url,omitempty"`
	AccessToken       string `json:"access_token,omitempty"`
	AccessTokenSecret string `json:"access_token_secret,omitempty"`
}

// ModifyUserResponse ...
type ModifyUserResponse struct {
	Success     bool   `json:"success,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// DeleteUserResponse ...
type DeleteUserResponse struct {
	Success     bool   `json:"success,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// GetUsers gets a list of users.
// queryValues may optionally contain CioParams.Email, CioParams.Status,
// CioParams.StatusOK, CioParams.Limit, CioParams.Offset
// 	https://context.io/docs/lite/users#get
func (cioLite *CioLite) GetUsers(queryValues CioParams) ([]GetUsersResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        "/users",
		queryValues: queryValues,
	}

	// Make response
	var response []GetUsersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUser get details about a given user.
// 	https://context.io/docs/lite/users#id-get
func (cioLite *CioLite) GetUser(userID string) (GetUsersResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s", userID),
	}

	// Make response
	var response GetUsersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUser create a new user.
// formValues may optionally contain CioParams.MigrateAccountID, CioParams.Email,
// CioParams.FirstName, CioParams.LastName, CioParams.Server, CioParams.Username,
// CioParams.UseSSL, CioParams.Port, CioParams.Type, CioParams.Password,
// CioParams.ProviderRefreshToken, CioParams.ProviderConsumerKey, CioParams.StatusCallbackURL
// 	https://context.io/docs/lite/users#post
func (cioLite *CioLite) CreateUser(formValues CioParams) (CreateUserResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       "/users",
		formValues: formValues,
	}

	// Make response
	var response CreateUserResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ModifyUser modifies a given user.
// formValues requires CioParams.FirstName, CioParams.LastName
// 	https://context.io/docs/lite/users#id-post
func (cioLite *CioLite) ModifyUser(userID string, formValues CioParams) (ModifyUserResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s", userID),
		formValues: formValues,
	}

	// Make response
	var response ModifyUserResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUser removes a given user.
// 	https://context.io/docs/lite/users#id-delete
func (cioLite *CioLite) DeleteUser(userID string) (DeleteUserResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/users/%s", userID),
	}

	// Make response
	var response DeleteUserResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

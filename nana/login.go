package nana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Nana ...
type Nana struct {
	ID           int
	Name         string
	Email        string
	password     string
	Token        string
	deviceID     string
	appsflyerID  string
	endpointBase *url.URL
	httpClient   *http.Client
}

// init ...
func (n *Nana) init() error {
	n.deviceID = strings.ToUpper(geneDeviceID())
	n.appsflyerID = strings.ToUpper(geneAppsflyerID())
	u, err := url.ParseRequestURI(APIEndpointBase)
	if err != nil {
		return err
	}
	n.endpointBase = u
	n.httpClient = http.DefaultClient
	return nil
}

// create ...
func (n *Nana) create() (*CreateResponse, error) {
	endpoint := APILegacyVersion + APIEndpointSignup
	params := struct {
		DeveiceID  string `json:"device_id"`
		ScreenName string `json:"screen_name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Type       string `json:"type"`
	}{
		n.deviceID,
		n.Name,
		n.Email,
		n.password,
		"nana",
	}
	paramsJSON, _ := json.Marshal(params)
	res, err := n.post(nil, endpoint, bytes.NewBuffer(paramsJSON))
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeCreateResponse(res)
}

// mailLogin ...
func (n *Nana) mailLogin() (*LoginResponse, error) {
	endpoint := APILegacyVersion + APIEndpointLogin
	params := struct {
		DeveiceID string `json:"device_id"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Type      string `json:"type"`
	}{
		n.deviceID,
		n.Email,
		n.password,
		"nana",
	}
	paramsJSON, _ := json.Marshal(params)
	res, err := n.post(nil, endpoint, bytes.NewBuffer(paramsJSON))
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeLoginResponse(res)
}

// UpdateUserInfo ...
func (n *Nana) UpdateUserInfo() (*Nana, error) {
	res, err := n.GetMyInfo()
	if err != nil {
		return nil, err
	}
	n.ID = res.UserID
	n.Name = res.ScreenName
	return n, nil
}

// CreateAccount ...
func CreateAccount(name string, password string) (*Nana, error) {
	if name == "" {
		name = randString(5)
	}
	if password == "" {
		password = randString(10)
	}
	email := fmt.Sprintf("%s@email.com", name)
	n := Nana{
		Name:     name,
		Email:    email,
		password: password,
	}
	if err := n.init(); err != nil {
		return nil, err
	}
	res, err := n.create()
	if err != nil {
		return nil, err
	}
	n.Token = res.Data.Token
	return n.UpdateUserInfo()
}

// Login ...
func Login(email string, password string) (*Nana, error) {
	n := Nana{
		Email:    email,
		password: password,
	}
	if err := n.init(); err != nil {
		return nil, err
	}
	res, err := n.mailLogin()
	if err != nil {
		return nil, err
	}
	n.Token = res.Data.Token
	return n.UpdateUserInfo()
}

// LoginByToken ...
func LoginByToken(token string) (*Nana, error) {
	n := Nana{
		Token: token,
	}
	if err := n.init(); err != nil {
		return nil, err
	}
	return n.UpdateUserInfo()
}

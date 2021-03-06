package nextengine

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"nextengine-sdk-go/entity"
	"nextengine-sdk-go/token"
)

const (
	authHost = "https://base.next-engine.org"
	apiHost  = "https://api.next-engine.org"
	// Success API Result
	Success = "success"
	// Error API Result
	Error = "error"
	// Redirect API Result
	Redirect = "redirect"
)

// TokenRepository is API token write/read interface
// If you want to change the storage location of API token to DB or session, you need to implement this interface
type TokenRepository interface {
	Token(context.Context) (*entity.Token, error)
	Save(context.Context, *entity.Token) error
}

// TokenGetter is interface
type TokenGetter interface {
	TokenValue() *entity.Token
}

// Client is Structure holding the settings of NextEngine API client
type Client struct {
	clientID        string
	clientSecret    string
	redirectURI     string
	httpClient      *http.Client
	TokenRepository TokenRepository
}

// NewDefaultClient Returns an instance of the API client with default settings
func NewDefaultClient(clientID string, clientSecret string, redirectURI string, accessToken string, refreshToken string) *Client {
	cli := &http.Client{}
	tr := token.NewMemoryTokenRepository(accessToken, refreshToken)
	return NewClient(clientID, clientSecret, redirectURI, cli, tr)
}

// NewClient Returns an instance of the API client
func NewClient(clientID string, clientSecret string, redirectURI string, httpClient *http.Client, tr TokenRepository) *Client {
	return &Client{
		clientID:        clientID,
		clientSecret:    clientSecret,
		redirectURI:     redirectURI,
		httpClient:      httpClient,
		TokenRepository: tr,
	}
}

// AuthURI Returns the URI of the authentication screen of Nexe Engine
func (c *Client) AuthURI(extraParam url.Values) string {
	v := url.Values{
		"client_id":    []string{c.clientID},
		"redirect_uri": []string{c.redirectURI},
	}
	for key, vals := range extraParam {
		for _, val := range vals {
			v.Add(key, val)
		}
	}

	u, _ := url.Parse(authHost + "/users/sign_in/")
	u.RawQuery = v.Encode()

	return u.String()
}

func newRequest(ctx context.Context, method string, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

// Authorize Fetch API token using uid and state
func (c *Client) Authorize(ctx context.Context, uid string, state string) (*entity.Authorize, error) {
	v := url.Values{
		"client_id":     []string{c.clientID},
		"client_secret": []string{c.clientSecret},
		"uid":           []string{uid},
		"state":         []string{state},
	}
	model := &entity.Authorize{}
	err := c.request(ctx, "/api_neauth", v, model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

// APIExecute is Execute the API and return the result
// Please specify a path starting with / for endpoint
func (c *Client) APIExecute(ctx context.Context, endpoint string, params map[string]string, model TokenGetter) error {
	tok, err := c.TokenRepository.Token(ctx)
	if err != nil {
		return err
	}

	v := url.Values{
		"access_token":  []string{tok.AccessToken},
		"refresh_token": []string{tok.RefreshToken},
	}

	return c.apiRequest(ctx, endpoint, v, params, model)
}

// APIExecuteNoRequiredLogin is Execute API that does not require login and return the result
// Please specify a path starting with / for endpoint
func (c *Client) APIExecuteNoRequiredLogin(ctx context.Context, endpoint string, params map[string]string, model TokenGetter) error {
	v := url.Values{
		"client_id":     []string{c.clientID},
		"client_secret": []string{c.clientSecret},
	}

	return c.apiRequest(ctx, endpoint, v, params, model)
}

func (c *Client) apiRequest(ctx context.Context, endpoint string, v url.Values, params map[string]string, model TokenGetter) error {
	for key, val := range params {
		v.Add(key, val)
	}

	err := c.request(ctx, endpoint, v, model)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) request(ctx context.Context, endpoint string, params url.Values, model TokenGetter) error {
	u, _ := url.Parse(apiHost + endpoint)

	httpRequest, err := newRequest(ctx, "POST", u.String(), bytes.NewBufferString(params.Encode()))
	if err != nil {
		return err
	}

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()
	return c.responseHandler(ctx, httpResponse.Body, model)
}

func (c *Client) responseHandler(ctx context.Context, body io.ReadCloser, model TokenGetter) error {
	if err := json.NewDecoder(body).Decode(model); err != nil {
		return err
	}

	tok := model.TokenValue()

	if tok.AccessToken != "" && tok.RefreshToken != "" {
		if err := c.TokenRepository.Save(ctx, tok); err != nil {
			return err
		}
	}

	return nil
}

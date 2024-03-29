package nextengine

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/takaaki-s/nextengine-sdk-go/entity"
	"github.com/takaaki-s/nextengine-sdk-go/repository"
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

// TokenReaderWriter provides read & write access to tokens
// If you want to change the storage location of API token to DB or session, you need to implement this interface
type TokenReaderWriter interface {
	TokenReader
	TokenWriter
}

// TokenReader provides retrieval for access tokens
type TokenReader interface {
	Read(context.Context) (*entity.Token, error)
}

// TokenWriter provides storage for access tokens
type TokenWriter interface {
	Write(context.Context, *entity.Token) error
}

// Client is Structure holding the settings of NextEngine API client
type Client struct {
	clientID        string
	clientSecret    string
	redirectURI     string
	httpClient      *http.Client
	TokenRepository TokenReaderWriter
}

// NewDefaultClient Returns an instance of the API client with default settings
func NewDefaultClient(clientID, clientSecret, redirectURI, accessToken, refreshToken string) *Client {
	cli := &http.Client{}
	tr := repository.NewMemoryTokenRepository(accessToken, refreshToken)
	return NewClient(clientID, clientSecret, redirectURI, cli, tr)
}

// NewClient Returns an instance of the API client
func NewClient(clientID, clientSecret, redirectURI string, httpClient *http.Client, tr TokenReaderWriter) *Client {
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

	u, _ := url.Parse(authHost)
	u.Path = path.Join(u.Path, "users", "sign_in")
	u.Path = u.Path + "/"
	u.RawQuery = v.Encode()

	return u.String()
}

func newRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

// Authorize Fetch API token using uid and state
func (c *Client) Authorize(ctx context.Context, uid, state string) (*entity.Authorize, error) {
	v := url.Values{
		"client_id":     []string{c.clientID},
		"client_secret": []string{c.clientSecret},
		"uid":           []string{uid},
		"state":         []string{state},
	}
	auth := &entity.Authorize{}
	err := c.request(ctx, "/api_neauth", v, auth)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// APIExecute is Execute the API and return the result
// Please specify a path starting with / for endpoint
func (c *Client) APIExecute(ctx context.Context, endpoint string, params map[string]string, entity TokenReader) error {
	tok, err := c.TokenRepository.Read(ctx)
	if err != nil {
		return err
	}

	v := url.Values{
		"access_token":  []string{tok.AccessToken},
		"refresh_token": []string{tok.RefreshToken},
	}

	return c.apiRequest(ctx, endpoint, v, params, entity)
}

// APIExecuteNoRequiredLogin is Execute API that does not require login and return the result
// Please specify a path starting with / for endpoint
func (c *Client) APIExecuteNoRequiredLogin(ctx context.Context, endpoint string, params map[string]string, entity TokenReader) error {
	v := url.Values{
		"client_id":     []string{c.clientID},
		"client_secret": []string{c.clientSecret},
	}

	return c.apiRequest(ctx, endpoint, v, params, entity)
}

func (c *Client) apiRequest(ctx context.Context, endpoint string, v url.Values, params map[string]string, entity TokenReader) error {
	for key, val := range params {
		v.Add(key, val)
	}

	err := c.request(ctx, endpoint, v, entity)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) request(ctx context.Context, endpoint string, params url.Values, entity TokenReader) error {
	u, _ := url.Parse(apiHost)
	u.Path = path.Join(u.Path, endpoint)

	httpRequest, err := newRequest(ctx, http.MethodPost, u.String(), bytes.NewBufferString(params.Encode()))
	if err != nil {
		return err
	}

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	if err := json.NewDecoder(httpResponse.Body).Decode(entity); err != nil {
		return err
	}

	tok, err := entity.Read(ctx)
	if err != nil {
		return err
	}

	if tok.AccessToken != "" && tok.RefreshToken != "" {
		if err := c.TokenRepository.Write(ctx, tok); err != nil {
			return err
		}
	}

	return nil
}

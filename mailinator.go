package mailinator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"time"
)

// Client .
type Client struct {
	apiToken   string
	baseURL    string
	HTTPClient *http.Client
}

// NewMailinatorClient creates new Mailinator client with given API Token
func NewMailinatorClient(apiToken string) *Client {
	return &Client{
		apiToken: apiToken,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		baseURL: "https://api.mailinator.com/api/v2",
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type successResponseWithContent struct {
	Bytes       []byte `json:"bytes"`
	ContentType string `json:"content-type"`
	FileName    string `json:"filename"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	return c.sendRequestWithOptions(req, v, false)
}

func (c *Client) sendRequestWithOptions(req *http.Request, v interface{}, returnBody bool) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	// Check if apiToken is provided before setting Authorization header
	if c.apiToken != "" {
		req.Header.Set("Authorization", c.apiToken)
	}

	// Set User-Agent header
	req.Header.Set("User-Agent", "Mailinator SDK - Go V1.0")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if returnBody {
		responseBodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		if len(responseBodyBytes) == 0 {
			return nil
		}

		*v.(*string) = string(responseBodyBytes)

		return nil
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.Unmarshal(responseBody, &errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")

	if contentType != "application/json" {
		disposition, params, err := mime.ParseMediaType(res.Header.Get("Content-Disposition"))

		if err != nil {
			return err
		}

		if disposition == "attachment" {
			filename := params["filename"]

			responseWithContent := successResponseWithContent{
				Bytes:       responseBody,
				ContentType: contentType,
				FileName:    filename,
			}

			jsonRes, err := json.Marshal(responseWithContent)

			if err != nil {
				return err
			}

			if err = json.Unmarshal(jsonRes, &v); err != nil {
				return err
			}
		}

		return nil
	}

	if err = json.Unmarshal(responseBody, &v); err != nil {
		return err
	}

	return nil
}

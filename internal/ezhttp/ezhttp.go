package ezhttp

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/topisenpai/gobin/gobin"
)

var defaultClient = &http.Client{
	Timeout: 10 * time.Second,
}

func Do(method string, path string, token string, body io.Reader) (*http.Response, error) {
	server := viper.GetString("server")
	request, err := http.NewRequest(method, server+path, body)
	if err != nil {
		return nil, err
	}
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}
	return defaultClient.Do(request)
}

func Get(path string) (*http.Response, error) {
	return Do(http.MethodGet, path, "", nil)
}

func Post(path string, body io.Reader) (*http.Response, error) {
	return Do(http.MethodPost, path, "", body)
}

func Patch(path string, token string, body io.Reader) (*http.Response, error) {
	return Do(http.MethodPatch, path, token, body)
}

func Delete(path string, token string) (*http.Response, error) {
	return Do(http.MethodDelete, path, token, nil)
}

func ProcessBody(cmd *cobra.Command, method string, rs *http.Response, body any) bool {
	if rs.StatusCode >= 200 && rs.StatusCode <= 299 {
		if err := json.NewDecoder(rs.Body).Decode(body); err != nil {
			cmd.PrintErrln("Failed to decode response:", err)
			return false
		}
		return true
	}
	var errRs gobin.ErrorResponse
	if err := json.NewDecoder(rs.Body).Decode(&errRs); err != nil {
		cmd.PrintErrln("Failed to decode error response:", err)
		return false
	}
	cmd.PrintErrf("Failed to %s: %s\n", method, errRs.Message)
	return false
}

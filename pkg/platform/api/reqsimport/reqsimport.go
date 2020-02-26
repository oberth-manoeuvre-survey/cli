package reqsimport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"path"
	"time"

	"github.com/ActiveState/cli/pkg/platform/api"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

const (
	translateContentType = "application/json"
)

// Opts ...
type Opts struct {
	translateURL string
}

// ReqsImport ...
type ReqsImport struct {
	opts   Opts
	client *http.Client
}

// New ...
func New(opts Opts) (*ReqsImport, error) {
	c := &http.Client{
		Timeout: 60 * time.Second,
	}

	ri := ReqsImport{
		opts:   opts,
		client: c,
	}

	return &ri, nil
}

// Init ...
func Init() *ReqsImport {
	svcURL := api.GetServiceURL(api.ServiceInventory)

	opts := Opts{
		translateURL: path.Join(svcURL.Host, svcURL.Path),
	}

	ri, err := New(opts)
	if err != nil {
		panic(err)
	}

	return ri
}

// ChangeRequest ...
func (ri *ReqsImport) ChangeRequest(data []byte) (*ChangeRequest, error) {
	reqMsg := ReqsTxtTranslateReqMsg{
		Data: string(data),
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(reqMsg); err != nil {
		return nil, err
	}

	url := ri.opts.translateURL

	resp, err := ri.client.Post(url, translateContentType, &buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint

	var respMsg ReqsTxtTranslateRespMsg

	if err := json.NewDecoder(resp.Body).Decode(&respMsg); err != nil {
		return nil, err
	}

	return respMsg.ChangeRequest, nil
}

// ChangeRequest ...
type ChangeRequest mono_models.CommitEditable

// ReqsTxtTranslateReqMsg ...
type ReqsTxtTranslateReqMsg struct {
	Data string `json:"requirements"`
}

// ReqsTxtTranslateRespMsg ...
type ReqsTxtTranslateRespMsg struct {
	*ChangeRequest
	Errors []string `json:"errors,omitempty"`
}
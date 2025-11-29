package restclient

import (
	"context"
	"encoding/json"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/hdlopez/clean-architecture-golang/apierror"
)

type restAPI struct {
	readClient *resty.Client
}

func (api *restAPI) get(ctx context.Context, url string, h http.Header, v interface{}) error {
	var r *resty.Response
	req := api.readClient.R()
	req.SetContext(ctx)
	req.SetError(&apierror.APIError{})

	r, err := req.Get(url)

	if err != nil {
		// returns API error
		return err
	}

	if r.StatusCode() != 200 {
		// returns API error
		return apierror.New(r.StatusCode(), "Status code was not 200")
	}

	if err = json.Unmarshal(r.Body(), v); err != nil {
		// returns API error
		return apierror.New(500, "Unmarshal error")
	}
	return nil
}

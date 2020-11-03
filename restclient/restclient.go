package restclient

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/hdlopez/clean-architecture-golang/apierror"
)

type restAPI struct {
	readClient *resty.Client
	//writeClient func() *resty.Client
	//logger   *apiLogger
}

func (api *restAPI) get(url string, h http.Header, v interface{}) (interface{}, error) {
	//start := time.Now()

	var r *resty.Response
	req := api.readClient.R()
	req.SetError(&apierror.APIError{})

	r, err := req.Get(url)

	if err != nil {
		// returns API error
	}

	//api.logger.httpGet(url, h, r, time.Since(start))
	if r.StatusCode() != 200 {
		// returns API error
	}

	if err = json.Unmarshal(r.Body(), v); err != nil {
		// returns API error
	}
	return v, nil
}

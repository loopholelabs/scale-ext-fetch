package extfetch

import (
	"fmt"
	"io"
	"net/http"
)

type FetchExtension struct {
}

type HostHttpConnector struct {
	Timeout int32
}

func (f *FetchExtension) New(hc *HttpConfig) (HttpConnector, error) {
	return &HostHttpConnector{
		Timeout: hc.Timeout,
	}, nil
}

func (hc *HostHttpConnector) Fetch(cd *ConnectionDetails) (HttpResponse, error) {
	fmt.Printf("HttpConnector.Fetch[%v] (HostHttpConnector %v)\n", cd, hc)
	resp := HttpResponse{}

	response, err := http.Get(cd.Url)
	if err != nil {
		return resp, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return resp, err
	}

	resp.StatusCode = int32(response.StatusCode)
	resp.Body = body
	resp.Headers = map[string]StringList{"Hello": {Values: []string{"World"}}}

	return resp, nil
}

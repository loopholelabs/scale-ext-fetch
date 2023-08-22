package extfetch

import (
	"fmt"
	"io"
	"net/http"
)

type HostHttpConnector struct {
	Timeout int32
}

func HostNew(hc *HttpConfig) (*HostHttpConnector, error) {
	return &HostHttpConnector{
		Timeout: hc.Timeout,
	}, nil
}

func (hc *HostHttpConnector) Fetch(cd *ConnectionDetails) (*HttpResponse, error) {
	fmt.Printf("HttpConnector.Fetch[%v]\n", cd)

	response, err := http.Get(cd.Url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resp := &HttpResponse{
		StatusCode: int32(response.StatusCode),
		Body:       body,
		Headers:    map[string]StringList{"Hello": {Values: []string{"World"}}},
	}
	return resp, nil
}

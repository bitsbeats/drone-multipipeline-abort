package plugin

import (
	"fmt"
	"net/http"
)

type transport struct {
	token string
}

func (t transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := new(http.Request)
	*req2 = *req

	req2.Header = make(http.Header, len(req.Header))

	for k, v := range req.Header {
		req2.Header[k] = append([]string(nil), v...)
	}

	req2.Header.Set("Authorization", fmt.Sprint("Bearer ", t.token))

	return http.DefaultTransport.RoundTrip(req2)
}

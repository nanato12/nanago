package nana

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"path"
)

func (n *Nana) url(base *url.URL, Endpoint string) string {
	u := *base
	u.Path = path.Join(u.Path, Endpoint)
	return u.String()
}

func (n *Nana) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if n.Token != "" {
		req.Header.Set("authorization", "token "+n.Token)
	}
	req.Header.Set("x-device-identifier", n.deviceID)
	req.Header.Set("x-appsflyer-identifier", n.appsflyerID)
	req.Header.Set("x-app-version", APPVersion)
	req.Header.Set("x-ios-version", OSVersion)
	req.Header.Set("user-agent", "nana/"+APPVersion+" (iPhone; iOS "+OSVersion+"; Scale/2.00)")
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return n.httpClient.Do(req)
}

func (n *Nana) get(ctx context.Context, Endpoint string, query url.Values) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, n.url(n.endpointBase, Endpoint), nil)
	if err != nil {
		return nil, err
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	return n.do(ctx, req)
}

func (n *Nana) post(ctx context.Context, Endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, n.url(n.endpointBase, Endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return n.do(ctx, req)
}

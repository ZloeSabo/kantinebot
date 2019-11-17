package wandel

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httputil"
)

type jsonParser func(*http.Response) error

func headersApp(app string) map[string]string {
	return map[string]string{
		"App-Brand": app,
	}
}

func headersAuthApp(app string, authorization string) map[string]string {
	res := headersApp(app)
	res["Authorization"] = authorization
	return res
}

func getResource(ctx context.Context, client httpClient, headers map[string]string, endpoint string, intf interface{}, dbg debuglogger) error {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	return performRequest(ctx, client, req, newJSONParser(intf), dbg)
}

func performRequest(ctx context.Context, client httpClient, req *http.Request, parser jsonParser, dbg debuglogger) error {
	req = req.WithContext(ctx)
	debugRequest(req, dbg)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	debugResponse(resp, dbg)

	return parser(resp)
}

func newJSONParser(dst interface{}) jsonParser {
	return func(resp *http.Response) error {
		return json.NewDecoder(resp.Body).Decode(dst)
	}
}

func debugRequest(req *http.Request, dbg debuglogger) {
	if dbg.DebugEnabled() {
		reqt, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			dbg.DebugPrintf("Cant dump request: %s", err)
		}
		dbg.DebugPrintf("Request: %s", string(reqt))
	}
}

func debugResponse(resp *http.Response, dbg debuglogger) {
	if dbg.DebugEnabled() {
		reqt, err := httputil.DumpResponse(resp, true)
		if err != nil {
			dbg.DebugPrintf("Cant dump response: %s", err)
		}
		dbg.DebugPrintf("Response: %s", string(reqt))
	}
}

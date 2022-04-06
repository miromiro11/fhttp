package http

import "net/http"

// StdlibTransport is a wrapper for using net/http transports with fhttp
type StdlibTransport struct {
	*Transport
}

func (txp *StdlibTransport) RoundTrip(stdReq *http.Request) (*http.Response, error) {
	req := &Request{
		Method:           stdReq.Method,
		URL:              stdReq.URL,
		Proto:            stdReq.Proto,
		ProtoMajor:       stdReq.ProtoMajor,
		ProtoMinor:       stdReq.ProtoMinor,
		Header:           Header(stdReq.Header),
		Body:             stdReq.Body,
		GetBody:          stdReq.GetBody,
		ContentLength:    stdReq.ContentLength,
		TransferEncoding: stdReq.TransferEncoding,
		Close:            stdReq.Close,
		Host:             stdReq.Host,
		Form:             stdReq.Form,
		PostForm:         stdReq.PostForm,
		MultipartForm:    stdReq.MultipartForm,
		Trailer:          Header(stdReq.Trailer),
		RemoteAddr:       stdReq.RemoteAddr,
		RequestURI:       stdReq.RequestURI,
		TLS:              stdReq.TLS,
		Cancel:           stdReq.Cancel,
		Response:         nil, // cannot assign this field
		ctx:              stdReq.Context(),
	}
	resp, err := txp.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	stdResp := &http.Response{
		Status:           resp.Status,
		StatusCode:       resp.StatusCode,
		Proto:            resp.Proto,
		ProtoMinor:       resp.ProtoMinor,
		ProtoMajor:       resp.ProtoMajor,
		Header:           http.Header(resp.Header),
		Body:             resp.Body,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
		Close:            resp.Close,
		Uncompressed:     resp.Uncompressed,
		Trailer:          http.Header(resp.Trailer),
		Request:          stdReq,
		TLS:              resp.TLS,
	}
	return stdResp, nil
}

// FHttpTransport is a wrapper for using net/http transports with fhttp
type FHttpTransport struct {
	*http.Transport
}

func (txp *FHttpTransport) RoundTrip(fhttpReq *Request) (*Response, error) {
	req := &http.Request{
		Method:           fhttpReq.Method,
		URL:              fhttpReq.URL,
		Proto:            fhttpReq.Proto,
		ProtoMajor:       fhttpReq.ProtoMajor,
		ProtoMinor:       fhttpReq.ProtoMinor,
		Header:           http.Header(fhttpReq.Header),
		Body:             fhttpReq.Body,
		GetBody:          fhttpReq.GetBody,
		ContentLength:    fhttpReq.ContentLength,
		TransferEncoding: fhttpReq.TransferEncoding,
		Close:            fhttpReq.Close,
		Host:             fhttpReq.Host,
		Form:             fhttpReq.Form,
		PostForm:         fhttpReq.PostForm,
		MultipartForm:    fhttpReq.MultipartForm,
		Trailer:          http.Header(fhttpReq.Trailer),
		RemoteAddr:       fhttpReq.RemoteAddr,
		RequestURI:       fhttpReq.RequestURI,
		TLS:              fhttpReq.TLS,
		Cancel:           fhttpReq.Cancel,
		Response:         nil, // cannot assign this field
	}
	req = req.WithContext(fhttpReq.ctx)

	resp, err := txp.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	stdResp := &Response{
		Status:           resp.Status,
		StatusCode:       resp.StatusCode,
		Proto:            resp.Proto,
		ProtoMinor:       resp.ProtoMinor,
		ProtoMajor:       resp.ProtoMajor,
		Header:           Header(resp.Header),
		Body:             resp.Body,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
		Close:            resp.Close,
		Uncompressed:     resp.Uncompressed,
		Trailer:          Header(resp.Trailer),
		Request:          fhttpReq,
		TLS:              resp.TLS,
	}
	return stdResp, nil
}

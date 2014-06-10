package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type PostBackHandler struct {
	Url          string
	EncodeOnPost bool
}

func (hnd *PostBackHandler) Deliver(message string) error {
	buff := strings.NewReader(hnd.getPostBody(message))

	_, err := http.Post(hnd.Url, "text/plain", buff)
	if err != nil {
		return fmt.Errorf("An error ocurred delivering a message. %q", err)
	}

	return nil
}

func (hnd *PostBackHandler) Describe() string {
	var redactedUrl string

	uri, err := url.Parse(hnd.Url)

	if err != nil {
		redactedUrl = hnd.Url
	} else {
		redactedUrl = uri.Scheme + "://" + uri.Host + uri.Path
	}

	return fmt.Sprintf("PostbackHandler (url=%s, encode=%t)", redactedUrl, hnd.EncodeOnPost)
}

func (hnd *PostBackHandler) getPostBody(data string) string {
	if hnd.EncodeOnPost == true {
		return url.QueryEscape(data)
	}

	return data
}

func NewPostBackHandler(postUrl string, encodeOnPost bool) *PostBackHandler {
	return &PostBackHandler{Url: postUrl, EncodeOnPost: encodeOnPost}
}

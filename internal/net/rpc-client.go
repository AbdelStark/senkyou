package net

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const contentTypeJson = "application/json"

type RpcClient interface {
	Call([]byte) ([]byte, error)
}

func NewRpcClient(url string) RpcClient {
	return rpcClient{
		url:    url,
		client: http.Client{},
	}
}

type rpcClient struct {
	url    string
	client http.Client
}

func (r rpcClient) Call(request []byte) ([]byte, error) {
	resp, err := r.client.Post(r.url, contentTypeJson, bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

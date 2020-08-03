package store

import (
	"context"
	"encoding/json"
	"fmt"

	"graphql-gen/graph/helper/client"
	"graphql-gen/graph/model"
)

type Client struct {
	httpClient client.Service
	host       string
	path       string
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	READ   = "read"
)

func NewClient(httpClient client.Service, host, path string, ) *Client {
	return &Client{
		httpClient: httpClient,
		host:       host,
		path:       path,
	}
}

func (c *Client) Create(_ context.Context, input *model.CreateInput) (*model.Response, error) {
	url := fmt.Sprintf("%s/%s", c.host, c.path)

	iResp, err := c.httpClient.CURL(POST, url, c.httpClient.WithBody(input))
	if err != nil {
		fmt.Printf("failed to create, errors: %+v", err)
		return &model.Response{}, err
	}

	return interfaceToResp(iResp)
}

func (c *Client) Update(_ context.Context, input *model.UpdateInput) (*model.Response, error) {
	url := fmt.Sprintf("%s/%s", c.host, c.path)

	iResp, err := c.httpClient.CURL(PUT, url, c.httpClient.WithBody(input))
	if err != nil {
		fmt.Printf("failed to update, errors: %+v", err)
		return &model.Response{}, err
	}

	return interfaceToResp(iResp)
}

func (c *Client) Get(_ context.Context, input string) (*model.Response, error) {
	url := fmt.Sprintf("%s/%s/%s?id=%s", c.host, c.path, READ, input)

	iResp, err := c.httpClient.CURL(GET, url)
	if err != nil {
		fmt.Printf("failed to create, errors: %+v", err)
		return &model.Response{}, err
	}
	return interfaceToResp(iResp)
}

func (c *Client) List(_ context.Context, offset, limit int) (int, []*model.Response, error) {
	var resp *model.ResponseList

	url := fmt.Sprintf("%s/%s?offset=%d&limit=%d", c.host, c.path, offset, limit)

	iListResp, err := c.httpClient.CURL(GET, url)
	if err != nil {
		fmt.Printf("failed to retrieve response, errors: %+v", err)
		return 0, []*model.Response{}, err
	}

	resp, err = interfaceToSliceResp(iListResp)
	if err != nil {
		fmt.Printf("failed to convert response, errors: %+v", err)
		return *resp.Total, resp.Response, err
	}

	return *resp.Total, resp.Response, err
}

func (c *Client) Delete(_ context.Context, input string) (*model.Response, error) {
	url := fmt.Sprintf("%s/%s?id=%s", c.host, c.path, input)

	iResp, err := c.httpClient.CURL(DELETE, url)
	if err != nil {
		fmt.Printf("failed to delete, errors: %+v", err)
		return &model.Response{}, err
	}
	return interfaceToResp(iResp)
}

func interfaceToResp(iResp interface{}) (*model.Response, error) {
	var resp model.Response
	byteData, err := json.Marshal(iResp)
	if err != nil {
		return &resp, err
	}

	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

func interfaceToSliceResp(iResp interface{}) (*model.ResponseList, error) {
	RespList := &model.ResponseList{}

	byteData, err := json.Marshal(iResp)
	if err != nil {
		return RespList, err
	}

	err = json.Unmarshal(byteData, &RespList)

	return RespList, err
}

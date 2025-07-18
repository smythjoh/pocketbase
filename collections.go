package pocketbase

import (
	"encoding/json"
	"fmt"

	"github.com/pocketbase/pocketbase/models"
)

type Collections struct {
	*Client
	BasePath string
}

func NewCollections(client *Client) *Collections {
	return &Collections{
		Client:   client,
		BasePath: client.url + "/api/collections",
	}
}

func (c *Collections) List(params ParamsList) (ResponseList[models.Collection], error) {
	var response ResponseList[models.Collection]

	if err := c.Authorize(); err != nil {
		return response, err
	}

	request := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("page", fmt.Sprintf("%v", params.Page)).
		SetQueryParam("perPage", fmt.Sprintf("%v", params.Size)).
		SetQueryParam("filters", params.Filters)

	resp, err := request.Get(c.BasePath)
	if err != nil {
		return response, fmt.Errorf("can't send list request to pocketbase, err %w", err)
	}

	if resp.IsError() {
		return response, fmt.Errorf("pocketbase returned status: %d, msg: %s, err %w",
			resp.StatusCode(),
			resp.String(),
			ErrInvalidResponse,
		)
	}
	if err := json.Unmarshal(resp.Body(), &response); err != nil {

		return response, fmt.Errorf("can't unmarshal response, err %w", err)
	}
	return response, nil
}

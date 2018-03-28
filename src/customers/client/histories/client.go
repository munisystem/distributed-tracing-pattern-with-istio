package histories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/client"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/model"
)

type clientImpl struct {
	url string
	cli *http.Client
}

func New(url string) client.Histories {
	return &clientImpl{
		url: url,
		cli: &http.Client{},
	}
}

type getCustomerHistoryResponse struct {
	Id         int64  `json:"id"`
	CustomerId int64  `json:"customer_id"`
	ItemName   string `json:"item_name"`
	Message    string `json:"message"`
}

func (c *clientImpl) GetCustomerHistory(ctx context.Context, customerId int64) (*model.History, error) {
	endpoint := fmt.Sprintf("/api/customers/%d", customerId)
	req, err := c.concreteRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get hisotry")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &getCustomerHistoryResponse{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}
	return &model.History{
		ItemName: r.ItemName,
		Message:  r.Message,
	}, nil
}

func (c *clientImpl) concreteRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, c.url+endpoint, body)
}

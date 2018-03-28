package items

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/client"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/model"
)

type clientImpl struct {
	url string
	cli *http.Client
}

func New(url string) client.Items {
	return &clientImpl{
		url: url,
		cli: &http.Client{},
	}
}

func (c *clientImpl) GetItem(ctx context.Context, itemId int64) (*model.Item, error) {
	endpoint := fmt.Sprintf("/api/items/%d", itemId)
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
		return nil, fmt.Errorf("failed to get item")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	i := &model.Item{}
	if err := json.Unmarshal(body, i); err != nil {
		return nil, err
	}
	return i, nil
}

func (c *clientImpl) concreteRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, c.url+endpoint, body)
}

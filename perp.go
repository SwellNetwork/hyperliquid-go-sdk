package hyperliquid

import (
	"context"
	"fmt"
)

func (c *HTTPClient) MetaAndAssetCtxs(ctx context.Context) ([]MetaAndAssetCtx, error) {
	body := map[string]any{"type": InfoTypeMetaAndAssetCtxs}

	var result metaAndAssetCtxsResult

	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(body).
		SetResult(&result).
		Post(PathInfo)
	if err != nil {
		return nil, fmt.Errorf("meta and asset ctxs request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("meta and asset ctxs request failed: status %d: %s", resp.StatusCode(), resp.String())
	}

	return transformMetaAndAssetCtxsResult(result)
}

package hyperliquid

import (
	"context"
	"fmt"
)

func (c *HTTPClient) SpotMetaAndAssetCtxs(ctx context.Context) (*SpotMetaAndAssetCtx, error) {
	body := map[string]any{"type": InfoTypeSpotMetaAndAssetCtxs}

	var result spotMetaAndAssetCtxsResult

	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(body).
		SetResult(&result).
		Post(PathInfo)
	if err != nil {
		return nil, fmt.Errorf("spot meta and asset ctxs request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("spot meta and asset ctxs request failed: status %d: %s", resp.StatusCode(), resp.String())
	}

	return transformSpotMetaAndAssetCtxsResult(result)

}

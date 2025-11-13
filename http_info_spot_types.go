package hyperliquid

import (
	"fmt"

	"github.com/goccy/go-json"
)

// Spot meta & asset ctxs

type (
	spotMetaAndAssetCtxsResult struct {
		Meta   spotMetaAndAssetCtxMeta
		Assets []spotAssetCtxPayload
	}

	spotMetaAndAssetCtxMeta struct {
		Universe []spotMetaAndAssetCtxUniverse `json:"universe"`
		Tokens   []SpotTokenInfo               `json:"tokens"`
	}

	spotMetaAndAssetCtxUniverse struct {
		Tokens      []int  `json:"tokens"`
		Name        string `json:"name"`
		Index       int    `json:"index"`
		IsCanonical bool   `json:"isCanonical"`
	}

	spotAssetCtxPayload struct {
		DayNtlVlm         string `json:"dayNtlVlm"`
		MarkPx            string `json:"markPx"`
		MidPx             string `json:"midPx"`
		PrevDayPx         string `json:"prevDayPx"`
		CirculatingSupply string `json:"circulatingSupply"`
		Coin              string `json:"coin"`
		TotalSupply       string `json:"totalSupply"`
		DayBaseVlm        string `json:"dayBaseVlm"`
	}
)

func (r *spotMetaAndAssetCtxsResult) UnmarshalJSON(data []byte) error {
	var payload [2]json.RawMessage
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	if err := json.Unmarshal(payload[0], &r.Meta); err != nil {
		return fmt.Errorf("decode spot meta section: %w", err)
	}

	if err := json.Unmarshal(payload[1], &r.Assets); err != nil {
		return fmt.Errorf("decode spot asset ctx section: %w", err)
	}

	return nil
}

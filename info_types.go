package hyperliquid

import (
	"fmt"

	"github.com/goccy/go-json"
)

// Predicted fundings

type predictedFundingsResult []predictedFundingsCoin

type predictedFundingsCoin struct {
	Coin      string
	Exchanges []predictedFundingsExchange
}

type predictedFundingsExchange struct {
	Exchange string
	Details  predictedFundingDetails
}

type predictedFundingDetails struct {
	FundingRate          string `json:"fundingRate"`
	NextFundingTime      int64  `json:"nextFundingTime"`
	FundingIntervalHours int64  `json:"fundingIntervalHours"`
}

func (p *predictedFundingsCoin) UnmarshalJSON(data []byte) error {
	var payload [2]json.RawMessage
	if err := json.Unmarshal(data, &payload); err != nil {
		return fmt.Errorf("decode predicted fundings coin: %w", err)
	}

	if err := json.Unmarshal(payload[0], &p.Coin); err != nil {
		return err
	}

	return json.Unmarshal(payload[1], &p.Exchanges)
}

func (p *predictedFundingsExchange) UnmarshalJSON(data []byte) error {
	var payload [2]json.RawMessage
	if err := json.Unmarshal(data, &payload); err != nil {
		return fmt.Errorf("decode predicted fundings exchange: %w", err)
	}

	if err := json.Unmarshal(payload[0], &p.Exchange); err != nil {
		return err
	}

	return json.Unmarshal(payload[1], &p.Details)
}

// Perp meta & asset ctxs

type metaAndAssetCtxsResult struct {
	Meta   metaAndAssetCtxMeta
	Assets []assetCtxPayload
}

func (m *metaAndAssetCtxsResult) UnmarshalJSON(data []byte) error {
	var payload [2]json.RawMessage
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	if err := json.Unmarshal(payload[0], &m.Meta); err != nil {
		return fmt.Errorf("decode meta section: %w", err)
	}

	if err := json.Unmarshal(payload[1], &m.Assets); err != nil {
		return fmt.Errorf("decode asset ctx section: %w", err)
	}

	return nil
}

type metaAndAssetCtxMeta struct {
	Universe        []metaAndAssetCtxUniverse `json:"universe"`
	MarginTables    marginTableMap            `json:"marginTables"`
	CollateralToken int                       `json:"collateralToken"`
}

type metaAndAssetCtxUniverse struct {
	SzDecimals    int    `json:"szDecimals"`
	Name          string `json:"name"`
	MaxLeverage   int    `json:"maxLeverage"`
	MarginTableID int    `json:"marginTableId"`
	OnlyIsolated  bool   `json:"onlyIsolated,omitempty"`
	IsDelisted    bool   `json:"isDelisted,omitempty"`
}

type marginTableMap map[int]marginTableDetails

func (m *marginTableMap) UnmarshalJSON(data []byte) error {
	var entries []json.RawMessage
	if err := json.Unmarshal(data, &entries); err != nil {
		return err
	}

	if len(entries) == 0 {
		*m = nil
		return nil
	}

	result := make(marginTableMap, len(entries))
	for _, raw := range entries {
		var payload [2]json.RawMessage
		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}

		var id int
		if err := json.Unmarshal(payload[0], &id); err != nil {
			return err
		}

		var details marginTableDetails
		if err := json.Unmarshal(payload[1], &details); err != nil {
			return err
		}

		result[id] = details
	}

	*m = result
	return nil
}

type marginTierPayload struct {
	LowerBound  string `json:"lowerBound"`
	MaxLeverage int    `json:"maxLeverage"`
}

type marginTableDetails struct {
	Description string              `json:"description"`
	MarginTiers []marginTierPayload `json:"marginTiers"`
}

type assetCtxPayload struct {
	Funding      string   `json:"funding"`
	OpenInterest string   `json:"openInterest"`
	PrevDayPx    string   `json:"prevDayPx"`
	DayNtlVlm    string   `json:"dayNtlVlm"`
	Premium      string   `json:"premium"`
	OraclePx     string   `json:"oraclePx"`
	MarkPx       string   `json:"markPx"`
	MidPx        string   `json:"midPx"`
	ImpactPxs    []string `json:"impactPxs"`
	DayBaseVlm   string   `json:"dayBaseVlm"`
}

// Spot meta & asset ctxs

type spotMetaAndAssetCtxsResult struct {
	Meta   spotMetaAndAssetCtxMeta
	Assets []spotAssetCtxPayload
}

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

type spotMetaAndAssetCtxMeta struct {
	Universe []spotMetaAndAssetCtxUniverse `json:"universe"`
	Tokens   []SpotTokenInfo               `json:"tokens"`
}

type spotMetaAndAssetCtxUniverse struct {
	Tokens      []int  `json:"tokens"`
	Name        string `json:"name"`
	Index       int    `json:"index"`
	IsCanonical bool   `json:"isCanonical"`
}

type spotAssetCtxPayload struct {
	DayNtlVlm         string `json:"dayNtlVlm"`
	MarkPx            string `json:"markPx"`
	MidPx             string `json:"midPx"`
	PrevDayPx         string `json:"prevDayPx"`
	CirculatingSupply string `json:"circulatingSupply"`
	Coin              string `json:"coin"`
	TotalSupply       string `json:"totalSupply"`
	DayBaseVlm        string `json:"dayBaseVlm"`
}

package hyperliquid

import (
	"fmt"

	"github.com/goccy/go-json"
)

// Predicted fundings

type (
	predictedFundingsResult []predictedFundingsCoin

	predictedFundingsCoin struct {
		Coin      string
		Exchanges []predictedFundingsExchange
	}

	predictedFundingsExchange struct {
		Exchange string
		Details  predictedFundingDetails
	}

	predictedFundingDetails struct {
		FundingRate          string `json:"fundingRate"`
		NextFundingTime      int64  `json:"nextFundingTime"`
		FundingIntervalHours int64  `json:"fundingIntervalHours"`
	}
)

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

type (
	metaAndAssetCtxsResult struct {
		Meta   metaAndAssetCtxMeta
		Assets []assetCtxPayload
	}

	metaAndAssetCtxMeta struct {
		Universe        []metaAndAssetCtxUniverse `json:"universe"`
		MarginTables    marginTableMap            `json:"marginTables"`
		CollateralToken int                       `json:"collateralToken"`
	}

	metaAndAssetCtxUniverse struct {
		SzDecimals    int    `json:"szDecimals"`
		Name          string `json:"name"`
		MaxLeverage   int    `json:"maxLeverage"`
		MarginTableID int    `json:"marginTableId"`
		OnlyIsolated  bool   `json:"onlyIsolated,omitempty"`
		IsDelisted    bool   `json:"isDelisted,omitempty"`
	}

	marginTableMap map[int]marginTableDetails

	marginTierPayload struct {
		LowerBound  string `json:"lowerBound"`
		MaxLeverage int    `json:"maxLeverage"`
	}

	marginTableDetails struct {
		Description string              `json:"description"`
		MarginTiers []marginTierPayload `json:"marginTiers"`
	}

	assetCtxPayload struct {
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
)

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

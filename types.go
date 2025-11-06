package hyperliquid

import (
	"fmt"

	"github.com/goccy/go-json"
)

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
	FundingIntervalHours int    `json:"fundingIntervalHours"`
}

func (p *predictedFundingsCoin) UnmarshalJSON(data []byte) error {
	var payload []json.RawMessage
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	if len(payload) != 2 {
		return fmt.Errorf("predicted fundings coin entry format invalid: expected 2 elements, got %d", len(payload))
	}

	if err := json.Unmarshal(payload[0], &p.Coin); err != nil {
		return err
	}

	if err := json.Unmarshal(payload[1], &p.Exchanges); err != nil {
		return err
	}

	return nil
}

func (p *predictedFundingsExchange) UnmarshalJSON(data []byte) error {
	var payload []json.RawMessage
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	if len(payload) != 2 {
		return fmt.Errorf("predicted fundings exchange entry format invalid: expected 2 elements, got %d", len(payload))
	}

	if err := json.Unmarshal(payload[0], &p.Exchange); err != nil {
		return err
	}

	if err := json.Unmarshal(payload[1], &p.Details); err != nil {
		return err
	}

	return nil
}

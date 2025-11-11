package hyperliquid

import (
	"context"
	"fmt"
)

func (c *HTTPClient) MetaAndAssetCtxs(ctx context.Context) (*MetaAndAssetCtx, error) {
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

//  MetaAndAssetCtxs sample response
// [
//     {
//         "universe": [
//             {
//                 "szDecimals": 5,
//                 "name": "BTC",
//                 "maxLeverage": 40,
//                 "marginTableId": 56
//             },
//             {
//                 "szDecimals": 4,
//                 "name": "ETH",
//                 "maxLeverage": 25,
//                 "marginTableId": 55
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "ATOM",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "MATIC",
//                 "maxLeverage": 20,
//                 "marginTableId": 20,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "DYDX",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "SOL",
//                 "maxLeverage": 20,
//                 "marginTableId": 54
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "AVAX",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 3,
//                 "name": "BNB",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "APE",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "OP",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "LTC",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "ARB",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "DOGE",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "INJ",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "SUI",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "kPEPE",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "CRV",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "LDO",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "LINK",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "STX",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "RNDR",
//                 "maxLeverage": 20,
//                 "marginTableId": 20,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "CFX",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "FTM",
//                 "maxLeverage": 10,
//                 "marginTableId": 10,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "GMX",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "SNX",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "XRP",
//                 "maxLeverage": 20,
//                 "marginTableId": 53
//             },
//             {
//                 "szDecimals": 3,
//                 "name": "BCH",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "APT",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "AAVE",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "COMP",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 4,
//                 "name": "MKR",
//                 "maxLeverage": 10,
//                 "marginTableId": 51,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "WLD",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "FXS",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "HPOS",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "isDelisted": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "RLB",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "isDelisted": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 3,
//                 "name": "UNIBOT",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "isDelisted": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "YGG",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "TRX",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "kSHIB",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "UNI",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "SEI",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "RUNE",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "OX",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "isDelisted": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "FRIEND",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "isDelisted": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "SHIA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "isDelisted": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "CYBER",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "ZRO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "BLZ",
//                 "maxLeverage": 5,
//                 "marginTableId": 5,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "DOT",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "BANANA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "TRB",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "FTT",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "LOOM",
//                 "maxLeverage": 10,
//                 "marginTableId": 10,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "OGN",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "RDNT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ARK",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "BNT",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "CANTO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "REQ",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "BIGTIME",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "KAS",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ORBS",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "BLUR",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "TIA",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "BSV",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ADA",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "TON",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MINA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "POLYX",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "GAS",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "PENDLE",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "STG",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "FET",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "STRAX",
//                 "maxLeverage": 10,
//                 "marginTableId": 10,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "NEAR",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MEME",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "ORDI",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "BADGER",
//                 "maxLeverage": 5,
//                 "marginTableId": 5,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "NEO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "ZEN",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "FIL",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "PYTH",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "SUSHI",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "ILV",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "IMX",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "kBONK",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "GMT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "SUPER",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "USTC",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "NFTI",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "isDelisted": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "JUP",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "kLUNC",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "RSR",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "GALA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "JTO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "NTRN",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "ACE",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MAV",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "WIF",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "CAKE",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "PEOPLE",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "ENS",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "ETC",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "XAI",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "MANTA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "UMA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ONDO",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ALT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "ZETA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "DYM",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "MAVIA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "W",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 5,
//                 "name": "PANDORA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "isDelisted": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "STRK",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "PIXEL",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "AI",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 3,
//                 "name": "TAO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "AR",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MYRO",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "kFLOKI",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "BOME",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "ETHFI",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ENA",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "MNT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "TNSR",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "SAGA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MERL",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "HBAR",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "POPCAT",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "OMNI",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "EIGEN",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "REZ",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "NOT",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "TURBO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "BRETT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "IO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ZK",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "BLAST",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "LISTA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MEW",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "RENDER",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "kDOGS",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "POL",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "CATI",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "CELO",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "HMSTR",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "SCR",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "NEIROETH",
//                 "maxLeverage": 5,
//                 "marginTableId": 5,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "kNEIRO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "GOAT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MOODENG",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "GRASS",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "PURR",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "PNUT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "XLM",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "CHILLGUY",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "SAND",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "IOTA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ALGO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "HYPE",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "ME",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MOVE",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "VIRTUAL",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "PENGU",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "USUAL",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "FARTCOIN",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "AI16Z",
//                 "maxLeverage": 5,
//                 "marginTableId": 5,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "AIXBT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ZEREBRO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "BIO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "GRIFFAIN",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "SPX",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "S",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "MORPHO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "TRUMP",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "MELANIA",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ANIME",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "VINE",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "VVV",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "JELLY",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "BERA",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "TST",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "LAYER",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "IP",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "OM",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "KAITO",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "NIL",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 3,
//                 "name": "PAXG",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "PROMPT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "BABY",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "WCT",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "HYPER",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ZORA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "INIT",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "DOOD",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "LAUNCHCOIN",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "isDelisted": true
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "NXPC",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "SOPH",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "RESOLV",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "SYRUP",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "PUMP",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "PROVE",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "YZY",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "XPL",
//                 "maxLeverage": 10,
//                 "marginTableId": 51
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "WLFI",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "LINEA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "SKY",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "ASTER",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "AVNT",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "STBL",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "0G",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "HEMI",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "APEX",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "2Z",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 2,
//                 "name": "ZEC",
//                 "maxLeverage": 10,
//                 "marginTableId": 52
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MON",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MET",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "MEGA",
//                 "maxLeverage": 3,
//                 "marginTableId": 3,
//                 "onlyIsolated": true,
//                 "marginMode": "strictIsolated"
//             },
//             {
//                 "szDecimals": 0,
//                 "name": "CC",
//                 "maxLeverage": 3,
//                 "marginTableId": 3
//             },
//             {
//                 "szDecimals": 1,
//                 "name": "ICP",
//                 "maxLeverage": 5,
//                 "marginTableId": 5
//             }
//         ],
//         "marginTables": [
//             [
//                 50,
//                 {
//                     "description": "",
//                     "marginTiers": [
//                         {
//                             "lowerBound": "0.0",
//                             "maxLeverage": 50
//                         }
//                     ]
//                 }
//             ],
//             [
//                 51,
//                 {
//                     "description": "tiered 10x",
//                     "marginTiers": [
//                         {
//                             "lowerBound": "0.0",
//                             "maxLeverage": 10
//                         },
//                         {
//                             "lowerBound": "3000000.0",
//                             "maxLeverage": 5
//                         }
//                     ]
//                 }
//             ],
//             [
//                 52,
//                 {
//                     "description": "tiered 10x (2)",
//                     "marginTiers": [
//                         {
//                             "lowerBound": "0.0",
//                             "maxLeverage": 10
//                         },
//                         {
//                             "lowerBound": "20000000.0",
//                             "maxLeverage": 5
//                         }
//                     ]
//                 }
//             ],
//             [
//                 53,
//                 {
//                     "description": "tiered 20x",
//                     "marginTiers": [
//                         {
//                             "lowerBound": "0.0",
//                             "maxLeverage": 20
//                         },
//                         {
//                             "lowerBound": "40000000.0",
//                             "maxLeverage": 10
//                         }
//                     ]
//                 }
//             ],
//             [
//                 54,
//                 {
//                     "description": "tiered 20x (2)",
//                     "marginTiers": [
//                         {
//                             "lowerBound": "0.0",
//                             "maxLeverage": 20
//                         },
//                         {
//                             "lowerBound": "70000000.0",
//                             "maxLeverage": 10
//                         }
//                     ]
//                 }
//             ],
//             [
//                 55,
//                 {
//                     "description": "tiered 25x",
//                     "marginTiers": [
//                         {
//                             "lowerBound": "0.0",
//                             "maxLeverage": 25
//                         },
//                         {
//                             "lowerBound": "100000000.0",
//                             "maxLeverage": 15
//                         }
//                     ]
//                 }
//             ],
//             [
//                 56,
//                 {
//                     "description": "tiered 40x",
//                     "marginTiers": [
//                         {
//                             "lowerBound": "0.0",
//                             "maxLeverage": 40
//                         },
//                         {
//                             "lowerBound": "150000000.0",
//                             "maxLeverage": 20
//                         }
//                     ]
//                 }
//             ]
//         ],
//         "collateralToken": 0
//     },
//     [
//         {
//             "funding": "0.0000125",
//             "openInterest": "29189.54088",
//             "prevDayPx": "101871.0",
//             "dayNtlVlm": "3167083091.9661107063",
//             "premium": "0.0002360294",
//             "oraclePx": "105919.0",
//             "markPx": "105942.0",
//             "midPx": "105944.5",
//             "impactPxs": [
//                 "105944.0",
//                 "105945.0"
//             ],
//             "dayBaseVlm": "30304.6659200001"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "506651.0913999989",
//             "prevDayPx": "3416.0",
//             "dayNtlVlm": "1688638783.8624215126",
//             "premium": "0.0000556313",
//             "oraclePx": "3595.1",
//             "markPx": "3595.3",
//             "midPx": "3595.35",
//             "impactPxs": [
//                 "3595.3",
//                 "3595.4"
//             ],
//             "dayBaseVlm": "475308.1713999998"
//         },
//         {
//             "funding": "0.0000004259",
//             "openInterest": "1652902.9400000002",
//             "prevDayPx": "2.8703",
//             "dayNtlVlm": "1111022.4857120002",
//             "premium": "-0.0010051935",
//             "oraclePx": "2.9845",
//             "markPx": "2.9809",
//             "midPx": "2.9807",
//             "impactPxs": [
//                 "2.9795",
//                 "2.9815"
//             ],
//             "dayBaseVlm": "377459.2399999999"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.37621",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.3754",
//             "markPx": "0.37621",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "16508594.8000000007",
//             "prevDayPx": "0.31781",
//             "dayNtlVlm": "241269.5203880001",
//             "premium": "0.0009051139",
//             "oraclePx": "0.33145",
//             "markPx": "0.33179",
//             "midPx": "0.33209",
//             "impactPxs": [
//                 "0.33175",
//                 "0.33247"
//             ],
//             "dayBaseVlm": "733102.1000000002"
//         },
//         {
//             "funding": "0.0000100487",
//             "openInterest": "3207738.2600000016",
//             "prevDayPx": "159.43",
//             "dayNtlVlm": "487949170.1526997089",
//             "premium": "-0.0005408004",
//             "oraclePx": "166.42",
//             "markPx": "166.34",
//             "midPx": "166.325",
//             "impactPxs": [
//                 "166.32",
//                 "166.33"
//             ],
//             "dayBaseVlm": "2970254.0799999936"
//         },
//         {
//             "funding": "-0.000007803",
//             "openInterest": "1071432.1199999996",
//             "prevDayPx": "17.043",
//             "dayNtlVlm": "9017710.9543300029",
//             "premium": "-0.0005593154",
//             "oraclePx": "17.879",
//             "markPx": "17.869",
//             "midPx": "17.8685",
//             "impactPxs": [
//                 "17.8677",
//                 "17.869"
//             ],
//             "dayBaseVlm": "506755.2999999999"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "53587.168",
//             "prevDayPx": "986.7",
//             "dayNtlVlm": "36861909.6594000161",
//             "premium": "-0.0007020007",
//             "oraclePx": "997.15",
//             "markPx": "996.56",
//             "midPx": "996.445",
//             "impactPxs": [
//                 "996.126",
//                 "996.45"
//             ],
//             "dayBaseVlm": "36851.914"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "4483295.2000000002",
//             "prevDayPx": "0.38532",
//             "dayNtlVlm": "771225.8939249999",
//             "premium": "-0.0005344871",
//             "oraclePx": "0.3929",
//             "markPx": "0.39257",
//             "midPx": "0.39259",
//             "impactPxs": [
//                 "0.3925",
//                 "0.39269"
//             ],
//             "dayBaseVlm": "1948034.2000000002"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "10877039.2000000011",
//             "prevDayPx": "0.40963",
//             "dayNtlVlm": "935331.0087060002",
//             "premium": "-0.0003712728",
//             "oraclePx": "0.43095",
//             "markPx": "0.43069",
//             "midPx": "0.43068",
//             "impactPxs": [
//                 "0.43058",
//                 "0.43079"
//             ],
//             "dayBaseVlm": "2202684.2000000007"
//         },
//         {
//             "funding": "0.0000143212",
//             "openInterest": "513225.0399999997",
//             "prevDayPx": "100.98",
//             "dayNtlVlm": "107707740.2102600783",
//             "premium": "0.0004707201",
//             "oraclePx": "107.07",
//             "markPx": "107.16",
//             "midPx": "107.14",
//             "impactPxs": [
//                 "107.1204",
//                 "107.1635"
//             ],
//             "dayBaseVlm": "1004218.7999999989"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "16200032.8000000007",
//             "prevDayPx": "0.28987",
//             "dayNtlVlm": "3098193.9676089999",
//             "premium": "-0.0006052455",
//             "oraclePx": "0.2974",
//             "markPx": "0.29719",
//             "midPx": "0.29717",
//             "impactPxs": [
//                 "0.29711",
//                 "0.29722"
//             ],
//             "dayBaseVlm": "10376960.0999999959"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "354794432.0",
//             "prevDayPx": "0.17413",
//             "dayNtlVlm": "61271258.6841000319",
//             "premium": "0.0000554139",
//             "oraclePx": "0.18046",
//             "markPx": "0.18048",
//             "midPx": "0.180475",
//             "impactPxs": [
//                 "0.18047",
//                 "0.180489"
//             ],
//             "dayBaseVlm": "342331784.0"
//         },
//         {
//             "funding": "-0.0000052679",
//             "openInterest": "270611.4",
//             "prevDayPx": "7.2193",
//             "dayNtlVlm": "2916513.0569799994",
//             "premium": "-0.0006065036",
//             "oraclePx": "7.7988",
//             "markPx": "7.7911",
//             "midPx": "7.79185",
//             "impactPxs": [
//                 "7.78962",
//                 "7.79407"
//             ],
//             "dayBaseVlm": "379965.8000000002"
//         },
//         {
//             "funding": "0.0000077118",
//             "openInterest": "10286953.3999999948",
//             "prevDayPx": "2.0913",
//             "dayNtlVlm": "11584149.950179996",
//             "premium": "-0.0004312943",
//             "oraclePx": "2.1563",
//             "markPx": "2.1553",
//             "midPx": "2.1552",
//             "impactPxs": [
//                 "2.15493",
//                 "2.15537"
//             ],
//             "dayBaseVlm": "5414582.8000000026"
//         },
//         {
//             "funding": "0.0000066445",
//             "openInterest": "4703313042.0",
//             "prevDayPx": "0.005973",
//             "dayNtlVlm": "6813899.3086299999",
//             "premium": "-0.0004922067",
//             "oraclePx": "0.006095",
//             "markPx": "0.006092",
//             "midPx": "0.006091",
//             "impactPxs": [
//                 "0.00609",
//                 "0.006092"
//             ],
//             "dayBaseVlm": "1121205950.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "17762727.4000000022",
//             "prevDayPx": "0.47377",
//             "dayNtlVlm": "3226132.8132730001",
//             "premium": "-0.0002054021",
//             "oraclePx": "0.48685",
//             "markPx": "0.48655",
//             "midPx": "0.48649",
//             "impactPxs": [
//                 "0.48631",
//                 "0.48675"
//             ],
//             "dayBaseVlm": "6667462.9000000013"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "8622005.3999999985",
//             "prevDayPx": "0.79676",
//             "dayNtlVlm": "1567045.900001999",
//             "premium": "-0.0000238138",
//             "oraclePx": "0.83985",
//             "markPx": "0.83979",
//             "midPx": "0.83971",
//             "impactPxs": [
//                 "0.83963",
//                 "0.83983"
//             ],
//             "dayBaseVlm": "1878468.7000000004"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "2469347.9999999995",
//             "prevDayPx": "15.321",
//             "dayNtlVlm": "17551926.3441000171",
//             "premium": "0.0004949215",
//             "oraclePx": "16.245",
//             "markPx": "16.255",
//             "midPx": "16.255",
//             "impactPxs": [
//                 "16.25304",
//                 "16.25679"
//             ],
//             "dayBaseVlm": "1097880.5999999999"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "2216445.2000000002",
//             "prevDayPx": "0.39578",
//             "dayNtlVlm": "675915.4611390004",
//             "premium": "-0.0006396285",
//             "oraclePx": "0.42212",
//             "markPx": "0.42178",
//             "midPx": "0.42169",
//             "impactPxs": [
//                 "0.42155",
//                 "0.42185"
//             ],
//             "dayBaseVlm": "1587956.2999999996"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "6.8946",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "7.03",
//             "markPx": "6.8952",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "-0.0000322324",
//             "openInterest": "29682084.0",
//             "prevDayPx": "0.098787",
//             "dayNtlVlm": "1572535.5511890007",
//             "premium": "-0.0012003833",
//             "oraclePx": "0.099135",
//             "markPx": "0.099017",
//             "midPx": "0.099001",
//             "impactPxs": [
//                 "0.098972",
//                 "0.099016"
//             ],
//             "dayBaseVlm": "15799773.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.73",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.712",
//             "markPx": "0.71688",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "-0.0000008638",
//             "openInterest": "101034.36",
//             "prevDayPx": "9.4066",
//             "dayNtlVlm": "506896.0547100001",
//             "premium": "0.0",
//             "oraclePx": "10.16",
//             "markPx": "10.156",
//             "midPx": "10.159",
//             "impactPxs": [
//                 "10.1552",
//                 "10.1618"
//             ],
//             "dayBaseVlm": "50648.5"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "11798267.4000000004",
//             "prevDayPx": "0.8218",
//             "dayNtlVlm": "1288261.9186720003",
//             "premium": "0.0",
//             "oraclePx": "0.8685",
//             "markPx": "0.86796",
//             "midPx": "0.86814",
//             "impactPxs": [
//                 "0.86766",
//                 "0.86865"
//             ],
//             "dayBaseVlm": "1517755.8000000005"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "36561556.0",
//             "prevDayPx": "2.2802",
//             "dayNtlVlm": "107220680.7165999264",
//             "premium": "0.0001210117",
//             "oraclePx": "2.4791",
//             "markPx": "2.4785",
//             "midPx": "2.47945",
//             "impactPxs": [
//                 "2.4794",
//                 "2.4795"
//             ],
//             "dayBaseVlm": "45248079.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "4821.43",
//             "prevDayPx": "491.53",
//             "dayNtlVlm": "1634619.5040499999",
//             "premium": "-0.0003454743",
//             "oraclePx": "506.55",
//             "markPx": "506.29",
//             "midPx": "506.22",
//             "impactPxs": [
//                 "506.109",
//                 "506.375"
//             ],
//             "dayBaseVlm": "3264.805"
//         },
//         {
//             "funding": "-0.0000320513",
//             "openInterest": "4829715.54",
//             "prevDayPx": "2.9968",
//             "dayNtlVlm": "11724009.046443997",
//             "premium": "-0.0009965743",
//             "oraclePx": "3.211",
//             "markPx": "3.2075",
//             "midPx": "3.207",
//             "impactPxs": [
//                 "3.2066",
//                 "3.2078"
//             ],
//             "dayBaseVlm": "3705335.6999999997"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "113378.22",
//             "prevDayPx": "200.04",
//             "dayNtlVlm": "7622406.3378000036",
//             "premium": "0.0",
//             "oraclePx": "219.71",
//             "markPx": "219.71",
//             "midPx": "219.715",
//             "impactPxs": [
//                 "219.6373",
//                 "219.7221"
//             ],
//             "dayBaseVlm": "35952.35"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "16969.26",
//             "prevDayPx": "33.473",
//             "dayNtlVlm": "139465.8232",
//             "premium": "-0.0004920658",
//             "oraclePx": "34.345",
//             "markPx": "34.318",
//             "midPx": "34.323",
//             "impactPxs": [
//                 "34.2963",
//                 "34.3281"
//             ],
//             "dayBaseVlm": "4042.15"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "1843.8",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "1828.8",
//             "markPx": "1831.5",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "7960093.5999999996",
//             "prevDayPx": "0.79286",
//             "dayNtlVlm": "3854452.0270530018",
//             "premium": "-0.00039807",
//             "oraclePx": "0.829",
//             "markPx": "0.82825",
//             "midPx": "0.82848",
//             "impactPxs": [
//                 "0.82845",
//                 "0.82867"
//             ],
//             "dayBaseVlm": "4731029.6000000015"
//         },
//         {
//             "funding": "0.0000034382",
//             "openInterest": "426600.2",
//             "prevDayPx": "1.2155",
//             "dayNtlVlm": "1288528.6089899996",
//             "premium": "-0.0013381701",
//             "oraclePx": "1.2405",
//             "markPx": "1.2387",
//             "midPx": "1.2385",
//             "impactPxs": [
//                 "1.23815",
//                 "1.23884"
//             ],
//             "dayBaseVlm": "1040928.4"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.041845",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.042203",
//             "markPx": "0.041873",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.071",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.071068",
//             "markPx": "0.071",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "7.636",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "6.053",
//             "markPx": "7.636",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "-0.000047941",
//             "openInterest": "1981190.0",
//             "prevDayPx": "0.11444",
//             "dayNtlVlm": "398581.4199199999",
//             "premium": "-0.0011082694",
//             "oraclePx": "0.1173",
//             "markPx": "0.11713",
//             "midPx": "0.11708",
//             "impactPxs": [
//                 "0.11703",
//                 "0.11717"
//             ],
//             "dayBaseVlm": "3389101.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "34101052.0",
//             "prevDayPx": "0.2894",
//             "dayNtlVlm": "1971943.2293500004",
//             "premium": "-0.0003269978",
//             "oraclePx": "0.29358",
//             "markPx": "0.29344",
//             "midPx": "0.293455",
//             "impactPxs": [
//                 "0.293433",
//                 "0.293484"
//             ],
//             "dayBaseVlm": "6759239.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "591099820.0",
//             "prevDayPx": "0.009841",
//             "dayNtlVlm": "1841287.1344079995",
//             "premium": "-0.0005020584",
//             "oraclePx": "0.009959",
//             "markPx": "0.009953",
//             "midPx": "0.009952",
//             "impactPxs": [
//                 "0.00995",
//                 "0.009954"
//             ],
//             "dayBaseVlm": "184482352.0"
//         },
//         {
//             "funding": "0.0000320324",
//             "openInterest": "5661998.5999999987",
//             "prevDayPx": "5.904",
//             "dayNtlVlm": "18381052.4485000037",
//             "premium": "0.0008215945",
//             "oraclePx": "6.7795",
//             "markPx": "6.7852",
//             "midPx": "6.78585",
//             "impactPxs": [
//                 "6.78507",
//                 "6.78861"
//             ],
//             "dayBaseVlm": "2793149.0999999992"
//         },
//         {
//             "funding": "-0.000043312",
//             "openInterest": "23389122.0",
//             "prevDayPx": "0.17351",
//             "dayNtlVlm": "3209150.276730001",
//             "premium": "-0.0013496251",
//             "oraclePx": "0.18005",
//             "markPx": "0.17981",
//             "midPx": "0.179775",
//             "impactPxs": [
//                 "0.179746",
//                 "0.179807"
//             ],
//             "dayBaseVlm": "17741437.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "2735848.7999999998",
//             "prevDayPx": "0.7994",
//             "dayNtlVlm": "417753.051024",
//             "premium": "-0.0005935796",
//             "oraclePx": "0.8255",
//             "markPx": "0.82501",
//             "midPx": "0.82475",
//             "impactPxs": [
//                 "0.82456",
//                 "0.82501"
//             ],
//             "dayBaseVlm": "512684.7000000001"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.007459",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.004982",
//             "markPx": "0.007459",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "4.72",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.47734",
//             "markPx": "4.72",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.000604",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.000209",
//             "markPx": "0.000621",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "1.0454",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "1.043",
//             "markPx": "1.0415",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "-0.0001654147",
//             "openInterest": "3992430.7999999984",
//             "prevDayPx": "1.6442",
//             "dayNtlVlm": "6975093.5763199963",
//             "premium": "-0.0015018623",
//             "oraclePx": "1.6646",
//             "markPx": "1.6617",
//             "midPx": "1.6619",
//             "impactPxs": [
//                 "1.66154",
//                 "1.6621"
//             ],
//             "dayBaseVlm": "4105891.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.070239",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.070028",
//             "markPx": "0.070043",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000091724",
//             "openInterest": "2000861.0",
//             "prevDayPx": "3.1145",
//             "dayNtlVlm": "3336800.7287700004",
//             "premium": "-0.0007705692",
//             "oraclePx": "3.2184",
//             "markPx": "3.2155",
//             "midPx": "3.21555",
//             "impactPxs": [
//                 "3.21412",
//                 "3.21592"
//             ],
//             "dayBaseVlm": "1042310.5000000003"
//         },
//         {
//             "funding": "-0.0000737926",
//             "openInterest": "32764.6",
//             "prevDayPx": "10.504",
//             "dayNtlVlm": "766790.7229000003",
//             "premium": "-0.0014004508",
//             "oraclePx": "9.9825",
//             "markPx": "9.9637",
//             "midPx": "9.96595",
//             "impactPxs": [
//                 "9.96297",
//                 "9.96852"
//             ],
//             "dayBaseVlm": "75361.3"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "12401.6",
//             "prevDayPx": "25.787",
//             "dayNtlVlm": "349789.4466999999",
//             "premium": "-0.0005717531",
//             "oraclePx": "26.41",
//             "markPx": "26.39",
//             "midPx": "26.39",
//             "impactPxs": [
//                 "26.3842",
//                 "26.3949"
//             ],
//             "dayBaseVlm": "13291.94"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "153997.6",
//             "prevDayPx": "0.78437",
//             "dayNtlVlm": "28463.512943",
//             "premium": "0.0",
//             "oraclePx": "0.801",
//             "markPx": "0.803",
//             "midPx": "0.80176",
//             "impactPxs": [
//                 "0.7985",
//                 "0.80532"
//             ],
//             "dayBaseVlm": "35778.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.07116",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.071385",
//             "markPx": "0.071161",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.0558",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.0558",
//             "markPx": "0.055683",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.018338",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.018338",
//             "markPx": "0.018328",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "-0.0000917057",
//             "openInterest": "740888.0",
//             "prevDayPx": "0.31123",
//             "dayNtlVlm": "91239.09962",
//             "premium": "-0.0006166829",
//             "oraclePx": "0.3081",
//             "markPx": "0.30771",
//             "midPx": "0.30781",
//             "impactPxs": [
//                 "0.307617",
//                 "0.30791"
//             ],
//             "dayBaseVlm": "295468.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.38457",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.38349",
//             "markPx": "0.38287",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.066",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.06598",
//             "markPx": "0.066",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.15473",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.15423",
//             "markPx": "0.15473",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "12456212.0",
//             "prevDayPx": "0.029579",
//             "dayNtlVlm": "167572.734631",
//             "premium": "-0.0008817766",
//             "oraclePx": "0.03062",
//             "markPx": "0.030589",
//             "midPx": "0.030583",
//             "impactPxs": [
//                 "0.030573",
//                 "0.030593"
//             ],
//             "dayBaseVlm": "5460367.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "12678898.0",
//             "prevDayPx": "0.049882",
//             "dayNtlVlm": "1197691.7379340006",
//             "premium": "0.0",
//             "oraclePx": "0.05263",
//             "markPx": "0.052621",
//             "midPx": "0.052617",
//             "impactPxs": [
//                 "0.052599",
//                 "0.052639"
//             ],
//             "dayBaseVlm": "23232271.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.01939",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.019393",
//             "markPx": "0.01939",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "19068242.0",
//             "prevDayPx": "0.045741",
//             "dayNtlVlm": "364325.129705",
//             "premium": "-0.0009743698",
//             "oraclePx": "0.04721",
//             "markPx": "0.047156",
//             "midPx": "0.04715",
//             "impactPxs": [
//                 "0.047136",
//                 "0.047164"
//             ],
//             "dayBaseVlm": "7733017.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "6377094.9999999991",
//             "prevDayPx": "0.99466",
//             "dayNtlVlm": "11380995.4166949969",
//             "premium": "0.0",
//             "oraclePx": "1.0084",
//             "markPx": "1.0083",
//             "midPx": "1.00845",
//             "impactPxs": [
//                 "1.00831",
//                 "1.00868"
//             ],
//             "dayBaseVlm": "11082494.0999999996"
//         },
//         {
//             "funding": "0.0000109636",
//             "openInterest": "52638.8",
//             "prevDayPx": "24.061",
//             "dayNtlVlm": "579133.1644399998",
//             "premium": "-0.0019140705",
//             "oraclePx": "24.555",
//             "markPx": "24.489",
//             "midPx": "24.487",
//             "impactPxs": [
//                 "24.4721",
//                 "24.508"
//             ],
//             "dayBaseVlm": "23365.61"
//         },
//         {
//             "funding": "0.0000106959",
//             "openInterest": "17184076.0",
//             "prevDayPx": "0.55979",
//             "dayNtlVlm": "10097326.0120200012",
//             "premium": "-0.0003330487",
//             "oraclePx": "0.5855",
//             "markPx": "0.58525",
//             "midPx": "0.58525",
//             "impactPxs": [
//                 "0.585143",
//                 "0.585305"
//             ],
//             "dayBaseVlm": "17587520.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "7179626.4000000004",
//             "prevDayPx": "2.0567",
//             "dayNtlVlm": "3685055.7632800001",
//             "premium": "0.000564573",
//             "oraclePx": "2.1255",
//             "markPx": "2.1265",
//             "midPx": "2.127",
//             "impactPxs": [
//                 "2.1267",
//                 "2.12729"
//             ],
//             "dayBaseVlm": "1739555.9999999995"
//         },
//         {
//             "funding": "0.0000088185",
//             "openInterest": "8574670.0",
//             "prevDayPx": "0.16027",
//             "dayNtlVlm": "854689.2538800003",
//             "premium": "-0.002008057",
//             "oraclePx": "0.16135",
//             "markPx": "0.16094",
//             "midPx": "0.16093",
//             "impactPxs": [
//                 "0.160802",
//                 "0.161026"
//             ],
//             "dayBaseVlm": "5334767.0"
//         },
//         {
//             "funding": "0.0000115861",
//             "openInterest": "3876956.0",
//             "prevDayPx": "0.076641",
//             "dayNtlVlm": "107741.3737169999",
//             "premium": "-0.0009949622",
//             "oraclePx": "0.0794",
//             "markPx": "0.07931",
//             "midPx": "0.079298",
//             "impactPxs": [
//                 "0.07926",
//                 "0.079321"
//             ],
//             "dayBaseVlm": "1370809.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "115980.4",
//             "prevDayPx": "2.4389",
//             "dayNtlVlm": "71552.64522",
//             "premium": "-0.0006070596",
//             "oraclePx": "2.4874",
//             "markPx": "2.4838",
//             "midPx": "2.48435",
//             "impactPxs": [
//                 "2.48288",
//                 "2.48589"
//             ],
//             "dayBaseVlm": "28753.4"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "2112910.0",
//             "prevDayPx": "2.7137",
//             "dayNtlVlm": "1059055.0070999998",
//             "premium": "-0.0010665016",
//             "oraclePx": "2.827",
//             "markPx": "2.8238",
//             "midPx": "2.82375",
//             "impactPxs": [
//                 "2.823316",
//                 "2.823985"
//             ],
//             "dayBaseVlm": "377462.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.1537",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.15385",
//             "markPx": "0.15387",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "-0.0000040628",
//             "openInterest": "5899258.0",
//             "prevDayPx": "0.29088",
//             "dayNtlVlm": "6835382.723260005",
//             "premium": "-0.0026495822",
//             "oraclePx": "0.35666",
//             "markPx": "0.35623",
//             "midPx": "0.35556",
//             "impactPxs": [
//                 "0.355292",
//                 "0.355715"
//             ],
//             "dayBaseVlm": "20900682.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "1.5848",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.0416",
//             "markPx": "1.5829",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "21285886.1999999993",
//             "prevDayPx": "2.7183",
//             "dayNtlVlm": "79146488.9699601084",
//             "premium": "0.0001623816",
//             "oraclePx": "2.956",
//             "markPx": "2.9559",
//             "midPx": "2.95745",
//             "impactPxs": [
//                 "2.95648",
//                 "2.95756"
//             ],
//             "dayBaseVlm": "26974371.3000000156"
//         },
//         {
//             "funding": "-0.0000015993",
//             "openInterest": "749338662.0",
//             "prevDayPx": "0.001589",
//             "dayNtlVlm": "379286.1559950002",
//             "premium": "-0.0012055455",
//             "oraclePx": "0.001659",
//             "markPx": "0.001657",
//             "midPx": "0.001656",
//             "impactPxs": [
//                 "0.001655",
//                 "0.001657"
//             ],
//             "dayBaseVlm": "232290653.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "170129.4",
//             "prevDayPx": "5.0647",
//             "dayNtlVlm": "774152.0975519996",
//             "premium": "-0.0004276827",
//             "oraclePx": "5.144",
//             "markPx": "5.1414",
//             "midPx": "5.1407",
//             "impactPxs": [
//                 "5.1398",
//                 "5.1418"
//             ],
//             "dayBaseVlm": "149336.07"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.89235",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.88959",
//             "markPx": "0.88823",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "100872.54",
//             "prevDayPx": "5.3583",
//             "dayNtlVlm": "307750.437547",
//             "premium": "-0.000225221",
//             "oraclePx": "5.3281",
//             "markPx": "5.3265",
//             "midPx": "5.3259",
//             "impactPxs": [
//                 "5.3248",
//                 "5.3269"
//             ],
//             "dayBaseVlm": "56848.57"
//         },
//         {
//             "funding": "0.0000014287",
//             "openInterest": "135999.6",
//             "prevDayPx": "14.72",
//             "dayNtlVlm": "7630213.8812700016",
//             "premium": "-0.0004779857",
//             "oraclePx": "14.854",
//             "markPx": "14.838",
//             "midPx": "14.844",
//             "impactPxs": [
//                 "14.8394",
//                 "14.8469"
//             ],
//             "dayBaseVlm": "481992.6500000003"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "4961929.0",
//             "prevDayPx": "2.5992",
//             "dayNtlVlm": "30792318.8575199731",
//             "premium": "0.0",
//             "oraclePx": "2.608",
//             "markPx": "2.6069",
//             "midPx": "2.6078",
//             "impactPxs": [
//                 "2.60627",
//                 "2.60865"
//             ],
//             "dayBaseVlm": "11161403.9000000078"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "11884784.0",
//             "prevDayPx": "0.099814",
//             "dayNtlVlm": "787901.0461999999",
//             "premium": "-0.0002978684",
//             "oraclePx": "0.10743",
//             "markPx": "0.10738",
//             "midPx": "0.107375",
//             "impactPxs": [
//                 "0.107355",
//                 "0.107398"
//             ],
//             "dayBaseVlm": "7411772.0"
//         },
//         {
//             "funding": "-0.0000015888",
//             "openInterest": "707054.6",
//             "prevDayPx": "0.47741",
//             "dayNtlVlm": "670620.3087909999",
//             "premium": "-0.0008457511",
//             "oraclePx": "0.4966",
//             "markPx": "0.4961",
//             "midPx": "0.49602",
//             "impactPxs": [
//                 "0.49583",
//                 "0.49618"
//             ],
//             "dayBaseVlm": "1351067.2000000004"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "12.128",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "12.14",
//             "markPx": "12.118",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "16175277.0",
//             "prevDayPx": "0.41844",
//             "dayNtlVlm": "313363.1862850002",
//             "premium": "0.0003172805",
//             "oraclePx": "0.44125",
//             "markPx": "0.44154",
//             "midPx": "0.44164",
//             "impactPxs": [
//                 "0.44139",
//                 "0.44168"
//             ],
//             "dayBaseVlm": "722135.8999999998"
//         },
//         {
//             "funding": "0.0000017183",
//             "openInterest": "423732338.0",
//             "prevDayPx": "0.012625",
//             "dayNtlVlm": "5176055.004062999",
//             "premium": "-0.0004570732",
//             "oraclePx": "0.013127",
//             "markPx": "0.013118",
//             "midPx": "0.013119",
//             "impactPxs": [
//                 "0.013117",
//                 "0.013121"
//             ],
//             "dayBaseVlm": "400953111.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "9252440.0",
//             "prevDayPx": "0.02366",
//             "dayNtlVlm": "105414.020646",
//             "premium": "-0.000707449",
//             "oraclePx": "0.02403",
//             "markPx": "0.024007",
//             "midPx": "0.024003",
//             "impactPxs": [
//                 "0.023994",
//                 "0.024013"
//             ],
//             "dayBaseVlm": "4375363.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "1968926.0",
//             "prevDayPx": "0.31006",
//             "dayNtlVlm": "144727.60538",
//             "premium": "-0.0006457565",
//             "oraclePx": "0.3252",
//             "markPx": "0.32494",
//             "midPx": "0.32488",
//             "impactPxs": [
//                 "0.32478",
//                 "0.32499"
//             ],
//             "dayBaseVlm": "450535.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "84078308.0",
//             "prevDayPx": "0.007264",
//             "dayNtlVlm": "183420.0802980001",
//             "premium": "0.0",
//             "oraclePx": "0.00742",
//             "markPx": "0.007418",
//             "midPx": "0.007419",
//             "impactPxs": [
//                 "0.007416",
//                 "0.007424"
//             ],
//             "dayBaseVlm": "24823966.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "6.6683",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "2518.4",
//             "markPx": "6.5609",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "-0.0000014573",
//             "openInterest": "18746238.0",
//             "prevDayPx": "0.34012",
//             "dayNtlVlm": "3302812.02324",
//             "premium": "-0.0005847624",
//             "oraclePx": "0.3557",
//             "markPx": "0.35535",
//             "midPx": "0.355415",
//             "impactPxs": [
//                 "0.355322",
//                 "0.355492"
//             ],
//             "dayBaseVlm": "9431958.0"
//         },
//         {
//             "funding": "-0.0000321659",
//             "openInterest": "6753024.0",
//             "prevDayPx": "0.037166",
//             "dayNtlVlm": "64440.2395",
//             "premium": "-0.0008263361",
//             "oraclePx": "0.037515",
//             "markPx": "0.037471",
//             "midPx": "0.037467",
//             "impactPxs": [
//                 "0.037452",
//                 "0.037484"
//             ],
//             "dayBaseVlm": "1718278.0"
//         },
//         {
//             "funding": "0.0000059069",
//             "openInterest": "58521494.0",
//             "prevDayPx": "0.004537",
//             "dayNtlVlm": "120442.569947",
//             "premium": "-0.0006407518",
//             "oraclePx": "0.004682",
//             "markPx": "0.004678",
//             "midPx": "0.004678",
//             "impactPxs": [
//                 "0.004676",
//                 "0.004679"
//             ],
//             "dayBaseVlm": "25821463.0"
//         },
//         {
//             "funding": "0.0000109541",
//             "openInterest": "92246276.0",
//             "prevDayPx": "0.00982",
//             "dayNtlVlm": "225373.0819429999",
//             "premium": "-0.0004852014",
//             "oraclePx": "0.010305",
//             "markPx": "0.010292",
//             "midPx": "0.010292",
//             "impactPxs": [
//                 "0.010285",
//                 "0.0103"
//             ],
//             "dayBaseVlm": "22074115.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "3411614.0",
//             "prevDayPx": "0.76928",
//             "dayNtlVlm": "934271.1670499999",
//             "premium": "-0.0006159351",
//             "oraclePx": "0.78255",
//             "markPx": "0.78188",
//             "midPx": "0.78176",
//             "impactPxs": [
//                 "0.781597",
//                 "0.782068"
//             ],
//             "dayBaseVlm": "1181942.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.12371",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.12385",
//             "markPx": "0.1238",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "854594.7999999999",
//             "prevDayPx": "0.2705",
//             "dayNtlVlm": "186754.7412249999",
//             "premium": "0.0007064641",
//             "oraclePx": "0.2831",
//             "markPx": "0.2835",
//             "midPx": "0.2834",
//             "impactPxs": [
//                 "0.2833",
//                 "0.2835"
//             ],
//             "dayBaseVlm": "658935.7099999998"
//         },
//         {
//             "funding": "-0.0000495604",
//             "openInterest": "16569168.0",
//             "prevDayPx": "0.032552",
//             "dayNtlVlm": "563170.6899920002",
//             "premium": "-0.0023830282",
//             "oraclePx": "0.03441",
//             "markPx": "0.03433",
//             "midPx": "0.034308",
//             "impactPxs": [
//                 "0.034292",
//                 "0.034328"
//             ],
//             "dayBaseVlm": "16305093.0"
//         },
//         {
//             "funding": "0.000012199",
//             "openInterest": "12013966.0",
//             "prevDayPx": "0.46508",
//             "dayNtlVlm": "4803054.468419998",
//             "premium": "-0.0005417096",
//             "oraclePx": "0.4855",
//             "markPx": "0.4851",
//             "midPx": "0.485095",
//             "impactPxs": [
//                 "0.48501",
//                 "0.485237"
//             ],
//             "dayBaseVlm": "10017613.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "1338968.8",
//             "prevDayPx": "2.5247",
//             "dayNtlVlm": "1362781.8145999997",
//             "premium": "-0.0002406312",
//             "oraclePx": "2.535",
//             "markPx": "2.5343",
//             "midPx": "2.5341",
//             "impactPxs": [
//                 "2.53382",
//                 "2.53439"
//             ],
//             "dayBaseVlm": "536600.3999999999"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "25403890.0",
//             "prevDayPx": "0.011933",
//             "dayNtlVlm": "247390.1707829999",
//             "premium": "0.0",
//             "oraclePx": "0.012073",
//             "markPx": "0.012073",
//             "midPx": "0.012077",
//             "impactPxs": [
//                 "0.012069",
//                 "0.012083"
//             ],
//             "dayBaseVlm": "19967404.0"
//         },
//         {
//             "funding": "0.0000081393",
//             "openInterest": "112184.94",
//             "prevDayPx": "13.445",
//             "dayNtlVlm": "240920.1081400001",
//             "premium": "-0.0009433962",
//             "oraclePx": "14.31",
//             "markPx": "14.297",
//             "midPx": "14.2935",
//             "impactPxs": [
//                 "14.2907",
//                 "14.2965"
//             ],
//             "dayBaseVlm": "17116.61"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "277784.7399999999",
//             "prevDayPx": "15.998",
//             "dayNtlVlm": "1178656.4889899998",
//             "premium": "0.0",
//             "oraclePx": "16.15",
//             "markPx": "16.148",
//             "midPx": "16.1475",
//             "impactPxs": [
//                 "16.1464",
//                 "16.1509"
//             ],
//             "dayBaseVlm": "73044.66"
//         },
//         {
//             "funding": "-0.0000346222",
//             "openInterest": "46801313.6000000015",
//             "prevDayPx": "0.02206",
//             "dayNtlVlm": "119753.469237",
//             "premium": "-0.0012760527",
//             "oraclePx": "0.02351",
//             "markPx": "0.02348",
//             "midPx": "0.02347",
//             "impactPxs": [
//                 "0.02347",
//                 "0.02348"
//             ],
//             "dayBaseVlm": "5206765.5999999987"
//         },
//         {
//             "funding": "-0.0000686067",
//             "openInterest": "1898206.2",
//             "prevDayPx": "0.10587",
//             "dayNtlVlm": "583485.0365669997",
//             "premium": "-0.0006151142",
//             "oraclePx": "0.1138",
//             "markPx": "0.11365",
//             "midPx": "0.11366",
//             "impactPxs": [
//                 "0.11359",
//                 "0.11373"
//             ],
//             "dayBaseVlm": "5081286.0999999987"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "1361571.5999999999",
//             "prevDayPx": "1.1934",
//             "dayNtlVlm": "435729.5909500001",
//             "premium": "0.0",
//             "oraclePx": "1.058",
//             "markPx": "1.0564",
//             "midPx": "1.0561",
//             "impactPxs": [
//                 "1.05305",
//                 "1.05825"
//             ],
//             "dayBaseVlm": "405929.6"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "7683306.0",
//             "prevDayPx": "0.64305",
//             "dayNtlVlm": "2889072.3380399994",
//             "premium": "-0.0002961672",
//             "oraclePx": "0.6888",
//             "markPx": "0.68852",
//             "midPx": "0.688445",
//             "impactPxs": [
//                 "0.688351",
//                 "0.688596"
//             ],
//             "dayBaseVlm": "4327055.0"
//         },
//         {
//             "funding": "0.0000104075",
//             "openInterest": "43357504.0",
//             "prevDayPx": "0.017017",
//             "dayNtlVlm": "233171.7055629999",
//             "premium": "-0.0003436426",
//             "oraclePx": "0.01746",
//             "markPx": "0.017453",
//             "midPx": "0.017444",
//             "impactPxs": [
//                 "0.017435",
//                 "0.017454"
//             ],
//             "dayBaseVlm": "13403687.0"
//         },
//         {
//             "funding": "-0.0000441843",
//             "openInterest": "8539916.8000000007",
//             "prevDayPx": "0.10728",
//             "dayNtlVlm": "202132.6813770001",
//             "premium": "-0.0013593113",
//             "oraclePx": "0.11035",
//             "markPx": "0.11017",
//             "midPx": "0.11015",
//             "impactPxs": [
//                 "0.11006",
//                 "0.1102"
//             ],
//             "dayBaseVlm": "1832486.7000000004"
//         },
//         {
//             "funding": "-0.000004839",
//             "openInterest": "8007885.4000000004",
//             "prevDayPx": "0.09526",
//             "dayNtlVlm": "304086.486048",
//             "premium": "-0.0008398068",
//             "oraclePx": "0.09526",
//             "markPx": "0.09516",
//             "midPx": "0.09516",
//             "impactPxs": [
//                 "0.09515",
//                 "0.09518"
//             ],
//             "dayBaseVlm": "3127022.3999999999"
//         },
//         {
//             "funding": "-0.0000548036",
//             "openInterest": "6869212.0",
//             "prevDayPx": "0.0906",
//             "dayNtlVlm": "437646.5494270001",
//             "premium": "-0.0012191067",
//             "oraclePx": "0.09023",
//             "markPx": "0.09006",
//             "midPx": "0.09004",
//             "impactPxs": [
//                 "0.08991",
//                 "0.09012"
//             ],
//             "dayBaseVlm": "4779646.4000000013"
//         },
//         {
//             "funding": "0.0000081274",
//             "openInterest": "39638492.799999997",
//             "prevDayPx": "0.06182",
//             "dayNtlVlm": "379145.923092",
//             "premium": "-0.000607718",
//             "oraclePx": "0.06582",
//             "markPx": "0.06576",
//             "midPx": "0.06576",
//             "impactPxs": [
//                 "0.06574",
//                 "0.06578"
//             ],
//             "dayBaseVlm": "5852555.4000000013"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "3012.0",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "1706.6",
//             "markPx": "3026.3",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0001130256",
//             "openInterest": "263593206.6000000238",
//             "prevDayPx": "0.13769",
//             "dayNtlVlm": "43710308.3242360428",
//             "premium": "0.0010067114",
//             "oraclePx": "0.1788",
//             "markPx": "0.17912",
//             "midPx": "0.17908",
//             "impactPxs": [
//                 "0.17898",
//                 "0.17925"
//             ],
//             "dayBaseVlm": "277582161.0999999642"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.024008",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.024025",
//             "markPx": "0.024027",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.12565",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.1257",
//             "markPx": "0.12555",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "67612.37",
//             "prevDayPx": "359.96",
//             "dayNtlVlm": "16225481.0748100039",
//             "premium": "-0.0003461906",
//             "oraclePx": "387.07",
//             "markPx": "387.07",
//             "midPx": "386.88",
//             "impactPxs": [
//                 "386.824",
//                 "386.936"
//             ],
//             "dayBaseVlm": "42775.173"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "434280.0999999999",
//             "prevDayPx": "5.4916",
//             "dayNtlVlm": "4968941.2378240013",
//             "premium": "-0.0011837456",
//             "oraclePx": "5.66",
//             "markPx": "5.652",
//             "midPx": "5.6509",
//             "impactPxs": [
//                 "5.6492",
//                 "5.6533"
//             ],
//             "dayBaseVlm": "871663.37"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.015813",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.015815",
//             "markPx": "0.015813",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "23995200.0",
//             "prevDayPx": "0.061896",
//             "dayNtlVlm": "1081454.3578820003",
//             "premium": "-0.0006766858",
//             "oraclePx": "0.063545",
//             "markPx": "0.06349",
//             "midPx": "0.063491",
//             "impactPxs": [
//                 "0.063484",
//                 "0.063502"
//             ],
//             "dayBaseVlm": "17144366.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "260995918.0",
//             "prevDayPx": "0.00091",
//             "dayNtlVlm": "113812.3883680001",
//             "premium": "-0.0010341262",
//             "oraclePx": "0.000967",
//             "markPx": "0.000966",
//             "midPx": "0.000965",
//             "impactPxs": [
//                 "0.000964",
//                 "0.000966"
//             ],
//             "dayBaseVlm": "120581836.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "6216037.2000000002",
//             "prevDayPx": "0.93925",
//             "dayNtlVlm": "887467.7818709998",
//             "premium": "-0.0005362473",
//             "oraclePx": "0.98835",
//             "markPx": "0.98765",
//             "midPx": "0.98753",
//             "impactPxs": [
//                 "0.9874",
//                 "0.98782"
//             ],
//             "dayBaseVlm": "898250.2000000002"
//         },
//         {
//             "funding": "0.0000068694",
//             "openInterest": "185973566.0",
//             "prevDayPx": "0.31848",
//             "dayNtlVlm": "24202758.5665300079",
//             "premium": "-0.0003234818",
//             "oraclePx": "0.34005",
//             "markPx": "0.33988",
//             "midPx": "0.339925",
//             "impactPxs": [
//                 "0.339834",
//                 "0.33994"
//             ],
//             "dayBaseVlm": "72623776.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "7368603.3999999994",
//             "prevDayPx": "1.3166",
//             "dayNtlVlm": "2904092.27495",
//             "premium": "-0.0000303905",
//             "oraclePx": "1.3162",
//             "markPx": "1.3156",
//             "midPx": "1.31595",
//             "impactPxs": [
//                 "1.31552",
//                 "1.31616"
//             ],
//             "dayBaseVlm": "2171937.2000000016"
//         },
//         {
//             "funding": "0.0000098042",
//             "openInterest": "4366139.2000000002",
//             "prevDayPx": "0.04963",
//             "dayNtlVlm": "210555.9963679999",
//             "premium": "-0.0009532888",
//             "oraclePx": "0.05245",
//             "markPx": "0.05237",
//             "midPx": "0.05238",
//             "impactPxs": [
//                 "0.05237",
//                 "0.0524"
//             ],
//             "dayBaseVlm": "4075172.0000000005"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "12310390.8000000007",
//             "prevDayPx": "0.09634",
//             "dayNtlVlm": "488731.4958649998",
//             "premium": "-0.0003929273",
//             "oraclePx": "0.1018",
//             "markPx": "0.1017",
//             "midPx": "0.10169",
//             "impactPxs": [
//                 "0.10164",
//                 "0.10176"
//             ],
//             "dayBaseVlm": "4831409.9999999991"
//         },
//         {
//             "funding": "-0.0002034832",
//             "openInterest": "9769614.0",
//             "prevDayPx": "0.39883",
//             "dayNtlVlm": "3781963.0435699993",
//             "premium": "-0.0020411582",
//             "oraclePx": "0.35813",
//             "markPx": "0.3575",
//             "midPx": "0.35725",
//             "impactPxs": [
//                 "0.357148",
//                 "0.357399"
//             ],
//             "dayBaseVlm": "10263563.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "13647854.0",
//             "prevDayPx": "0.17117",
//             "dayNtlVlm": "1173315.1314099985",
//             "premium": "0.0",
//             "oraclePx": "0.18736",
//             "markPx": "0.1873",
//             "midPx": "0.18735",
//             "impactPxs": [
//                 "0.18731",
//                 "0.187369"
//             ],
//             "dayBaseVlm": "6610921.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "18753150.0",
//             "prevDayPx": "0.14065",
//             "dayNtlVlm": "3375225.9698599987",
//             "premium": "0.0",
//             "oraclePx": "0.14375",
//             "markPx": "0.14374",
//             "midPx": "0.14375",
//             "impactPxs": [
//                 "0.143677",
//                 "0.143807"
//             ],
//             "dayBaseVlm": "23288266.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "4.1915",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "4.1948",
//             "markPx": "4.1915",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "8038543.9800000004",
//             "prevDayPx": "0.7753",
//             "dayNtlVlm": "7460183.8795589982",
//             "premium": "-0.0004719764",
//             "oraclePx": "0.8475",
//             "markPx": "0.847",
//             "midPx": "0.847",
//             "impactPxs": [
//                 "0.8468",
//                 "0.8471"
//             ],
//             "dayBaseVlm": "8817431.4999999963"
//         },
//         {
//             "funding": "-0.0000363084",
//             "openInterest": "49607886.0",
//             "prevDayPx": "0.007836",
//             "dayNtlVlm": "182173.2398739999",
//             "premium": "-0.0007575758",
//             "oraclePx": "0.00792",
//             "markPx": "0.007912",
//             "midPx": "0.007913",
//             "impactPxs": [
//                 "0.00791",
//                 "0.007914"
//             ],
//             "dayBaseVlm": "22910197.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "373623674.0",
//             "prevDayPx": "0.000754",
//             "dayNtlVlm": "87034.87786",
//             "premium": "0.0",
//             "oraclePx": "0.000743",
//             "markPx": "0.000742",
//             "midPx": "0.000742",
//             "impactPxs": [
//                 "0.000742",
//                 "0.000743"
//             ],
//             "dayBaseVlm": "114097272.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "668380436.0",
//             "prevDayPx": "0.002099",
//             "dayNtlVlm": "362029.3263990002",
//             "premium": "-0.0004638219",
//             "oraclePx": "0.002156",
//             "markPx": "0.002155",
//             "midPx": "0.002154",
//             "impactPxs": [
//                 "0.002154",
//                 "0.002155"
//             ],
//             "dayBaseVlm": "167988025.0"
//         },
//         {
//             "funding": "-0.0000153186",
//             "openInterest": "56016054.0",
//             "prevDayPx": "0.022143",
//             "dayNtlVlm": "1476885.9361019998",
//             "premium": "-0.0012396694",
//             "oraclePx": "0.0242",
//             "markPx": "0.024165",
//             "midPx": "0.024161",
//             "impactPxs": [
//                 "0.024154",
//                 "0.02417"
//             ],
//             "dayBaseVlm": "62008130.0"
//         },
//         {
//             "funding": "-0.0001163132",
//             "openInterest": "2384255.1999999997",
//             "prevDayPx": "0.30319",
//             "dayNtlVlm": "606263.146715",
//             "premium": "-0.0018554023",
//             "oraclePx": "0.32338",
//             "markPx": "0.32261",
//             "midPx": "0.32268",
//             "impactPxs": [
//                 "0.32255",
//                 "0.32278"
//             ],
//             "dayBaseVlm": "1918666.8000000005"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "101238398.0",
//             "prevDayPx": "0.056637",
//             "dayNtlVlm": "7832319.845084006",
//             "premium": "0.0",
//             "oraclePx": "0.06526",
//             "markPx": "0.065125",
//             "midPx": "0.065175",
//             "impactPxs": [
//                 "0.065067",
//                 "0.065261"
//             ],
//             "dayBaseVlm": "133354188.0"
//         },
//         {
//             "funding": "-0.0001118649",
//             "openInterest": "964905624.0",
//             "prevDayPx": "0.001217",
//             "dayNtlVlm": "1229451.9883309999",
//             "premium": "-0.0015847861",
//             "oraclePx": "0.001262",
//             "markPx": "0.001258",
//             "midPx": "0.001259",
//             "impactPxs": [
//                 "0.001257",
//                 "0.00126"
//             ],
//             "dayBaseVlm": "1013819771.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.12493",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.1247",
//             "markPx": "0.12466",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "158334508.0",
//             "prevDayPx": "0.001576",
//             "dayNtlVlm": "110938.188782",
//             "premium": "0.0",
//             "oraclePx": "0.001647",
//             "markPx": "0.001645",
//             "midPx": "0.001646",
//             "impactPxs": [
//                 "0.001644",
//                 "0.001647"
//             ],
//             "dayBaseVlm": "68361634.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "1156205.8",
//             "prevDayPx": "2.231",
//             "dayNtlVlm": "2653884.5104999989",
//             "premium": "0.0",
//             "oraclePx": "2.4615",
//             "markPx": "2.4598",
//             "midPx": "2.46005",
//             "impactPxs": [
//                 "2.4589",
//                 "2.46165"
//             ],
//             "dayBaseVlm": "1104346.4000000001"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.1367",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.13675",
//             "markPx": "0.13659",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "12539526.0",
//             "prevDayPx": "0.17427",
//             "dayNtlVlm": "735118.2435999997",
//             "premium": "-0.0004954077",
//             "oraclePx": "0.17965",
//             "markPx": "0.17956",
//             "midPx": "0.179525",
//             "impactPxs": [
//                 "0.179496",
//                 "0.179561"
//             ],
//             "dayBaseVlm": "4091554.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.083021",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.08275",
//             "markPx": "0.082821",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "727198.0",
//             "prevDayPx": "0.23981",
//             "dayNtlVlm": "181272.85306",
//             "premium": "-0.000710597",
//             "oraclePx": "0.24205",
//             "markPx": "0.24185",
//             "midPx": "0.241755",
//             "impactPxs": [
//                 "0.241615",
//                 "0.241878"
//             ],
//             "dayBaseVlm": "743007.0"
//         },
//         {
//             "funding": "-0.0000047978",
//             "openInterest": "717377100.0",
//             "prevDayPx": "0.000327",
//             "dayNtlVlm": "201190.0070379999",
//             "premium": "-0.0030487805",
//             "oraclePx": "0.000328",
//             "markPx": "0.000327",
//             "midPx": "0.000326",
//             "impactPxs": [
//                 "0.000326",
//                 "0.000327"
//             ],
//             "dayBaseVlm": "610461152.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "1678312.8",
//             "prevDayPx": "0.15795",
//             "dayNtlVlm": "184785.9490889999",
//             "premium": "-0.000243383",
//             "oraclePx": "0.16435",
//             "markPx": "0.1643",
//             "midPx": "0.16421",
//             "impactPxs": [
//                 "0.16415",
//                 "0.16431"
//             ],
//             "dayBaseVlm": "1127949.3000000003"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.018168",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.017705",
//             "markPx": "0.01808",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "1700242.0000000002",
//             "prevDayPx": "0.16541",
//             "dayNtlVlm": "657496.136267",
//             "premium": "0.0",
//             "oraclePx": "0.17249",
//             "markPx": "0.1724",
//             "midPx": "0.17241",
//             "impactPxs": [
//                 "0.17234",
//                 "0.17249"
//             ],
//             "dayBaseVlm": "3797960.4000000008"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "11704496.0",
//             "prevDayPx": "0.047187",
//             "dayNtlVlm": "1183333.3097039997",
//             "premium": "-0.0009463722",
//             "oraclePx": "0.05072",
//             "markPx": "0.050671",
//             "midPx": "0.050655",
//             "impactPxs": [
//                 "0.050624",
//                 "0.050672"
//             ],
//             "dayBaseVlm": "23397899.0"
//         },
//         {
//             "funding": "-0.0000910425",
//             "openInterest": "20866204.0",
//             "prevDayPx": "0.090434",
//             "dayNtlVlm": "2169691.0656609996",
//             "premium": "-0.0008565178",
//             "oraclePx": "0.096904",
//             "markPx": "0.09679",
//             "midPx": "0.096792",
//             "impactPxs": [
//                 "0.096769",
//                 "0.096821"
//             ],
//             "dayBaseVlm": "21755398.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "3769247.7999999998",
//             "prevDayPx": "0.35623",
//             "dayNtlVlm": "553997.3560110002",
//             "premium": "0.0",
//             "oraclePx": "0.35248",
//             "markPx": "0.35245",
//             "midPx": "0.35247",
//             "impactPxs": [
//                 "0.35236",
//                 "0.3526"
//             ],
//             "dayBaseVlm": "1489660.3000000003"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "46971094.0",
//             "prevDayPx": "0.10942",
//             "dayNtlVlm": "1234861.7184499998",
//             "premium": "0.0",
//             "oraclePx": "0.1147",
//             "markPx": "0.1148",
//             "midPx": "0.11482",
//             "impactPxs": [
//                 "0.114622",
//                 "0.115104"
//             ],
//             "dayBaseVlm": "10832999.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "12292753.0",
//             "prevDayPx": "0.11566",
//             "dayNtlVlm": "1280773.2602260001",
//             "premium": "-0.0005847953",
//             "oraclePx": "0.1197",
//             "markPx": "0.11962",
//             "midPx": "0.1196",
//             "impactPxs": [
//                 "0.11957",
//                 "0.11963"
//             ],
//             "dayBaseVlm": "10678403.2999999989"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "17984896.0",
//             "prevDayPx": "0.27732",
//             "dayNtlVlm": "2395136.9188300003",
//             "premium": "-0.000160562",
//             "oraclePx": "0.29895",
//             "markPx": "0.29884",
//             "midPx": "0.298855",
//             "impactPxs": [
//                 "0.298807",
//                 "0.298902"
//             ],
//             "dayBaseVlm": "8354238.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "22956276.0",
//             "prevDayPx": "0.02239",
//             "dayNtlVlm": "347884.9282090001",
//             "premium": "-0.0005141388",
//             "oraclePx": "0.02334",
//             "markPx": "0.023315",
//             "midPx": "0.023308",
//             "impactPxs": [
//                 "0.023291",
//                 "0.023328"
//             ],
//             "dayBaseVlm": "14987495.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "4338476.0",
//             "prevDayPx": "0.20111",
//             "dayNtlVlm": "142662.1673800001",
//             "premium": "-0.0006335283",
//             "oraclePx": "0.2052",
//             "markPx": "0.20504",
//             "midPx": "0.20501",
//             "impactPxs": [
//                 "0.204958",
//                 "0.20507"
//             ],
//             "dayBaseVlm": "695953.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "5168158.0",
//             "prevDayPx": "0.13819",
//             "dayNtlVlm": "236489.6670400001",
//             "premium": "-0.0003539209",
//             "oraclePx": "0.1441",
//             "markPx": "0.14401",
//             "midPx": "0.143965",
//             "impactPxs": [
//                 "0.143893",
//                 "0.144049"
//             ],
//             "dayBaseVlm": "1682489.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "9420932.0",
//             "prevDayPx": "0.17234",
//             "dayNtlVlm": "673450.5753900001",
//             "premium": "-0.0001049724",
//             "oraclePx": "0.181",
//             "markPx": "0.18088",
//             "midPx": "0.1809",
//             "impactPxs": [
//                 "0.180829",
//                 "0.180981"
//             ],
//             "dayBaseVlm": "3761012.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "21106042.3000000231",
//             "prevDayPx": "40.2",
//             "dayNtlVlm": "341153432.0524590611",
//             "premium": "-0.0002834199",
//             "oraclePx": "42.34",
//             "markPx": "42.316",
//             "midPx": "42.3275",
//             "impactPxs": [
//                 "42.3241",
//                 "42.328"
//             ],
//             "dayBaseVlm": "8183110.9199999925"
//         },
//         {
//             "funding": "-0.0000888824",
//             "openInterest": "2352939.1999999997",
//             "prevDayPx": "0.41755",
//             "dayNtlVlm": "237299.096545",
//             "premium": "-0.0013443396",
//             "oraclePx": "0.424",
//             "markPx": "0.42335",
//             "midPx": "0.42329",
//             "impactPxs": [
//                 "0.4232",
//                 "0.42343"
//             ],
//             "dayBaseVlm": "560357.6999999998"
//         },
//         {
//             "funding": "-0.0000062779",
//             "openInterest": "19192596.0",
//             "prevDayPx": "0.061486",
//             "dayNtlVlm": "402121.2838009999",
//             "premium": "-0.0013453992",
//             "oraclePx": "0.062435",
//             "markPx": "0.062355",
//             "midPx": "0.062335",
//             "impactPxs": [
//                 "0.062321",
//                 "0.062351"
//             ],
//             "dayBaseVlm": "6418375.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "16694028.799999997",
//             "prevDayPx": "1.3314",
//             "dayNtlVlm": "21548836.3614099808",
//             "premium": "-0.000145705",
//             "oraclePx": "1.5099",
//             "markPx": "1.5095",
//             "midPx": "1.5095",
//             "impactPxs": [
//                 "1.50932",
//                 "1.50968"
//             ],
//             "dayBaseVlm": "14699245.7000000086"
//         },
//         {
//             "funding": "0.000011952",
//             "openInterest": "358234212.0",
//             "prevDayPx": "0.014375",
//             "dayNtlVlm": "4119232.1323700002",
//             "premium": "-0.0004456044",
//             "oraclePx": "0.015709",
//             "markPx": "0.015699",
//             "midPx": "0.0157",
//             "impactPxs": [
//                 "0.015697",
//                 "0.015702"
//             ],
//             "dayBaseVlm": "269376666.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "18990368.3999999985",
//             "prevDayPx": "0.03315",
//             "dayNtlVlm": "194475.0050140001",
//             "premium": "0.0",
//             "oraclePx": "0.0336",
//             "markPx": "0.0336",
//             "midPx": "0.03358",
//             "impactPxs": [
//                 "0.03357",
//                 "0.0336"
//             ],
//             "dayBaseVlm": "5753200.6999999993"
//         },
//         {
//             "funding": "0.0000151992",
//             "openInterest": "175572154.1999999881",
//             "prevDayPx": "0.30608",
//             "dayNtlVlm": "75728847.004034996",
//             "premium": "0.0002384856",
//             "oraclePx": "0.33545",
//             "markPx": "0.33571",
//             "midPx": "0.33561",
//             "impactPxs": [
//                 "0.33553",
//                 "0.33571"
//             ],
//             "dayBaseVlm": "232945290.400000006"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.06206",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.06153",
//             "markPx": "0.06157",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "18063564.0",
//             "prevDayPx": "0.06311",
//             "dayNtlVlm": "844431.5956389999",
//             "premium": "-0.0003106509",
//             "oraclePx": "0.0676",
//             "markPx": "0.06756",
//             "midPx": "0.067552",
//             "impactPxs": [
//                 "0.067505",
//                 "0.067579"
//             ],
//             "dayBaseVlm": "12711282.0"
//         },
//         {
//             "funding": "0.0000219602",
//             "openInterest": "231513520.0",
//             "prevDayPx": "0.044474",
//             "dayNtlVlm": "580182.8223790003",
//             "premium": "0.0006737247",
//             "oraclePx": "0.04156",
//             "markPx": "0.041627",
//             "midPx": "0.041629",
//             "impactPxs": [
//                 "0.041588",
//                 "0.041648"
//             ],
//             "dayBaseVlm": "13471364.0"
//         },
//         {
//             "funding": "-0.0000273998",
//             "openInterest": "20216936.0",
//             "prevDayPx": "0.078617",
//             "dayNtlVlm": "426776.4096690001",
//             "premium": "-0.0003566816",
//             "oraclePx": "0.081305",
//             "markPx": "0.081245",
//             "midPx": "0.081253",
//             "impactPxs": [
//                 "0.081231",
//                 "0.081276"
//             ],
//             "dayBaseVlm": "5262072.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "41276428.0",
//             "prevDayPx": "0.017577",
//             "dayNtlVlm": "272507.9491530001",
//             "premium": "0.0001097394",
//             "oraclePx": "0.018225",
//             "markPx": "0.018236",
//             "midPx": "0.01823",
//             "impactPxs": [
//                 "0.018227",
//                 "0.018248"
//             ],
//             "dayBaseVlm": "14748103.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "11078438.5999999978",
//             "prevDayPx": "0.67467",
//             "dayNtlVlm": "4631972.8822300015",
//             "premium": "0.0003056617",
//             "oraclePx": "0.71975",
//             "markPx": "0.72016",
//             "midPx": "0.72014",
//             "impactPxs": [
//                 "0.71997",
//                 "0.72035"
//             ],
//             "dayBaseVlm": "6487706.3999999976"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "23398704.0",
//             "prevDayPx": "0.15853",
//             "dayNtlVlm": "1997889.1623599993",
//             "premium": "-0.0009342676",
//             "oraclePx": "0.14985",
//             "markPx": "0.14975",
//             "midPx": "0.149635",
//             "impactPxs": [
//                 "0.149584",
//                 "0.14971"
//             ],
//             "dayBaseVlm": "12680718.0"
//         },
//         {
//             "funding": "-0.0000756072",
//             "openInterest": "1304198.2000000002",
//             "prevDayPx": "1.8439",
//             "dayNtlVlm": "1065945.4888100002",
//             "premium": "-0.0012449799",
//             "oraclePx": "1.992",
//             "markPx": "1.9889",
//             "midPx": "1.98895",
//             "impactPxs": [
//                 "1.98847",
//                 "1.98952"
//             ],
//             "dayBaseVlm": "545597.0999999996"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "5798859.0",
//             "prevDayPx": "7.5103",
//             "dayNtlVlm": "29536808.081280008",
//             "premium": "-0.0001984127",
//             "oraclePx": "8.064",
//             "markPx": "8.0644",
//             "midPx": "8.0619",
//             "impactPxs": [
//                 "8.06093",
//                 "8.0624"
//             ],
//             "dayBaseVlm": "3743839.5999999996"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "26895044.4000000022",
//             "prevDayPx": "0.10986",
//             "dayNtlVlm": "1353323.7720900001",
//             "premium": "-0.0001702273",
//             "oraclePx": "0.11749",
//             "markPx": "0.11744",
//             "midPx": "0.11739",
//             "impactPxs": [
//                 "0.11729",
//                 "0.11747"
//             ],
//             "dayBaseVlm": "11601433.4999999981"
//         },
//         {
//             "funding": "0.0000108418",
//             "openInterest": "69328726.0",
//             "prevDayPx": "0.007534",
//             "dayNtlVlm": "442640.5026059999",
//             "premium": "-0.0001282051",
//             "oraclePx": "0.0078",
//             "markPx": "0.007797",
//             "midPx": "0.007795",
//             "impactPxs": [
//                 "0.007792",
//                 "0.007799"
//             ],
//             "dayBaseVlm": "56266976.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "23783440.0",
//             "prevDayPx": "0.041975",
//             "dayNtlVlm": "559631.1148339999",
//             "premium": "0.0",
//             "oraclePx": "0.043605",
//             "markPx": "0.043597",
//             "midPx": "0.043611",
//             "impactPxs": [
//                 "0.043584",
//                 "0.043636"
//             ],
//             "dayBaseVlm": "12532869.0"
//         },
//         {
//             "funding": "-0.000018175",
//             "openInterest": "717237.3199999999",
//             "prevDayPx": "1.2583",
//             "dayNtlVlm": "273764.5933889999",
//             "premium": "-0.0001480933",
//             "oraclePx": "1.3505",
//             "markPx": "1.35",
//             "midPx": "1.3496",
//             "impactPxs": [
//                 "1.3489",
//                 "1.3503"
//             ],
//             "dayBaseVlm": "210424.2900000001"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.037555",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.03719",
//             "markPx": "0.037555",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "-0.0002681931",
//             "openInterest": "3925824.1999999997",
//             "prevDayPx": "1.5972",
//             "dayNtlVlm": "1178170.0254399995",
//             "premium": "-0.0026933779",
//             "oraclePx": "1.797",
//             "markPx": "1.7914",
//             "midPx": "1.7911",
//             "impactPxs": [
//                 "1.79073",
//                 "1.79216"
//             ],
//             "dayBaseVlm": "696279.9"
//         },
//         {
//             "funding": "0.0000099372",
//             "openInterest": "13355704.0",
//             "prevDayPx": "0.01839",
//             "dayNtlVlm": "472465.9762930001",
//             "premium": "-0.0011326861",
//             "oraclePx": "0.01854",
//             "markPx": "0.018513",
//             "midPx": "0.018505",
//             "impactPxs": [
//                 "0.018497",
//                 "0.018519"
//             ],
//             "dayBaseVlm": "25083202.0"
//         },
//         {
//             "funding": "-0.000149626",
//             "openInterest": "7460158.0",
//             "prevDayPx": "0.24938",
//             "dayNtlVlm": "1779670.282240001",
//             "premium": "-0.0026347305",
//             "oraclePx": "0.2505",
//             "markPx": "0.24967",
//             "midPx": "0.249635",
//             "impactPxs": [
//                 "0.249476",
//                 "0.24984"
//             ],
//             "dayBaseVlm": "7099353.0"
//         },
//         {
//             "funding": "-0.0000540989",
//             "openInterest": "1011628.2000000001",
//             "prevDayPx": "3.722",
//             "dayNtlVlm": "1951933.3651199993",
//             "premium": "-0.0010067559",
//             "oraclePx": "3.7745",
//             "markPx": "3.7693",
//             "midPx": "3.7699",
//             "impactPxs": [
//                 "3.76861",
//                 "3.7707"
//             ],
//             "dayBaseVlm": "527639.6999999998"
//         },
//         {
//             "funding": "-0.0001640399",
//             "openInterest": "21979400.8000000007",
//             "prevDayPx": "0.095",
//             "dayNtlVlm": "517056.0354160001",
//             "premium": "-0.0017128463",
//             "oraclePx": "0.09925",
//             "markPx": "0.09907",
//             "midPx": "0.09903",
//             "impactPxs": [
//                 "0.09901",
//                 "0.09908"
//             ],
//             "dayBaseVlm": "5264970.9000000004"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "2165118.0",
//             "prevDayPx": "0.86901",
//             "dayNtlVlm": "444999.8137999998",
//             "premium": "-0.0002716905",
//             "oraclePx": "0.8944",
//             "markPx": "0.894",
//             "midPx": "0.89385",
//             "impactPxs": [
//                 "0.893609",
//                 "0.894157"
//             ],
//             "dayBaseVlm": "501652.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "2937320.0",
//             "prevDayPx": "0.19038",
//             "dayNtlVlm": "178914.66836",
//             "premium": "-0.0001646938",
//             "oraclePx": "0.1943",
//             "markPx": "0.1943",
//             "midPx": "0.19417",
//             "impactPxs": [
//                 "0.194065",
//                 "0.194268"
//             ],
//             "dayBaseVlm": "919607.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "6371.916",
//             "prevDayPx": "3989.7",
//             "dayNtlVlm": "4310264.6748000011",
//             "premium": "0.0",
//             "oraclePx": "4077.2",
//             "markPx": "4077.2",
//             "midPx": "4077.25",
//             "impactPxs": [
//                 "4077.2",
//                 "4077.319"
//             ],
//             "dayBaseVlm": "1070.695"
//         },
//         {
//             "funding": "-0.0000215313",
//             "openInterest": "6716118.0",
//             "prevDayPx": "0.091278",
//             "dayNtlVlm": "551502.8770710002",
//             "premium": "-0.0009947338",
//             "oraclePx": "0.08545",
//             "markPx": "0.085358",
//             "midPx": "0.08535",
//             "impactPxs": [
//                 "0.085328",
//                 "0.085365"
//             ],
//             "dayBaseVlm": "6225675.0"
//         },
//         {
//             "funding": "-0.0000635999",
//             "openInterest": "15043244.0",
//             "prevDayPx": "0.026037",
//             "dayNtlVlm": "128985.2265270001",
//             "premium": "-0.0014609204",
//             "oraclePx": "0.02738",
//             "markPx": "0.027329",
//             "midPx": "0.027327",
//             "impactPxs": [
//                 "0.027312",
//                 "0.02734"
//             ],
//             "dayBaseVlm": "4806874.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "3932094.0",
//             "prevDayPx": "0.12335",
//             "dayNtlVlm": "161817.07953",
//             "premium": "-0.0009940358",
//             "oraclePx": "0.12575",
//             "markPx": "0.12547",
//             "midPx": "0.12551",
//             "impactPxs": [
//                 "0.125401",
//                 "0.125625"
//             ],
//             "dayBaseVlm": "1296363.0"
//         },
//         {
//             "funding": "-0.0000078339",
//             "openInterest": "3325150.0",
//             "prevDayPx": "0.16991",
//             "dayNtlVlm": "269274.85288",
//             "premium": "-0.0006144928",
//             "oraclePx": "0.1725",
//             "markPx": "0.1723",
//             "midPx": "0.172295",
//             "impactPxs": [
//                 "0.17222",
//                 "0.172394"
//             ],
//             "dayBaseVlm": "1535927.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "116332152.0",
//             "prevDayPx": "0.059397",
//             "dayNtlVlm": "3365310.6626189994",
//             "premium": "-0.0002726544",
//             "oraclePx": "0.06235",
//             "markPx": "0.06231",
//             "midPx": "0.062303",
//             "impactPxs": [
//                 "0.062284",
//                 "0.062333"
//             ],
//             "dayBaseVlm": "53447314.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "9081046.0",
//             "prevDayPx": "0.13146",
//             "dayNtlVlm": "228686.3801799999",
//             "premium": "0.0",
//             "oraclePx": "0.13055",
//             "markPx": "0.13041",
//             "midPx": "0.130465",
//             "impactPxs": [
//                 "0.130379",
//                 "0.13055"
//             ],
//             "dayBaseVlm": "1713016.0"
//         },
//         {
//             "funding": "-0.0000466739",
//             "openInterest": "154803878.0",
//             "prevDayPx": "0.006069",
//             "dayNtlVlm": "458225.966846",
//             "premium": "-0.0015867978",
//             "oraclePx": "0.006302",
//             "markPx": "0.006288",
//             "midPx": "0.006288",
//             "impactPxs": [
//                 "0.006285",
//                 "0.006292"
//             ],
//             "dayBaseVlm": "72399883.0"
//         },
//         {
//             "funding": "0.0",
//             "openInterest": "0.0",
//             "prevDayPx": "0.060054",
//             "dayNtlVlm": "0.0",
//             "premium": null,
//             "oraclePx": "0.06012",
//             "markPx": "0.060367",
//             "midPx": null,
//             "impactPxs": null,
//             "dayBaseVlm": "0.0"
//         },
//         {
//             "funding": "-0.0000593246",
//             "openInterest": "2114432.0",
//             "prevDayPx": "0.38604",
//             "dayNtlVlm": "250074.8229100001",
//             "premium": "-0.0015291098",
//             "oraclePx": "0.37015",
//             "markPx": "0.36952",
//             "midPx": "0.36947",
//             "impactPxs": [
//                 "0.36937",
//                 "0.369584"
//             ],
//             "dayBaseVlm": "657076.0"
//         },
//         {
//             "funding": "-0.0000471903",
//             "openInterest": "30187494.0",
//             "prevDayPx": "0.02261",
//             "dayNtlVlm": "381865.7763449999",
//             "premium": "-0.0012893244",
//             "oraclePx": "0.023268",
//             "markPx": "0.023232",
//             "midPx": "0.023233",
//             "impactPxs": [
//                 "0.023228",
//                 "0.023238"
//             ],
//             "dayBaseVlm": "16484102.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "26918794.0",
//             "prevDayPx": "0.113",
//             "dayNtlVlm": "4908270.6357300002",
//             "premium": "-0.0022978269",
//             "oraclePx": "0.11228",
//             "markPx": "0.1121",
//             "midPx": "0.111975",
//             "impactPxs": [
//                 "0.111943",
//                 "0.112022"
//             ],
//             "dayBaseVlm": "43148420.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "6345656.0",
//             "prevDayPx": "0.44516",
//             "dayNtlVlm": "1555580.5840700001",
//             "premium": "0.0003468812",
//             "oraclePx": "0.47855",
//             "markPx": "0.47878",
//             "midPx": "0.478845",
//             "impactPxs": [
//                 "0.478716",
//                 "0.479048"
//             ],
//             "dayBaseVlm": "3278807.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "28840007468.0",
//             "prevDayPx": "0.003794",
//             "dayNtlVlm": "92092034.2344749421",
//             "premium": "0.0002292001",
//             "oraclePx": "0.004363",
//             "markPx": "0.004364",
//             "midPx": "0.004366",
//             "impactPxs": [
//                 "0.004364",
//                 "0.004367"
//             ],
//             "dayBaseVlm": "22266428223.0"
//         },
//         {
//             "funding": "-0.0001086969",
//             "openInterest": "1952486.0",
//             "prevDayPx": "0.56934",
//             "dayNtlVlm": "302054.92499",
//             "premium": "-0.0014347421",
//             "oraclePx": "0.59035",
//             "markPx": "0.58916",
//             "midPx": "0.589195",
//             "impactPxs": [
//                 "0.588961",
//                 "0.589503"
//             ],
//             "dayBaseVlm": "512585.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "1881106.0",
//             "prevDayPx": "0.38292",
//             "dayNtlVlm": "2293403.7465999993",
//             "premium": "0.0",
//             "oraclePx": "0.3865",
//             "markPx": "0.38645",
//             "midPx": "0.38634",
//             "impactPxs": [
//                 "0.386127",
//                 "0.386616"
//             ],
//             "dayBaseVlm": "5985191.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "302700860.0",
//             "prevDayPx": "0.28111",
//             "dayNtlVlm": "102040782.5658800155",
//             "premium": "-0.0008720113",
//             "oraclePx": "0.31995",
//             "markPx": "0.31979",
//             "midPx": "0.319665",
//             "impactPxs": [
//                 "0.319516",
//                 "0.319671"
//             ],
//             "dayBaseVlm": "332889639.0"
//         },
//         {
//             "funding": "-0.0000521466",
//             "openInterest": "121729558.0",
//             "prevDayPx": "0.12103",
//             "dayNtlVlm": "27522201.5827900097",
//             "premium": "0.0",
//             "oraclePx": "0.1507",
//             "markPx": "0.15054",
//             "midPx": "0.15063",
//             "impactPxs": [
//                 "0.150583",
//                 "0.150714"
//             ],
//             "dayBaseVlm": "183593520.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "215903660.0",
//             "prevDayPx": "0.012046",
//             "dayNtlVlm": "1120985.643075",
//             "premium": "0.0",
//             "oraclePx": "0.01321",
//             "markPx": "0.013204",
//             "midPx": "0.013211",
//             "impactPxs": [
//                 "0.013202",
//                 "0.013218"
//             ],
//             "dayBaseVlm": "88765505.0"
//         },
//         {
//             "funding": "0.0000042688",
//             "openInterest": "21075484.0",
//             "prevDayPx": "0.052271",
//             "dayNtlVlm": "336438.896024",
//             "premium": "-0.0001618414",
//             "oraclePx": "0.05561",
//             "markPx": "0.055534",
//             "midPx": "0.05557",
//             "impactPxs": [
//                 "0.05554",
//                 "0.055601"
//             ],
//             "dayBaseVlm": "6183157.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "127048070.0",
//             "prevDayPx": "1.0507",
//             "dayNtlVlm": "47159630.6610999778",
//             "premium": "-0.0003590664",
//             "oraclePx": "1.114",
//             "markPx": "1.1139",
//             "midPx": "1.11355",
//             "impactPxs": [
//                 "1.113488",
//                 "1.1136"
//             ],
//             "dayBaseVlm": "42273535.0"
//         },
//         {
//             "funding": "0.0000109406",
//             "openInterest": "11255512.0",
//             "prevDayPx": "0.51815",
//             "dayNtlVlm": "2353538.0379400011",
//             "premium": "-0.0002708016",
//             "oraclePx": "0.53914",
//             "markPx": "0.53861",
//             "midPx": "0.53879",
//             "impactPxs": [
//                 "0.538652",
//                 "0.538994"
//             ],
//             "dayBaseVlm": "4352620.0"
//         },
//         {
//             "funding": "0.0000167474",
//             "openInterest": "43654976.0",
//             "prevDayPx": "0.073632",
//             "dayNtlVlm": "883685.661783",
//             "premium": "0.0",
//             "oraclePx": "0.076076",
//             "markPx": "0.076149",
//             "midPx": "0.076059",
//             "impactPxs": [
//                 "0.076015",
//                 "0.076103"
//             ],
//             "dayBaseVlm": "11558676.0"
//         },
//         {
//             "funding": "-0.0021432228",
//             "openInterest": "2504328.0",
//             "prevDayPx": "1.5494",
//             "dayNtlVlm": "5838383.649299996",
//             "premium": "-0.0172651781",
//             "oraclePx": "1.6257",
//             "markPx": "1.5975",
//             "midPx": "1.5968",
//             "impactPxs": [
//                 "1.596076",
//                 "1.597632"
//             ],
//             "dayBaseVlm": "3576297.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "9478832.0",
//             "prevDayPx": "0.033811",
//             "dayNtlVlm": "1847427.8528829999",
//             "premium": "-0.0025139665",
//             "oraclePx": "0.0358",
//             "markPx": "0.035793",
//             "midPx": "0.035703",
//             "impactPxs": [
//                 "0.035688",
//                 "0.03571"
//             ],
//             "dayBaseVlm": "49088995.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "6054704.0",
//             "prevDayPx": "0.83716",
//             "dayNtlVlm": "2050912.4336300003",
//             "premium": "0.0014881923",
//             "oraclePx": "0.82785",
//             "markPx": "0.82888",
//             "midPx": "0.82956",
//             "impactPxs": [
//                 "0.829082",
//                 "0.830077"
//             ],
//             "dayBaseVlm": "2428259.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "24846278.0",
//             "prevDayPx": "0.18753",
//             "dayNtlVlm": "1467684.4274100005",
//             "premium": "-0.0008316951",
//             "oraclePx": "0.19839",
//             "markPx": "0.19821",
//             "midPx": "0.19818",
//             "impactPxs": [
//                 "0.198139",
//                 "0.198225"
//             ],
//             "dayBaseVlm": "7394019.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "561649.3400000001",
//             "prevDayPx": "581.52",
//             "dayNtlVlm": "433130208.3447001576",
//             "premium": "-0.0006126802",
//             "oraclePx": "661.03",
//             "markPx": "660.32",
//             "midPx": "660.44",
//             "impactPxs": [
//                 "660.3094",
//                 "660.625"
//             ],
//             "dayBaseVlm": "684185.5599999984"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "194436054.0",
//             "prevDayPx": "0.059032",
//             "dayNtlVlm": "407068.1080330001",
//             "premium": "0.0001335415",
//             "oraclePx": "0.061479",
//             "markPx": "0.062627",
//             "midPx": "0.06273",
//             "impactPxs": [
//                 "0.0623",
//                 "0.062946"
//             ],
//             "dayBaseVlm": "6713103.0"
//         },
//         {
//             "funding": "0.0000263443",
//             "openInterest": "27541730.0",
//             "prevDayPx": "0.35653",
//             "dayNtlVlm": "1883200.2078299997",
//             "premium": "0.0009645944",
//             "oraclePx": "0.3587",
//             "markPx": "0.3591",
//             "midPx": "0.359225",
//             "impactPxs": [
//                 "0.359046",
//                 "0.359448"
//             ],
//             "dayBaseVlm": "5256493.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "14934748.0",
//             "prevDayPx": "0.4852",
//             "dayNtlVlm": "706592.1970599999",
//             "premium": "-0.0001187786",
//             "oraclePx": "0.48241",
//             "markPx": "0.46767",
//             "midPx": "0.468315",
//             "impactPxs": [
//                 "0.461507",
//                 "0.47668"
//             ],
//             "dayBaseVlm": "1456595.0"
//         },
//         {
//             "funding": "0.0000742879",
//             "openInterest": "38034730.0",
//             "prevDayPx": "0.15474",
//             "dayNtlVlm": "6379343.1192100039",
//             "premium": "0.0027279496",
//             "oraclePx": "0.14773",
//             "markPx": "0.148",
//             "midPx": "0.148215",
//             "impactPxs": [
//                 "0.148133",
//                 "0.148562"
//             ],
//             "dayBaseVlm": "42344185.0"
//         },
//         {
//             "funding": "0.0000125",
//             "openInterest": "609245.8",
//             "prevDayPx": "7.9466",
//             "dayNtlVlm": "22677992.6545000039",
//             "premium": "0.0",
//             "oraclePx": "7.226",
//             "markPx": "7.2178",
//             "midPx": "7.22695",
//             "impactPxs": [
//                 "7.22511",
//                 "7.22935"
//             ],
//             "dayBaseVlm": "3059820.0999999982"
//         }
//     ]
// ]

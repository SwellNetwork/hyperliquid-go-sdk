package hyperliquid

func transformSpotMetaAndAssetCtxsResult(result spotMetaAndAssetCtxsResult) (*SpotMetaAndAssetCtx, error) {
	meta := SpotMeta{
		Universe: make([]SpotAssetInfo, len(result.Meta.Universe)),
		Tokens:   result.Meta.Tokens,
	}

	for i, uni := range result.Meta.Universe {
		meta.Universe[i] = SpotAssetInfo{
			Tokens:      append([]int(nil), uni.Tokens...),
			Name:        uni.Name,
			Index:       uni.Index,
			IsCanonical: uni.IsCanonical,
		}
	}

	ctxs := make([]SpotAssetCtx, len(result.Assets))
	for i, asset := range result.Assets {
		ctxs[i] = SpotAssetCtx(asset)
	}

	return &SpotMetaAndAssetCtx{
		Meta: meta,
		Ctxs: ctxs,
	}, nil
}

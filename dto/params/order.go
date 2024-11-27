package params

type OrderParam struct {
	Offset int64 `form:"offset" json:"offset"`
	Limit  int64 `form:"limit" json:"limit"`
	// 1-card 2-fragment
	NftType        int32  `form:"nftType" json:"nftType"`
	MaxCardTokenId int64  `form:"maxCardTokenId" json:"maxCardTokenId"`
	NftAddr        string `form:"nftAddr" json:"nftAddr" binding:"required"`
	Seller         string `form:"seller" json:"seller"`
}

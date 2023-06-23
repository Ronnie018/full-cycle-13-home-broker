package entity

type Investor struct {
	ID       			string  `json:"id"`
	Name     			string `json:"name"`
	AssetPositon 	[]*InvestorAssetPosition
}

type InvestorAssetPosition struct {
	AssetID string
	Shares 	int
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID: 					id,
		AssetPositon: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPositon = append(i.AssetPositon, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, amountShares int) { // qtdShares
	assetPositon := i.GetAssetPosition(assetID)
	if assetPositon == nil {
		i.AssetPositon = append(i.AssetPositon, NewInvestorAssetPosition(assetID, amountShares))
	}else {
		assetPositon.Shares += amountShares
	}
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return	&InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, current := range i.AssetPositon {
		if current.AssetID == assetID {
			return current
		}
	}
	return nil
}

package entity

import (
	"time"
	// "container/heap"
	"github.com/google/uuid" // go mod tidy - get the missing dependencies
)

type Transaction struct {
	ID 						string
	SellingOrder	*Order
	BuyingOrder  	*Order
	Shares 				int
	Price 				float64
	Total 				float64
	datetime 			time.Time
}

func NewTransaction (sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price
	// generate go uuid

	return &Transaction{
		ID: 						uuid.New().String(),
		SellingOrder: 	sellingOrder,
		BuyingOrder: 		buyingOrder,
		Shares: 				shares,
		Price: 					price,
		Total: 					total,
		datetime:				time.Now(),
	}
}

func (t *Transaction) minShares () int {
	if (t.BuyingOrder.PendingShares < t.SellingOrder.PendingShares) {
		return t.BuyingOrder.PendingShares
	}
	return t.SellingOrder.PendingShares
}

func (t *Transaction) updateBuyerAssets(shareAmount int) {
	order := t.BuyingOrder

	order.Investor.UpdateAssetPosition(order.Asset.ID, shareAmount)

	order.PendingShares -= shareAmount
}

func (t *Transaction) updateSellerAssets(shareAmount int) {
	order := t.SellingOrder

	order.Investor.UpdateAssetPosition(order.Asset.ID, -shareAmount)

	order.PendingShares -= shareAmount
}


func (t *Transaction) updateTransactionTotal() {
	t.Total = float64(t.Shares) * t.BuyingOrder.Price
}

func (t *Transaction) updateOrderStatus() {
	if t.SellingOrder.PendingShares == 0 {
		t.SellingOrder.Status = "CLOSED"
	}
	if t.BuyingOrder.PendingShares == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
}
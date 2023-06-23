package entity

import (
	"container/heap"
	"sync"
)

type Book struct {
	Order 				[]*Order
	Transactions 	[]*Transaction
	OrdersChan 		chan *Order
	OrdersChanOut chan *Order
	Wg 						*sync.WaitGroup
}

func NewBook(orderChan chan *Order, orderChanOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order: 					[]*Order{},
		Transactions: 	[]*Transaction{},
		OrdersChan: 		orderChan,
		OrdersChanOut: 	orderChanOut,
		Wg: 						wg,
	}
}

func (b *Book) Trade() {
	// buyOrders := NewOrderQueue()
	// sellOrders := NewOrderQueue()

	buyOrders := make(map[string]*OrderQueue)
	sellOrders := make(map[string]*OrderQueue)

	// heap.Init(buyOrders)
	// heap.Init(sellOrders)

	for order := range b.OrdersChan {

		asset := order.Asset.ID

		if buyOrders[asset] == nil {
			buyOrders[asset] = NewOrderQueue()
			heap.Init(buyOrders[asset])
		}

		if sellOrders[asset] == nil {
			sellOrders[asset] = NewOrderQueue()
			heap.Init(sellOrders[asset])
		}

		switch order.OrderType {

			case "BUY":
				buyOrders[asset].Push(order)

				if sellOrders[asset].Len() > 0 && sellOrders[asset].Orders[0].Price <= order.Price {
					sellOrder := sellOrders[asset].Pop().(*Order)
					if sellOrder.PendingShares > 0 {
						transaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)

						b.addTransaction(transaction, b.Wg)

						sellOrder.Transactions = append(sellOrder.Transactions, transaction)
						order.Transactions = append(order.Transactions, transaction)

						b.OrdersChanOut <- sellOrder
						b.OrdersChanOut <- order

						if sellOrder.PendingShares > 0 {
							 sellOrders[asset].Push(sellOrder)
						}
					}
				}


			case "SELL":
				sellOrders[asset].Push(order)

				if buyOrders[asset].Len() > 0 && buyOrders[asset].Orders[0].Price >= order.Price {
					buyOrder := buyOrders[asset].Pop().(*Order)
					if buyOrder.PendingShares > 0 {
						transaction := NewTransaction(order, buyOrder, order.Shares, buyOrder.Price)
						b.addTransaction(transaction, b.Wg)

						buyOrder.Transactions = append(buyOrder.Transactions, transaction)
						order.Transactions = append(order.Transactions, transaction)

						b.OrdersChanOut <- buyOrder
						b.OrdersChanOut <- order

						if buyOrder.PendingShares > 0 {
							buyOrders[asset].Push(buyOrder)
						}
					}
				}

		}
	}
}

func (b *Book) addTransaction(transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	minShares := transaction.minShares()
	
	transaction.updateSellerAssets(minShares)

	transaction.updateBuyerAssets(minShares)

	transaction.updateTransactionTotal()
	
	transaction.updateOrderStatus()

	b.Transactions = append(b.Transactions, transaction)
}

/*
	func (b *Book) addTransaction(transaction *Transaction, wg *sync.WaitGroup) {
		defer wg.Done()

		minShares := transaction.minShares()
		
		// transaction.SellingOrder.Investor.UpdateAssetPosition(transaction.SellingOrder.Asset.ID, -minShares)
		// transaction.SellingOrder.PendingShares -= minShares
		transaction.updateSellerAssets(minShares)
		
		// transaction.BuyingOrder.Investor.UpdateAssetPosition(transaction.BuyingOrder.Asset.ID, minShares)
		// transaction.BuyingOrder.PendingShares -= minShares
		transaction.updateBuyerAssets(minShares)

		// transaction.Total = float64(transaction.Shares)  * transaction.BuyingOrder.Price
		transaction.updateTransactionTotal()

		// if transaction.BuyingOrder.PendingShares == 0 {
		// 	transaction.BuyingOrder.Status = "CLOSED"
		// }
		// if transaction.SellingOrder.PendingShares == 0 {
		// 	transaction.SellingOrder.Status = "CLOSED"
		// }
		
		transaction.updateOrderStatus()

		b.Transactions = append(b.Transactions, transaction)
	}
*/
package entity

type	OrderQueue struct {
	Orders []*Order
}

func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

func (oq OrderQueue) Len() int { 
	// we dont need the * here cause its just reading the length and nothing need to change on the original array
	// "*" means that the value passed to the function will be applyied to the memory, ie, if i change a value in the object, the original value will change, otherwise, if i dont put * in there im only changing the static value 
	//
	return len(oq.Orders)
}

func (oq *OrderQueue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

func (oq *OrderQueue) Pop() any {
	old := oq.Orders
	last := len(old) - 1
	item := old[last]
	oq.Orders = old[0 : last]

	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
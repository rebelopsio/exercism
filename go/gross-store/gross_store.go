package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}
	if _, ok := bill[item]; ok {
		bill[item] += units[unit]
		return true
	} else {
		bill[item] = units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := bill[item]; !ok {
		return false
	}
	if _, ok := units[unit]; !ok {
		return false
	}
	if units[unit] > bill[item] {
		return false
	}
	bill[item] -= units[unit]
	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	return qty, ok
}

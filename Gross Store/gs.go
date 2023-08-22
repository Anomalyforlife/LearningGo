package gross

func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

func NewBill() map[string]int {
	return map[string]int{}
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	toAdd, exists := units[unit]
	if exists {
		bill[item] += toAdd
		return true
	} else {
		return false
	}
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitPieces, unitExists := units[unit]
	billPieces, billExists := bill[item]
	if !unitExists || !billExists || unitPieces > billPieces { // if the unit doesn't exist or the item doesn't exist or the unit is greater than the item
		return false
	} else if unitPieces == billPieces {
		delete(bill, item)
		return true
	} else {
		bill[item] -= unitPieces
		return true
	}
}

func GetItem(bill map[string]int, item string) (int, bool) {
	pieces, exists := bill[item]
	return pieces, exists
}

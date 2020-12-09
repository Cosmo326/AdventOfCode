package _7

type BagContent struct {
	Name string
	Count int
}

type Bag struct {
	Name string
	Content [] BagContent
}

func Where(Bags []Bag, clause func(Bag) bool) []Bag {
	retVal := make([]Bag, 0)
	for _, bag := range Bags {
		if clause(bag) {
			retVal = append(retVal, bag)
		}
	}
	return retVal
}

func Contains(bagContents []BagContent, clause func(bagContent BagContent) bool) bool {
	for _, bagContent := range bagContents {
		if clause(bagContent) {
			return true
		}
	}
	return false
}
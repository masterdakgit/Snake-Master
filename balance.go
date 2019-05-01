package SnakeMasters

func (w *World) setBalance() {
	curentBalance := w.calcBalance()
	if curentBalance < w.balance {
		w.addEatN(w.balance - curentBalance)
	}
	if curentBalance > w.balance {
		w.delEatN(curentBalance - w.balance)
	}
}

func (w *World) calcBalance() int {
	result := 0
	for x := range w.area {
		for y := range w.area[x] {
			if w.area[x][y] == elEat {
				result++
			}
		}
	}
	for n := range w.snakes {
		result += len(w.snakes[n].body)
	}
	return result
}

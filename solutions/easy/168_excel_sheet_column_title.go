package easy

func ConvertToTitle(columnNumber int) string {
	var result string

	//пока(while) columnNumber не станет равен 0
	for columnNumber > 0 {
		columnNumber-- // Уменьшаем на 1, так как индексация букв начинается с 1
		offset := columnNumber % 26
		result = string('A'+offset) + result // Добавляем букву в начало строки
		columnNumber /= 26
	}

	return result
}

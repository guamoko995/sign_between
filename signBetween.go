package main

import "fmt"

// Возводит целое число в целую положительную степень
func Pow(namber, power int) int {
	p := 1
	for i := 0; i < power; i++ {
		p *= namber
	}
	return p
}

// Выражение
type expression struct {
	// Парамтрами выражения являются числа (цифры) стоящие
	// слева от знака равно, между которыми можно поставить
	// знак +, -, или ничего.
	// *По условию [9 8 7 6 5 4 3 2 1 0]
	parameters []int

	// Стратегии расстоновки знаков, выраженные числом.
	// Если это число представить в троичной системе
	// счисления, младший разряд будет отчечать за первый
	// знак, следующий за второй и т.д. причем:
	// значение разряда равное 0 - это пропуск знака ()
	// значение разряда равное 1 - это знак (+)
	// значение разряда равное 2 - это знак (-)
	// Пример для трех чисел (цифр) [1 2 3]:
	// strategy = 0: 123
	// strategy = 1: 1+23
	// strategy = 2: 1-23
	// strategy = 3: 12+3
	// strategy = 4: 1+2+3
	// strategy = 5: 1-2+3
	// strategy = 6: 12-3
	// strategy = 7: 1+2-3
	// strategy = 8: 1-2-3
	strateges []int

	// Значение справа от знака равно.
	// *По условию 200
	value int

	// Разрешимость выражения.
	// solvable=true - выражение разрешимо.
	solvable bool
}

//Отображает выражение.
func (ex *expression) String() string {
	// Индекс последнего элемента.
	l := len(ex.parameters) - 1
	var str string

	// Если выражение не разрешимо или не решено, вместо возможных
	// знвков отобразится символ "_", а после выражения
	// метка " - не разрешимо"
	if !ex.solvable {
		for _, p := range ex.parameters[:l] {
			str += fmt.Sprint(p) + "_"
		}
		return str + fmt.Sprint(ex.parameters[l]) + "=" + fmt.Sprint(ex.value) + " - не разрешимо"
	}

	// Если выражение разрешимо и решено, отображаются решения
	for _, strategy := range ex.strateges {
		for i, p := range ex.parameters[:l] {
			// К строке добавляется параметр
			str += fmt.Sprint(p)

			// Получает знак (или ничего) стоящий после i-того
			// параметра, согласно стратегии расстановки
			o := ReadeOperator(strategy, i)
			// К строке добавляется соответствующий знак знак (или ничего)
			switch o {
			case 1:
				str += "+"
			case 2:
				str += "-"
			}
		}
		// К строке добавляется последний параметр, знак равно и значение выражения.
		str += fmt.Sprint(ex.parameters[l]) + "=" + fmt.Sprint(ex.value) + "\n"
	}
	return str
}

// Считывает операнд соединяя числа, начиная с числа с индексом cursor
// до ближайшего знака или конца массива в соответствии с переданной strategy.
func ReadeOperand(strategy int, nambers []int, coursore int) (int, int) {
	operand := nambers[coursore]
	l := len(nambers)
	var i int
	for i = coursore + 1; i < l; i++ {
		// Получаен знак (или ничего) 0(), 1(+), или 2(-) согласно strategy
		operator := ReadeOperator(strategy, i-1)
		if operator == 0 { //()
			operand = operand*10 + nambers[i]
		} else {
			return operand, i
		}
	}
	return operand, i
}

// Считывает "знак" после i-того числа (цифры) в соответствии с
// переданной strategy. 0 соответствует (), 1 с-ет (+), 2 с-ет (-)
func ReadeOperator(strategy, i int) byte {
	return byte((strategy / Pow(3, i)) % 3)
}

// Разрешает выражение, перебирая стратегии расстановки знаков
func (ex *expression) Solve() {
	l := len(ex.parameters)

	// Переберает все стратегии по-порядку
	for strategy := 0; strategy < Pow(3, l-1); strategy++ {
		// Получает первый операнд и переставляет курсор на начало
		// следующего операнда.
		firstOperand, coursore := ReadeOperand(strategy, ex.parameters, 0)

		// Сворачивает выражение, пока числа слева от равно не закончатся
		for coursore < l {
			var operator byte
			var secondOperand int

			// Получает оператор после первого операнда (гарантированно + или -)
			operator = ReadeOperator(strategy, coursore-1)

			//Получает второй операнд и переставляет курсов на начало
			// следующего операнда.
			secondOperand, coursore = ReadeOperand(strategy, ex.parameters, coursore)

			// Вычисляет первое действие (операцию над первыми двумя операндами)
			// и записывает результат в первый операнд для следующего действия
			switch operator {
			case 1: // +
				firstOperand += secondOperand
			case 2: // -
				firstOperand -= secondOperand
			default:
				panic("Недопустимое значение оператора: " + fmt.Sprint(operator))
			}
		}

		// Если выражение истино, то в выражении сохраняется стратегия
		// расстановки операторов и выставляется флаг разрешимости выражения
		if firstOperand == ex.value {
			ex.strateges = append(ex.strateges, strategy)
			ex.solvable = true
		}
	}
}

func main() {
	// Создаем выражение, соответствующее условию
	s := &expression{
		parameters: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		strateges:  make([]int, 0),
		value:      200,
	}

	// Подбираем стратегию расстановки знаков
	s.Solve()

	// Выводим результат
	fmt.Print(s)
}

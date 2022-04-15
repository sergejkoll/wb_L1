package main

var justString string

//переслайсинг (re-slicing) среза не создаёт копию массива в основании. Массив полностью будет существовать в памяти,
//пока на него не перестанут ссылаться. Иногда это вызывает хранение всех данных в памяти,
//когда нужна только их небольшая часть. (https://go.dev/blog/slices-intro)
// Так как слайс ссылается на исходный массив, пока слайс есть в памяти, сборщик мусора не сможет очистить массив
func someFunc() {
	v := createHugeString(1 << 10)
	tmp := make([]byte, 100)
	copy(tmp, v)
	justString = string(tmp)
}

func someFunc2() {
	v := createHugeString(1 << 10)
	var tmp []byte
	justString = string(append(tmp, v[:100]...))
}

func main() {
	someFunc()
}

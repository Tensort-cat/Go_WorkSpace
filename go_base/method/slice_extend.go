package main

type IntSlice []int

func (s IntSlice) Get(i int) int {
	return s[i]
}

func (s IntSlice) Set(i int, v int) {
	s[i] = v
}

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) AppendByValue(v int) { // 并不影响原切片
	s = append(s, v)
}

func (s *IntSlice) AppendByPointer(v int) { // 使用指针接收者，会影响原切片
	*s = append(*s, v)
}

func main() {
	s := IntSlice{1, 2, 3}
	s.Set(1, 4)
	s.Set(2, 5)
	s.AppendByValue(6)
	for i := 0; i < s.Len(); i++ {
		print(s.Get(i), " ")
	}
	println()
	s.AppendByPointer(6)
	for i := 0; i < s.Len(); i++ {
		print(s.Get(i), " ")
	}
}

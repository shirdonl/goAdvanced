package main

func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			println(i)
		}()
	}
}

//3
//3
//3

//func main() {
//	for i := 0; i < 3; i++ {
//		defer func(i int) {
//			println(i)
//		}(i)
//	}
//}
//2
//1
//0

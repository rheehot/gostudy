package main

func main() {
	var k int = 10
	var p = &k  //k의 주소를 할당
	println(*p) //p가 가리키는 주소에 있는 실제 내용을 출력

	if val := i * 2; val < max {
		println(val)
	}
	// 아래 처럼 사용하면 Scope 벗어나 에러
	val++
}

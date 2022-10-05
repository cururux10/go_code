package main // 컴파일용
import (
	"fmt"

	functwo "main.go/FuncTwo"
	"main.go/something"
	"rsc.io/quote"
)

func main() {
	name := "Go Developers"
	// len, uppername := funcOne.
	// fmt.Println("len for", len, "upper ", uppername)
	fmt.Println(something.Publicfunc())
	len, uppername := something.LenAndUpper(name)
	totalsum := functwo.Sumintall(4, 10)
	Avrfloat32 := functwo.Avrfloat32(4, 10)
	deferints, lenints := functwo.Intdefer(4, 10)
	fmt.Println("- len for", len, "\n- upper ", uppername, "\n- totalsum", totalsum, "\n- Avrfloat32", Avrfloat32, "\n- deferints", deferints, "\n- lenints", lenints)

	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}

package main
import(
	"golang.org/x/tour/wc"
	//"fmt"
	"strings"
)

func WordCount(s string) map[string]int{
	words := make(map[string]int)
	for _, word := range strings.Fields(s) {
		words[word] += 1
	}
	return words
}

func main(){
	//WordCount("I have a pen and a pencil.")
	wc.Test(WordCount)
}

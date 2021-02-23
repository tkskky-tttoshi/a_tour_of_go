package main

import (
	"golang.org/x/tour/pic"
	//"code.google.com/p/go-tour/pic"
)


func Pic(dx, dy int) [][]uint8 {
	//↓エラー
	//GOでは配列の初期化は変数ではできない
	//var slice [dx][dy]uint8
	//slice := make([][]uint8, dx, dy)

	//function x * y
	slice := make([][]uint8, dx, dy)
	for index := range slice{
		slice[index] = make([]uint8, dy)
	}
	for i := range slice{
		for j := range slice[i]{
			slice[i][j] = uint8(i * j)
		}
	}
	return slice

}

func main(){
	pic.Show(Pic)
}

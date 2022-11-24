// Go言語の構造体の埋め込み型について

package main

import "fmt"

type Dog struct{}

func (d *Dog) Bark() string {
	return "Bow"
}

//Dogを部品として持つ
type BUllDog struct{ Dog }

//Dogを部品として持つ
type ShibaInu struct{ Dog }

func (s *ShibaInu) Bark() string {
	return "ワン"
}

func DogVoice(d *Dog) string {
	return d.Bark()
}

func main() {

	//BUllDogのインスタンス
	bd := &BUllDog{}
	fmt.Println(bd.Bark())

	//ShibaInuのインスタンス
	si := &ShibaInu{}
	fmt.Println(si.Bark())

	// Go言語の構造体の埋め込み型はオブジェクト指向の継承ではなくコンポジションに過ぎないので、以下の処理は実行できない
	// fmt.Println(si)
}

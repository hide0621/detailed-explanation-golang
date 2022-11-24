package main

import "fmt"

type MyErr struct{}

func (me *MyErr) Error() string {
	return ""
}

func Apply() error {
	var err *MyErr = nil
	// 戻り値の「err」はerror型で返すようになっているが、実際の変数errはMyErrのポインタ型なので(MyErrのアドレスにnilと入れている)、戻り値の型と異なっているので「false」となる
	return err
}

func Apply2() error {
	var err error = nil
	// 戻り値の「err」の型と変数errの型が同じなので「true」となる
	// 以下の書き方よりも「return nil」と直接書いた方がミスを防げるので良い
	return err
}

func main() {
	fmt.Println(Apply() == nil)  // false
	fmt.Println(Apply2() == nil) // true
}

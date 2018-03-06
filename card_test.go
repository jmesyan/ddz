package ddz

import (
	"fmt"
	"testing"
)

func TestCardFromString(t *testing.T) {
	//type args struct {
	//	s string
	//}
	//tests := []struct {
	//	name    string
	//	args    args
	//	want    Card
	//	wantErr bool
	//}{
	//// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, err := CardFromString(tt.args.s)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("CardFromString() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if got != tt.want {
	//			t.Errorf("CardFromString() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	cs1 := CardSet()
	rk := RankCount{}
	rk.CountRanks(cs1)
	//for _, v := range cs1 {
	//	r := v.Rank()>>8 - 1
	//	rk.ranks[r]++
	//}
	fmt.Println(rk)
}

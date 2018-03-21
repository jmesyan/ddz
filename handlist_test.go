package ddz

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHandContext(t *testing.T) {
	//type args struct {
	//	cs CardSlice
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want *handContext
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := newHandContext(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("newHandContext() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	str := "♣3 ♣4 ♠5 ♦6 ♠6 ♥7 ♠7 ♦7 ♦8 ♣8 ♣9 ♦9 ♦T"
	cs, err := CardSliceFromString(str, " ")
	if err != nil {
		t.Error(err)
	}
	stdHands := StandardAnalyze(cs)
	for _, h := range stdHands {
		fmt.Println(h)
	}

	fmt.Println("---------------------")

	advHands := AdvancedAnalyze(cs)
	for _, h := range advHands {
		fmt.Println(h)
	}
}

func TestHandContext_Update(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *handContext
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.update(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandContext_searchPrimal(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat    *Hand
		primalNum int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchPrimal(tt.args.toBeat, tt.args.primalNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchPrimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandContext_searchBomb(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat *Hand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchBomb(tt.args.toBeat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchBomb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandContext_searchTrioKicker(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat    *Hand
		kickerNum int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchTrioKicker(tt.args.toBeat, tt.args.kickerNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchTrioKicker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandContext_searchChain(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat    *Hand
		duplicate int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchChain(tt.args.toBeat, tt.args.duplicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextComb(t *testing.T) {
	type args struct {
		comb []int
		k    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextComb(tt.args.comb, tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("nextComb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandContext_searchTrioKickerChain(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat *Hand
		kc     int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchTrioKickerChain(tt.args.toBeat, tt.args.kc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchTrioKickerChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

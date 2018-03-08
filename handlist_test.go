package ddz

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHandContext(t *testing.T) {
	cs := CardSet()
	cs.Sort()
	c1 := CardSlice{
		Club6, Heart6, Spade6, Spade4,
	}
	h := HandParse(c1)
	fmt.Println(h)
	cs = cs.RemoveRank(Rank6)
	ctx := NewHandContext(cs)
	beat := ctx.searchTrioKicker(h, 1)
	fmt.Println(beat)
	//type args struct {
	//	cs CardSlice
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want *HandContext
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := NewHandContext(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("NewHandContext() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
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
		want   *HandContext
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &HandContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.Update(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandContext.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

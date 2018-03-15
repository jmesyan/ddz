package ddz

import (
	"reflect"
	"testing"
)

func TestNewHandContext(t *testing.T) {
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name string
		args args
		want *HandContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandContext(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandContext() = %v, want %v", got, tt.want)
			}
		})
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
			ctx := &HandContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchPrimal(tt.args.toBeat, tt.args.primalNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandContext.searchPrimal() = %v, want %v", got, tt.want)
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
			ctx := &HandContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchBomb(tt.args.toBeat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandContext.searchBomb() = %v, want %v", got, tt.want)
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
			ctx := &HandContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchTrioKicker(tt.args.toBeat, tt.args.kickerNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandContext.searchTrioKicker() = %v, want %v", got, tt.want)
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
			ctx := &HandContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchChain(tt.args.toBeat, tt.args.duplicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandContext.searchChain() = %v, want %v", got, tt.want)
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
			ctx := &HandContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchTrioKickerChain(tt.args.toBeat, tt.args.kc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandContext.searchTrioKickerChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

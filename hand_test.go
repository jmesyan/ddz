package ddz

import (
	"reflect"
	"testing"
)

func TestHand_Copy(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hand.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsChain(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.IsChain(); got != tt.want {
				t.Errorf("Hand.IsChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsNuke(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.IsNuke(); got != tt.want {
				t.Errorf("Hand.IsNuke() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_IsBomb(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.IsBomb(); got != tt.want {
				t.Errorf("Hand.IsBomb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_Primal(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.Primal(); got != tt.want {
				t.Errorf("Hand.Primal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_Kicker(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.Kicker(); got != tt.want {
				t.Errorf("Hand.Kicker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPatternMatch(t *testing.T) {
	type args struct {
		sorted  RankCount
		pattern int
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
			if got := isPatternMatch(tt.args.sorted, tt.args.pattern); got != tt.want {
				t.Errorf("isPatternMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_distribute(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	type args struct {
		cs     CardSlice
		rc     RankCount
		d1     int
		d2     int
		length int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			h.distribute(tt.args.cs, tt.args.rc, tt.args.d1, tt.args.d2, tt.args.length)
		})
	}
}

func TestHand_checkNuke(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.checkNuke(tt.args.cs); got != tt.want {
				t.Errorf("Hand.checkNuke() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_checkBomb(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	type args struct {
		cs     CardSlice
		sorted RankCount
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.checkBomb(tt.args.cs, tt.args.sorted); got != tt.want {
				t.Errorf("Hand.checkBomb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandParse(t *testing.T) {
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name string
		args args
		want *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandParse(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_compareBomb(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	type args struct {
		rhs Hand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   HandCompareResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.compareBomb(tt.args.rhs); got != tt.want {
				t.Errorf("Hand.compareBomb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_Compare(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	type args struct {
		rhs Hand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   HandCompareResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.Compare(tt.args.rhs); got != tt.want {
				t.Errorf("Hand.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_String(t *testing.T) {
	type fields struct {
		Type  byte
		Cards CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				Type:  tt.fields.Type,
				Cards: tt.fields.Cards,
			}
			if got := h.String(); got != tt.want {
				t.Errorf("Hand.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

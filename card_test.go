package ddz

import (
	"reflect"
	"testing"
)

func TestCard_Prime(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want uint32
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Prime(); got != tt.want {
				t.Errorf("Card.Prime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Rank(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want Rank
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Rank(); got != tt.want {
				t.Errorf("Card.Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Suit(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want Suit
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Suit(); got != tt.want {
				t.Errorf("Card.Suit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Bits(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want uint32
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Bits(); got != tt.want {
				t.Errorf("Card.Bits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsBlack(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsBlack(); got != tt.want {
				t.Errorf("Card.IsBlack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsRed(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsRed(); got != tt.want {
				t.Errorf("Card.IsRed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsJoker(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsJoker(); got != tt.want {
				t.Errorf("Card.IsJoker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_String(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Card.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardMake(t *testing.T) {
	type args struct {
		p uint32
		r uint32
		s uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want Card
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CardMake(tt.args.p, tt.args.r, tt.args.s, tt.args.b); got != tt.want {
				t.Errorf("CardMake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Card
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CardFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("CardFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CardFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSet(t *testing.T) {
	tests := []struct {
		name string
		want CardSlice
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CardSet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSliceFromString(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    CardSlice
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CardSliceFromString(tt.args.s, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("CardSliceFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSliceFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_String(t *testing.T) {
	tests := []struct {
		name string
		cs   CardSlice
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.String(); got != tt.want {
				t.Errorf("CardSlice.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_Sort(t *testing.T) {
	tests := []struct {
		name string
		cs   CardSlice
		want CardSlice
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_Reverse(t *testing.T) {
	tests := []struct {
		name string
		cs   CardSlice
		want CardSlice
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_Shuffle(t *testing.T) {
	tests := []struct {
		name string
		cs   CardSlice
		want CardSlice
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.Shuffle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_Subtract(t *testing.T) {
	type args struct {
		rhs CardSlice
	}
	tests := []struct {
		name string
		cs   CardSlice
		args args
		want CardSlice
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.Subtract(tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_CopyRank(t *testing.T) {
	type args struct {
		r Rank
	}
	tests := []struct {
		name string
		cs   CardSlice
		args args
		want CardSlice
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.CopyRank(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.CopyRank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_RemoveRank(t *testing.T) {
	type args struct {
		r Rank
	}
	tests := []struct {
		name string
		cs   CardSlice
		args args
		want CardSlice
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.RemoveRank(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.RemoveRank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_Copy(t *testing.T) {
	tests := []struct {
		name string
		cs   CardSlice
		want CardSlice
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_Ranks(t *testing.T) {
	tests := []struct {
		name string
		cs   CardSlice
		want RankCount
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.Ranks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.Ranks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankCount_Count(t *testing.T) {
	type args struct {
		r Rank
	}
	tests := []struct {
		name string
		rc   RankCount
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rc.Count(tt.args.r); got != tt.want {
				t.Errorf("RankCount.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankCount_Copy(t *testing.T) {
	tests := []struct {
		name string
		rc   RankCount
		want RankCount
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rc.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RankCount.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankCount_Update(t *testing.T) {
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name string
		rc   *RankCount
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rc.Update(tt.args.cs)
		})
	}
}

func TestRankCount_Sort(t *testing.T) {
	tests := []struct {
		name string
		rc   RankCount
		want RankCount
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rc.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RankCount.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankCount_Equals(t *testing.T) {
	type args struct {
		rhs RankCount
	}
	tests := []struct {
		name string
		rc   RankCount
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rc.Equals(tt.args.rhs); got != tt.want {
				t.Errorf("RankCount.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankCount_IsChain(t *testing.T) {
	type args struct {
		duplicate    int
		expectLength int
	}
	tests := []struct {
		name string
		rc   RankCount
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rc.IsChain(tt.args.duplicate, tt.args.expectLength); got != tt.want {
				t.Errorf("RankCount.IsChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankCount_String(t *testing.T) {
	tests := []struct {
		name string
		rc   RankCount
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rc.String(); got != tt.want {
				t.Errorf("RankCount.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

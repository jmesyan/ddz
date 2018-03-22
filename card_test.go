package ddz

import (
	"reflect"
	"testing"
)

func TestCard_Rank(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want Rank
	}{
		{"rank3", Club3, Rank3},
		{"rank4", Heart4, Rank4},
		{"rank5", Diamond5, Rank5},
		{"rank6", Spade6, Rank6},
		{"rank7", Club7, Rank7},
		{"rank8", Heart8, Rank8},
		{"rank9", Diamond9, Rank9},
		{"rankT", SpadeT, RankT},
		{"rankJ", ClubJ, RankJ},
		{"rankQ", HeartQ, RankQ},
		{"rankK", DiamondK, RankK},
		{"rankA", SpadeA, RankA},
		{"rank2", Club2, Rank2},
		{"rankr", Jokerr, Rankr},
		{"rankR", JokerR, RankR},
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
		{"club", Club3, SuitClub},
		{"heart", Heart4, SuitHeart},
		{"diamond", Diamond5, SuitDiamond},
		{"spade", Spade6, SuitSpade},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Suit(); got != tt.want {
				t.Errorf("Card.Suit() = %v, want %v", got, tt.want)
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
		{"club", Club3, true},
		{"heart", Heart4, false},
		{"diamond", Diamond5, false},
		{"spade", Spade6, true},
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
		{"club", Club3, false},
		{"heart", Heart4, true},
		{"diamond", Diamond5, true},
		{"spade", Spade6, false},
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
		{"Joker color", JokerR, true},
		{"Joker black", Jokerr, true},
		{"other", Club3, false},
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
		{"♣3", Club3, "♣3"},
		{"♣4", Club4, "♣4"},
		{"♣5", Club5, "♣5"},
		{"♣6", Club6, "♣6"},
		{"♣7", Club7, "♣7"},
		{"♣8", Club8, "♣8"},
		{"♣9", Club9, "♣9"},
		{"♣T", ClubT, "♣T"},
		{"♣J", ClubJ, "♣J"},
		{"♣Q", ClubQ, "♣Q"},
		{"♣K", ClubK, "♣K"},
		{"♣A", ClubA, "♣A"},
		{"♣2", Club2, "♣2"},
		{"♦3", Diamond3, "♦3"},
		{"♦4", Diamond4, "♦4"},
		{"♦5", Diamond5, "♦5"},
		{"♦6", Diamond6, "♦6"},
		{"♦7", Diamond7, "♦7"},
		{"♦8", Diamond8, "♦8"},
		{"♦9", Diamond9, "♦9"},
		{"♦T", DiamondT, "♦T"},
		{"♦J", DiamondJ, "♦J"},
		{"♦Q", DiamondQ, "♦Q"},
		{"♦K", DiamondK, "♦K"},
		{"♦A", DiamondA, "♦A"},
		{"♦2", Diamond2, "♦2"},
		{"♥3", Heart3, "♥3"},
		{"♥4", Heart4, "♥4"},
		{"♥5", Heart5, "♥5"},
		{"♥6", Heart6, "♥6"},
		{"♥7", Heart7, "♥7"},
		{"♥8", Heart8, "♥8"},
		{"♥9", Heart9, "♥9"},
		{"♥T", HeartT, "♥T"},
		{"♥J", HeartJ, "♥J"},
		{"♥Q", HeartQ, "♥Q"},
		{"♥K", HeartK, "♥K"},
		{"♥A", HeartA, "♥A"},
		{"♥2", Heart2, "♥2"},
		{"♠3", Spade3, "♠3"},
		{"♠4", Spade4, "♠4"},
		{"♠5", Spade5, "♠5"},
		{"♠6", Spade6, "♠6"},
		{"♠7", Spade7, "♠7"},
		{"♠8", Spade8, "♠8"},
		{"♠9", Spade9, "♠9"},
		{"♠T", SpadeT, "♠T"},
		{"♠J", SpadeJ, "♠J"},
		{"♠Q", SpadeQ, "♠Q"},
		{"♠K", SpadeK, "♠K"},
		{"♠A", SpadeA, "♠A"},
		{"♠2", Spade2, "♠2"},
		{"♣r", Jokerr, "♣r"},
		{"♦R", JokerR, "♦R"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Card.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeCard(t *testing.T) {
	type args struct {
		suit Suit
		rank Rank
	}
	tests := []struct {
		name string
		args args
		want Card
	}{
		{"club3", args{suit: SuitClub, rank: Rank3}, Club3},
		{"heart4", args{suit: SuitHeart, rank: Rank4}, Heart4},
		{"diamond5", args{suit: SuitDiamond, rank: Rank5}, Diamond5},
		{"spade6", args{suit: SuitSpade, rank: Rank6}, Spade6},
		{"jokerr", args{suit: SuitSpade, rank: Rankr}, Jokerr},
		{"jokerR", args{suit: SuitHeart, rank: RankR}, JokerR},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeCard(tt.args.suit, tt.args.rank); got != tt.want {
				t.Errorf("MakeCard() = %v, want %v", got, tt.want)
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
		{"♣3", args{s: "♣3"}, Club3, false},
		{"♣4", args{s: "♣4"}, Club4, false},
		{"♣5", args{s: "♣5"}, Club5, false},
		{"♣6", args{s: "♣6"}, Club6, false},
		{"♣7", args{s: "♣7"}, Club7, false},
		{"♣8", args{s: "♣8"}, Club8, false},
		{"♣9", args{s: "♣9"}, Club9, false},
		{"♣T", args{s: "♣T"}, ClubT, false},
		{"♣J", args{s: "♣J"}, ClubJ, false},
		{"♣Q", args{s: "♣Q"}, ClubQ, false},
		{"♣K", args{s: "♣K"}, ClubK, false},
		{"♣A", args{s: "♣A"}, ClubA, false},
		{"♣2", args{s: "♣2"}, Club2, false},
		{"♦3", args{s: "♦3"}, Diamond3, false},
		{"♦4", args{s: "♦4"}, Diamond4, false},
		{"♦5", args{s: "♦5"}, Diamond5, false},
		{"♦6", args{s: "♦6"}, Diamond6, false},
		{"♦7", args{s: "♦7"}, Diamond7, false},
		{"♦8", args{s: "♦8"}, Diamond8, false},
		{"♦9", args{s: "♦9"}, Diamond9, false},
		{"♦T", args{s: "♦T"}, DiamondT, false},
		{"♦J", args{s: "♦J"}, DiamondJ, false},
		{"♦Q", args{s: "♦Q"}, DiamondQ, false},
		{"♦K", args{s: "♦K"}, DiamondK, false},
		{"♦A", args{s: "♦A"}, DiamondA, false},
		{"♦2", args{s: "♦2"}, Diamond2, false},
		{"♥3", args{s: "♥3"}, Heart3, false},
		{"♥4", args{s: "♥4"}, Heart4, false},
		{"♥5", args{s: "♥5"}, Heart5, false},
		{"♥6", args{s: "♥6"}, Heart6, false},
		{"♥7", args{s: "♥7"}, Heart7, false},
		{"♥8", args{s: "♥8"}, Heart8, false},
		{"♥9", args{s: "♥9"}, Heart9, false},
		{"♥T", args{s: "♥T"}, HeartT, false},
		{"♥J", args{s: "♥J"}, HeartJ, false},
		{"♥Q", args{s: "♥Q"}, HeartQ, false},
		{"♥K", args{s: "♥K"}, HeartK, false},
		{"♥A", args{s: "♥A"}, HeartA, false},
		{"♥2", args{s: "♥2"}, Heart2, false},
		{"♠3", args{s: "♠3"}, Spade3, false},
		{"♠4", args{s: "♠4"}, Spade4, false},
		{"♠5", args{s: "♠5"}, Spade5, false},
		{"♠6", args{s: "♠6"}, Spade6, false},
		{"♠7", args{s: "♠7"}, Spade7, false},
		{"♠8", args{s: "♠8"}, Spade8, false},
		{"♠9", args{s: "♠9"}, Spade9, false},
		{"♠T", args{s: "♠T"}, SpadeT, false},
		{"♠J", args{s: "♠J"}, SpadeJ, false},
		{"♠Q", args{s: "♠Q"}, SpadeQ, false},
		{"♠K", args{s: "♠K"}, SpadeK, false},
		{"♠A", args{s: "♠A"}, SpadeA, false},
		{"♠2", args{s: "♠2"}, Spade2, false},
		{"♣r", args{s: "♣r"}, Jokerr, false},
		{"♦R", args{s: "♦R"}, JokerR, false},
		{"err", args{s: "err"}, Card(0), true},
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

func TestCardSlice_search(t *testing.T) {
	type args struct {
		c Card
		f func(int) bool
	}
	tests := []struct {
		name string
		cs   CardSlice
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.search(tt.args.c, tt.args.f); got != tt.want {
				t.Errorf("CardSlice.search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_Search(t *testing.T) {
	type args struct {
		c      Card
		ascend bool
	}
	tests := []struct {
		name string
		cs   CardSlice
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.Search(tt.args.c, tt.args.ascend); got != tt.want {
				t.Errorf("CardSlice.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_Contains(t *testing.T) {
	type args struct {
		rhs    CardSlice
		ascend bool
	}
	tests := []struct {
		name string
		cs   CardSlice
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.Contains(tt.args.rhs, tt.args.ascend); got != tt.want {
				t.Errorf("CardSlice.Contains() = %v, want %v", got, tt.want)
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

package ddz

import (
	"errors"
	"math/rand"
	"sort"
)

// Card represent poker card with uint32 number
// xbbbbbbb|bbbbbbbb|cdhsrrrr|xxpppppp
// p: prime number (duce=2, trey=3, ace=41)
// r: rank (duce=0, trey=1, ace=12)
// cdhs: suits (c=club, d=diamond, h=heart, s=spade)
// b: rank bitmask
type Card uint32

const (
	maskPrime = 0x000000FF
	maskRank  = 0x00000F00
	maskSuit  = 0x0000F000
	maskBits  = 0xFFFF0000

	Prime3 = 0x00000002
	Prime4 = 0x00000003
	Prime5 = 0x00000005
	Prime6 = 0x00000007
	Prime7 = 0x0000000B
	Prime8 = 0x0000000D
	Prime9 = 0x00000011
	PrimeT = 0x00000013
	PrimeJ = 0x00000017
	PrimeQ = 0x0000001D
	PrimeK = 0x0000001F
	PrimeA = 0x00000025
	Prime2 = 0x00000029
	Primer = 0x0000002B
	PrimeR = 0x0000002F

	Rank3 = 0x000000100
	Rank4 = 0x000000200
	Rank5 = 0x000000300
	Rank6 = 0x000000400
	Rank7 = 0x000000500
	Rank8 = 0x000000600
	Rank9 = 0x000000700
	RankT = 0x000000800
	RankJ = 0x000000900
	RankQ = 0x000000A00
	RankK = 0x000000B00
	RankA = 0x000000C00
	Rank2 = 0x000000D00
	Rankr = 0x000000E00
	RankR = 0x000000F00

	SuitClub    = 0x00008000
	SuitDiamond = 0x00004000
	SuitHeart   = 0x00002000
	SuitSpade   = 0x00001000

	Bits3 = 0x00010000
	Bits4 = 0x00020000
	Bits5 = 0x00040000
	Bits6 = 0x00080000
	Bits7 = 0x00100000
	Bits8 = 0x00200000
	Bits9 = 0x00400000
	BitsT = 0x00800000
	BitsJ = 0x01000000
	BitsQ = 0x02000000
	BitsK = 0x04000000
	BitsA = 0x08000000
	Bits2 = 0x10000000
	Bitsr = 0x20000000
	BitsR = 0x40000000

	Club3    = Card(0x00018102)
	Club4    = Card(0x00028203)
	Club5    = Card(0x00048305)
	Club6    = Card(0x00088407)
	Club7    = Card(0x0010850B)
	Club8    = Card(0x0020860D)
	Club9    = Card(0x00408711)
	ClubT    = Card(0x00808813)
	ClubJ    = Card(0x01008917)
	ClubQ    = Card(0x02008A1D)
	ClubK    = Card(0x04008B1F)
	ClubA    = Card(0x08008C25)
	Club2    = Card(0x10008D29)
	Diamond3 = Card(0x00014102)
	Diamond4 = Card(0x00024203)
	Diamond5 = Card(0x00044305)
	Diamond6 = Card(0x00084407)
	Diamond7 = Card(0x0010450B)
	Diamond8 = Card(0x0020460D)
	Diamond9 = Card(0x00404711)
	DiamondT = Card(0x00804813)
	DiamondJ = Card(0x01004917)
	DiamondQ = Card(0x02004A1D)
	DiamondK = Card(0x04004B1F)
	DiamondA = Card(0x08004C25)
	Diamond2 = Card(0x10004D29)
	Heart3   = Card(0x00012102)
	Heart4   = Card(0x00022203)
	Heart5   = Card(0x00042305)
	Heart6   = Card(0x00082407)
	Heart7   = Card(0x0010250B)
	Heart8   = Card(0x0020260D)
	Heart9   = Card(0x00402711)
	HeartT   = Card(0x00802813)
	HeartJ   = Card(0x01002917)
	HeartQ   = Card(0x02002A1D)
	HeartK   = Card(0x04002B1F)
	HeartA   = Card(0x08002C25)
	Heart2   = Card(0x10002D29)
	Spade3   = Card(0x00011102)
	Spade4   = Card(0x00021203)
	Spade5   = Card(0x00041305)
	Spade6   = Card(0x00081407)
	Spade7   = Card(0x0010150B)
	Spade8   = Card(0x0020160D)
	Spade9   = Card(0x00401711)
	SpadeT   = Card(0x00801813)
	SpadeJ   = Card(0x01001917)
	SpadeQ   = Card(0x02001A1D)
	SpadeK   = Card(0x04001B1F)
	SpadeA   = Card(0x08001C25)
	Spade2   = Card(0x10001D29)
	Jokerr   = Card(0x20008E2B)
	JokerR   = Card(0x40004F2F)
)

var (
	ErrorInvalidFormat = errors.New("invalid card string format")
	suitMap            map[uint32]string
	rankMap            map[uint32]string
	cardMap            map[string]Card
)

func init() {
	suitMap = make(map[uint32]string)
	suitMap[SuitClub] = "♣"
	suitMap[SuitDiamond] = "♦"
	suitMap[SuitHeart] = "♥"
	suitMap[SuitSpade] = "♠"

	rankMap = make(map[uint32]string)
	rankMap[Rank3] = "3"
	rankMap[Rank4] = "4"
	rankMap[Rank5] = "5"
	rankMap[Rank6] = "6"
	rankMap[Rank7] = "7"
	rankMap[Rank8] = "8"
	rankMap[Rank9] = "9"
	rankMap[RankT] = "T"
	rankMap[RankJ] = "J"
	rankMap[RankQ] = "Q"
	rankMap[RankK] = "K"
	rankMap[RankA] = "A"
	rankMap[Rank2] = "2"
	rankMap[Rankr] = "r"
	rankMap[RankR] = "R"

	cardMap = make(map[string]Card)
	cardMap["♣3"] = Club3
	cardMap["♣4"] = Club4
	cardMap["♣5"] = Club5
	cardMap["♣6"] = Club6
	cardMap["♣7"] = Club7
	cardMap["♣8"] = Club8
	cardMap["♣9"] = Club9
	cardMap["♣T"] = ClubT
	cardMap["♣J"] = ClubJ
	cardMap["♣Q"] = ClubQ
	cardMap["♣K"] = ClubK
	cardMap["♣A"] = ClubA
	cardMap["♣2"] = Club2
	cardMap["♦3"] = Diamond3
	cardMap["♦4"] = Diamond4
	cardMap["♦5"] = Diamond5
	cardMap["♦6"] = Diamond6
	cardMap["♦7"] = Diamond7
	cardMap["♦8"] = Diamond8
	cardMap["♦9"] = Diamond9
	cardMap["♦T"] = DiamondT
	cardMap["♦J"] = DiamondJ
	cardMap["♦Q"] = DiamondQ
	cardMap["♦K"] = DiamondK
	cardMap["♦A"] = DiamondA
	cardMap["♦2"] = Diamond2
	cardMap["♥3"] = Heart3
	cardMap["♥4"] = Heart4
	cardMap["♥5"] = Heart5
	cardMap["♥6"] = Heart6
	cardMap["♥7"] = Heart7
	cardMap["♥8"] = Heart8
	cardMap["♥9"] = Heart9
	cardMap["♥T"] = HeartT
	cardMap["♥J"] = HeartJ
	cardMap["♥Q"] = HeartQ
	cardMap["♥K"] = HeartK
	cardMap["♥A"] = HeartA
	cardMap["♥2"] = Heart2
	cardMap["♠3"] = Spade3
	cardMap["♠4"] = Spade4
	cardMap["♠5"] = Spade5
	cardMap["♠6"] = Spade6
	cardMap["♠7"] = Spade7
	cardMap["♠8"] = Spade8
	cardMap["♠9"] = Spade9
	cardMap["♠T"] = SpadeT
	cardMap["♠J"] = SpadeJ
	cardMap["♠Q"] = SpadeQ
	cardMap["♠K"] = SpadeK
	cardMap["♠A"] = SpadeA
	cardMap["♠2"] = Spade2
	cardMap["♣r"] = Jokerr
	cardMap["♦R"] = JokerR
}

// Prime returns card's prime bits
func (c Card) Prime() uint32 {
	return uint32(c & maskPrime)
}

// Rank returns card's rank bits
func (c Card) Rank() uint32 {
	return uint32(c & maskRank)
}

// Suit returns card's suit bits
func (c Card) Suit() uint32 {
	return uint32(c & maskSuit)
}

// Suit returns card's bits part
func (c Card) Bits() uint32 {
	return uint32(c & maskBits)
}

// IsBlack true if card's suit is club or spade
func (c Card) IsBlack() bool {
	return c.Suit() == SuitClub || c.Suit() == SuitSpade
}

// IsRed true if card's suit is diamond or heart
func (c Card) IsRed() bool {
	return c.Suit() == SuitDiamond || c.Suit() == SuitHeart
}

// IsJoker true if card's rank is joker
func (c Card) IsJoker() bool {
	return c.Rank() == Rankr || c.Rank() == RankR
}

// String returns card's unicode representation
func (c Card) String() string {
	if str, ok := suitMap[c.Suit()]; ok {
		if post, ok := rankMap[c.Rank()]; ok {
			return str + post
		}
	}

	return ""
}

// MakeCard from bits|suit|rank|prime
func CardMake(p, r, s, b uint32) Card {
	return Card(p | r | s | b)
}

// CardFromString parse string to card
func CardFromString(s string) (Card, error) {
	if c, ok := cardMap[s]; ok {
		return c, nil
	}

	return 0, ErrorInvalidFormat
}

// CardSlice wrapper of card slice
type CardSlice []Card

// CardSet returns a full set of pokers
func CardSet() CardSlice {
	return CardSlice{
		JokerR,
		Jokerr,
		Spade2,
		SpadeA,
		SpadeK,
		SpadeQ,
		SpadeJ,
		SpadeT,
		Spade9,
		Spade8,
		Spade7,
		Spade6,
		Spade5,
		Spade4,
		Spade3,
		Heart2,
		HeartA,
		HeartK,
		HeartQ,
		HeartJ,
		HeartT,
		Heart9,
		Heart8,
		Heart7,
		Heart6,
		Heart5,
		Heart4,
		Heart3,
		Diamond2,
		DiamondA,
		DiamondK,
		DiamondQ,
		DiamondJ,
		DiamondT,
		Diamond9,
		Diamond8,
		Diamond7,
		Diamond6,
		Diamond5,
		Diamond4,
		Diamond3,
		Club2,
		ClubA,
		ClubK,
		ClubQ,
		ClubJ,
		ClubT,
		Club9,
		Club8,
		Club7,
		Club6,
		Club5,
		Club4,
		Club3,
	}
}

// Sort cards from Joker to 3
func (cs CardSlice) Sort() {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i] > cs[j]
	})
}

// Reverse cards
func (cs CardSlice) Reverse() {
	for i := len(cs)/2 - 1; i >= 0; i-- {
		opp := len(cs) - 1 - i
		cs[i], cs[opp] = cs[opp], cs[i]
	}
}

// Shuffle cards using Fisher-Yates algorithm
func (cs CardSlice) Shuffle() {
	for i := len(cs) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cs[i], cs[j] = cs[j], cs[i]
	}
}

package ddz

import (
	"errors"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// Ranks
const (
	Rank3   Rank = 0x01
	Rank4   Rank = 0x02
	Rank5   Rank = 0x03
	Rank6   Rank = 0x04
	Rank7   Rank = 0x05
	Rank8   Rank = 0x06
	Rank9   Rank = 0x07
	RankT   Rank = 0x08
	RankJ   Rank = 0x09
	RankQ   Rank = 0x0A
	RankK   Rank = 0x0B
	RankA   Rank = 0x0C
	Rank2   Rank = 0x0D
	Rankr   Rank = 0x0E
	RankR   Rank = 0x0F
	RankInc Rank = 0x01

	RankCountSize = int(RankR + RankInc)
)

// Suits
const (
	SuitSpade   Suit = 0x10
	SuitHeart   Suit = 0x20
	SuitDiamond Suit = 0x40
	SuitClub    Suit = 0x80
)

const (
	maskRank = 0x0F
	maskSuit = 0xF0
)

// Card set
const (
	Club3    Card = 0x11
	Club4    Card = 0x12
	Club5    Card = 0x13
	Club6    Card = 0x14
	Club7    Card = 0x15
	Club8    Card = 0x16
	Club9    Card = 0x17
	ClubT    Card = 0x18
	ClubJ    Card = 0x19
	ClubQ    Card = 0x1A
	ClubK    Card = 0x1B
	ClubA    Card = 0x1C
	Club2    Card = 0x1D
	Diamond3 Card = 0x21
	Diamond4 Card = 0x22
	Diamond5 Card = 0x23
	Diamond6 Card = 0x24
	Diamond7 Card = 0x25
	Diamond8 Card = 0x26
	Diamond9 Card = 0x27
	DiamondT Card = 0x28
	DiamondJ Card = 0x29
	DiamondQ Card = 0x2A
	DiamondK Card = 0x2B
	DiamondA Card = 0x2C
	Diamond2 Card = 0x2D
	Heart3   Card = 0x41
	Heart4   Card = 0x42
	Heart5   Card = 0x43
	Heart6   Card = 0x44
	Heart7   Card = 0x45
	Heart8   Card = 0x46
	Heart9   Card = 0x47
	HeartT   Card = 0x48
	HeartJ   Card = 0x49
	HeartQ   Card = 0x4A
	HeartK   Card = 0x4B
	HeartA   Card = 0x4C
	Heart2   Card = 0x4D
	Spade3   Card = 0x81
	Spade4   Card = 0x82
	Spade5   Card = 0x83
	Spade6   Card = 0x84
	Spade7   Card = 0x85
	Spade8   Card = 0x86
	Spade9   Card = 0x87
	SpadeT   Card = 0x88
	SpadeJ   Card = 0x89
	SpadeQ   Card = 0x8A
	SpadeK   Card = 0x8B
	SpadeA   Card = 0x8C
	Spade2   Card = 0x8D
	Jokerr   Card = 0x1E
	JokerR   Card = 0x2F
)

// Card represent card with a byte
// |cdhs|rrrr|
// r: rank bits
// cdhs: suit bits (c=club, d=diamond, h=heart, s=spade)
type Card uint8

// Suit represent card's suit with a byte
type Suit uint8

// Rank represent card's rank with a byte
type Rank uint8

// RankCount holds ranks counts in card slice
type RankCount [RankCountSize]int

var (
	// ErrorInvalidFormat error
	ErrorInvalidFormat = errors.New("invalid card string format")
	suitMap            map[Suit]string
	rankMap            map[Rank]string
	cardMap            map[string]Card
)

func init() {
	suitMap = make(map[Suit]string)
	suitMap[SuitClub] = "♣"
	suitMap[SuitDiamond] = "♦"
	suitMap[SuitHeart] = "♥"
	suitMap[SuitSpade] = "♠"

	rankMap = make(map[Rank]string)
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

// Rank returns card's rank bits
func (c Card) Rank() Rank {
	return Rank(c & maskRank)
}

// Suit returns card's suit bits
func (c Card) Suit() Suit {
	return Suit(c & maskSuit)
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
func MakeCard(suit Suit, rank Rank) Card {
	return Card(uint8(suit) | uint8(rank))
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

// CardSliceFromString parse card slice from string
func CardSliceFromString(s, sep string) (CardSlice, error) {
	if sep == "" {
		sep = " "
	}

	segs := strings.Split(s, sep)
	cs := CardSlice{}
	for _, v := range segs {
		c, err := CardFromString(v)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}

	return cs, nil
}

// String interface
func (cs CardSlice) String() string {
	str := ""
	for i, v := range cs {
		str += v.String()
		if i < len(cs)-1 {
			str += " "
		}
	}

	return str
}

// Sort cards in ascending order
func (cs CardSlice) Sort() CardSlice {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i] < cs[j]
	})

	return cs
}

// Reverse cards
func (cs CardSlice) Reverse() CardSlice {
	for i := len(cs)/2 - 1; i >= 0; i-- {
		opp := len(cs) - 1 - i
		cs[i], cs[opp] = cs[opp], cs[i]
	}

	return cs
}

// Shuffle cards using Fisher-Yates algorithm
func (cs CardSlice) Shuffle() CardSlice {
	for i := len(cs) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cs[i], cs[j] = cs[j], cs[i]
	}

	return cs
}

// Subtract remove cards that appear in rhs and return new card slice
func (cs CardSlice) Subtract(rhs CardSlice) CardSlice {
	n := CardSlice{}
	for _, x := range cs {
		found := false
		for _, y := range rhs {
			if x == y {
				found = true
				break
			}
		}
		if !found {
			n = append(n, x)
		}
	}

	return n
}

// CopyRank cards with rank from slice
func (cs CardSlice) CopyRank(r Rank) CardSlice {
	ret := CardSlice{}
	for _, v := range cs {
		if v.Rank() == r {
			ret = append(ret, v)
		}
	}

	return ret
}

// RemoveRank removes cards with rank, and return new card slice
func (cs CardSlice) RemoveRank(r Rank) CardSlice {
	n := CardSlice{}
	for _, x := range cs {
		if x.Rank() != r {
			n = append(n, x)
		}
	}

	return n
}

// Copy card slice
func (cs CardSlice) Copy() CardSlice {
	n := make(CardSlice, len(cs))
	copy(n, cs)
	return n
}

// Ranks returns rank count from 3~Joker
func (cs CardSlice) Ranks() RankCount {
	rc := RankCount{}
	rc.Update(cs)
	return rc
}

// Search return position of card
// cs must be sorted
// if c is not existed, a position of c might appear in cs will be returned
//
func (cs CardSlice) search(c Card, f func(int) bool) int {
	i, j := 0, len(cs)
	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

// Search card c in card slice cs, cs must be sorted
func (cs CardSlice) Search(c Card, ascend bool) int {
	if ascend {
		return cs.search(c, func(i int) bool {
			return cs[i] >= c
		})
	}
	return cs.search(c, func(i int) bool {
		return cs[i] <= c
	})
}

// Contains returns true if card slice contains another
// cs must be sorted, rhs can be unsorted
//
func (cs CardSlice) Contains(rhs CardSlice, ascend bool) bool {
	length := len(rhs)
	if len(cs) < length {
		return false
	}

	for _, r := range rhs {
		i := cs.Search(r, ascend)
		if i < len(cs) && cs[i] == r {
			length--
		}
	}

	return length == 0
}

// Copy returns a copy of rank count
func (rc RankCount) Copy() RankCount {
	n := RankCount{}
	copy(n[:], rc[:])
	return n
}

// Update count ranks in card slice
func (rc *RankCount) Update(cs CardSlice) {
	for _, v := range cs {
		rc[v.Rank()]++
	}
}

// Sort rank count in descending order and return a new rank count
func (rc RankCount) Sort() RankCount {
	n := rc.Copy()
	sort.Slice(n[:], func(i, j int) bool {
		return n[i] > n[j]
	})
	return n
}

// Equals return true if two rank count are identical
func (rc RankCount) Equals(rhs RankCount) bool {
	for i := 0; i < RankCountSize; i++ {
		if rc[i] != rhs[i] {
			return false
		}
	}
	return true
}

// IsChain checks pattern like 334455 666777 etc
// | 666 | 777 | 888 | 999 |
// | 123 |                   duplicate: 3
// |  1     2     3     4  | expectLength: 4
func (rc RankCount) IsChain(duplicate, expectLength int) bool {
	marker := 0
	length := 0
	// joker and 2 cannot chained up
	for i := Rank3; i <= Rank2; i += RankInc {
		// found first match
		if rc[i] == duplicate && marker == 0 {
			marker = int(i)
			continue
		}

		// matches end
		if rc[i] != duplicate && marker != 0 || i == Rank2 {
			length = (int(i) - marker) / int(RankInc)
			break
		}
	}
	return length == expectLength
}

// String ify rank count
func (rc RankCount) String() string {
	s := ""
	for i, v := range rc {
		s += strconv.Itoa(v)
		if i < len(rc) {
			s += " "
		}
	}

	return s
}

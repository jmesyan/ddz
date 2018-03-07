package ddz

/*
 * brief about Hand_Parse function
 * --------------------------------------
 *
 * Landlord has few types of hands,
 * for example, solo-chain as 34567, trio-solo-chain as 333444555A26
 * their patterns are 11111, 333111
 * which can help us to determine the type of a specific hand
 *
 * for example, a 6 cards can have 4 types of hands
 * solo-chain, pair-chain, trio-chain and four-dual-solo
 *
 * the parse process can be simply describe as
 * 1. rank count            --  count every rank in the card array
 * 2. rank count sort       --  sort the count result
 * 3. pattern match         --  iterate through possible patterns
 *
 *
 *
 *  hand specification contains 5 elements
 *  --------------------------------------
 *  pattern
 *  length
 *  primal
 *  kicker
 *  chain
 *
 *  except solo/pair/trio chains
 *  hands that had same length may have 2 variations at most
 */

// Hand is a valid card set that can play.
// cards format must be like 12345/112233/1112223344/11122234 etc
type Hand struct {
	Type  byte      // hand type
	Cards CardSlice // cards
}

// HandCompareResult represent the compare result between hands
type HandCompareResult int

const (
	HandMinLength          = 1
	HandMaxLength          = 20
	HandSoloChainMinLength = 5
	HandPairChainMinLength = 6
	HandTrioChainMinLength = 6
	HandFourChainMinLength = 8

	HandPrimalNone byte = 0x00
	HandPrimalNuke byte = 0x06
	HandPrimalBomb byte = 0x05
	HandPrimalFour byte = 0x04
	HandPrimalTrio byte = 0x03
	HandPrimalPair byte = 0x02
	HandPrimalSolo byte = 0x01

	HandKickerNone     byte = 0x00
	HandKickerSolo     byte = 0x10
	HandKickerPair     byte = 0x20
	HandKickerDualSolo byte = 0x30
	HandKickerDualPair byte = 0x40

	HandChainless byte = 0x00
	HandChain     byte = 0x80

	HandNone       byte = 0x00
	HandSearchMask byte = 0xFF

	handPatternNone = 0  // place holder
	handPattern1    = 1  // 1, solo
	handPattern2a   = 2  // 2, pair
	handPattern2b   = 3  // 2, nuke
	handPattern3    = 4  // 3, trio
	handPattern4a   = 5  // bomb
	handPattern4b   = 6  // trio solo
	handPattern5a   = 7  // solo chain
	handPattern5b   = 8  // trio pair
	handPattern6a   = 9  // solo chain
	handPattern6b   = 10 // pair chain
	handPattern6c   = 11 // trio chain
	handPattern6d   = 12 // four dual solo
	handPattern7    = 13 // solo chain
	handPattern8a   = 14 // solo chain
	handPattern8b   = 15 // pair chain
	handPattern8c   = 16 // trio solo chain
	handPattern8d   = 17 // four dual pair
	handPattern8e   = 18 // four chain
	handPattern9a   = 19 // solo chain
	handPattern9b   = 20 // trio chain
	handPattern10a  = 21 // solo chain
	handPattern10b  = 22 // pair chain
	handPattern10c  = 23 // trio pair chain
	handPattern11   = 24 // solo chain
	handPattern12a  = 25 // solo chain
	handPattern12b  = 26 // pair chain
	handPattern12c  = 27 // trio chain
	handPattern12d  = 28 // trio solo chain
	handPattern12e  = 29 // four chain
	handPattern12f  = 30 // four dual solo chain
	handPattern14   = 31 // pair chain
	handPattern15   = 32 // trio chain
	handPattern16a  = 33 // pair chain
	handPattern16b  = 34 // trio solo chain
	handPattern16c  = 35 // four chain
	handPattern16d  = 36 // four dual pair chain
	handPattern18a  = 37 // pair chain
	handPattern18b  = 38 // trio chain
	handPattern18c  = 39 // four dual solo chain
	handPattern20a  = 40 // pair chain
	handPattern20b  = 41 // trio solo chain
	handPattern20c  = 42 // four chain
	handPatternEnd  = handPattern20c + 1

	patternLength = 12 // 3~A
	handVariation = 2  // except chains, hands with same length will have 2 variation at most
	handSpecNum   = 4  // pattern, primal, kicker, chain

	HandCompareIllegal = -2
	HandCompareLess    = -1
	HandCompareEqual   = 0
	HandCompareGreater = 1
)

var (
	handPattern     [handPatternEnd][patternLength]int
	handSpecs       [HandMaxLength + 1][handVariation][handSpecNum]byte
	primalStringMap map[byte]string
	kickerStringMap map[byte]string
)

func init() {
	handPattern = [handPatternEnd][patternLength]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // place holder
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 1, solo
		{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 2, pair
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 2, nuke
		{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 3, trio
		{4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 4, bomb
		{3, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 4, trio solo
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0}, // 5, solo chain
		{3, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 5, trio pair
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0}, // 6, solo chain
		{2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 6, pair chain
		{3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 6, trio chain
		{4, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 6, four dual solo
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0}, // 7, solo chain
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}, // 8, solo chain
		{2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0}, // 8, pair chain
		{3, 3, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, // 8, trio solo chain
		{4, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 8, four dual pair
		{4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 8, four chain
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}, // 9, solo chain
		{3, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 9, trio chain
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, // 10, solo chain
		{2, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0}, // 10, pair chain
		{3, 3, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0}, // 10, trio pair chain
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}, // 11, solo chain
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // 12, solo chain
		{2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0}, // 12, pair chain
		{3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0}, // 12, trio chain
		{3, 3, 3, 1, 1, 1, 0, 0, 0, 0, 0, 0}, // 12, trio solo chain
		{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 12, four chain
		{4, 4, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0}, // 12, four dual solo chain
		{2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 0}, // 14, pair chain
		{3, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0}, // 15, trio chain
		{2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0}, // 16, pair chain
		{3, 3, 3, 3, 1, 1, 1, 1, 0, 0, 0, 0}, // 16, trio solo chain
		{4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0}, // 16, four chain
		{4, 4, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0}, // 16, four dual pair chain
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0}, // 18, pair chain
		{3, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0}, // 18, trio chain
		{4, 4, 4, 1, 1, 1, 1, 1, 1, 0, 0, 0}, // 18, four dual solo chain
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0}, // 20, pair chain
		{3, 3, 3, 3, 3, 1, 1, 1, 1, 1, 0, 0}, // 20, trio solo chain
		{4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0}, // 20, four chain
	}

	handSpecs = [HandMaxLength + 1][handVariation][handSpecNum]byte{
		{ // place holder
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{
			{handPattern1, HandPrimalSolo, 0, 0},
			{0, 0, 0, 0}},
		{ // 2
			{handPattern2a, HandPrimalPair, 0, 0},
			{0, 0, 0, 0}},
		{ // 3
			{handPattern3, HandPrimalTrio, 0, 0},
			{0, 0, 0, 0}},
		{ // 4
			{handPattern4b, HandPrimalTrio, HandKickerSolo, 0},
			{0, 0, 0, 0}},
		{ // 5
			{handPattern5b, HandPrimalTrio, HandKickerPair, 0},
			{0, 0, 0, 0}},
		{ // 6
			{handPattern6d, HandPrimalFour, HandKickerDualSolo, 0},
			{0, 0, 0, 0}},
		{ // 7
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{ // 8
			{handPattern8c, HandPrimalTrio, HandKickerSolo, HandChain},
			{handPattern8d, HandPrimalFour, HandKickerDualPair, 0}},
		{ // 9
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{ // 10
			{handPattern10c, HandPrimalTrio, HandKickerPair, HandChain},
			{0, 0, 0, 0}},
		{ // 11
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{ // 12
			{handPattern12d, HandPrimalTrio, HandKickerSolo, HandChain},
			{handPattern12f, HandPrimalFour, HandKickerDualSolo, HandChain}},
		{ // 13
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{ // 14
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{ // 15
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{ // 16
			{handPattern16b, HandPrimalTrio, HandKickerSolo, HandChain},
			{handPattern16d, HandPrimalFour, HandKickerDualPair, HandChain}},
		{ // 17
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{ // 18
			{handPattern18c, HandPrimalFour, HandKickerDualSolo, HandChain},
			{0, 0, 0, 0}},
		{ // 19
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{ // 20
			{handPattern20b, HandPrimalTrio, HandKickerSolo, HandChain},
			{0, 0, 0, 0}},
	}

	primalStringMap = map[byte]string{
		HandPrimalNone: "none",
		HandPrimalSolo: "solo",
		HandPrimalPair: "pair",
		HandPrimalTrio: "trio",
		HandPrimalFour: "four",
		HandPrimalBomb: "bomb",
		HandPrimalNuke: "nuke",
	}

	kickerStringMap = map[byte]string{
		HandKickerNone:     "none",
		HandKickerSolo:     "solo",
		HandKickerPair:     "pair",
		HandKickerDualSolo: "dual solo",
		HandKickerDualPair: "dual pair",
	}

}

// Copy returns a copy of hand
func (h Hand) Copy() Hand {
	return Hand{
		Type:  h.Type,
		Cards: h.Cards.Copy(),
	}
}

// IsChain true if hand is chained
func (h Hand) IsChain() bool {
	return h.Type&HandChain != 0
}

// IsNuke true if hand is nuke
func (h Hand) IsNuke() bool {
	return h.Type&HandPrimalNuke != 0
}

// IsBomb ture if hand is bomb
func (h Hand) IsBomb() bool {
	return h.Type&HandPrimalBomb != 0
}

// Primal get primal
func (h Hand) Primal() byte {
	return h.Type & 0x0F
}

// Kicker get kicker
func (h Hand) Kicker() byte {
	return h.Type & 0x70
}

func isPatternMatch(sorted RankCount, pattern int) bool {
	for i := 0; i < len(handPattern[pattern]); i++ {
		if handPattern[pattern][i] != sorted[i] {
			return false
		}
	}
	return true
}

// distribute cards
// for example, distribute(xxx, 88666644, 422, 4, 2, 8)
// hand will be 66668844
func (h *Hand) distribute(cs CardSlice, rc RankCount, d1, d2, length int) {
	temp := CardSlice{}
	for _, card := range cs {
		num := rc.Count(card.Rank())
		if num == d1 {
			h.Cards = append(h.Cards, card)
		} else if num == d2 {
			temp = append(temp, card)
		}

		if len(h.Cards)+len(temp) >= length {
			h.Cards = append(h.Cards, temp...)
			break
		}
	}
}

func (h *Hand) checkNuke(cs CardSlice) bool {
	h.Type = HandNone
	if len(cs) == 2 && cs[0].Rank() == RankR && cs[1].Rank() == Rankr {
		h.Type = HandPrimalNuke | HandKickerNone | HandChainless
		copy(h.Cards, cs)
	}

	return h.IsNuke()
}

func (h *Hand) checkBomb(cs CardSlice, sorted RankCount) bool {
	h.Type = HandNone
	if len(cs) == 4 && isPatternMatch(sorted, handPattern4a) {
		// bomb, 4
		h.Type = HandPrimalBomb | HandKickerNone | HandChainless
		copy(h.Cards, cs)
	}

	return h.IsBomb()
}

func HandParse(cs CardSlice) *Hand {
	if len(cs) < HandMinLength || len(cs) > HandMaxLength {
		return nil
	}

	chainLength := [HandPrimalFour + 1]int{
		0,
		HandSoloChainMinLength,
		HandPairChainMinLength,
		HandTrioChainMinLength,
		HandFourChainMinLength,
	}

	// result
	hand := &Hand{}
	// sort cards
	cs.Sort()
	// count ranks
	rc := cs.Ranks()
	sorted := rc.Sort()
	// card length
	cardLen := len(cs)

	// nuke
	if hand.checkNuke(cs) {
		return hand
	}
	// bomb
	if hand.checkBomb(cs, sorted) {
		return hand
	}
	// chains
	for i := HandPrimalSolo; i < HandPrimalFour+1; i++ {
		chainMinLen := chainLength[i]
		dup := int(i)
		if cardLen >= chainMinLen && cardLen%dup == 0 && rc.IsChain(dup, cardLen/dup) {
			hand.Type = i | HandKickerNone | HandChain
			copy(hand.Cards, cs)
			break
		}
	}

	// chain or other type
	if hand.Type == HandNone {
		d2 := []int{0, 1, 2, 1, 2}

		for i := 0; i < 2; i++ {
			pattern := handSpecs[cardLen][i][0]
			primal := handSpecs[cardLen][i][1]
			kicker := handSpecs[cardLen][i][2]
			chain := handSpecs[cardLen][i][3]

			if pattern == 0 {
				hand.Type = HandNone
				break
			}

			if isPatternMatch(sorted, int(pattern)) {
				d1 := primal
				hand.distribute(cs, rc, int(d1), d2[kicker>>4], cardLen)
				hand.Type = primal | kicker | chain
				break
			}
		}
	}

	return hand
}

// compare between hands, one of the hands must be bomb or nuke
func (h Hand) compareBomb(rhs Hand) HandCompareResult {
	if h.Type == rhs.Type && h.Cards[0].Rank() == rhs.Cards[0].Rank() {
		// same type same ranks, equal
		return HandCompareEqual
	} else if h.Type == HandPrimalBomb && rhs.Type == HandPrimalBomb {
		// both are bombs, compare by card ranks
		if h.Cards[0].Rank() > rhs.Cards[0].Rank() {
			return HandCompareGreater
		} else {
			return HandCompareLess
		}
	} else {
		// nuke > bomb
		if h.Primal() > rhs.Primal() {
			return HandCompareGreater
		} else {
			return HandCompareLess
		}
	}
}

// Compare between hands
// hands must be same type unless there are bombs or nukes
func (h Hand) Compare(rhs Hand) HandCompareResult {
	if h.Type != rhs.Type {
		// different types, check for bombs or nukes
		if !h.IsBomb() && !h.IsNuke() && !rhs.IsBomb() && !rhs.IsNuke() {
			return HandCompareIllegal
		} else {
			return h.compareBomb(rhs)
		}
	} else {
		// same types and not bombs or nukes
		if len(h.Cards) != len(rhs.Cards) {
			// different lengths
			return HandCompareIllegal
		} else {
			// same types and same lengths
			if h.Cards[0].Rank() == rhs.Cards[0].Rank() {
				return HandCompareEqual
			} else {
				if h.Cards[0].Rank() > rhs.Cards[0].Rank() {
					return HandCompareGreater
				} else {
					return HandCompareLess
				}
			}
		}
	}
}

// String ify
func (h Hand) String() string {
	str := primalStringMap[h.Primal()]
	if h.Kicker() != HandKickerNone {
		str += " " + kickerStringMap[h.Kicker()]
	}
	if h.IsChain() {
		str += " " + "chain"
	}

	str += "\n" + h.Cards.String()

	return str
}

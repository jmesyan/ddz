package ddz

// HandList wrapper for hand slice
type HandList []Hand

// HandContext context of beat search
type HandContext struct {
	ranks    RankCount
	cards    CardSlice
	reversed CardSlice
}

// NewHandContext returns new hand context from card slice
func NewHandContext(cs CardSlice) *HandContext {
	ctx := HandContext{}
	ctx.Update(cs)
	return &ctx
}

// Update hand context with sorted card slice
func (ctx *HandContext) Update(cs CardSlice) *HandContext {
	ctx.ranks.Update(cs)
	ctx.cards = cs.Copy()
	ctx.reversed = cs.Copy().Reverse()
	return ctx
}

func (ctx *HandContext) searchPrimal(toBeat *Hand, primalNum int) *Hand {
	var beat *Hand
	rank := toBeat.Cards[0].Rank()
	// search for primal, from low to high rank
	for i := 0; i < len(ctx.cards); {
		v := ctx.cards[i]
		num := ctx.ranks[v.Rank()]
		if v.Rank() > rank && num >= primalNum {
			beat = &Hand{}
			beat.Type = toBeat.Type
			beat.Cards = make(CardSlice, primalNum)
			copy(beat.Cards, ctx.cards[i:i+primalNum])
			break
		}
		i += num
	}

	return beat
}

func (ctx *HandContext) searchBomb(toBeat *Hand) *Hand {
	if toBeat.Type == HandPrimalNuke {
		// cannot beat nuke
		return nil
	}

	var beat *Hand

	// search for a higher rank bomb
	if toBeat.Type == HandPrimalBomb {
		beat = ctx.searchPrimal(toBeat, 4)
	} else {
		// to beat is not a nuke or bomb, search for a bomb to beat
		for i := 0; i < len(ctx.cards); {
			v := ctx.cards[i]
			num := ctx.ranks[v.Rank()]
			if num == 4 {
				beat = &Hand{
					Cards: ctx.cards.CopyRank(v.Rank()),
				}
				break
			} else {
				i += num
			}
		}
	}

	// search for nuke
	if beat == nil {
		if ctx.ranks[Rankr] != 0 && ctx.ranks[RankR] != 0 {
			beat = &Hand{
				Type:  HandPrimalNuke,
				Cards: CardSlice{JokerR, Jokerr},
			}
		}
	} else {
		beat.Type = HandPrimalBomb
	}

	return beat
}

// for a standard 54 card set, each rank has only four cards (except jokers)
// so a trio cannot beat a trio with same rank
// like player one play 333, and it is impossible for player two to beat it with another 333
// BUT when iterate through player's own cards
// we assume trio with same ranks but with different
// kickers can compare with each other
// like 33355 > 33344
func (ctx *HandContext) searchTrioKicker(toBeat *Hand, kickerNum int) *Hand {
	hTrioBeat := &Hand{}
	hKickBeat := &Hand{}
	canBeat := false

	temp := ctx.cards.Copy()

	// copy hands
	hTrio := &Hand{
		Cards: make(CardSlice, 3),
	}
	hKick := &Hand{
		Cards: make(CardSlice, kickerNum),
	}
	copy(hTrio.Cards, toBeat.Cards[0:3])
	copy(hKick.Cards, toBeat.Cards[3:3+kickerNum])

	// same rank trio
	if temp.Contains(hTrio.Cards, true) {
		// keep trio beat
		hTrioBeat.Cards = hTrio.Cards.Copy()
		temp = temp.RemoveRank(hTrioBeat.Cards[0].Rank())

		// search for a higher rank kicker
		// round 1: only search those count[rank] == kick
		for i := 0; i < len(temp); {
			v := temp[i]
			num := ctx.ranks[v.Rank()]
			if num >= kickerNum && v.Rank() > hKick.Cards[0].Rank() {
				hKickBeat.Cards = append(hKickBeat.Cards, temp[i:i+kickerNum]...)
				canBeat = true
				break
			} else {
				i += num
			}
		}

		// if kicker can't beat, restore trio
		if !canBeat {
			hTrioBeat.Cards = CardSlice{}
			temp = ctx.cards.Copy()
		}
	}

	// same rank trio not found
	// OR
	// same rank trio found, but kicker can't beat
	if !canBeat {
		hTrioBeat = ctx.searchPrimal(hTrio, int(HandPrimalTrio))
		if hTrioBeat != nil {
			// trio beat found, search for kicker beat
			// remove trio from temp
			temp = temp.RemoveRank(hTrioBeat.Cards[0].Rank())
			// search for a kicker
			for i := 0; i < len(temp); {
				v := temp[i]
				num := ctx.ranks[v.Rank()]
				if num >= kickerNum {
					hKickBeat.Cards = append(hKickBeat.Cards, temp[i:i+kickerNum]...)
					canBeat = true
					break
				} else {
					i += num
				}
			}
		}
	}

	// beat
	if canBeat {
		beat := &Hand{
			Type:  toBeat.Type,
			Cards: make(CardSlice, 3+kickerNum),
		}
		copy(beat.Cards, hTrioBeat.Cards)
		copy(beat.Cards[3:], hKickBeat.Cards)

		return beat
	}

	return nil
}

func (ctx *HandContext) searchFourKicker(toBeat *Hand, kickerNum int) *Hand {
	hFourBeat := &Hand{}
	hKickBeat := &Hand{}
	canBeat := false

	temp := ctx.cards.Copy()

	// copy hands
	hFour := &Hand{
		Cards: make(CardSlice, 4),
	}
	hKick1 := &Hand{
		Cards: make(CardSlice, kickerNum),
	}
	hKick2 := &Hand{
		Cards: make(CardSlice, kickerNum),
	}
	copy(hFour.Cards, toBeat.Cards[0:4])
	copy(hKick1.Cards, toBeat.Cards[4:4+kickerNum])
	copy(hKick2.Cards, toBeat.Cards[4+kickerNum:4+kickerNum*2])

	// same rank four
	if temp.Contains(hFour.Cards, true) {
		// keep four beat
		hFourBeat.Cards = hFour.Cards.Copy()
		temp = temp.RemoveRank(hFourBeat.Cards[0].Rank())

		// search for a higher rank kicker
		for i := 0; i < len(temp); {
			v := temp[i]
			num := ctx.ranks[v.Rank()]
			if num >= kickerNum && v.Rank() > hKick1.Cards[0].Rank() && v.Rank() != hKick2.Cards[0].Rank() {
				if v.Rank() < hKick2.Cards[0].Rank() {
					hKickBeat.Cards = append(hKickBeat.Cards, temp[i:i+kickerNum]...)
					hKickBeat.Cards = append(hKickBeat.Cards, hKick2.Cards...)
				} else {
					hKickBeat.Cards = append(hKickBeat.Cards, hKick2.Cards...)
					hKickBeat.Cards = append(hKickBeat.Cards, temp[i:i+kickerNum]...)
				}
				canBeat = true
				break
			} else {
				i += num
			}
		}

		// if kicker can't beat, restore four
		if !canBeat {
			hFourBeat.Cards = CardSlice{}
			temp = ctx.cards.Copy()
		}
	}

	// same rank trio not found
	// OR
	// same rank four found, but kicker can't beat
	if !canBeat {
		hFourBeat = ctx.searchPrimal(hFour, int(HandPrimalFour))
		if hFourBeat != nil {
			// trio beat found, search for kicker beat
			// remove trio from temp
			temp = temp.RemoveRank(hFourBeat.Cards[0].Rank())
			// search for a kicker
			for i := 0; i < len(temp); {
				v := temp[i]
				num := ctx.ranks[v.Rank()]
				if num >= kickerNum {
					hKickBeat.Cards = append(hKickBeat.Cards, temp[i:i+kickerNum]...)
					if len(hKickBeat.Cards) >= kickerNum*2 {
						canBeat = true
						break
					}
				}
				i += num
			}
		}
	}

	// beat
	if canBeat {
		beat := &Hand{
			Type:  toBeat.Type,
			Cards: make(CardSlice, 4+kickerNum*2),
		}
		copy(beat.Cards, hFourBeat.Cards)
		copy(beat.Cards[4:], hKickBeat.Cards)

		return beat
	}

	return nil
}

func (ctx *HandContext) searchChain(toBeat *Hand, duplicate int) *Hand {
	chainLen := len(toBeat.Cards) / duplicate

	// this is ugly, but it seems to be the best way to iterate ranks
	rankLen := Rank(chainLen * int(RankInc))
	footer := toBeat.Cards[0].Rank()
	temp := CardSlice{}
	var found bool

	// search for chain in rank count
	for i := footer + RankInc; i <= Rank2-rankLen; i += RankInc {
		found = true
		for j := Rank(0); j < rankLen; j += RankInc {
			// check if chain breaks
			if ctx.ranks[i+j] < duplicate {
				found = false
				break
			}
		}

		if found {
			footer = i     // beat footer rank
			k := duplicate // how many cards needed for each rank
			for i := 0; i < len(ctx.cards) && chainLen > 0; i++ {
				if ctx.cards[i].Rank() == footer {
					temp = append(temp, ctx.cards[i])
					k--
					if k == 0 {
						k = duplicate
						chainLen--
						footer += RankInc
					}
				}
			}
			break
		}
	}

	if found {
		return &Hand{
			Type:  toBeat.Type,
			Cards: temp,
		}
	}

	return nil
}

// https://compprog.wordpress.com/2007/10/17/generating-combinations-1/
// next_comb(int comb[], int k, int n)
// Generates the next combination of n elements as k after comb
//
// comb => the previous combination ( use (0, 1, 2, ..., k) for first)
// k => the size of the subsets to generate
// n => the size of the original set
//
// Returns: 1 if a valid combination was found
// 0, otherwise
func nextComb(comb []int, k, n int) bool {
	i := k - 1
	comb[i]++
	for i > 0 && comb[i] >= n-k+1+i {
		i--
		comb[i]++
	}
	if comb[0] > n-k {
		// Combination (n-k, n-k+1, ..., n) reached
		// No more combinations can be generated
		return false
	}

	// comb now looks like (..., x, n, n, n, ..., n).
	// Turn it into (..., x, x + 1, x + 2, ...)
	for i = i + 1; i < k; i++ {
		comb[i] = comb[i-1] + 1
	}

	return true
}

func (ctx *HandContext) searchTrioKickerChain(toBeat *Hand, kc int) *Hand {
	chainLen := len(toBeat.Cards) / (3 + kc)
	hTrio := &Hand{
		Cards: toBeat.Cards[0 : 3*chainLen],
		Type:  HandPrimalTrio | HandKickerNone | HandChain,
	}
	hTrioBeat := &Hand{
		Cards: make(CardSlice, 0),
	}
	hKick := &Hand{
		Cards: toBeat.Cards[3*chainLen : 3*chainLen+kc*chainLen],
	}
	hKickBeat := &Hand{
		Cards: make(CardSlice, 0),
	}

	canBeat := false

	// self beat
	temp := ctx.reversed.Copy()
	if temp.Contains(hTrio.Cards, false) {
		n := 0
		kickCount := ctx.ranks.Copy()
		// remove trio from kick count
		for i := 0; i < len(hTrio.Cards); i += 3 {
			kickCount[hTrio.Cards[i].Rank()] = 0
		}

		comb2rank := make([]int, RankCountSize)
		rank2comb := make([]int, RankCountSize)
		j := 0
		// remove count < kc and calculate n
		for i := Rank3; i < Rank2; i += RankInc {
			if kickCount[i] < kc {
				kickCount[i] = 0
			} else {
				n++

				// combination index to rank, and vice versa
				//
				// ranks count [x,0,x,0 ...] might have zeros between available ranks
				// which can not apply next_comb directly
				// use a comb-to-rank map to compress rank count array
				comb2rank[j] = int(i)
				rank2comb[i] = j
				j++
			}
		}

		// combination
		comb := make([]int, RankCountSize)
		for i := 0; i < len(hKick.Cards); i += kc {
			comb[j] = rank2comb[hKick.Cards[i].Rank()]
			j++
		}

		// find next combination
		if nextComb(comb, chainLen, n) {
			for i := 0; i < chainLen; i++ {
				rank := comb2rank[i]
				for j := 0; j < len(temp); j++ {
					if temp[j].Rank() == Rank(rank) {
						hKickBeat.Cards = append(hKickBeat.Cards, temp[j:j+kc]...)
						break
					}
				}
			}
			canBeat = true
			// copy trio to beat
			hTrioBeat.Cards = append(hTrioBeat.Cards, hTrio.Cards...)
			hKickBeat.Cards.Sort()
		}
	}

	// cannot find same rank trio chain, search for higher rank trio
	if !canBeat {
		hTrioBeat = ctx.searchChain(hTrio, 3)
		if hTrioBeat != nil {
			// higher rank trio chain found, search for kickers
			count := ctx.ranks.Copy()
			// remove trio from rank count
			for i := 0; i < len(hTrioBeat.Cards); i += 3 {
				temp = temp.RemoveRank(hTrioBeat.Cards[i].Rank())
				count[hTrioBeat.Cards[i].Rank()] = 0
			}

			for i := 0; i < chainLen; i++ {
				for j := 0; j < len(temp); j++ {
					if count[temp[i].Rank()] >= kc {
						hKickBeat.Cards = append(hKickBeat.Cards, temp[i:i+kc]...)
						temp = temp.RemoveRank(temp[i].Rank())
						break
					}
				}
			}

			if len(hKickBeat.Cards) == chainLen*kc {
				canBeat = true
			}
		}
	}

	// final
	if canBeat {
		beat := &Hand{
			Cards: hTrioBeat.Cards.Copy(),
			Type:  toBeat.Type,
		}
		beat.Cards = append(beat.Cards, hKickBeat.Cards...)
		return beat
	}

	return nil
}

const (
	handTypeSolo          byte = HandPrimalSolo | HandKickerNone | HandChainless
	handTypePair          byte = HandPrimalPair | HandKickerNone | HandChainless
	handTypeTrio          byte = HandPrimalTrio | HandKickerNone | HandChainless
	handTypeTrioPair      byte = HandPrimalTrio | HandKickerPair | HandChainless
	handTypeTrioSolo      byte = HandPrimalTrio | HandKickerSolo | HandChainless
	handTypeFourDualSolo  byte = HandPrimalFour | HandKickerDualSolo | HandChainless
	handTypeFourDualPair  byte = HandPrimalFour | HandKickerDualPair | HandChainless
	handTypeSoloChain     byte = HandPrimalSolo | HandKickerNone | HandChain
	handTypePairChain     byte = HandPrimalPair | HandKickerNone | HandChain
	handTypeTrioChain     byte = HandPrimalTrio | HandKickerNone | HandChain
	handTypeTrioPairChain byte = HandPrimalTrio | HandKickerPair | HandChain
	handTypeTrioSoloChain byte = HandPrimalTrio | HandKickerSolo | HandChain
)

// FindBeat in card slice, return nil if there's no beat
func (cs CardSlice) FindBeat(toBeat *Hand) *Hand {
	// setup search context
	ctx := NewHandContext(cs)

	var beat *Hand

	switch toBeat.Type {
	case handTypeSolo, handTypePair, handTypeTrio:
		beat = ctx.searchPrimal(toBeat, int(toBeat.Primal()))
	case handTypeTrioSolo, handTypeTrioPair:
		beat = ctx.searchTrioKicker(toBeat, int(toBeat.Kicker()>>4))
	case handTypeFourDualSolo, handTypeFourDualPair:
		beat = ctx.searchFourKicker(toBeat, int((toBeat.Kicker()-HandKickerDualSolo)>>4)+1)
	case handTypeSoloChain, handTypePairChain, handTypeTrioChain:
		beat = ctx.searchChain(toBeat, int(toBeat.Primal()))
	case handTypeTrioSoloChain, handTypeTrioPairChain:
		beat = ctx.searchTrioKickerChain(toBeat, int(toBeat.Kicker()>>4))
	}

	if beat == nil {
		beat = ctx.searchBomb(toBeat)
	}

	return beat
}

// FindBeatList finds all beats in card slice, return nil if there's no beat
func (cs *CardSlice) FindBeatList(toBeat *Hand) []*Hand {
	beatList := make([]*Hand, 0)
	for {
		beat := cs.FindBeat(toBeat)
		if beat != nil {
			beatList = append([]*Hand{beat}, beatList...)
		} else {
			break
		}
	}

	if len(beatList) == 0 {
		return nil
	}
	return beatList
}

// extractConsecutive hands like 34567 / 334455 / 333444555 etc
// array is a processed card array holds count[rank] == duplicate
//
func extractConsecutive(cs CardSlice, duplicate int) (CardSlice, []*Hand) {
	primals := []byte{0, HandPrimalSolo, HandPrimalPair, HandPrimalTrio}
	chainLen := []int{0, HandSoloChainMinLength, HandPairChainMinLength, HandTrioChainMinLength}

	if duplicate < 1 || duplicate > 3 || len(cs) == 0 {
		return cs, nil
	}

	handList := make([]*Hand, 0)

	i := duplicate
	cardNum := len(cs) / duplicate
	lastRank := cs[0].Rank()
	for cardNum > 0 {
		if lastRank+RankInc != cs[i].Rank() {
			if i >= chainLen[duplicate] {
				// chain break
				hand := &Hand{
					Cards: cs[0:i],
					Type:  primals[duplicate] | HandChain,
				}
				cs = cs[i:]
				handList = append([]*Hand{hand}, handList...)
			} else {
				// not a chain
				for j := 0; j < i/duplicate; j++ {
					hand := &Hand{
						Cards: cs[0:duplicate],
						Type:  primals[duplicate],
					}
					cs = cs[duplicate:]
					handList = append([]*Hand{hand}, handList...)
				}
			}

			if len(cs) == 0 {
				break
			}
		} else {
			// chain intact
			lastRank = cs[i].Rank()
			i += duplicate
		}
	}
	k := i - duplicate // step back
	if k != 0 && k == len(cs) {
		// all chained up
		if k >= chainLen[duplicate] {
			// can chain up
			hand := &Hand{
				Cards: cs[0 : i-duplicate],
				Type:  primals[duplicate] | HandChain,
			}
			cs = cs[i-duplicate:]
			handList = append([]*Hand{hand}, handList...)
		} else {
			for j := 0; j < k/duplicate; j++ {
				hand := &Hand{
					Cards: cs[0:duplicate],
					Type:  primals[duplicate],
				}
				cs = cs[duplicate:]
				handList = append([]*Hand{hand}, handList...)
			}
		}
	}
	if len(handList) == 0 {
		handList = nil
	}
	return cs, handList
}

// extractNukeBombDeuce extract nuke/bomb/2 from card slice
// card slice, rank count with these cards removed will be returned
func extractNukeBombDeuce(cs CardSlice, rc RankCount) (CardSlice, RankCount, []*Hand) {
	handList := make([]*Hand, 0)
	if rc[Rankr] != 0 && rc[RankR] != 0 {
		// nuke
		hand := &Hand{
			Cards: CardSlice{Jokerr, JokerR},
			Type:  HandPrimalNuke,
		}
		handList = append([]*Hand{hand}, handList...)
		rc[Rankr] = 0
		rc[RankR] = 0
		cs = cs.RemoveRank(Rankr).RemoveRank(RankR)
	}

	for i := Rank3; i < Rankr; i += RankInc {
		// bomb
		if rc[i] == 4 {
			hand := &Hand{
				Cards: cs.CopyRank(i),
				Type:  HandPrimalBomb,
			}
			handList = append([]*Hand{hand}, handList...)
			rc[i] = 0
			cs = cs.RemoveRank(i)
		}
	}

	if rc[Rankr] != 0 || rc[RankR] != 0 {
		var c Card
		var r Rank
		if rc[Rankr] != 0 {
			c = Jokerr
			r = Rankr
		} else {
			c = JokerR
			r = RankR
		}
		// joker
		hand := &Hand{
			Cards: CardSlice{c},
			Type:  HandPrimalSolo,
		}
		handList = append([]*Hand{hand}, handList...)
		rc[r] = 0
		cs = cs.RemoveRank(r)
	}

	if rc[Rank2] != 0 {
		// 2
		hand := &Hand{
			Cards: cs.CopyRank(Rank2),
		}
		switch rc[Rank2] {
		case 1:
			hand.Type = HandPrimalSolo
		case 2:
			hand.Type = HandPrimalPair
		case 3:
			hand.Type = HandPrimalTrio
		}
		rc[Rank2] = 0
		cs = cs.RemoveRank(Rank2)
		handList = append([]*Hand{hand}, handList...)
	}

	if len(handList) == 0 {
		handList = nil
	}
	return cs, rc, handList
}

func StandardAnalyze(cs CardSlice) []*Hand {
	cs = cs.Sort()
	rc := cs.Ranks()

	soloSlice := make(CardSlice, 0)
	pairSlice := make(CardSlice, 0)
	trioSlice := make(CardSlice, 0)
	slices := []CardSlice{soloSlice, pairSlice, trioSlice}

	// nuke, bomb and 2
	cs, rc, handList := extractNukeBombDeuce(cs, rc)

	// copy cards into different slices by their number
	for i := 0; i < len(cs); {
		c := rc[cs[i].Rank()]
		if c != 0 {
			slices[c-1] = append(slices[c-1], cs[i:i+int(c)]...)
			i += c
		} else {
			i++
		}
	}

	// extract chains
	for i := 2; i >= 0; i-- {
		_, l := extractConsecutive(slices[i], i+1)
		if l != nil && len(l) > 0 {
			handList = append(l, handList...)
		}
	}

	if len(handList) == 0 {
		handList = nil
	}

	return handList
}

func (ctx *HandContext) findLongestConsecutive(duplicate int) *Hand {
	// early break
	if duplicate < 1 || duplicate > 3 {
		return nil
	}

	primals := []byte{0, HandPrimalSolo, HandPrimalPair, HandPrimalTrio}
	chainLen := []int{0, HandSoloChainMinLength, HandPairChainMinLength, HandTrioChainMinLength}

	if len(ctx.cards) < chainLen[duplicate] {
		return nil
	}

	rankStart := Rank(0)
	chain := make(CardSlice, 0)

	// i <= Rank2
	// but count[Rank2] must be 0
	// for 2/bomb/nuke has been removed before calling this method
	for i := Rank3; i <= Rank2; i += RankInc {
		// find start of a possible chain
		if rankStart == 0 {
			if ctx.ranks[i] >= duplicate {
				rankStart = i
			}
			continue
		}

		if ctx.ranks[i] < duplicate {
			// chain breaks, extract chain and set new possible start
			if int(i-rankStart)*duplicate >= chainLen[duplicate] && int(i-rankStart) > len(chain) {
				// valid chain, store rank

			}

		}
	}
}

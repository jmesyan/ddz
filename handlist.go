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
		num := ctx.ranks.Count(v.Rank())
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
		if ctx.ranks.Count(Rankr) != 0 && ctx.ranks.Count(RankR) != 0 {
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
			num := ctx.ranks.Count(v.Rank())
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
				num := ctx.ranks.Count(v.Rank())
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
			if ctx.ranks.Count(i+j) < duplicate {
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

func (ctx *HandContext) searchTrioKickerChain(toBeat *Hand, kc int) *Hand {
	chainLen := len(toBeat.Cards) / (3 + kc)
	hTrio := &Hand{
		Cards: toBeat.Cards[0 : 3*chainLen],
		Type:  HandPrimalTrio | HandKickerNone | HandChain,
	}
	hKick := &Hand{
		Cards: toBeat.Cards[3*chainLen : 3*chainLen+kc*chainLen],
	}

	// self beat
	temp := ctx.reversed.Copy()
	if temp.Contains(hTrio.Cards, false) {
		n := 0
		kickCount := ctx.ranks.Copy()
		// remove trio from kick count
		for i := 0; i < len(hTrio.Cards); i += 3 {
			kickCount[hTrio.Cards[i].Rank()] = 0
		}

		// remove count < kc and calculate n

	}

	return nil
}

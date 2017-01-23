package ddz

type HandList []*Hand

func (hl HandList) Push(h *Hand) HandList {
	return append(hl, h)
}

func (hl HandList) Pop() (*Hand, HandList) {
	if len(hl) == 0 {
		return nil, hl
	}
	return hl[len(hl)-1], hl[:len(hl)-1]
}

func (hl HandList) Shift() (*Hand, HandList) {
	if len(hl) == 0 {
		return nil, hl
	}

	return hl[0], hl[1:]
}

func (hl HandList) Unshift(h *Hand) HandList {
	return append(HandList{h}, hl...)
}

func (hl HandList) Remove(h *Hand) HandList {
	for i := 0; i < len(hl); i++ {
		if hl[i] == h {
			return append(hl[:i], hl[i+1:]...)
		}
	}

	return hl
}

func (hl HandList) FindKind(kind int) *Hand {
	for i := 0; i < len(hl); i++ {
		if hl[i].kind == kind {
			return hl[i]
		}
	}

	return nil
}

// Beat search context
type handContext struct {
	count  []int     // rank count
	cards  CardSlice // original cards
	rcards CardSlice // reverse sorted cards
}

// Setup search context from a CardSlice
func newHandContext(array CardSlice) *handContext {
	ctx := new(handContext)
	ctx.count = array.CountRank()
	ctx.cards = array.Clone()
	ctx.rcards = array.Clone().Sort().Reverse()

	return ctx
}

// Search beat to primal
func (ctx *handContext) searchBeatPrimal(tobeat, beat *Hand, primal int) bool {
	rank := CardRank(tobeat.cards[0])
	count := ctx.count
	temp := ctx.rcards
	for i := 0; i < len(temp); i++ {
		if CardRank(temp[i]) > rank && count[CardRank(temp[i])] >= primal {
			beat.kind = tobeat.kind
			beat.cards = temp[i : i+primal]
			return true
		}
	}

	return false
}

// Search beat to bomb
func (ctx *handContext) searchBeatBomb(tobeat, beat *Hand) bool {
	canbeat := false
	// Can't beat nuke
	if tobeat.kind == HandFormat(HandPrimalNuke, HandKickerNone, HandChainless) {
		return false
	}

	if tobeat.kind == HandFormat(HandPrimalBomb, HandKickerNone, HandChainless) {
		// Search for a bomb with higher rank
		canbeat = ctx.searchBeatPrimal(tobeat, beat, 4)
	} else {
		// tobeat is neither a nuke nor a bomb, search a bomb to beat it
		for i := 0; i < len(ctx.cards); i++ {
			if ctx.count[CardRank(ctx.cards[i])] == 4 {
				canbeat = true
				beat.cards = ctx.cards.ExtractRank(CardRank(ctx.cards[i]))
				break
			}
		}
	}

	// Search for a nuke
	if !canbeat {
		if ctx.count[CardRankr] != 0 && ctx.count[CardRankR] != 0 {
			canbeat = true
			beat.kind = HandFormat(HandPrimalNuke, HandKickerNone, HandChainless)
			beat.cards = CardSlice{CardRedJoker, CardBlackJoker}
		}
	} else {
		beat.kind = HandFormat(HandPrimalBomb, HandKickerNone, HandChainless)
	}

	return canbeat
}

// Search a TrioKicker beat
// for a standard 54 card set, each rank has four cards
// so it is impossible to have two trio with same rank at a time
// a) player1 SEARCH BEAT player2 : impossible for 333aa vs 333bb
// but
// b) player1 SEARCH_BEAT_LOOP player1 : possible for 333aa vs 333bb
func (ctx *handContext) searchBeatTrioKicker(tobeat, beat *Hand, kick int) bool {
	trioHand := new(Hand)
	kickHand := new(Hand)
	trioBeat := new(Hand)
	kickBeat := new(Hand)

	canBeat := false
	count := ctx.count
	temp := ctx.rcards.Clone()

	// Copy hands
	trioHand.cards = tobeat.cards[:3]
	kickHand.cards = tobeat.cards[3 : 3+kick]

	// Same rank trio, case b
	if temp.ContainsAll(trioHand.cards) {
		// Keep trio beat
		trioBeat.cards = trioHand.cards.Clone()
		temp = temp.RemoveRank(CardRank(trioBeat.cards[0]))

		// Search for a higher rank kicker
		// Round 1: only search for those count[rank] == kick
		for i := 0; i < len(temp); i++ {
			if count[CardRank(temp[i])] >= kick && CardRank(temp[i]) > CardRank(kickHand.cards[0]) {
				kickBeat.cards = temp.Clone()[i : i+kick]
				canBeat = true
				break
			}
		}

		// If kicker can't beat, restore trio
		if !canBeat {
			trioBeat.cards = CardSlice{}
			temp = ctx.rcards.Clone()
		}
	}

	// Same rank trio not found
	// or
	// Same rank trio found, but kicker can't beat
	if !canBeat {
		canTrioBeat := ctx.searchBeatPrimal(trioHand, trioBeat, HandPrimalTrio)
		// Trio beat found, search for kicker beat
		if canTrioBeat {
			// Remove trio from temp
			temp = temp.RemoveRank(CardRank(trioBeat.cards[0]))

			// Search for a kicker
			for i := 0; i < len(temp); i++ {
				if count[CardRank(temp[i])] >= kick {
					kickBeat.cards = kickBeat.cards.Concat(temp[i : i+kick])
					canBeat = true
					break
				}
			}
		}
	}

	// Beat
	if canBeat {
		beat.cards = trioBeat.cards.Concat(kickBeat.cards)
		beat.kind = tobeat.kind
	}

	return canBeat
}

func (ctx *handContext) searchBeatChain(tobeat, beat *Hand, duplicate int) bool {
	found := false
	temp := CardSlice{}
	chainLen := len(tobeat.cards) / duplicate
	footer := CardRank(tobeat.cards[len(tobeat.cards)-1])

	// Search for beat chain in rank counts
	for i := footer + 1; i <= CardRank2-chainLen; i++ {
		found := true
		for j := 0; j < chainLen; j++ {
			// Check if chain breaks
			if ctx.count[i+j] < duplicate {
				found = false
				break
			}
		}

		if found {
			footer = i     // beat footer rank
			k := duplicate // how many cards needed for each rank

			for i := len(ctx.cards); i >= 0 && chainLen > 0; i-- {
				if CardRank(ctx.cards[i]) == footer {
					temp = temp.Unshift(ctx.cards[i])
					k--

					if k == 0 {
						k = duplicate
						chainLen--
						footer++
					}
				}
			}
			break
		}
	}

	if found {
		beat.kind = tobeat.kind
		beat.cards = temp
		return true
	} else {
		return false
	}
}

func (ctx *handContext) searchBeatTrioKickerChain(tobeat, beat *Hand, kc int) bool {
	canBeat := false
	trioHand := new(Hand)
	kickHand := new(Hand)
	trioBeat := new(Hand)
	kickBeat := new(Hand)

	count := make([]int, len(ctx.count))
	copy(count, ctx.count)
	temp := ctx.rcards.Clone()
	chainLen := len(tobeat.cards) / (HandPrimalTrio + kc)

	// Copy tobeat cards
	trioHand.cards = tobeat.cards.Clone()[:3*chainLen]
	kickHand.cards = tobeat.cards.Clone()[3*chainLen : 3*chainLen+chainLen*kc]
	trioHand.kind = HandFormat(HandPrimalTrio, HandKickerNone, HandChain)

	// Self beat, see searchBeatTrioKicker
	if temp.ContainsAll(trioHand.cards) {
		// Combination total
		n := 0
		// Remove trio from kickCount
		kickCount := make([]int, len(ctx.count))
		copy(kickCount, ctx.count)

		for i := 0; i < len(trioHand.cards); i += 3 {
			kickCount[CardRank(trioHand.cards[i])] = 0
		}

		// Remove count < kc and calculate n
		for i := CardRankBeg; i < CardRankEnd; i++ {
			if kickCount[i] < kc {
				kickCount[i] = 0
			} else {
				n++
			}
		}

		// Setup comb-rank and rank-comb map
		j := 0
		combRank := make([]int, CardRankEnd)
		rankComb := make([]int, CardRankEnd)

		for i := 0; i < CardRankEnd; i++ {
			combRank[i], rankComb[i] = -1, -1
		}

		for i := 0; i < CardRankEnd; i++ {
			if kickCount[i] != 0 {
				combRank[j] = i
				rankComb[i] = j
				j++
			}
		}

		// Setup combination
		comb := make([]int, CardRankEnd)
		for i := 0; i < CardRankEnd; i++ {
			comb[i] = -1
		}

		j = 0
		for i := 0; i < len(kickHand.cards); i += kc {
			comb[j] = rankComb[CardRank(kickHand.cards[i])]
			j++
		}

		// Find next combination
		if NextCombination(comb, chainLen, n) {
			// Next combination found, copy kickers
			for i := 0; i < chainLen; i++ {
				rank := combRank[comb[i]]
				for j = 0; j < len(temp); j++ {
					if CardRank(temp[j]) == rank {
						kickBeat.cards = kickBeat.cards.Concat(temp[j : j+kc])
						break
					}
				}
			}

			canBeat = true

			// Copy trio to beat
			trioBeat.cards = trioBeat.cards.Concat(trioHand.cards).Sort()
		}
	}

	// Can't find same rank trio chain, search for higher rank trio
	if !canBeat {
		// restore rank count
		copy(count, ctx.count)
		canTrioBeat := ctx.searchBeatChain(trioHand, trioBeat, 3)

		// Higher rank trio chain found, search for kickers
		if canTrioBeat {
			// Remove trio from temp
			for i := 0; i < len(trioBeat.cards); i += 3 {
				temp = temp.RemoveRank(CardRank(trioBeat.cards[i]))
				count[CardRank(trioBeat.cards[0])] = 0
			}

			for j := 0; j < chainLen; j++ {
				for i := 0; i < len(temp); i++ {
					if count[CardRank(temp[i])] >= kc {
						trioBeat.cards = trioBeat.cards.Concat(temp[i : i+kc])
						temp = temp.RemoveRank(CardRank(temp[i]))
						break
					}
				}
			}

			if len(kickBeat.cards) == kc*chainLen {
				canBeat = true
			}
		}
	}

	if canBeat {
		beat.kind = tobeat.kind
		beat.cards = trioBeat.cards.Concat(kickBeat.cards)
	}

	return canBeat
}

func searchBeat(cards CardSlice, tobeat, beat *Hand) bool {
    return false
}

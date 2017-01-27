/*
 * Am I writing the right go code ?
 * should I use slice like this ?
 * review the code when the translation from C is done
 */
package ddz

type HandList []*Hand

// ----------------------------------------------------------------------------
// HandList
// ----------------------------------------------------------------------------

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

func (hl HandList) Concat(rhs HandList) HandList {
	return append(hl, rhs...)
}

// ----------------------------------------------------------------------------
// Beat search
// ----------------------------------------------------------------------------

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

func (ctx *handContext) clone() *handContext {
	newContext := new(handContext)
	newContext.count = make([]int, len(ctx.count))
	copy(newContext.count, ctx.count)
	newContext.cards = ctx.cards.Clone()
	newContext.rcards = ctx.rcards.Clone()
	return newContext
}

// Search beat to primal
func (ctx *handContext) searchBeatPrimal(tobeat, beat *Hand, primal int) bool {
	rank := tobeat.cards.RankAt(0)
	count := ctx.count
	temp := ctx.rcards
	for i := 0; i < len(temp); i++ {
		if temp.RankAt(i) > rank && count[temp.RankAt(i)] >= primal {
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
			if ctx.count[ctx.cards.RankAt(i)] == 4 {
				canbeat = true
				beat.cards, _ = ctx.cards.ExtractRank(ctx.cards.RankAt(i))
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
		temp = temp.RemoveRank(trioBeat.cards.RankAt(0))

		// Search for a higher rank kicker
		// Round 1: only search for those count[rank] == kick
		for i := 0; i < len(temp); i++ {
			if count[temp.RankAt(i)] >= kick && temp.RankAt(i) > kickHand.cards.RankAt(0) {
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
			temp = temp.RemoveRank(trioBeat.cards.RankAt(0))

			// Search for a kicker
			for i := 0; i < len(temp); i++ {
				if count[temp.RankAt(i)] >= kick {
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
	footer := tobeat.cards.RankAt(len(tobeat.cards) - 1)

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
				if ctx.cards.RankAt(i) == footer {
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
			kickCount[trioHand.cards.RankAt(i)] = 0
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
			comb[j] = rankComb[kickHand.cards.RankAt(i)]
			j++
		}

		// Find next combination
		if NextCombination(comb, chainLen, n) {
			// Next combination found, copy kickers
			for i := 0; i < chainLen; i++ {
				rank := combRank[comb[i]]
				for j = 0; j < len(temp); j++ {
					if temp.RankAt(j) == rank {
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
				temp = temp.RemoveRank(trioBeat.cards.RankAt(i))
				count[trioBeat.cards.RankAt(0)] = 0
			}

			for j := 0; j < chainLen; j++ {
				for i := 0; i < len(temp); i++ {
					if count[temp.RankAt(i)] >= kc {
						trioBeat.cards = trioBeat.cards.Concat(temp[i : i+kc])
						temp = temp.RemoveRank(temp.RankAt(i))
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
	// Setup search context
	canBeat := false
	ctx := newHandContext(cards)

	// Start search
	switch tobeat.kind {
	case HandFormat(HandPrimalSolo, HandKickerNone, HandCompareEqual):
		canBeat = ctx.searchBeatPrimal(tobeat, beat, HandPrimalSolo)
	case HandFormat(HandPrimalPair, HandKickerNone, HandChainless):
		canBeat = ctx.searchBeatPrimal(tobeat, beat, HandPrimalPair)
	case HandFormat(HandPrimalTrio, HandKickerNone, HandChainless):
		canBeat = ctx.searchBeatPrimal(tobeat, beat, HandPrimalTrio)
	case HandFormat(HandPrimalTrio, HandKickerPair, HandChainless):
		canBeat = ctx.searchBeatTrioKicker(tobeat, beat, HandPrimalPair)
	case HandFormat(HandPrimalTrio, HandKickerSolo, HandChainless):
		canBeat = ctx.searchBeatTrioKicker(tobeat, beat, HandPrimalSolo)
	case HandFormat(HandPrimalSolo, HandKickerNone, HandChain):
		canBeat = ctx.searchBeatChain(tobeat, beat, HandPrimalSolo)
	case HandFormat(HandPrimalPair, HandKickerNone, HandChain):
		canBeat = ctx.searchBeatChain(tobeat, beat, HandPrimalPair)
	case HandFormat(HandPrimalTrio, HandKickerNone, HandChain):
		canBeat = ctx.searchBeatChain(tobeat, beat, HandPrimalTrio)
	case HandFormat(HandPrimalFour, HandKickerNone, HandChain):
		canBeat = ctx.searchBeatChain(tobeat, beat, HandPrimalFour)
	case HandFormat(HandPrimalTrio, HandKickerPair, HandChain):
		canBeat = ctx.searchBeatTrioKickerChain(tobeat, beat, HandPrimalPair)
	case HandFormat(HandPrimalTrio, HandKickerSolo, HandChain):
		canBeat = ctx.searchBeatTrioKickerChain(tobeat, beat, HandPrimalSolo)
	}

	// Final solution, search for bomb and nuke
	if !canBeat {
		canBeat = ctx.searchBeatBomb(tobeat, beat)
	}

	return canBeat
}

// Search for beat, result will be store in beat
// 1, if [beat->type] != 0, then search [new beat] > [beat]
// 2, search [beat] > [tobeat], then store in [beat]
func (cards CardSlice) SearchBeat(tobeat, beat *Hand) bool {
	// Already in search loop, continue
	if beat.kind != HandNone {
		return searchBeat(cards, beat, beat)
	} else {
		return searchBeat(cards, tobeat, beat)
	}
}

func (cards CardSlice) SearchBeatList(tobeat *Hand) HandList {
	canBeat := true
	beat := new(Hand)
	handToBeat := tobeat.Clone()
	handList := HandList{}

	for canBeat {
		canBeat = searchBeat(cards, handToBeat, beat)
		if canBeat {
			handToBeat = beat.Clone()
			handList = handList.Push(beat.Clone())
		}
	}

	return handList
}

// ----------------------------------------------------------------------------
// Hand analyzer
// ----------------------------------------------------------------------------

var primalArray = [...]int{0, HandPrimalSolo, HandPrimalPair, HandPrimalTrio}
var chainLength = [...]int{0, HandSoloChainMinLen, HandPairChainMinLen, HandTrioChainMinLen}

// Extract hands like 34567 / 334455 / 333444555 etc
// array is a processed card array holds count[rank] == duplicate
func extractConsecutive(array CardSlice, duplicate int) HandList {
	hl := HandList{}
	hand := new(Hand)

	if duplicate < 1 || duplicate > 3 || len(array) == 0 {
		return hl
	}

	lastRank := array.RankAt(0)
	i := duplicate

	for cardNum, steps := 0, len(array)/duplicate-1; cardNum < steps; cardNum++ {
		if lastRank-1 != array.RankAt(i) {
			// Chain breaks
			if i >= chainLength[duplicate] {
				// Chain
				hand.kind = HandFormat(primalArray[duplicate], HandKickerNone, HandChain)
				hand.cards, array = array[:i], array[i:]
				hl = hl.Push(hand.Clone())
			} else {
				// Not a chain
				for j := 0; j < i/duplicate; j++ {
					hand.kind = HandFormat(primalArray[duplicate], HandKickerNone, HandChainless)
					hand.cards, array = array[:duplicate], array[duplicate:]
					hl = hl.Push(hand.Clone())
				}
			}

			if len(array) == 0 {
				break
			}

			lastRank = array.RankAt(0)
			i = duplicate
		} else {
			// Chain intact
			lastRank = array.RankAt(i)
			i += duplicate
		}
	}

	// All chained up
	if i != 0 && i == len(array) {
		// Can chain up
		if i >= chainLength[duplicate] {
			hand.kind = HandFormat(primalArray[duplicate], HandKickerNone, HandChain)
			hand.cards, array = array[:i], array[i:]
			hl = hl.Push(hand.Clone())
		} else {
			for j := 0; j < i/duplicate; j++ {
				hand.kind = HandFormat(primalArray[duplicate], HandKickerNone, HandChainless)
				hand.cards, array = array[:duplicate], array[duplicate:]
				hl = hl.Push(hand.Clone())
			}
		}
	}

	return hl
}

// Extract nuke/bomb/2 from array, these cards will be remove from array
func extractNukeBome2(hl HandList, array CardSlice, count []int) (HandList, CardSlice, []int) {
	hand := new(Hand)
	// Nuke
	if count[CardRankr] != 0 && count[CardRankR] != 0 {
		hand.kind = HandFormat(HandPrimalNuke, HandKickerNone, HandChainless)
		hand.cards = CardSlice{CardRedJoker, CardBlackJoker}
		count[CardRankr], count[CardRankR] = 0, 0
		array = array.RemoveRank(CardRankr).RemoveRank(CardRankR)
		hl = hl.Push(hand.Clone())
	}

	// Bomb
	for i := CardRank2; i >= CardRank3; i-- {
		if count[i] == 4 {
			hand.kind = HandFormat(HandPrimalBomb, HandKickerNone, HandChainless)
			hand.cards, array = array.ExtractRank(i)
			count[i] = 0
			hl = hl.Push(hand.Clone())
		}
	}

	// Joker
	if count[CardRankr] != 0 || count[CardRankR] != 0 {
		if count[CardRankr] != 0 {
			hand.cards = CardSlice{CardBlackJoker}
			array = array.RemoveRank(CardRankr)
			count[CardRankr] = 0
		} else {
			hand.cards = CardSlice{CardRedJoker}
			array = array.RemoveRank(CardRankR)
			count[CardRankR] = 0
		}
		hand.kind = HandFormat(HandPrimalSolo, HandKickerNone, HandChainless)
		hl = hl.Push(hand.Clone())
	}

	// 2
	if count[CardRank2] != 0 {
		hand.kind = HandFormat(primalArray[count[CardRank2]], HandKickerNone, HandChainless)
		hand.cards, array = array.ExtractRank(CardRank2)
		count[CardRank2] = 0

		hl = hl.Push(hand.Clone())
	}

	return hl, array, count
}

func (cards CardSlice) StandardAnalyze() HandList {
	chainArrays := []CardSlice{{}, {}, {}}
	array := cards.Clone().Sort()
	count := array.CountRank()
	hl := HandList{}
	// Nuke, bombs and 2
	hl, array, count = extractNukeBome2(hl, array, count)
	// Chains
	for i := 0; i < len(array); {
		c := count[array.RankAt(i)]
		if c != 0 {
			chainArrays[c-1] = chainArrays[c-1].Concat(array[i : i+c])
			i += c
		} else {
			i++
		}
	}

	for i := 3; i > 0; i-- {
		hl = hl.Concat(extractConsecutive(chainArrays[i-1], i))
	}

	return hl
}

// ----------------------------------------------------------------------------
// Hand Evaluator
// ----------------------------------------------------------------------------

// Calculate the length of hand list analyze from a card array
// In the origin C version, the analyze function was rewrite with no memory allocation
// But in go, we choose the easy way
func (array CardSlice) standardEvaluator() int {
	return len(array.StandardAnalyze())
}

// ----------------------------------------------------------------------------
// Advanced Analyze
// ----------------------------------------------------------------------------

func (ctx *handContext) searchLongestConsecutive(duplicate int) *Hand {
	hand := new(Hand)
	cards := ctx.rcards
	count := ctx.count

	// Early break
	if duplicate < 1 || duplicate > 3 || len(cards) < chainLength[duplicate] {
		return hand
	}

	// Setup
	rankChain := []int{}
	rankStart := 0

	// i <= CARD_RANK_2
	// but count[CARD_RANK_2] must be 0
	// for 2/bomb/nuke has been removed before calling this function
	for i := CardRank3; i < CardRank2; i++ {
		// Find start of a possible chain
		if rankStart == 0 && count[i] >= duplicate {
			rankStart = i
			continue
		}

		if count[i] < duplicate {
			// Chain break, extract chain and set a new possible start
			if ((i-rankStart)*duplicate >= chainLength[duplicate]) && (i-rankStart) > len(rankChain) {
				// Valid chain, store rank in card slice
				for j := rankStart; j < i; j++ {
					rankChain = append(rankChain, j)
				}
			}

			rankStart = 0
		}
	}

	// Convert rank chain to card slice
	if len(rankChain) > 0 {
		for i := len(rankChain) - 1; i >= 0; i-- {
			lastRank := rankChain[i]
			k := duplicate

			for j := 0; j < len(rankChain); j++ {
				if cards.RankAt(j) == lastRank {
					hand.cards = hand.cards.Push(cards[j])
					k--

					if k == 0 {
						break
					}
				}
			}
		}

		hand.kind = HandFormat(primalArray[duplicate], HandKickerNone, HandChain)
	}

	return hand
}

func (ctx *handContext) searchPrimal(primal int) *Hand {
	hand := new(Hand)
	count := ctx.count
	rcards := ctx.rcards

	if primal < 1 || primal > 3 {
		return hand
	}

	// Search count[rank] >= primal
	for i := 0; i < len(rcards); i++ {
		if count[rcards.RankAt(i)] >= primal {
			// Found
			hand.kind = HandFormat(primalArray[primal], HandKickerNone, HandChain)
			hand.cards = rcards.Clone()[i : i+primal]
			break
		}
	}

	return hand
}

const (
	handSearchTypes = 3
)

// Pass a empty hand to start traverse
func (ctx *handContext) traverseChains(begin int, hand *Hand) (bool, int) {
	if len(ctx.cards) == 0 || begin >= handSearchTypes {
		return false, 0
	}

	found := false

	// Initialize search
	if hand.kind == HandNone {
		i := begin
		for i < handSearchTypes && hand.kind == 0 {
			tmp := ctx.searchLongestConsecutive(primalArray[i])
			hand.Set(tmp)

			if hand.kind != HandNone {
				found = true
				break
			} else {
				i++
				begin = i
			}
		}
		// If found == false, should PANIC
	} else {
		// Continue search via beat
		found = ctx.cards.SearchBeat(hand, hand)
	}

	return found, begin
}

// Extract all chains or primal hands in hand context
func (ctx *handContext) extractAllChains() HandList {
	handList := HandList{}

	workingHand := new(Hand)
	lastHand := new(Hand)
	lastSearch := 0

	found, lastSearch := ctx.traverseChains(lastSearch, lastHand)
	for found {
		handList.Unshift(lastHand.Clone())
		workingHand.Set(lastHand)
		for {
			found, lastSearch = ctx.traverseChains(lastSearch, workingHand)
			if found {
				handList.Push(workingHand.Clone())
			} else {
				break
			}
		}

		// Can't find any more hands, try to reduce chain length
		if lastHand.kind != HandNone {
			primal := lastHand.Primal()
			if len(lastHand.cards) > chainLength[primal] {
				lastHand.cards = lastHand.cards[1:]
				found = true
			} else {
				lastHand.kind = HandNone
			}

			// Still can't found, loop through hand types for more
			if !found {
				lastSearch++
				lastHand.Clear()
				found, lastSearch = ctx.traverseChains(lastSearch, lastHand)
			}
		}
	}
	return handList
}

type searchPayload struct {
	context *handContext
	hand    *Hand
	weight  int
}

// Advance search tree
type searchTree struct {
	payload *searchPayload

	child   *searchTree
	sibling *searchTree
	parent  *searchTree
}

func newSearchPayload(context *handContext, hand *Hand, weight int) *searchPayload {
	payload := new(searchPayload)
	payload.context = context
	payload.hand = hand
	payload.weight = weight
	return payload
}

func newSearchTree(payload *searchPayload) *searchTree {
	tree := new(searchTree)
	tree.payload = payload
	return tree
}

func (tree *searchTree) addChild(newNode *searchTree) *searchTree {
	newNode.sibling = tree.sibling
	tree.sibling = newNode
	newNode.parent = tree.parent
	return newNode
}

func (tree *searchTree) dumpLeaf() []*searchTree {
	leaf := make([]*searchTree, 0)
	stack := make([]*searchTree, 0)
	stack = append(stack, tree)

	for len(stack) > 0 {
		node, stack := stack[len(stack)-1], stack[:len(stack)-1]
		temp := node.child

		for temp != nil {
			stack = append(stack, temp)
			temp = temp.sibling
		}

		if node.child == nil {
			leaf = append(leaf, node)
		}
	}

	return leaf
}

func (tree *searchTree) addHand(hand *Hand) *searchTree {
	oldPayload := tree.payload
	newPayload := new(searchPayload)

	// Make diff
	newPayload.context = oldPayload.context.clone()
	newPayload.hand = hand.Clone()
	newPayload.context.cards = newPayload.context.cards.Subtract(hand.cards)
	newPayload.context.rcards = newPayload.context.cards.Clone().Reverse()
	newPayload.context.count = newPayload.context.cards.CountRank()
	newPayload.weight = oldPayload.weight + 1

	// Tree expansion
	return tree.addChild(newSearchTree(newPayload))
}

// Search hand via shortest hand list
func (array CardSlice) AdvanceAnalyze() HandList {
    var shortest *searchTree = nil

	handList := make(HandList, 0)

	// Setup search context
	ctx := newHandContext(array)
	// Extract bombs and 2
	handList, ctx.cards, ctx.count = extractNukeBome2(handList, ctx.cards, ctx.count)
	// Finish building beat search context
	ctx.rcards = ctx.cards.Clone().Reverse()

	// Magic goes here

	// Root
	payload := newSearchPayload(ctx, nil, 0)
	grandTree := newSearchTree(payload)

	// First expansion
	chains := ctx.extractAllChains()

	// No chains, fallback to standard analyze
	if len(chains) == 0 {
		return array.StandardAnalyze()
	}

	// Got chains, make first expand
	stack := make([]*searchTree, 0)
	for i := 0; i < len(chains); i++ {
		treeNode := grandTree.addHand(chains[i])
		stack = append(stack, treeNode)
	}

	// Loop start
	for len(stack) != 0 {
		// Pop stack
		workingTree := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// Expansion
		chains = workingTree.payload.context.extractAllChains()
		if len(chains) != 0 {
			// Push new nodes
			for i := 0; i < len(chains); i++ {
				treeNode := workingTree.addHand(chains[i])
				stack = append(stack, treeNode)
			}
		}
	}

	// Tree construction complete
	leaves := grandTree.dumpLeaf()

	// Find shortest path
	for len(leaves) != 0 {
		// Pop stack
		workingTree := leaves[len(leaves)-1]
		leaves = leaves[:len(leaves)-1]
        payload = workingTree.payload
        // Calculate other hands weight
        payload.weight += payload.context.cards.standardEvaluator()

        if shortest == nil || payload.weight < shortest.payload.weight {
            shortest = workingTree
        }
	}

    // Extract shortest  node's other hands
    others := shortest.payload.context.cards.StandardAnalyze()

    for shortest != nil && shortest.payload.weight != 0 {
        others = others.Unshift(shortest.payload.hand.Clone())
        shortest = shortest.parent
    }

    others = others.Concat(handList)

	return others
}

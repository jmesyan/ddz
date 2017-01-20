package ddz

import (
    "sort"
)

type Hand struct {
    kind  uint8
    cards CardSlice
}

func HandGetPrimal(h uint8) uint8 {
    return h & 0x0F
}

func HandGetKicker(h uint8) uint8 {
    return h & 0x70
}

func HandGetChain(h uint8) uint8 {
    return h & 0x80
}

func HandSetMask(h *uint8, m uint8) uint8 {
    *h |= m
    return *h
}

func HandFormat(primal, kicker, chain uint8) uint8 {
    return primal | kicker | chain
}

func HandClearPrimal(h *uint8) uint8 {
    *h &= 0xF0
    return *h
}

func HandClearKicker(h *uint8) uint8 {
    *h &= 0x8F
    return *h
}

func HandClearChain(h *uint8) uint8 {
    *h &= 0x7F
    return *h
}

func (h *Hand) Primal() uint8 {
    return HandGetPrimal(h.kind)
}

func (h *Hand) SetPrimal(primal uint8) *Hand {
    HandSetMask(&h.kind, primal)
    return h
}

func (h *Hand) Kicker() uint8 {
    return HandGetKicker(h.kind)
}

func (h *Hand) SetKicker(kicker uint8) *Hand {
    HandSetMask(&h.kind, kicker)
    return h
}

func (h *Hand) Chain() uint8 {
    return HandGetChain(h.kind)
}

func (h *Hand) SetChain(chain uint8) *Hand {
    HandSetMask(&h.kind, chain)
    return h
}

func (h *Hand) Kind() uint8 {
    return h.kind
}

func (h *Hand) SetKind(kind uint8) *Hand {
    h.kind = kind
    return h
}

func (h *Hand) Format(primal, kicker, chain uint8) *Hand {
    h.kind = HandFormat(primal, kicker, chain)
    return h
}

func (h *Hand) SetCards(cards CardSlice) *Hand {
    h.cards = cards.Clone()
    return h
}

func (h *Hand) Clear() *Hand {
    h.kind = HandNone
    h.cards = make(CardSlice, 0)
    return h
}

func (h *Hand) Clone() *Hand {
    cpy := new(Hand)
    cpy.kind = h.kind
    cpy.cards = h.cards.Clone()
    return cpy
}

func (h *Hand) Set(o *Hand) *Hand {
    h.kind = o.kind
    h.cards = o.cards.Clone()
    return h
}

func (h *Hand) CountRank() ([]int, []int) {
    count := h.cards.CountRank()
    sorted := make([]int, len(count))
    copy(sorted, count)
    sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
    return count, sorted
}

// private

const (
    handVariation = 2
    handSpec = 4

    patternLen = 12

    handPatternNone = 0  // place holder
    handPattern_1 = 1  // 1, solo
    handPattern_2_1 = 2  // 2, pair
    handPattern_2_2 = 3  // 2, nuke
    handPattern_3 = 4  // 3, trio
    handPattern_4_1 = 5  // bomb
    handPattern_4_2 = 6  // trio solo
    handPattern_5_1 = 7  // solo chain
    handPattern_5_2 = 8  // trio pair
    handPattern_6_1 = 9  // solo chain
    handPattern_6_2 = 10 // pair chain
    handPattern_6_3 = 11 // trio chain
    handPattern_6_4 = 12 // four dual solo
    handPattern_7 = 13 // solo chain
    handPattern_8_1 = 14 // solo chain
    handPattern_8_2 = 15 // pair chain
    handPattern_8_3 = 16 // trio solo chain
    handPattern_8_4 = 17 // four dual pair
    handPattern_8_5 = 18 // four chain
    handPattern_9_1 = 19 // solo chain
    handPattern_9_2 = 20 // trio chain
    handPattern_10_1 = 21 // solo chain
    handPattern_10_2 = 22 // pair chain
    handPattern_10_3 = 23 // trio pair chain
    handPattern_11 = 24 // solo chain
    handPattern_12_1 = 25 // solo chain
    handPattern_12_2 = 26 // pair chain
    handPattern_12_3 = 27 // trio chain
    handPattern_12_4 = 28 // trio solo chain
    handPattern_12_5 = 29 // four chain
    handPattern_12_6 = 30 // four dual solo chain
    handPattern_14 = 31 // pair chain
    handPattern_15 = 32 // trio chain
    handPattern_16_1 = 33 // pair chain
    handPattern_16_2 = 34 // trio solo chain
    handPattern_16_3 = 35 // four chain
    handPattern_16_4 = 36 // four dual pair chain
    handPattern_18_1 = 37 // pair chain
    handPattern_18_2 = 38 // trio chain
    handPattern_18_3 = 39 // four dual solo chain
    handPattern_20_1 = 40 // pair chain
    handPattern_20_2 = 41 // trio solo chain
    handPattern_20_3 = 42 // four chain
    handPatternEnd = handPattern_20_3
)

var handPatterns = [][]int{
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

var handSpecs = [][][]int{
    {
        // place holder
        {0, 0, 0, 0}, {0, 0, 0, 0},
    }, {
        // 1
        {handPattern_1, HandPrimalSolo, 0, 0},
        {0, 0, 0, 0},
    }, {
        // 2
        {handPattern_2_1, HandPrimalPair, 0, 0},
        {0, 0, 0, 0},
    }, {
        // 3
        {handPattern_3, HandPrimalTrio, 0, 0},
        {0, 0, 0, 0},
    }, {
        // 4
        {handPattern_4_2, HandPrimalTrio, HandKickerSolo, 0},
        {0, 0, 0, 0},
    }, {
        // 5
        {handPattern_5_2, HandPrimalTrio, HandKickerPair, 0},
        {0, 0, 0, 0},
    }, {
        // 6
        {handPattern_6_4, HandPrimalFour, HandKickerDualSolo, 0},
        {0, 0, 0, 0},
    }, {
        // 7
        {0, 0, 0, 0}, {0, 0, 0, 0},
    }, {
        // 8
        {handPattern_8_3, HandPrimalTrio, HandKickerSolo, HandChain},
        {handPattern_8_4, HandPrimalFour, HandKickerDualPair, 0},
    }, {
        // 9
        {0, 0, 0, 0}, {0, 0, 0, 0},
    }, {
        // 10
        {handPattern_10_3, HandPrimalTrio, HandKickerPair, HandChain},
        {0, 0, 0, 0},
    }, {
        // 11
        {0, 0, 0, 0}, {0, 0, 0, 0},
    }, {
        // 12
        {handPattern_12_4, HandPrimalTrio, HandKickerSolo, HandChain},
        {handPattern_12_6, HandPrimalFour, HandKickerDualSolo, HandChain},
    }, {
        // 13
        {0, 0, 0, 0}, {0, 0, 0, 0},
    }, {
        // 14
        {0, 0, 0, 0}, {0, 0, 0, 0},
    }, {
        // 15
        {0, 0, 0, 0}, {0, 0, 0, 0},
    }, {
        // 16
        {handPattern_16_2, HandPrimalTrio, HandKickerSolo, HandChain},
        {handPattern_16_4, HandPrimalFour, HandKickerDualPair, HandChain},
    }, {
        // 17
        {0, 0, 0, 0}, {0, 0, 0, 0},
    }, {
        // 18
        {handPattern_18_3, HandPrimalFour, HandKickerDualSolo, HandChain},
        {0, 0, 0, 0},
    }, {
        // 19
        {0, 0, 0, 0}, {0, 0, 0, 0},
    }, {
        // 20
        {handPattern_20_2, HandPrimalTrio, HandKickerSolo, HandChain},
        {0, 0, 0, 0},
    },
};

func compareSlice(a, b []int) bool {
    if a == nil || b == nil {
        return false
    }

    if len(a) != len(b) {
        return false
    }

    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

func patterMatch(sorted []int, pattern int) bool {
    return compareSlice(sorted, handPatterns[pattern])
}

func checkChain(count []int, duplicate, expectLen int) bool {
    marker, length := 0, 0
    // joker and 2 can't chain up
    for i := CardRank3; i < CardRank2; i++ {
        if count[i] == duplicate && marker == 0 {
            marker = i
            continue
        }

        // matches end
        if count[i] != duplicate && marker != 0 {
            length = i - marker
            break
        }
    }
    return length == expectLen
}

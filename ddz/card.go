package ddz

import (
    "strings"
    "sort"
    "math/rand"
)

type CardSlice []uint8

func CardRank(card uint8) uint8 {
    return card & 0x0F
}

func CardSuit(card uint8) uint8 {
    return card & 0xF0
}

func Card(suit uint8, rank uint8) uint8 {
    return suit | rank
}

func MakeCardSet() CardSlice {
    slice := CardSlice{
        0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D,
        0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D,
        0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D,
        0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D,
        0x1E, 0x2F}

    return slice
}

func CardToStr(card uint8) string {
    var str string
    suit := CardSuit(card)
    rank := CardRank(card)

    switch suit {
    case CardSuitSpade:
        str += "\u2660"
    case CardSuitClub:
        str += "\u2663"
    case CardSuitHeart:
        str += "\u2665"
    case CardSuitDiamond:
        str += "\u2666"
    }

    switch {
    case rank == CardRankT:
        str += "T"
    case rank == CardRankJ:
        str += "J"
    case rank == CardRankQ:
        str += "Q"
    case rank == CardRankK:
        str += "K"
    case rank == CardRankA:
        str += "A"
    case rank == CardRank2:
        str += "2"
    case rank == CardRankr:
        str += "r"
    case rank == CardRankR:
        str += "R"
    case rank <= CardRank9 && rank >= CardRank3:
        str += string(uint8('3') + rank - CardRank3)
    }

    return str
}

func CardSliceFromString(str string) CardSlice {
    slice := make(CardSlice, 0)
    var suit, rank uint8
    for _, v := range (str) {
        switch {
        case v == 9824:
            suit = CardSuitSpade
        case v == 9827:
            suit = CardSuitClub
        case v == 9829:
            suit = CardSuitHeart
        case v == 9830:
            suit = CardSuitDiamond
        case v == 'T':
            rank = CardRankT
        case v == 'J':
            rank = CardRankJ
        case v == 'Q':
            rank = CardRankQ
        case v == 'K':
            rank = CardRankK
        case v == 'A':
            rank = CardRankA
        case v == '2':
            rank = CardRank2
        case v == 'r':
            rank = CardRankr
        case v == 'R':
            rank = CardRankR
        case v <= '9' && v >= '3':
            rank = CardRank3 + uint8(v - '3')
        }

        if suit != 0 && rank != 0 {
            slice = append(slice, Card(suit, rank))
            suit, rank = 0, 0
        }
    }
    return slice
}

// stringify with white space
func (cs CardSlice) ToString() string {
    return cs.ToString2(" ")
}

// stringify with separator
func (cs CardSlice) ToString2(sep string) string {
    cards := make([]string, 0)
    for _, v := range (cs) {
        cards = append(cards, CardToStr(v))
    }

    if len(sep) == 0 {
        return strings.Join(cards, " ")
    } else {
        return strings.Join(cards, sep)
    }
}

// sort len
func (cs CardSlice) Len() int {
    return len(cs)
}

// sort swap
func (cs CardSlice) Swap(i, j int) {
    cs[i], cs[j] = cs[j], cs[i]
}

// sort less
func (cs CardSlice) Less(i, j int) bool {
    var ra, rb uint8

    // rotation
    ra = (cs[i] & 0xF0) >> 4 | (cs[i] & 0x0F) << 4
    rb = (cs[j] & 0xF0) >> 4 | (cs[j] & 0x0F) << 4

    return rb < ra
}

// sort
func (cs CardSlice) Sort() CardSlice {
    sort.Sort(cs)
    return cs
}

func (cs CardSlice) Clone() CardSlice {
    cpy := make(CardSlice, len(cs))
    copy(cpy, cs)
    return cpy
}

func (cs CardSlice) Reverse() CardSlice {
    for i, j := 0, len(cs) - 1; i < j; i, j = i + 1, j - 1 {
        cs[i], cs[j] = cs[j], cs[i]
    }

    return cs
}

func (cs CardSlice) Concat(s CardSlice) CardSlice {
    return append(cs, s...)
}

func (cs CardSlice) Subtract(sub CardSlice) CardSlice {
    if sub == nil {
        return cs
    }

    diff := make(CardSlice, 0)
    for i := 0; i < len(cs); i++ {
        card := cs[i]
        if sub.Contains(card) {
            card = 0
        }

        if card != 0 {
            diff = append(diff, card)
        }
    }

    return diff
}

func (cs CardSlice) Contains(c uint8) bool {
    for i := 0; i < len(cs); i++ {
        if c == cs[i] {
            return true
        }
    }

    return false
}

func (cs CardSlice) ContainsAll(s CardSlice) bool {
    if s == nil || len(s) == 0 {
        return true
    }

    for i := 0; i < len(s); i++ {
        if !cs.Contains(s[i]) {
            return false
        }
    }

    return true
}

func (cs CardSlice) Equals(s CardSlice) bool {
    if s == nil || len(cs) != len(s) {
        return false
    }

    a := cs.Sort()
    b := s.Sort()

    for i := 0; i < len(cs); i++ {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

func (cs CardSlice) Push(card uint8) CardSlice {
    return append(cs, card)
}

func (cs CardSlice) Pop() (uint8, CardSlice) {
    if len(cs) == 0 {
        return 0, cs
    }

    return cs[len(cs) - 1], cs[:len(cs) - 1]
}

func (cs CardSlice) Shift(card uint8) (uint8, CardSlice) {
    if len(cs) == 0 {
        return 0, cs
    }

    return cs[0], cs[1:]
}

func (cs CardSlice) Unshift(card uint8) CardSlice {
    return append(CardSlice{card}, cs...)
}

func (cs CardSlice) DropFront(count int) CardSlice {
    if count >= len(cs) {
        count = len(cs)
    }
    return cs[count:]
}

func (cs CardSlice) DropBack(count int) CardSlice {
    if count >= len(cs) {
        count = len(cs)
    }
    return cs[:len(cs) - count]
}

func (cs CardSlice) Insert(i int, card uint8) CardSlice {
    if i < 0 || i >= len(cs) {
        return cs
    }

    ret := append(cs, 0)
    copy(ret[i + 1:], cs[i:])
    ret[i] = card
    return ret
}

func (cs CardSlice) Remove(i int) (uint8, CardSlice) {
    if i < 0 || i >= len(cs) {
        return 0, cs
    }

    card := cs[i]
    ret := cs.Clone()
    return card, append(ret[:i], ret[i + 1:]...)
}

func (cs CardSlice) RemoveCard(card uint8) (bool, CardSlice) {
    for i := 0; i < len(cs); i++ {
        if cs[i] == card {
            _, s := cs.Remove(i)
            return true, s
        }
    }

    return false, cs
}

func (cs CardSlice) ExtractRank(rank uint8) CardSlice {
    ret := make(CardSlice, 0)
    for i := 0; i < len(cs); i++ {
        if CardRank(cs[i]) == rank {
            ret = append(ret, cs[i])
        }
    }

    return ret
}

func (cs CardSlice) RemoveRank(rank uint8) CardSlice {
    left := make(CardSlice, 0)
    for i := 0; i < len(cs); i++ {
        if CardRank(cs[i]) != rank {
            left = append(left, cs[i])
        }
    }

    return left
}

func (cs CardSlice) CountRank() []int {
    count := make([]int, CardRankEnd)
    for i := 0; i < len(cs); i++ {
        count[CardRank(cs[i])]++
    }

    return count
}

func (cs CardSlice) CountSortRank() ([]int, []int) {
    count := cs.CountRank()
    sorted := make([]int, len(count))
    copy(sorted, count)
    sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
    return count, sorted
}

func (cs CardSlice) Shuffle(seed int64) CardSlice {
    rand.Seed(seed)
    for i := range cs {
        j := rand.Intn(i + 1)
        cs[i], cs[j] = cs[j], cs[i]
    }

    return cs
}

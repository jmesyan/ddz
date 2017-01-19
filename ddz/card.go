package ddz

import (
    "strings"
    "sort"
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
    slice := make(CardSlice, 0);
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

func (cs CardSlice) ToString() string {
    return cs.ToString2(" ")
}

func (cs CardSlice) ToString2(sep string) string {
    cards := make([]string, 0)
    for _, v := range(cs) {
         cards = append(cards, CardToStr(v))
    }

    if len(sep) == 0 {
        return strings.Join(cards, " ")
    } else {
        return strings.Join(cards, sep)
    }
}

func (cs CardSlice) Len() int {
    return len(cs)
}

func (cs CardSlice) Swap(i, j int) {
    cs[i], cs[j] = cs[j], cs[i]
}

func (cs CardSlice) Less(i, j int) bool {
    var ra, rb uint8

    // rotation
    ra = (cs[i] & 0xF0) >> 4 | (cs[i] & 0x0F) << 4
    rb = (cs[j] & 0xF0) >> 4 | (cs[j] & 0x0F) << 4

    return rb < ra
}

func (cs CardSlice) Sort() {
    sort.Sort(cs)
}

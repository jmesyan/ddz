package golandlord

const (
    CardRank3 = 0x01
    CardRank4 = 0x02
    CardRank5 = 0x03
    CardRank6 = 0x04
    CardRank7 = 0x05
    CardRank8 = 0x06
    CardRank9 = 0x07
    CardRankT = 0x08
    CardRankJ = 0x09
    CardRankQ = 0x0A
    CardRankK = 0x0B
    CardRankA = 0x0C
    CardRank2 = 0x0D
    CardRankr = 0x0E
    CardRankR = 0x0F

    CardRankBeg = CardRank3
    CardRankEnd = CardRankR

    CardSuitClub = 0x10
    CardSuitDiamond = 0x20
    CardSuitHeart = 0x30
    CardSuitSpade = 0x40

    CardSetLength = 54
)

func CardRank(card uint8) {
    return card & 0x0F
}

func CardSuit(card uint8) {
    return card & 0xF0
}

func Card(suit uint8, rank uint8) {
    return suit | rank
}

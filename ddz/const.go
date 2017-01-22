package ddz

const (
    // card
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
    CardRankEnd = CardRankR + 1

    CardSuitClub = 0x10
    CardSuitDiamond = 0x20
    CardSuitHeart = 0x30
    CardSuitSpade = 0x40

    CardSetLen = 54

    // hand
    HandMinLen = 1
    HandMaxLen = 20
    HandSoloChainMinLen = 5
    HandPairChainMinLen = 6
    HandTrioChainMinLen = 6
    HandFourChainMinLen = 8

    HandPrimalNone = 0x00
    HandPrimalNuke = 0x06
    HandPrimalBomb = 0x05
    HandPrimalFour = 0x04
    HandPrimalTrio = 0x03
    HandPrimalPair = 0x02
    HandPrimalSolo = 0x01

    HandKickerNone = 0x00
    HandKickerSolo = 0x10
    HandKickerPair = 0x20
    HandKickerDualSolo = 0x30
    HandKickerDualPair = 0x40

    HandChainless = 0x00
    HandChain = 0x80
    HandNone = 0
    HandSearchMask = 0xFF

    HandCompareIllegal = -3
    HandCompareLess = -1
    HandCompareEqual = 0
    HandCompareGreater = 1
)

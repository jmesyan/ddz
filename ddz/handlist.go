package ddz

type HandList []*Hand

func (hl HandList) Push(h *Hand) HandList {
    return append(hl, h)
}

func (hl HandList) Pop() (*Hand, HandList) {
    if len(hl) == 0 {
        return nil, hl
    }
    return hl[len(hl) - 1], hl[:len(hl) - 1]
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
    for i:= 0; i < len(hl); i++ {
        if hl[i] == h {
            return append(hl[:i], hl[i+1:]...)
        }
    }

    return hl
}

func (hl HandList) FindKind(kind uint8) *Hand {
    for i:= 0; i < len(hl); i++ {
        if hl[i].kind == kind {
            return hl[i]
        }
    }

    return nil
}

type handContext struct {
    count []int
    cards CardSlice
    rcards CardSlice
}

func newHandContext(array CardSlice) *handContext {
    hctx := new(handContext)
    hctx.count = array.CountRank()
    hctx.cards = array.Clone()
    hctx.rcards = array.Clone().Reverse()

    return hctx
}

func (ctx *handContext) searchBeatPrimal(tobeat, beat *Hand, primal uint8) bool {
    rank := CardRank(tobeat.cards[0])
    count := ctx.count
    temp := ctx.rcards
    for i:= 0; i < len(temp); i++ {
        if CardRank(temp[i]) > rank && count[CardRank(temp[i])] >= int(primal) {
            beat.kind = tobeat.kind
            beat.cards = temp[i:i+int(primal)]
            return true
        }
    }

    return false
}

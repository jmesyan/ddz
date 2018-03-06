package ddz

// HandType enum
type HandType int

// Hand is a valid card set that can play.
// cards format must be like 12345/112233/1112223344/11122234 etc
type Hand struct {
	Type  HandType  // hand type
	Cards CardSlice // cards
}

// Copy returns a copy of hand
func (h Hand) Copy() Hand {
	return Hand{
		Type:  h.Type,
		Cards: h.Cards.Copy(),
	}
}

package ddz

type PlayerIdentity int

const (
	PlayerIdentityPeasant  PlayerIdentity = 0
	PlayerIdentityLandlord PlayerIdentity = 1
)

type BidAction int

const (
	BidAbstain BidAction = 0
	BidBid     BidAction = 1
)

type Status int

const (
	StatusHalt  Status = 0
	StatusBid   Status = 1
	StatusReady Status = 2
	StatusPause Status = 3
	StatusOver  Status = 4
)

type Phase int

const (
	PhasePlay  Phase = 0
	PhaseQuery Phase = 1
	PhasePass  Phase = 2
)

type Game interface {
	GetNextPlayer() Player
	LastHand() *Hand
	KittyCards() CardSlice
	HighestBidder() Player
	Landlord() Player
	Status() Status
	Phase() Phase
}

type Player interface {
	GetReady(g *Game)
	Bid(g *Game) BidAction
	Start(g *Game)
	Play(g *Game) *Hand
	Beat(g *Game) *Hand
}

type SimpleGame struct {
	Game
	players       []*Player // players
	deck          CardSlice // deck
	lastHand      *Hand     // last hand played
	kittyCards    CardSlice // kitty cards
	bid           int       // current bid
	highestBidder *Player   // the highest bidder
	landlord      *Player   // landlord player
	winner        *Player   // player wins last game
	status        Status    // game status
	phase         Phase     // game phase
}

type SimplePlayer struct {
	Player
	cards     CardSlice      // card slice, will change during game play
	record    CardSlice      // card record
	handList  []*Hand        // analyze result of card slice
	identity  PlayerIdentity // player identity
	seatIndex int
	bid       int
}

func (p *SimplePlayer) GetReady(g *Game) {

}

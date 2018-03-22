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

const ()

type Game struct {
	players       []*Player // players
	deck          CardSlice // deck
	lastHand      *Hand     // last hand played
	kittyCards    CardSlice // kitty cards
	bid           int       // current bid
	highestBidder *Player   // the highest bidder
	landlord      *Player   // landlord player
	winner        *Player   // player wins last game
	status        int
}

type Player interface {
	GetReady()
	Bid()
	Start()
	Play()
	Beat()
}

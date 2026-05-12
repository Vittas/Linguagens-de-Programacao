package function

type User struct {
	CardName                string
	CardNo                  string
	CardStatus              string
	CardType                string
	CitizenIDNo             string
	Doors                   []string
	DynamicCheckCode        string
	FirstEnter              bool
	Handicap                bool
	IsValid                 bool
	Password                string
	RecNo                   int
	RepeatEnterRouteTimeout uint64
	TimeSections            []int
	UseTime                 int
	UserID                  int
	UserType                int
	VTOPosition             string
	ValidDateEnd            string
	ValidDateStart          string
}

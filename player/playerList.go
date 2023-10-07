package player

type PlayerList struct {
	Players map[string]*Player
}

var players = make(map[string]*Player, 0)

func init() {
	players = LoadPlayerList()
}

func LoadPlayerList() map[string]*Player {
	//TODO load the already existing playerlist
	return make(map[string]*Player, 0)
}

func RegisterPlayer(ownerId string, player *Player) bool {
	if players[ownerId] != nil {
		return false
	}
	players[ownerId] = player
	return true
}

func GetPlayerByOwnerId(ownerId string) *Player {
	return players[ownerId]
}

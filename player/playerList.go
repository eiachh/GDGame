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

func RegisterPlayer(ownerId string, player *Player) (bool, string) {
	if players[ownerId] != nil {
		return false, "You already have a character."
	}
	for existingPlayerName := range players {
		if existingPlayerName == player.PlayerName {
			return false, "Character name is taken!"
		}
	}
	players[ownerId] = player
	return true, "Successfully registered as: " + player.PlayerName
}

func GetPlayerByOwnerId(ownerId string) *Player {
	return players[ownerId]
}

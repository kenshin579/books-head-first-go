package gadget

func playList(device TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func Example_playList() {
	player := TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	playList(player, mixtape)

	//Output:
	//Playing Jessie's Girl
	//Playing Whip it
	//Playing 9 to 5
	//Stopped!

}

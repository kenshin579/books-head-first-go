package main

import "github.com/headfirstgo/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func Example_player() {
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	//Output:
	//Playing Jessie's Girl
	//Playing Whip it
	//Playing 9 to 5
	//Stopped!
	//Playing Jessie's Girl
	//Playing Whip it
	//Playing 9 to 5
	//Stopped!
}

func Example_tryout() {
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})

	//Output:
	//Playing Test Track
	//Stopped!
	//Recording
	//Playing Test Track
	//Stopped!
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader   = bufio.NewReader(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	P        int
	M        int
	RoomList []Room
)

type User struct {
	Level int
	Name  string
}

type Room struct {
	BenchLevel int
	Users      *[]User
}

func (r *Room) IsStarted() bool {
	if len(*r.Users) == M {
		return true
	} else {
		return false
	}
}

func (r *Room) CanEnter(level int) bool {
	if r.BenchLevel-10 <= level && level <= r.BenchLevel+10 && len(*r.Users) < M{
		return true
	} else {
		return false
	}
}

func EnterRoom(level int, name string) {
	user := User{
		Level: level,
		Name:  name,
	}

	for _, room := range RoomList {
		if room.CanEnter(level) {
			*room.Users = append(*room.Users, user)
			return
		}
	}

	RoomList = append(RoomList, Room{
		BenchLevel: level,
		Users:      &[]User{user},
	})
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &P, &M)
	for i := 0; i < P; i++ {
		var l int
		var n string
		fmt.Fscan(reader, &l, &n)

		EnterRoom(l, n)
	}

	for _, room := range RoomList {
		if room.IsStarted() {
			fmt.Fprintln(writer, "Started!")
		} else {
			fmt.Fprintln(writer, "Waiting!")
		}

		sort.Slice(*room.Users, func(i, j int) bool {
			return (*room.Users)[i].Name < (*room.Users)[j].Name
		})
		for _, user := range *room.Users {
			
			fmt.Fprintln(writer, user.Level, user.Name)
		}
	}
}

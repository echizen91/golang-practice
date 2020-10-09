package linkedlist

import (
	"fmt"
	"strings"
)

/*
	Credits: https://medium.com/backendarmy/linked-lists-in-go-f7a7b27a03b9
	[Singly Linked List Practice]
	We are going to build the music playlist now.
	We know that the music playlist will have a sequence of songs and we can listen to the next song by clicking the next track button.

	There are primarily two entities here. The playlist itself and the songs in the playlist. We need to define them in our code.

	Our song type has the fields
	***
		name,
		artist,
		next — a link to the next song.
	***

	Fields in our playlist are the
	***
		name,
		start, — a pointer to the head of our linked list or first song in the playlist
		nowPlaying - which points to currently playing song.
	***

	We need five main operations here to implement a basic playlist.
	1. Create a playlist
	2. Add song
	3. Show all songs
	4. Start playing
	5. Next/Previous Song
*/

// Song : structure for song
type Song struct {
	name   string
	artist string
	next   *Song
}

// Playlist : collection of songs
type Playlist struct {
	name       string
	start      *Song
	nowPlaying *Song
}

// CreatePlaylist : creation of playlist (assume creating a playlist with no songs)
func CreatePlaylist(name string) *Playlist {
	p := &Playlist{
		name: name,
	}
	return p
}

// AddSong : add a song to the playlist
func (p *Playlist) AddSong(name string, artist string) error {
	// Create song
	newSong := &Song{
		name:   name,
		artist: artist,
	}

	// Check if playlist is nil, assign to start
	if p.start == nil {
		p.start = newSong
	} else {
		// Else check for the last link in start, and add song there
		holder := p.start
		for holder.next != nil {
			holder = holder.next
		}
		holder.next = newSong
	}

	return nil
}

// ShowAllSongs : print out all songs
func (p *Playlist) ShowAllSongs() []string {
	var result []string

	// Check if playlist is nil
	if p.start != nil {
		// Counter for #
		counter := 1

		// Declare head
		holder := p.start

		// assign first song
		result = append(result, fmt.Sprintf("#%v %v by %v", counter, holder.name, holder.artist))
		counter++

		for holder.next != nil {
			holder = holder.next
			result = append(result, fmt.Sprintf("#%v %v by %v", counter, holder.name, holder.artist))
			counter++
		}
	}

	return result
}

// StartPlaying : start playing song
func (p *Playlist) StartPlaying() {
	// If nowPlaying is not nil -> start playing
	if p.nowPlaying == nil {
		p.nowPlaying = p.start
	} else { // If nowPlaying is active -> stop playing
		p.nowPlaying = nil
	}
}

// NextPrevSong : Next song or Previous song selector
func (p *Playlist) NextPrevSong(cursor string) {
	if p.nowPlaying != nil {
		// next -> go next
		if strings.Contains(`next`, strings.ToLower(cursor)) {
			p.nowPlaying = p.nowPlaying.next
		}
		// previous -> go previous
		if strings.Contains(`previous`, strings.ToLower(cursor)) {
			holder := p.start
			if holder.name != p.nowPlaying.name {
				x := true
				for x {
					if holder.next.name == p.nowPlaying.name {
						x = false
					} else {
						holder = holder.next
					}
				}
			}
			p.nowPlaying = holder
		}
	}
}

package linkedlist

import (
	"reflect"
	"testing"
)

func TestPlaylist_CreatePlaylist(t *testing.T) {
	p := CreatePlaylist("September")
	compareP := &Playlist{
		name: "September",
	}

	if reflect.DeepEqual(p, compareP) != true {
		t.Errorf("Create Playlist is Incorrect: %v ; wants: %v", p, compareP)
	}
}

func TestPlaylist_AddSong(t *testing.T) {
	p := &Playlist{
		name: "September",
	}

	compareP := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
		},
	}

	p.AddSong("Bleeding Love", "Leona Lewis")

	{
		if reflect.DeepEqual(p.name, compareP.name) != true {
			t.Errorf("Add Song is Incorrect: %v ; wants: %v", p.name, compareP.name)
		}

		if reflect.DeepEqual(p.start, compareP.start) != true {
			t.Errorf("Add Song is Incorrect: %v ; wants: %v", p.start, compareP.start)
		}

		if reflect.DeepEqual(p.start, compareP.start) != true {
			t.Errorf("Add Song is Incorrect: %v ; wants: %v", p.start, compareP.start)
		}

		if reflect.DeepEqual(p.nowPlaying, compareP.nowPlaying) != true {
			t.Errorf("Add Song is Incorrect: %v ; wants: %v", p.nowPlaying, compareP.nowPlaying)
		}
	}

	p2 := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Billie Jeans",
				artist: "Michael Jackson",
			},
		},
	}

	compareP2 := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Billie Jeans",
				artist: "Michael Jackson",
				next: &Song{
					name:   "Perfect",
					artist: "Simple Plan",
				},
			},
		},
	}

	p2.AddSong("Perfect", "Simple Plan")

	{
		if reflect.DeepEqual(p2.name, compareP2.name) != true {
			t.Errorf("Add Song is Incorrect: %v ; wants: %v", p2.name, compareP2.name)
		}
		if reflect.DeepEqual(p2.start.name, compareP2.start.name) != true {
			t.Errorf("Add Song is Incorrect: %v ; wants: %v", p2.start.name, compareP2.start.name)
		}
		if reflect.DeepEqual(p2.nowPlaying, compareP2.nowPlaying) != true {
			t.Errorf("Add Song is Incorrect: %v ; wants: %v", p2.nowPlaying, compareP2.nowPlaying)
		}

		if reflect.DeepEqual(p2.start.next, compareP2.start.next) != true {
			t.Errorf("Add Song is Incorrect: %v ; wants: %v", p2.start.next, compareP2.start.next)
		}

	}

}

func TestPlaylist_ShowAllSongs(t *testing.T) {
	p := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Perfect",
				artist: "Simple Plan",
			},
		},
	}

	compareP := []string{
		"#1 Bleeding Love by Leona Lewis",
		"#2 Perfect by Simple Plan",
	}

	result := p.ShowAllSongs()

	if reflect.DeepEqual(result, compareP) != true {
		t.Errorf("Show All Songs Incorrect: %v ; wants: %v", result, compareP)
	}

	p2 := &Playlist{
		name: "September",
	}

	var compareP2 []string
	result2 := p2.ShowAllSongs()

	if reflect.DeepEqual(result2, compareP2) != true {
		t.Errorf("Show All Songs Incorrect: %v ; wants: %v", result2, compareP2)
	}

}

func TestPlaylist_StartPlaying(t *testing.T) {
	p := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Perfect",
				artist: "Simple Plan",
			},
		},
	}

	p.StartPlaying()

	compareP := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Perfect",
				artist: "Simple Plan",
			},
		},
		nowPlaying: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Perfect",
				artist: "Simple Plan",
			},
		},
	}

	if reflect.DeepEqual(p.nowPlaying, compareP.nowPlaying) != true {
		t.Errorf("Start Playing Incorrect: %v ; wants: %v", p.nowPlaying, compareP.nowPlaying)
	}

	p2 := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Perfect",
				artist: "Simple Plan",
			},
		},
		nowPlaying: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Perfect",
				artist: "Simple Plan",
			},
		},
	}

	compareP2 := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Perfect",
				artist: "Simple Plan",
			},
		},
	}

	p2.StartPlaying()

	if reflect.DeepEqual(p2.nowPlaying, compareP2.nowPlaying) != true {
		t.Errorf("Start Playing Incorrect: %v ; wants: %v", p2.nowPlaying, compareP2.nowPlaying)
	}
}

func TestPlaylist_NextPrevSong(t *testing.T) {
	p := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Billie Jeans",
				artist: "Michael Jackson",
				next: &Song{
					name:   "Perfect",
					artist: "Simple Plan",
				},
			},
		},
		nowPlaying: &Song{
			name:   "Billie Jeans",
			artist: "Michael Jackson",
			next: &Song{
				name:   "Perfect",
				artist: "Simple Plan",
			},
		},
	}

	// compare for next
	compareP := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Billie Jeans",
				artist: "Michael Jackson",
				next: &Song{
					name:   "Perfect",
					artist: "Simple Plan",
				},
			},
		},
		nowPlaying: &Song{
			name:   "Perfect",
			artist: "Simple Plan",
		},
	}

	p.NextPrevSong("next")

	if reflect.DeepEqual(p.nowPlaying.name, compareP.nowPlaying.name) != true {
		t.Errorf("Next Previous Song Incorrect: %v ; wants: %v", p.nowPlaying.name, compareP.nowPlaying.name)
	}

	// Compare for prev
	compareP2 := &Playlist{
		name: "September",
		start: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
			next: &Song{
				name:   "Billie Jeans",
				artist: "Michael Jackson",
				next: &Song{
					name:   "Perfect",
					artist: "Simple Plan",
				},
			},
		},
		nowPlaying: &Song{
			name:   "Bleeding Love",
			artist: "Leona Lewis",
		},
	}

	p.NextPrevSong("previous")
	p.NextPrevSong("previous")

	if reflect.DeepEqual(p.nowPlaying.name, compareP2.nowPlaying.name) != true {
		t.Errorf("Next Previous Song Incorrect: %v ; wants: %v", p.nowPlaying.name, compareP2.nowPlaying.name)
	}
}

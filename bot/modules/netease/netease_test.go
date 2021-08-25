package netease

import (
	"testing"
)

func TestSearch(t *testing.T) {
	res, err := Search("挪威的森林")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("search: %v", res)
	}
}

func TestSongRandom(t *testing.T) {
	res, err := GetRandomSong()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("search: %v", res)
	}
}

func TestSongDetail(t *testing.T) {
	res, err := GetSongDetail(157288)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("search: %v", res)
	}
}

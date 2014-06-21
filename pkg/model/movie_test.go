package model

import "testing"

func TestNewMovie(t *testing.T) {
	movie, err := NewMovie("/path/to/movie.mkv")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if movie.Filename != "/path/to/movie.mkv" {
		t.Errorf("expected movie.Filename /path/to/movie.mkv, got %v", movie.Filename)
	}
}

func TestNewMovieEmptyFilename(t *testing.T) {
	_, err := NewMovie("")
	if err == nil {
		t.Errorf("expected 'empty filename' error, got %v", err)
	}
}

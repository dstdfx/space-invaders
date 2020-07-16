package main

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/hajimehoshi/ebiten"
)

// sequence represents a set of spites that handles its changing.
type sequence struct {
	currentFrame int
	frames       []*ebiten.Image
	sampleRate   float64
	loop         bool
}

func newSequence(filepath string, sampleRate float64, loop bool) (*sequence, error) {
	var seq sequence

	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v: %v", filepath, err)
	}

	for _, file := range files {
		filename := path.Join(filepath, file.Name())

		f, err := loadImage(filename)
		if err != nil {
			return nil, fmt.Errorf("loading sequence frame: %v", err)
		}
		seq.frames = append(seq.frames, f)
	}

	seq.sampleRate = sampleRate
	seq.loop = loop

	return &seq, nil
}

func (seq *sequence) image() *ebiten.Image {
	return seq.frames[seq.currentFrame]
}

func (seq *sequence) nextFrame() bool {
	if seq.currentFrame == len(seq.frames)-1 {
		if seq.loop {
			seq.currentFrame = 0
		} else {
			return true
		}
	} else {
		seq.currentFrame++
	}

	return false
}

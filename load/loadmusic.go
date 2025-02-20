package load

import (
	"bytes"
	"embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

func LoadMusic(ctx *audio.Context, fs embed.FS, path string) *audio.Player {
	musicFile, err := fs.ReadFile(path)
	if err != nil {
		panic(err)
	}
	musicStream, err := vorbis.DecodeWithSampleRate(44100, bytes.NewReader(musicFile))
	if err != nil {
		panic(err)
	}

	// Crear un loop infinito de la música
	loopStream := audio.NewInfiniteLoop(musicStream, musicStream.Length())

	player, err := audio.NewPlayer(ctx, loopStream)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.5)
	return player
}

func LoadHitSound(ctx *audio.Context, fs embed.FS, path string) *audio.Player {
	soundFile, err := fs.ReadFile(path)
	if err != nil {
		panic(err)
	}
	soundStream, err := vorbis.DecodeWithSampleRate(44100, bytes.NewReader(soundFile))
	if err != nil {
		panic(err)
	}

	player, err := audio.NewPlayer(ctx, soundStream)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.5)
	return player
}

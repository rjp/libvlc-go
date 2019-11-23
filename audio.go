package vlc

// #cgo LDFLAGS: -L/Applications/VLC.app/Contents/MacOS/lib/ -lvlc
// #cgo CFLAGS: -I/Applications/VLC.app/Contents/MacOS/include/
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import "errors"

// AudioOutput is an abstraction for rendering decoded (or pass-through)
// audio samples.
type AudioOutput struct {
	Name        string
	Description string
}

// AudioOutputList returns a list of audio output devices that can be used
// with an instance of a player.
func AudioOutputList() ([]*AudioOutput, error) {
	if inst == nil {
		return nil, errors.New("module must be initialized first")
	}

	outputs := C.libvlc_audio_output_list_get(inst.handle)
	if outputs == nil {
		return nil, getError()
	}
	defer C.libvlc_audio_output_list_release(outputs)

	audioOutputs := []*AudioOutput{}
	for p := outputs; p != nil; p = (*C.libvlc_audio_output_t)(p.p_next) {
		audioOutput := &AudioOutput{
			Name:        C.GoString(p.psz_name),
			Description: C.GoString(p.psz_description),
		}

		audioOutputs = append(audioOutputs, audioOutput)
	}

	return audioOutputs, getError()
}

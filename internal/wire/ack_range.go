package wire

import "github.com/mutdroco/mpquic_nonml_for_video_stream_backend/internal/protocol"

// AckRange is an ACK range
type AckRange struct {
	First protocol.PacketNumber
	Last  protocol.PacketNumber
}

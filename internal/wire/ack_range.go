package wire

import "github.com/mutdroco/mpquic_for_video_stream_backend/internal/protocol"

// AckRange is an ACK range
type AckRange struct {
	First protocol.PacketNumber
	Last  protocol.PacketNumber
}

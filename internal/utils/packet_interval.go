package utils

import "github.com/mutdroco/mpquic_for_video_stream_backend/internal/protocol"

// PacketInterval is an interval from one PacketNumber to the other
// +gen linkedlist
type PacketInterval struct {
	Start protocol.PacketNumber
	End   protocol.PacketNumber
}

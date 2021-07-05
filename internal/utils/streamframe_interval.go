package utils

import "github.com/mutdroco/mpquic_for_video_stream_backend/internal/protocol"

// ByteInterval is an interval from one ByteCount to the other
// +gen linkedlist
type ByteInterval struct {
	Start protocol.ByteCount
	End   protocol.ByteCount
}

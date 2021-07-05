package congestion

import "github.com/mutdroco/mpquic_nonml_for_video_stream_backend/internal/protocol"

type connectionStats struct {
	slowstartPacketsLost protocol.PacketNumber
	slowstartBytesLost   protocol.ByteCount
}

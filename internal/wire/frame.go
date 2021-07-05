package wire

import (
	"bytes"

	"github.com/mutdroco/mpquic_nonml_for_video_stream_backend/internal/protocol"
)

// A Frame in QUIC
type Frame interface {
	Write(b *bytes.Buffer, version protocol.VersionNumber) error
	MinLength(version protocol.VersionNumber) (protocol.ByteCount, error)
}

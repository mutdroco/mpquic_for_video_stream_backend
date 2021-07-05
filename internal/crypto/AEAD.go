package crypto

import "github.com/mutdroco/mpquic_nonml_for_video_stream_backend/internal/protocol"

// An AEAD implements QUIC's authenticated encryption and associated data
type AEAD interface {
	Open(dst, src []byte, packetNumber protocol.PacketNumber, associatedData []byte) ([]byte, error)
	Seal(dst, src []byte, packetNumber protocol.PacketNumber, associatedData []byte) []byte
	Overhead() int
}

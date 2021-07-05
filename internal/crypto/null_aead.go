package crypto

import "github.com/mutdroco/mpquic_nonml_for_video_stream_backend/internal/protocol"

// NewNullAEAD creates a NullAEAD
func NewNullAEAD(p protocol.Perspective, v protocol.VersionNumber) AEAD {
	if v.UsesTLS() {
		return &nullAEADFNV64a{}
	}
	return &nullAEADFNV128a{perspective: p}
}

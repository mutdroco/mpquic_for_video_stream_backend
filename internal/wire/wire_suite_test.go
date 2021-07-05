package wire

import (
	"github.com/mutdroco/mpquic_for_video_stream_backend/internal/protocol"
	"github.com/mutdroco/mpquic_for_video_stream_backend/internal/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCrypto(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wire Suite")
}

const (
	versionLittleEndian = protocol.Version37 // a QUIC version that uses little endian encoding
	versionBigEndian    = protocol.Version39 // a QUIC version that uses big endian encoding
)

var _ = BeforeSuite(func() {
	Expect(utils.GetByteOrder(versionLittleEndian)).To(Equal(utils.LittleEndian))
	Expect(utils.GetByteOrder(versionBigEndian)).To(Equal(utils.BigEndian))
})

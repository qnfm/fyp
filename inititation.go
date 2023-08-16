type MessageInitiation struct {
Type        uint32
Sender      uint32
Ephemeral   [NoisePublicKeySize+chacha20poly1305.Overhead]byte
//Server Static Encapsulation
CipherTextS [kyber512.CiphertextSize]byte
Static      [blake2s.Size+chacha20poly1305.Overhead]byte
Timestamp   [tai64n.TimestampSize+chacha20poly1305.Overhead]byte
MAC1        [blake2s.Size128]byte
MAC2        [blake2s.Size128]byte
}
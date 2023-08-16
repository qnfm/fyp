type MessageResponse struct {
Type        uint32
Sender      uint32
Receiver    uint32
//Ephemeral Encapsulation
CipherTextE [kyber512.CiphertextSize]byte
//Client Static Encapsulation
CipherTextC [kyber512.CiphertextSize
+chacha20poly1305.Overhead]byte 
MAC1        [blake2s.Size128]byte
MAC2        [blake2s.Size128]byte
}
func BenchmarkHandshakeServer(b *testing.B) {
for i := 0; i < b.N; i++ {
	b.StopTimer()
	dev1 := randBDevice(b)
	dev2 := randBDevice(b)
	
	peer1, _ := dev2.NewPeer(dev1.staticIdentity.privateKey.publicKey())
	peer2, _ := dev1.NewPeer(dev2.staticIdentity.privateKey.publicKey())
	peer1.Start()
	peer2.Start()

	msg1, _ := dev1.CreateMessageInitiation(peer2)

	packet := make([]byte, 0, MessageInitiationSize)
	writer := bytes.NewBuffer(packet)
	binary.Write(writer, binary.LittleEndian, msg1)
	b.StartTimer()
	dev2.ConsumeMessageInitiation(msg1)
	msg2, _ := dev2.CreateMessageResponse(peer1)
	b.StopTimer()
	dev1.ConsumeMessageResponse(msg2)
	dev1.Close()
	dev2.Close()
	b.StartTimer()
}
}

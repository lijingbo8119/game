package client

func Main() {
	go viewProgramStart()
	connInit("localhost:8081")
	websocketClient()
}

package client

func Main() {
	go func() {
		viewsInit()
		viewInit()
		viewProgram.Start()
	}()
	connInit("localhost:8081")
	websocketClient()
}

let ws = new WebSocket("ws://localhost:3000/ws?user=jaimesoad&recepient=foo")

ws.onmessage = (event) => {
    console.log(event.data)
}

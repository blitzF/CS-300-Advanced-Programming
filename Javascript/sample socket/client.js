new Vue({
	data: {
		text: "pop",
		message: [],
		name: "",
		ws: new WebSocket("ws://localhost:8080")
		

	},
	methods: {
		sendMessage(num) {
			this.ws.send(JSON.stringify({
				type: "resp",
				data: num
			})) 
			//this.message.push(`You: ${this.number}`)
		}

	},
	mounted() {
		this.ws.onopen = () => {
			this.name = prompt("Enter your name?")
			this.ws.send(JSON.stringify({
				type: "name",
				data: this.name
			})) 
		}

		this.ws.onmessage = recvMsg => {
			recvdMsg = JSON.parse(recvMsg.data)
			if(recvdMsg.name == "first"){
				document.getElementById(recvdMsg.place).style.backgroundColor = "red"
			}
			else{
				document.getElementById(recvdMsg.place).style.backgroundColor = "purple"
			}
			//console.log(recvMsg)
		}
	}
}).$mount(`#root`)
const webSocket = require('ws')

const wss = new webSocket.Server({port:8080})
console.log("WebScoekt Server Started...")
let first


wss.on('connection', ws => {
	if (first == undefined){
		first = ws
		first.on('message', async msg =>{
			msgg = JSON.parse(msg)
			if(msgg.type == "name") {
			ws.PersonName = msgg.data
			console.log(`${ws.PersonName} is Connected`) 
			}
		})
		

	}else{
		while(seclock){
			//nothing
		}
		first.on('message', async msg =>{
			msgg = JSON.parse(msg)
			if(msgg.type == "name") {
			ws.PersonName = msgg.data
			console.log(`${ws.PersonName} is Connected`) 
			}

		if(msgg.type == "resp") {
			wss.clients.forEach(client => {
				client.send(JSON.stringify({name: "first", place: msgg.data}))
				})
			//console.log(`${msgg.data}`)
			}
		})
		ws.on('message',async msg => {
		msgg = JSON.parse(msg)
		if(msgg.type == "name") {
			ws.PersonName = msgg.data
			console.log(`${ws.PersonName} is Connected`)
		}

		if(msgg.type == "resp") {
			wss.clients.forEach(client => {
				client.send(JSON.stringify({name: "second", place: msgg.data}))
				})
			//console.log(`${msgg.data}`)
			}
		})

	}
})
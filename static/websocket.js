var message = {Type: "", Timestamp:null, Data:null};
var webSocket;
var jSessionState;
var url;
var playerId;
var gameId;

function connect(){
	webSocket= new WebSocket(url);
	jSessionState.append("<p>Connecting...</p>");
	webSocket.onopen= function(event){
		jSessionState.append("<p>Connected</p>");
	}
	
	webSocket.onmessage = function(event){
		var msg = JSON.parse(event.data);
		console.log(msg);
		if(msg.Type == "Connection ok")
		{
			message.Type = "Broker new game";
			var d = new Date();
			message.Timestamp = d;
			webSocket.send(JSON.stringify(message));
		}
		if(msg.Type == "Join game ok"){
			playerId = msg.Data.PlayerId;
-			gameId = msg.Data.GameId;
			jSessionState.append("<p>Joined game</p>");
			jSessionState.append("<p>Waiting for other player...</p>");
			
			
		}
		if(msg.Type == "Player won game") {
			jSessionState.append("<p>Congratulations, you won!</p>")
		}
	}	
	
	
	webSocket.onclose = function(e) {
		console.log('Socket is closed. Reconnect will be attempted in 1 second.', e.reason);
		setTimeout(function() {
			message.data = {"GameId":gameId, "PlayerId":playerId};
			connect();
		}, 1000)
	};

	webSocket.onerror = function(err) {
		console.error('Socket encountered error: ', err.message, 'Closing socket');
		webSocket.close();
	};
		
}



var message = {Type: "", Timestamp:null, Data:null};
var webSocket;
var jSessionState;
var url;
var playerId = "";
var gameId = "";

function addToLog(text){
	jSessionState.append(text);
	jSessionState.stop().animate({
		scrollTop: jSessionState[0].scrollHeight}, 800);
}


function connect(){
	webSocket= new WebSocket(url);
	addToLog("Connecting...</br>");
	webSocket.onopen= function(event){
		addToLog("Connected</br>");
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
			gameId = msg.Data.GameId;
			
			addToLog("Joined game</br>");
			addToLog("Waiting for other player...</br>");			
		}
		
		if(msg.Type == "Placement phase")
		{
			//let player place ships and hit ready once done
		}
		
		if(msg.Type == "Player won game") {
			addToLog("<p>Congratulations, you won!</p>")
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



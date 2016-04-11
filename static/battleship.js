var battleship;


function battleshipGame(canvasId){
	var self = this;
	this.canvas = document.getElementById(canvasId);
	this.jCanvas = $("#"+canvasId);
	this.context = this.canvas.getContext("2d");
	
	this.bw = this.jCanvas.width()-25;
	this.bh = this.jCanvas.height()-40;
	this.halfBh = this.bh/2;
	this.size = 10;
	this.leftMargin = 4;
	this.topMargin = 4;
	
	this.letters = ['a','b','c','d','e','f','g','h','i','j'];
	
	this.enemyField = [];
	this.friendlyField = [];
	this.ships = [];
	
	this.initBoard = function(){
		var context = this.context;
		var bh = this.bh;
		var bw = this.bw;
		var size = this.size;
		var leftMargin = this.leftMargin;
		var halfBh = this.halfBh;
		var topMargin = this.topMargin;
		var letters = this.letters;
		var enemyField = this.enemyField;
		var friendlyField = this.friendlyField;
		
		var countX = bw/size;
		var countY = bh/(size*2);
		
		//fill field arrays
		
		for(var col = 1; col <= size; ++col)
		{
			enemyField[col-1] = [];
			friendlyField[col-1] = []
			for(var row = 1; row <= size; ++row)
			{
				var enemySquare = {};
				enemySquare.left = leftMargin + (col-1)*countX + size;
				enemySquare.right = leftMargin + (col)*countX + size;
				enemySquare.top = 0.5 + (row-1)*countY + size + topMargin;
				enemySquare.bottom = 0.5 + (row)*countY + size + topMargin;
				enemySquare.fillStyle = "white";
				enemyField[col-1][row-1] = enemySquare;
				
				
				var friendlySquare = {};
				friendlySquare.left = leftMargin + (col-1)*countX + size;
				friendlySquare.right = leftMargin + col*countX + size;
				friendlySquare.top = 0.5 + (row-1)*countY + size*2 + topMargin*2 + halfBh ;
				friendlySquare.bottom = 0.5 + row*countY + size*2 + topMargin*2 + halfBh;
				friendlySquare.fillStyle = "white";
				friendlyField[col-1][row-1] = friendlySquare;
				
			}
		}
		
	}
	
	this.drawBoard = function(){
		var context = this.context;
		var bh = this.bh;
		var bw = this.bw;
		var size = this.size;
		var leftMargin = this.leftMargin;
		var halfBh = this.halfBh;
		var topMargin = this.topMargin;
		var letters = this.letters;
		var enemyField = this.enemyField;
		var friendlyField = this.friendlyField;
		
		context.strokeStyle = "gray";
		context.stroke();
		
		
		this.drawGridLines();
		this.drawText();
		this.drawShips();
	}
	
	this.drawGridLines = function(){
		var context = this.context;
		var bw = this.bw;
		var bh = this.bh;
		var size = this.size;
		var halfBh = bh/2;
		var leftMargin = this.leftMargin;
		var topMargin = this.topMargin;
		var countX = bw/size;
		var countY = bh/(size*2);
		for (var x = 0; x <= bw; x += countX) {
			context.moveTo(leftMargin + x + size, size + topMargin);
			context.lineTo(leftMargin + x + size, halfBh + size + topMargin);
			
			context.moveTo(leftMargin + x + size, size + halfBh + size + topMargin*2);
			context.lineTo(leftMargin + x + size, halfBh + size + halfBh + size + topMargin*2);
		}
		for (var y = 0; y <= halfBh; y += countY) {
			context.moveTo(leftMargin + size, 0.5 + y + size + topMargin);
			context.lineTo(leftMargin + bw + size, 0.5 + y + size + topMargin);
			
			context.moveTo(leftMargin + size, 0.5 + y + size + halfBh + size + topMargin*2);
			context.lineTo(leftMargin + bw + size, 0.5 + y + size + halfBh + size + topMargin*2);
		}
		
		
		context.strokeStyle = "black";
		context.stroke();
		
	}
	
	this.drawText = function(){
		var context = this.context;
		var bw = this.bw;
		var bh = this.bh;
		var size = this.size;
		var topMargin = this.topMargin;
		var countX = bw/size;
		var countY = bh/(size*2);
		var letters = this.letters;
		
		for(var i = 1; i <= size; ++i)
		{
			context.fillText(i,0.5,0.5+i*countY);
			context.fillText(i,0.5,0.5+(i+size)*countY + topMargin);
			
			
			context.fillText(letters[i-1], 0.5+i*countY, 10);
			context.fillText(letters[i-1], 0.5+i*countY, 25 + size*countY);
		}
	}
	
	this.fillSquares = function(){
		var size = this.size;
		var enemyField = this.enemyField;
		var friendlyField = this.friendlyField;
		var fillSquare = this.fillSquare;
		for(var col = 1; col <= size; ++col)
		{
			for(var row = 1; row <= size; ++row)
			{
				var enemySquare = enemyField[col-1][row-1];
				
				fillSquare(enemySquare);
				
				
				var friendlySquare = friendlyField[col-1][row-1];
				fillSquare(friendlySquare);
			}
			
		}
	}
	
	this.fillSquare = function(square){
		var size = self.size;
		var bw = self.bw;
		var bh = self.bh;
		var context = self.context;
		context.beginPath();
		context.rect(square.left+4, square.top+4, (bw/size)-8, (bh/(size*2))-8);
		context.fillStyle = square.fillStyle;

		context.fill();
	}
	
	
	this.drawShips = function(){
		var ships = self.ships;
		var ctx = self.context;
		
		for(var i = 0; i <= ships.length-1; ++i){
			
			var img = new Image;
			img.src = ships[i].src;
			ctx.drawImage(img,ships[i].x, ships[i].y);
		}
	}
	
	this.mouseHover = function(e){
		
		var size = self.size;
		var bw = self.bw;
		var bh = self.bh;
		var enemyField = self.enemyField;
		var friendlyField = self.friendlyField;
		var context = self.context;
		var rect = this.getBoundingClientRect();
		var x = e.clientX-rect.left;
		var y = e.clientY-rect.top;

		for(var col = 1; col <= size; ++col)
		{
			for(var row = 1; row <= size; ++row)
			{
				var square = enemyField[col-1][row-1];
				
				context.beginPath();
				context.rect(square.left, square.top, bw/size, bh/(size*2));
				
				if(context.isPointInPath(x,y))
				{
					square.fillStyle = "red";
				}
				else
					square.fillStyle = "white"
					
					
				var square = friendlyField[col-1][row-1];
				
				context.beginPath();
				context.rect(square.left, square.top, bw/size, bh/(size*2));
				
				if(context.isPointInPath(x,y))
					square.fillStyle = "red";
				else
					square.fillStyle = "white"
			}
		}
		self.fillSquares();
		self.drawShips();
		
	}
	
	canvas.onmousemove = this.mouseHover;
	
	// To enable drag and drop
	canvas.addEventListener("dragover", function (evt) {
		evt.preventDefault();
	}, false);
	
	canvas.addEventListener("drop", function (evt) {
		evt.preventDefault();
		var id = evt.dataTransfer.getData("text");
		var img = document.getElementById(id);
		//TODO security: check if src url is from correct source
		
		var ship = {};
		ship.src = img.src;
		ship.x = evt.pageX - self.canvas.offsetLeft - $(this).offset().left;
		ship.y = evt.pageY - self.canvas.offsetTop - $(this).offset().top;
				
		var point = snapToFriendlyGrid(ship.x, ship.y);
		ship.x = point.x;
		ship.y = point.y;
		
		
		//todo: move left/down if not completely within grid (depends on size/rotation)
		self.ships.push(ship);
		console.log(ship);
		img.style.display = "none";
		
		//redraw
		self.drawBoard();
		
	});
	
	function snapToFriendlyGrid(x,y){
		var friendlyField = self.friendlyField;
		var currentX = -10000;
		var currentY = -10000;
		console.log(friendlyField);
		for(var i = 0; i < friendlyField.length; ++i){
			for(var j=0; j < friendlyField[i].length; ++j)
			{
				var sq = friendlyField[i][j];
				if(Math.abs(x-friendlyField[i][j].left) < Math.abs(x-currentX)  ){
					currentX = friendlyField[i][j].left;
				}
				if(Math.abs(y-friendlyField[i][j].top) < Math.abs(y-currentY)){
					currentY = friendlyField[i][j].top;
				}
			}
		}
		
		var point = {};
		point.x = currentX;
		point.y = currentY;
		
		
		return point;
	}
	
	this.initBoard();
	this.drawBoard();
	
	
}
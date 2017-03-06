var RED = require(process.env.NODE_RED_HOME+"/red/red");
var Go = require('gonode').Go;

module.exports = function(RED) {
	var go;

	function  Modbus(n){
		RED.nodes.createNode(this, n)
		var node = this;
		node.status({fill:"grey",shape:"ring",text:"Waiting"});
		var json_input = {
			addrModbus: n.addrModbus,
			portModbus: n.portModbus,
			configFile: n.configFile,
			text: 	"Initialize"
		}

		go = new Go({path: 'main/Modbus.go', initAtOnce: true}, function(err) {
			if(err) {
				console.log('Error initializing.\n');
				return;
			}
			go.on('error', function(err) {
				if(err.parser) {
					console.log("Golang program says: " + err.data.toString());
				} else {
					console.log("Gonodo error: " + err.data.toString());
				}
			});

			go.execute(json_input, function(result, response) {
				console.log("OK: " + result.ok);
				console.log("Timeout: " + result.timeout);
				console.log("Terminated: " + result.terminated);
				if(result.ok) {
					node.status({fill:"green",shape:"dot",text:"Done"});
					console.log('Go finished correctly.');
					console.log('Go responded: ' + response.text + '\n\n');
				} else {
					node.status({fill:"red",shape:"dot",text:"Error"});
					console.log('Something went wrong with the \'Create\' command.');
					console.log('Go responded: ' + response.text + '\n\n');
				}
			});

		});


		node.on("input", function(msg) {
			node.status({fill:"green",shape:"ring",text:"Processing"});
			msg.payload.text = "Input";

			go.execute(msg.payload, function(result, response) {
				console.log("OK: " + result.ok);
				console.log("Timeout: " + result.timeout);
				console.log("Terminated: " + result.terminated);
				if(result.ok) {
					node.status({fill:"green",shape:"dot",text:"Done"});
					console.log('Go finished correctly.');
					console.log('Go responded: ' + response.text + '\n\n');
				} else {
					node.status({fill:"red",shape:"dot",text:"Error"});
					console.log('Something went wrong with the \'Input\' command.');
					console.log('Go responded: ' + response.text + '\n\n');
				}
			});
		})
	}
	RED.nodes.registerType("Modbus", Modbus);
}

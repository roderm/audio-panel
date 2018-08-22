import React from "react";
import AvrReceiver from "./AvrReceiver.jsx"

class App extends React.Component {
    constructor(props){
        super(props)
        this.state = {
            AvrDevices: [],
            err: null
        }
    }
    componentDidMount() {
        
        let ws = new WebSocket("ws://localhost:3000/api")
        console.log("asdf")

        /*this.ws = new WebSocket('ws://localhost:3000/devices')
        this.ws.onmessage = e => this.setState({ AvrDevices: Object.values(JSON.parse(e.data)) })
        this.ws.onerror = e => this.setState({ error: 'WebSocket error' })
        this.ws.onclose = e => !e.wasClean && this.setState({ error: `WebSocket error: ${e.code} ${e.reason}` })*/
      }    
    componentWillUnmount() {
      this.ws.close()
    }    
    render() {
        return (
        <div>
            <h2>My AVR</h2>
            {
                this.state.AvrDevices.map(function(avr){
                return <AvrReceiver receiver={avr} />;
                })
            }
        </div>
        );
    }
};

export default App;
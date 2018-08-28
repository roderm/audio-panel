import React from "react";
import AvrReceiver from "./AvrReceiver.jsx"
import WebsocketApi from "../js/websocket.api.js"
import { Container } from 'semantic-ui-react'
class App extends React.Component {
    constructor(props){
        super(props)
        this.state = {
            AvrDevices: new Map(),
            err: null
        }
        
    }
    componentDidMount() {
        var that = this;
        this.ws = new WebsocketApi("ws://localhost:3000/api")
        this.ws.rx("get_devices", {}).then(function(res){
            res.Devices.every((d)=> that.state.AvrDevices.set(d.Id, d))
            that.forceUpdate()
        })
        this.ws.sub("subscribe_update").subscribe("onsubscribe_update", function(dev){
            console.log(dev)
            that.state.AvrDevices.set(dev.Id, dev)
            that.forceUpdate()
        })
      }    
    componentWillUnmount() {
      this.ws.close()
    }    
    render() {
        let that = this

        let avrs = (() => {
            let ret = [];
            for(let [k, d] of that.state.AvrDevices) {
                ret.push(<AvrReceiver receiver={d} key={k} devId={k}/>)
            }
            return ret
        })()
        //that.state.AvrDevices.keys().map((k) => <AvrReceiver receiver={that.state.AvrDevices[k]} key={k} devId={k}/> )
        return (
        <div>
            <h2>My AVR</h2>
            <Container>
                {avrs}
            </Container>
            
        </div>
        );
    }
};

export default App;
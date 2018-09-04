import React from "react";
import DeviceComponent from "./Device.jsx"
import WebsocketApi from "../js/websocket.api.js"
import { Container } from 'semantic-ui-react'

class App extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            Devices: new Map(),
            err: null
        }

    }
    componentDidMount() {
        var that = this;
        this.ws = new WebsocketApi("ws://" + location.hostname + ":3000/api")
        this.ws.rx("get_devices", {}).then(function (res) {
            if (res.Devices) {
                res.Devices.every((d) => that.state.Devices.set(d.Identifier, d))
                console.log(res.Devices)
                that.forceUpdate()
            } else {
                console.log("no devices received")
            }
        })
        this.ws.sub("subscribe_update").subscribe("onsubscribe_update", (update) => {
            let device = that.state.Devices.get(update.DeviceIdentifier)
            if(!device){
                return
            }
            let item = device.Items.find(item => item.Identifier == update.ItemIdentifier)
            let itemIndex = device.Items.indexOf(item)
            let propIndex = item.Properties.indexOf(item.Properties.find( p => p.Name == update.Property.Name))
            if (propIndex != -1) {
                item.Properties[propIndex] = update.Property
            }else{
                item.Properties.push(update.Property)
            }
            if(itemIndex != -1) {
                device.Items[itemIndex] = item
            }else{
                device.Items.push(item)
            }
            that.state.Devices.set(update.DeviceIdentifier, device)
            that.forceUpdate()
        })
    }
    sendProperty(prop){
        this.ws.rx("set", [prop]).then(console.log)
    }
    componentWillUnmount() {
        this.ws.close()
    }
    render() {
        let that = this

        let devs = (() => {
            let ret = [];
            for (let [k, d] of that.state.Devices) {
                ret.push(<DeviceComponent device={d} key={k} devId={k} onUpdate={this.sendProperty.bind(this)}/>)
            }
            return ret
        })()
        //that.state.AvrDevices.keys().map((k) => <AvrReceiver receiver={that.state.AvrDevices[k]} key={k} devId={k}/> )
        return (
            <div>
                <h2>My Devices</h2>
                <Container>
                    {devs}
                </Container>

            </div>
        );
    }
};

export default App;
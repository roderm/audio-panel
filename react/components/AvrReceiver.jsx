import React from "react";
import { Segment, Button, Grid, Icon, Input } from 'semantic-ui-react'
import Slider from 'rc-slider';
import WebsocketApi from "../js/websocket.api"
import 'rc-slider/assets/index.css';

class SmartSlider extends React.Component {
    constructor(props) {
        super(props)
        const { value } = this.props
        this.state = { value }
    }
    componentWillReceiveProps(nextProps) {
        this.setValue(nextProps.value)
    }
    setValue(e) {
        this.setState({ value: e })
    }
    update() {
        this.props.onUpdate(this.state)
    }
    render() {
        return <Slider value={this.state.value} onChange={this.setValue.bind(this)} onAfterChange={this.update.bind(this)} />
    }
}
class AvrZone extends React.Component {
    constructor(props) {
        super(props)
        this.state = props.zone
    }
    componentWillReceiveProps(nextProps) {
        this.setState(nextProps.zone)
    }
    updateVolume(e) {
        let ws = new WebsocketApi("ws://localhost:3000/api")
        ws.rx("set_volume", [{ "Volume": e.value, "Zone": this.props.zoneId, "Device": this.props.devId }]).then(console.log)
    }
    toggleMute(e) {
        let o = { zone: this.state }
        o.zone.Muted = !this.state.Muted
        this.componentWillReceiveProps(o)
        let ws = new WebsocketApi("ws://localhost:3000/api")
        ws.rx("set_mute", [{ "Mute": this.state.Muted, "Zone": this.props.zoneId, "Device": this.props.devId }]).then(console.log)
    }
    togglePwr(e) {
        let o = { zone: this.state }
        o.zone.power = !this.state.power
        this.componentWillReceiveProps(o)
        //let ws = new WebsocketApi("ws://localhost:3000/api")
        //ws.rx("set_mute", [{ "Mute": this.state.State, "Zone": this.props.zoneId, "Device": this.props.devId }]).then(console.log)
    }
    render() {
        return (
            <Segment>
                <h4>{this.props.zone.name}</h4>
                <Grid>
                    <Grid.Row>
                        <Grid.Column width={3}>
                            <Button toggle active={!this.state.power} onClick={this.togglePwr.bind(this)}><Icon name='volume off' /></Button>
                        </Grid.Column>
                        <Grid.Column width={10}>
                            <SmartSlider value={this.state.volume} min={1} max={50} onUpdate={this.updateVolume.bind(this)} />
                        </Grid.Column>
                        <Grid.Column width={3}>
                            <Button toggle active={!this.state.mute} onClick={this.toggleMute.bind(this)}><Icon name='volume off' /></Button>
                        </Grid.Column>
                    </Grid.Row>
                </Grid>

            </Segment>
        )
    }
}
class AvrReceiver extends React.Component {
    constructor(props) {
        super(props)
        this.receiver = props.receiver
    }
    componentWillReceiveProps(nextProps) {
        this.receiver = nextProps.receiver
    }
    render() {
        var zones = this.receiver.zones;

        return (
            <div>
                <h4>({this.props.receiver.id})</h4>
                <Segment.Group>
                    {Object.keys(zones).map((k) => <AvrZone key={k} zone={zones[k]} zoneId={k} devId={this.receiver.id} />)}
                </Segment.Group>
            </div>
        )
    }
}

export default AvrReceiver;
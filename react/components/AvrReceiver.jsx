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
        this.render()
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
        this.state = nextProps.zone
        this.render()
    }
    setVolume(e) {
        this.state.Volume = e
    }
    updateVolume(e) {
        let ws = new WebsocketApi("ws://localhost:3000/api")
        ws.rx("set_volume", [{ "Volume": e.value, "Zone": this.props.zoneId, "Device": this.props.devId }]).then(console.log)
    }
    render() {
        console.log("render zone")
        return (
            <Segment>
                <h4>{this.props.zone.Name}</h4>
                <Grid>
                    <Grid.Row>
                        <Grid.Column width={13}>
                            <SmartSlider value={this.state.Volume} min={1} max={50} onUpdate={this.updateVolume.bind(this)} />
                        </Grid.Column>
                        <Grid.Column width={3}>
                            <Button toggle active={!this.state.Muted}><Icon name='volume off' /></Button>
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
        this.render()
    }
    render() {
        var zones = this.receiver.Zones;
        let myZones = []
        if (zones[0] != undefined) {
            myZones.push(<AvrZone key={0} zone={zones[0]} zoneId={0} devId={this.receiver.Id} />)
        }
        if (zones[1] != undefined) {
            myZones.push(<AvrZone key={1} zone={zones[1]} zoneId={1} devId={this.receiver.Id} />)
        }
        if (zones[2] != undefined) {
            myZones.push(<AvrZone key={2} zone={zones[2]} zoneId={2} devId={this.receiver.Id} />)
        }
        return (
            <div>
                <h4>{this.props.receiver.IP}</h4>
                <Segment.Group>
                    {myZones}
                </Segment.Group>
            </div>
        )
    }
}

export default AvrReceiver;
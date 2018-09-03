import React from "react";
import { Segment, Button, Grid, Icon, Input } from 'semantic-ui-react'
import Slider from 'rc-slider';
import WebsocketApi from "../js/websocket.api"
import 'rc-slider/assets/index.css';

class VolumeSlider extends React.Component {
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

class PowerToggle extends React.Component {
    constructor(props) {
        super(props)
        const { property, ...other } = props
        this.state = { property, other }
    }
    componentWillReceiveProps(nextProps) {
        const { property, ...other } = nextProps
        this.setState({ property, other })
    }
    onClick(e) {
        this.state.property.Value.Boolean = !this.state.property.Value.Boolean
        this.props.onUpdate(this.state.property)
    }
    render() {
        let text = this.state.property.Value.Boolean ? "On": "OFF"
        return <Button toggle active={this.state.property.Value.Boolean} onClick={this.onClick.bind(this)}>{text}</Button>
    }
}
class Empty extends React.Component {
    render() {
        return <p>No component</p>
    }
}
class AvrItem extends React.Component {
    constructor(props) {
        super(props)
        const { is, ...other } = props
        this.state = { is, other }
        this.components = {
            "auto": Empty,
            "power": PowerToggle,
            "mute": PowerToggle,
        }
    }
    componentWillReceiveProps(nextProps) {
        const { is, ...other } = nextProps
        this.setState({is, other})
    }
    render() {
        let TagName = this.components[this.state.is || 'auto'];
        if (TagName == undefined) {
            TagName = this.components["auto"]
        }
        return <TagName {...this.state.other} />
    }
}
export class AvrZone extends React.Component {
    constructor(props) {
        super(props)
        this.state = props.item
    }
    componentWillReceiveProps(nextProps) {
        this.setState(nextProps.zone)
    }
    onUpdate(property) {
        this.props.onUpdate({
            ItemIdentifier: this.state.Identifier,
            Property: property
        })
    }
    render() {
        let Properties = this.state.Properties
        if (Properties == null) {
            return <p>No properties</p>
        }
        return (
            <Segment>
                <h4>{this.state.Name}</h4>
                <div style={{ display: "flex" }}>
                    {Object.keys(Properties).map((p) => <AvrItem is={Properties[p].Name} key={this.state.Identifier + Properties[p].Name} property={Properties[p]} onUpdate={this.onUpdate.bind(this)} />)}
                </div>
            </Segment>
        )
    }
}

export class AvrReceiver extends React.Component {
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

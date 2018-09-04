import React from "react";
import { Segment, Button, Grid, Icon, Input } from 'semantic-ui-react'
import Slider from 'rc-slider';
import WebsocketApi from "../js/websocket.api"
import 'rc-slider/assets/index.css';

class PercentageSlider extends React.Component {
    constructor(props) {
        super(props)
        let { property, style, ...other } = props
        if (!property.Value) {
            property.Value = { Decimal: 0 }
        }
        this.state = { property, style, other }
    }
    componentWillReceiveProps(nextProps) {
        const { property, style, ...other } = nextProps
        this.setState({ property, style, other })
    }
    setValue(e) {
        this.state.property.Value.Decimal = e
        this.setState(this.state)
    }
    update() {
        this.props.onUpdate(this.state.property)
    }
    render() {
        return <Slider value={this.state.property.Value.Decimal} onChange={this.setValue.bind(this)} onAfterChange={this.update.bind(this)} style={this.state.style} />
    }
}

class ToggleButton extends React.Component {
    constructor(props) {
        super(props)
        let { property, style, ...other } = props
        if (!property.Value) {
            property.Value = { Boolean: false }
        }
        this.state = { property, style, other }
    }
    componentWillReceiveProps(nextProps) {
        const { property } = nextProps
        this.setState({ property })
    }
    onClick(e) {
        this.state.property.Value.Boolean = !this.state.property.Value.Boolean
        this.props.onUpdate(this.state.property)
    }
    render() {
        let text = this.state.property.Value.Boolean ? "On" : "OFF"
        return <Button toggle active={this.state.property.Value.Boolean} onClick={this.onClick.bind(this)} style={this.state.style}>{text}</Button>
    }
}
class Empty extends React.Component {
    constructor(props) {
        super(props)
        this.state = props
    }
    render() {
        return <p style={{hidden: true}}></p>
    }
}
class AvrItem extends React.Component {
    constructor(props) {
        super(props)
        const { is, ...other } = props
        this.state = { is, other }
        this.components = {
            "auto": {
                "component": Empty, "css": { order: 100 }
            },
            "power": {
                "component": ToggleButton, "css": { order: 1, "flex-basis": "5rem" }
            },
            "mute": {
                "component": ToggleButton, "css": { order: 10, "flex-basis": "5rem" }
            },
            "volume": {
                "component": PercentageSlider, "css": { order: 2, "flex-basis": "30rem" }
            },
            "bass": {
                "component": PercentageSlider, "css": { order: 30, "flex-basis": "10rem" }
            },
            "treble": {
                "component": PercentageSlider, "css": { order: 31, "flex-basis": "10rem" }
            }
        }
    }
    componentWillReceiveProps(nextProps) {
        const { is, ...other } = nextProps
        this.setState({ is, other })
    }
    render() {
        let c = this.components[this.state.is || 'auto'];
        if (c == undefined) {
            c = this.components["auto"]
        }
        let TagName = c.component
        return <TagName {...this.state.other} style={c.css} />
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
                <div style={{ display: "flex", "flex-wrap": "wrap", width: "100%" }}>
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

import React from "react";
import { AvrZone } from "./AvrReceiver.jsx"
import { Segment, Button, Grid, Icon, Input } from 'semantic-ui-react'

class DynamicRow extends React.Component {
    constructor(props) {
        super(props)
        const { is, ...other } = props
        this.state = { is, other }
        this.components = {
            audio_zone: AvrZone,
            audio_amplifier: AvrZone,
            auto: AvrZone,
        }
    }
    componentWillReceiveProps(nextProps) {
        const { is, ...other } = nextProps
        this.setState({is, other})
    }
    render() {
        const TagName = this.components[this.state.is || 'auto'];
        return <TagName {...this.state.other}/>
    }
}
class DeviceComponent extends React.Component {
    constructor(props) {
        super(props)
        this.device = props.device
    }
    componentWillReceiveProps(nextProps) {
        this.device = nextProps.device
    }
    onUpdate(u){
        u.DeviceIdentifier = this.props.device.Identifier
        this.props.onUpdate(u)
    }
    render() {
        var Items = this.device.Items;
        return (
            <div>
                <h4>({this.props.device.Name})</h4>
            <Segment.Group>
                {Object.keys(Items).map((k) => <DynamicRow is={k.Type} key={k} item={Items[k]} onUpdate={this.onUpdate.bind(this)} />)}
            </Segment.Group>
            </div>
        )
    }
}
export default DeviceComponent;
import React from "react";

class AvrReceiver extends React.Component {
    constructor(props) {
        super(props)
        
    }
    render() {
        return (
            <div>
                <h4>{this.props.receiver.IP}</h4>
                <p>Zones</p>
            </div>
        )
    }
}
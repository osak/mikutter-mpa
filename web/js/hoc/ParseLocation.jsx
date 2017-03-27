import React from 'react';
import * as QueryString from 'query-string';

export default (Base) => class extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        const newProps = {
            ...this.props,
            location: {
                ...this.props.location,
                query: QueryString.parse(this.props.location.search)
            }
        };
        return <Base {...newProps} />
    }
}

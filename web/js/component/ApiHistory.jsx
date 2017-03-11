import React from 'react';

export default class ApiHistory extends React.Component {
    constructor(props) {
        super();
    }

    render() {
        return (
            <div className="apihistory">
                <p>API calls issued to render this page:</p>
                {this.apiCallList()}
            </div>
        );
    }

    apiCallList() {
        const lis = this.props.apiCalls.map((c) => {
            return (
                <li>
                    <code><i>{c.method}</i> {c.url}</code>
                    <code className="apihistory__payload">
                        {c.payload}
                    </code>
                </li>
            );
        });
        return (
            <ul className="apihistory__list">
                {lis}
            </ul>
        );
    }
}

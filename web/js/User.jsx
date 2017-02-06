import React from 'react';
import * as Api from './Api.js';

export default class User extends React.Component {
    constructor() {
        super();
        this.state = {
            user: null
        };
    }

    async componentDidMount() {
        const userId = this.props.params.id;
        const user = await Api.User.individual(userId).get();
        this.setState({user});
    }

    render() {
        if (this.state.user === null) {
            return null;
        }
        return (
            <table>
                <tbody>
                <tr>
                    <th>ID</th>
                    <td>{this.state.user.id}</td>
                </tr>
                <tr>
                    <th>Name</th>
                    <td>{this.state.user.name}</td>
                </tr>
                {this.loginToken()}
                </tbody>
            </table>
        )
    }

    loginToken() {
        const token = this.state.user.loginToken;
        if (token !== undefined) {
            return (
                <tr>
                    <th>Token</th>
                    <td><code>{token}</code></td>
                </tr>
            );
        } else {
            return null;
        }
    }
}
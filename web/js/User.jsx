import React from 'react';
import * as Api from './Api.js';

export default class User extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            user: null
        };
    }

    async componentDidMount() {
        if (this.renderingMe()) {
            const user = await Api.Me.get();
            this.setState({user});
        } else {
            const userId = this.props.match.params.id;
            const user = await Api.User.individual(userId).get();
            this.setState({user});
        }
    }

    renderingMe() {
        return this.props.match.path === '/me';
    }

    render() {
        if (this.state.user === null) {
            return null;
        }
        return (
            <table className="user">
                <tbody>
                <tr>
                    <th>ID</th>
                    <td>{this.state.user.id}</td>
                </tr>
                <tr>
                    <th>Name</th>
                    <td>{this.state.user.name}</td>
                </tr>
                {this.renderingMe() ? this.loginToken() : null}
                </tbody>
            </table>
        )
    }

    loginToken() {
        return (
            <tr>
                <th>Token</th>
                <td className="form-group">
                    <input type="text" value={localStorage.getItem('AUTH_TOKEN')} className="form-control user__login-token" onClick={(e) => e.target.select()} readOnly />
                </td>
            </tr>
        );
    }
}
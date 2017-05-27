import React from "react";
import * as Api from "./Api.js";

export default class User extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            user: null,
            reset: false
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

    async _resetToken() {
        await Api.Token.delete();
        this.setState({
            reset: true
        });
    }

    renderingMe() {
        return this.props.match.path === '/me';
    }

    render() {
        if (this.state.user === null) {
            return null;
        }
        return (
            <div>
                {this.state.reset ? <div className="alert alert-success" role="alert">Your token is successfully reset. <a href="/api/auth/login">Login again.</a></div>: null}
                <table className="user">
                    <tbody>
                    <tr>
                        <th>Name</th>
                        <td>{this.state.user.name}</td>
                    </tr>
                    {this.renderingMe() ? this.loginToken() : null}
                    </tbody>
                </table>
            </div>
        )
    }

    loginToken() {
        return (
            <tr>
                <th>Token</th>
                <td className="form-group user__login-token">
                    <input type="text" value={localStorage.getItem('AUTH_TOKEN')}
                           className="form-control user__login-token-input" onClick={(e) => e.target.select()} readOnly/>
                    <button className="btn btn-inline btn-danger" onClick={this._resetToken.bind(this)}>Reset token</button>
                </td>
            </tr>
        );
    }
}
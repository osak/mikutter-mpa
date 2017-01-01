import React from 'react';
import * as Api from './Api.js';

export default class Navbar extends React.Component {
    constructor() {
        super();
        this.state = {
            user: null,
        };
    }

    async componentDidMount() {
        let user = await Api.Me.get();
        this.setState({
            user: user
        });
    }

    render() {
        return (
            <nav className="container">
                <div className="container-fluid">
                    <div className="navbar-header">
                        <a className="navbar-brand" href="/">MPA</a>
                    </div>
                    <ul className="nav navbar-nav navbar-right">
                        <li>{this.currentUser()}</li>
                    </ul>
                </div>
            </nav>
        );
    }

    currentUser() {
        if (this.state.user) {
            return (<a href="#">{this.state.user.login}</a>);
        } else {
            return (<a href="/api/auth/login">Login</a>);
        }
    }
}

Navbar.properties = {
}

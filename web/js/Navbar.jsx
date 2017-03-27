import React from 'react';
import {Link} from 'react-router-dom';
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
                        <Link className="navbar-brand" to="/">MPA</Link>
                    </div>
                    <ul className="nav navbar-nav">
                        {this.pluginAdd()}
                    </ul>
                    <ul className="nav navbar-nav navbar-right">
                        <li>{this.currentUser()}</li>
                    </ul>
                </div>
            </nav>
        );
    }

    pluginAdd() {
        if (this.state.user) {
            return (<li><Link to="/plugin/add">Add plugin</Link></li>);
        } else {
            return null;
        }
    }
    currentUser() {
        if (this.state.user) {
            return (<Link to="/me">{this.state.user.login}</Link>);
        } else {
            return (<a href="/api/auth/login">Login</a>);
        }
    }
}

Navbar.properties = {
}

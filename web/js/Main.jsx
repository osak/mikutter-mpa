import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter, Route} from 'react-router-dom';

import Top from './Top.jsx';
import Navbar from './Navbar.jsx';
import Plugin from './Plugin.jsx';
import Search from './Search.jsx';
import RegisterPlugin from './RegisterPlugin.jsx';
import User from './User.jsx';
import ApiHistory from './component/ApiHistory.jsx';
import {registerAjaxHook} from './Ajax.jsx';

class Main extends React.Component {
    constructor() {
        super();
        this.state = {
            apiCalls: []
        };
        registerAjaxHook((url, method, params) => this.setState({
            apiCalls: this.state.apiCalls.concat({url, method, payload:params}),
        }));
    }

    componentWillReceiveProps() {
        this.setState({
            apiCalls: []
        });
    }

    render() {
        return (
            <div>
                <header className="navbar navbar-inverse">
                    <Navbar />
                </header>
                <div className="container">
                    <ApiHistory apiCalls={this.state.apiCalls} />
                    <Route exact path="/" component={Top} />
                    <Route path="/plugin" component={Search} />
                    <Route path="/plugin/add" component={RegisterPlugin} />
                    <Route path="/plugin/:name" component={Plugin} />
                    <Route path="/user/:id" component={User} />
                    <Route path="/me" component={User} />
                </div>
            </div>
        );
    }
}

function render() {
    ReactDOM.render((
        <BrowserRouter>
            <Route path="/" component={Main} />
        </BrowserRouter>
    ), document.getElementById('main'));
}

function loadAuthCookie() {
    let cookies = document.cookie.split(';').reduce((h, p) => {
        let [key, val] = p.split('=');
        h[key] = val;
        return h;
    }, {});
    if (cookies['AUTH_TOKEN']) {
        localStorage.setItem('AUTH_TOKEN', cookies['AUTH_TOKEN']);
        // Remove cookie
        document.cookie = 'AUTH_TOKEN=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path="/"';
    }
}

window.addEventListener('DOMContentLoaded', () => {
    loadAuthCookie();
    render();
});

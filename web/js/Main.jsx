import React from 'react';
import ReactDOM from 'react-dom';
import {Router, Route, IndexRoute, browserHistory} from 'react-router';

import Top from './Top.jsx';
import Navbar from './Navbar.jsx';
import Plugin from './Plugin.jsx';
import Search from './Search.jsx';

class Main extends React.Component {
    render() {
        return (
            <div>
                <header className="navbar navbar-inverse">
                    <Navbar />
                </header>
                <div className="container">
                    {this.props.children}
                </div>
            </div>
        );
    }
}

function render() {
    ReactDOM.render((
        <Router history={browserHistory}>
            <Route path="/" component={Main}>
                <IndexRoute component={Top} />
                <Route path="/plugin" component={Search} />
                <Route path="/plugin/:name" component={Plugin} />
            </Route>
        </Router>
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

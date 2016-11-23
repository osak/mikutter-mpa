import React from 'react';
import ReactDOM from 'react-dom';

import Navbar from './Navbar.jsx';
import Plugin from './Plugin.jsx';

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
    let testPlugin = {
        name: 'Test',
        version: '3.9.3',
        description: 'Test plugin'
    };
    ReactDOM.render(<Main><Plugin spec={testPlugin} /></Main>, document.getElementById('main'));
}

window.addEventListener('DOMContentLoaded', () => {
    render();
});

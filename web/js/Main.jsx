import React from 'react';
import ReactDOM from 'react-dom';

import Navbar from './Navbar.jsx';

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
    ReactDOM.render(<Main></Main>, document.getElementById('main'));
}

window.addEventListener('DOMContentLoaded', () => {
    render();
});

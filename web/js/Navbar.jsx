import React from 'react';

export default class Navbar extends React.Component {
    render() {
        return (
            <nav className="container">
                <div className="container-fluid">
                    <div className="navbar-header">
                        <a className="navbar-brand" href="/">MPA</a>
                    </div>
                </div>
            </nav>
        );
    }
}

Navbar.properties = {
}

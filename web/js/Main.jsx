import React from 'react';
import ReactDOM from 'react-dom';

function render() {
    ReactDOM.render(<h1>Test</h1>, document.getElementById('main'));
}

window.addEventListener('DOMContentLoaded', () => {
    render();
});

import React from 'react';

export default class SearchBox extends React.Component {
    constructor(props) {
        super();
        this.onTextboxChange = this.onTextboxChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.state = {
            searchText: props.searchText || ''
        };
    }

    onTextboxChange(e) {
        this.setState({
            searchText: e.target.value
        });
    }

    handleSubmit(e) {
        e.preventDefault();
        this.props.onSubmit(this.state.searchText);
    }

    render() {
        let placeholder = this.props.placeholder || 'Enter search query';
        return (
            <form className="form-inline" onSubmit={this.handleSubmit}>
                <div className="form-group">
                    <input type="text" className="form-control" placeholder={this.props.placeholder} value={this.state.searchText} onChange={this.onTextboxChange} />
                    <button className="btn btn-primary">Search</button>
                </div>
            </form>
        );
    }
}

SearchBox.properties = {
    onSubmit: React.PropTypes.func.isRequired,
    placeholder: React.PropTypes.string
};

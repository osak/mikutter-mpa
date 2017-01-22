import React from 'react';

export default class SearchBox extends React.Component {
    constructor(props) {
        super();
        this.onTextboxChange = this.onTextboxChange.bind(this);
        this.onSearchButtonClick = this.onSearchButtonClick.bind(this);
        this.state = {
            searchText: props.searchText || ''
        };
    }

    onTextboxChange(e) {
        this.setState({
            searchText: e.target.value
        });
    }

    onSearchButtonClick(e) {
        this.props.onSubmit(this.state.searchText);
    }

    render() {
        let placeholder = this.props.placeholder || 'Enter search query';
        return (
            <div className="form-inline">
                <div className="form-group">
                    <input type="text" className="form-control" placeholder={this.props.placeholder} value={this.state.searchText} onChange={this.onTextboxChange} />
                    <button className="btn btn-primary" onClick={this.onSearchButtonClick}>Search</button>
                </div>
            </div>
        );
    }
}

SearchBox.properties = {
    onSubmit: React.PropTypes.func.isRequired,
    placeholder: React.PropTypes.string
};

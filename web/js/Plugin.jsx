import React from 'react';

export default class Plugin extends React.Component {
    constructor() {
        super();
        this.onTextboxClick = this.onTextboxClick.bind(this);
    }

    render() {
        let repository = this.props.spec.repositoryUrl || '(unpublished)';
        return (
            <div className="plugin">
                <h1 className="plugin__name">{this.props.spec.name}</h1>
                <div className="plugin__version">{this.props.spec.version}</div>
                <div className="plugin__description">{this.props.spec.description}</div>
                <div className="plugin__repository form-inline">
                    <div className="form-group">
                        <label htmlFor="repository" className="plugin__repository__label">URL</label>
                        <input name="repository" className="form-control" type="text" value={repository} readOnly onClick={this.onTextboxClick}/>
                    </div>
                </div>
            </div>
        );
    }

    onTextboxClick(e) {
        e.target.select();
    }
}

Plugin.properties = {
    spec: React.PropTypes.shape({
        name: React.PropTypes.string.isRequired,
        version: React.PropTypes.string.isRequired,
        description: React.PropTypes.string.isRequired,
        repositoryUrl: React.PropTypes.string
    }).isRequired
}

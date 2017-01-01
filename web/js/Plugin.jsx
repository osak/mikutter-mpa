import React from 'react';
import * as Api from './Api.js';

export default class Plugin extends React.Component {
    constructor() {
        super();
        this.onTextboxClick = this.onTextboxClick.bind(this);
        this.state = {
            spec: {
                name: null,
                repositoryUrl: null,
                version: null,
                description: null
            }
        };
    }

    async componentDidMount() {
        let pluginName = this.props.params.name;
        let spec = await Api.Plugin.individual(pluginName).get();
        this.setState({
            spec: spec
        });
    }

    render() {
        let repository = this.state.spec.repositoryUrl || '(unpublished)';
        return (
            <div className="plugin">
                <h1 className="plugin__name">{this.state.spec.name}</h1>
                <div className="plugin__version">{this.state.spec.version}</div>
                <div className="plugin__description">{this.state.spec.description}</div>
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
    params: React.PropTypes.shape({
        name: React.PropTypes.string.isRequired
    }).isRequired
}

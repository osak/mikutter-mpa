import React from 'react';
import * as Api from './Api.js';

export default class Plugin extends React.Component {
    constructor() {
        super();
        this.onTextboxClick = this.onTextboxClick.bind(this);
        this.state = {
            spec: {
                name: null,
                url: null,
                repoUrl: null,
                version: null,
                description: null
            },
            user: {
                name: null
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
        return (
            <div>
                <section className="plugin">
                    <h1 className="plugin__name">{this.state.spec.name}</h1>
                    <div className="plugin__version">{this.state.spec.version}</div>
                    <div className="plugin__description">{this.state.spec.description}</div>
                    <div className="plugin__download">
                        <span className="glyphicon glyphicon-download"></span><a href={this.state.spec.url}>Download</a>
                    </div>
                    {this.repoLink()}
                </section>
                <section className="user">
                    <h2 className="user__name">{this.state.user.name}</h2>
                </section>
            </div>
        );
    }

    repoLink() {
        if (this.state.spec.repoUrl) {
            return (
                <div className="plugin__repository">
                    <span className="glyphicon glyphicon-globe"></span>
                    {this.props.spec.repoUrl}
                </div>
            );
        } else {
            return "";
        }
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

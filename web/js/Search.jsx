import React from 'react';
import {Link} from 'react-router';
import SearchBox from './component/SearchBox.jsx';
import * as Api from './Api.js';

export default class Search extends React.Component {
    constructor() {
        super();
        this.state = {
            searchResult: []
        }
    }

    async componentDidMount() {
        this.update();
    }

    async componentDidUpdate(prevProps, prevState) {
        if (this.props.location.query.filter != prevProps.location.query.filter) {
            this.update();
        }
    }

    async update() {
        let query = this.props.location.query.filter;
        let result = await Api.Plugin.get({
            filter: query
        });
        this.setState({
            searchResult: result
        });
    }

    render() {
        console.log(this.state);
        let searchFunc = (query) => {
            this.props.router.push(`/plugin?filter=${query}`);
        };

        let results = this.state.searchResult.map((spec) => {
            return (
                <li className="search__result" key={spec.name}>
                    <div className="search__result__name"><Link to={`/plugin/${spec.name}`}>{spec.name}</Link></div>
                    <div className="search__result__version">{spec.version}</div>
                    <div className="search__result__description">{spec.description}</div>
                </li>
            );
        });

        return (
            <div>
                <SearchBox searchText={this.props.location.query.filter} onSubmit={searchFunc} />
                <ul className="search">
                    {results}
                </ul>
            </div>
        );
    }
}

Search.properties = {
}

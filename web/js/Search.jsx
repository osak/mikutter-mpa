import React from 'react';
import SearchBox from './component/SearchBox.jsx';
import {get} from './Ajax.jsx';

export default class Search extends React.Component {
    constructor() {
        super();
        this.state = {
            searchResult: []
        }
    }

    async componentDidMount() {
        let query = this.props.location.query.filter;
        let result = await get(`/api/plugin?filter=${query}`);
        this.setState({
            searchResult: result
        });
    }

    render() {
        let searchFunc = (query) => {
            router.push(`/plugin?filter=${query}`);
        };

        let results = this.state.searchResult.map((spec) => {
            return (
                <li className="search__result" key={spec.name}>
                    <div className="search__result__name"><a href={`/plugin/${spec.name}`}>{spec.name}</a></div>
                    <div className="search__result__version">{spec.version}</div>
                    <div className="search__result__description">{spec.description}</div>
                </li>
            );
        });

        return (
            <div>
                <SearchBox onSubmit={searchFunc} />
                <ul className="search">
                    {results}
                </ul>
            </div>
        );
    }
}

Search.properties = {
}
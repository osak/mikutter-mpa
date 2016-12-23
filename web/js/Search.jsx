import React from 'react';

const Search = ({router}) => {
    let searchBox = (<input type="text" className="form-control" placeholder="Plugin name (e.g. mpa)" />);
    let searchFunc = () => {
        let name = searchBox.value;
        router.push(`/plugin/${name}`);
    };
    return (
        <div className="jumbotron">
            <h1>Mikutter Plugin Archives</h1>
            <p>Welcome! Try search:</p>
            <div className="form-inline">
                <div className="form-group">
                    {searchBox}
                    <button className="btn btn-primary" onClick={searchFunc}>Search</button>
                </div>
            </div>
        </div>
    );
}

Search.properties = {
}

export default Search;

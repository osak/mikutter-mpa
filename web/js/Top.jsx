import React from 'react';
import SearchBox from './component/SearchBox.jsx';

const Top = ({history}) => {
    let searchFunc = (query) => {
        let name = query;
        history.push(`/plugin?filter=${name}`);
    };
    return (
        <div className="jumbotron">
            <h1>Mikutter Plugin Archives</h1>
            <p>Welcome! Try search:</p>
            <div className="form-inline">
                <div className="form-group">
                    <SearchBox placeholder="Plugin name (e.g. mpa)" onSubmit={searchFunc} />
                </div>
            </div>
        </div>
    );
}

Top.properties = {
}

export default Top;

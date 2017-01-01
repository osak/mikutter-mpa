import React from 'react';
import * as Api from './Api.js';

const RegisterPlugin = ({router}) => {
    function onSubmitButtonClicked(e) {
        e.preventDefault();
        let name = document.getElementById('plugin-name').value;
        let repo = document.getElementById('plugin-repository').value;
        let version = document.getElementById('plugin-version').value;
        let description = document.getElementById('plugin-description').value;

        (async () => {
            try {
                let result = await Api.Plugin.post({
                    name: name,
                    url: repo,
                    version: version,
                    description: description
                });
                let alertBox = document.getElementById('alert');
                alertBox.className = 'alert alert-success';
                alertBox.innerHTML = 'Successfully created';
            } catch (e) {
                console.log(e);
                let alertBox = document.getElementById('alert');
                alertBox.className = 'alert alert-danger';
                alertBox.innerHTML = 'Internal error';
            }
        })();
        return false;
    }

    return (
        <div>
            <div id="alert" className="alert" role="alert"></div>
            <form className="form-horizontal">
                <div className="form-group">
                    <label htmlFor="plugin-name" className="control-label col-sm-2">Plugin Name</label>
                    <div className="col-sm-10">
                        <input type="text" id="plugin-name" className="form-control" />
                    </div>
                </div>
                <div className="form-group">
                    <label htmlFor="plugin-repository" className="control-label col-sm-2">Repository</label>
                    <div className="col-sm-10">
                        <input type="text" id="plugin-repository" className="form-control" />
                    </div>
                </div>
                <div className="form-group">
                    <label htmlFor="plugin-version" className="control-label col-sm-2">Version</label>
                    <div className="col-sm-10">
                        <input type="text" id="plugin-version" className="form-control" />
                    </div>
                </div>
                <div className="form-group">
                    <label htmlFor="plugin-description" className="control-label col-sm-2">Description</label>
                    <div className="col-sm-10">
                        <input type="text" id="plugin-description" className="form-control" />
                    </div>
                </div>


                <div className="form-group">
                    <div className="col-sm-offset-2 col-sm-10">
                        <button className="btn btn-info" onClick={onSubmitButtonClicked}>Submit</button>
                    </div>
                </div>
            </form>
        </div>
    );
}

export default RegisterPlugin;

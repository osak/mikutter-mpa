import React from 'react';
import * as Api from './Api.js';

const RegisterPlugin = ({router}) => {
    function onSubmitButtonClicked(e) {
        e.preventDefault();

        (async () => {
            try {
                let result = await Api.Plugin.post(new FormData(document.getElementById('plugin-form')));
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
            <form id="plugin-form" className="form-horizontal" encType="multipart/form-data">
                <div className="form-group">
                    <label htmlFor="plugin-archive" className="control-label col-sm-2">Plugin Name</label>
                    <div className="col-sm-10">
                        <input type="file" id="plugin-archive" name="plugin-archive" className="form-control" />
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

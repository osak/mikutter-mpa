import {get, post} from './Ajax.jsx';

const BASE_PATH = '/api';

class Endpoint {
    constructor(path, needAuth = false) {
        this.path = path;
        this.needAuth = needAuth;
    }

    fullUri() {
        return BASE_PATH + this.path;
    }

    get() {
        let headers = this.buildHeaders();
        return get(this.fullUri(), '', headers);
    }

    post(payload) {
        let headers = this.buildHeaders();
        return post(this.fullUri(), payload, headers);
    }

    buildHeaders() {
        let headers = new Map();
        if (this.needAuth) {
            headers.set('Authorization', `Bearer ${localStorage.getItem('AUTH_TOKEN')}`);
        }
        return headers;
    }
}

export const Me = new Endpoint('/me', true);
export const Plugin = new Endpoint('/plugin', true);

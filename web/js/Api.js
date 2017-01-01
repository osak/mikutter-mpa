import {get} from './Ajax.jsx';

const BASE_PATH = '/api';

class Endpoint {
    constructor(path, needAuth = false) {
        this.path = path;
        this.needAuth = needAuth;
    }

    fullUri() {
        return BASE_PATH + this.path;
    }

    call() {
        let headers = new Map();
        if (this.needAuth) {
            headers.set('Authorization', `Bearer ${localStorage.getItem('AUTH_TOKEN')}`);
        }
        return get(this.fullUri(), '', headers);
    }
}

export const Me = new Endpoint('/me', true);

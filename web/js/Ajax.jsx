function ajax(url, method, params, headers, callback) {
    let xhr = new XMLHttpRequest();
    xhr.onreadystatechange = () => {
        if (xhr.readyState < 4 || xhr.status !== 200) {
            return new Error();
        }
        if (xhr.readyState == 4) {
            callback(xhr.response);
        }
    }
    xhr.open(method, url, true);
    for (let [key, val] of headers) {
        xhr.setRequestHeader(key, val);
    }
    if (method == 'GET') {
        xhr.send('');
    } else if (method == 'POST') {
        xhr.send(JSON.stringify(params));
    }
}

function get(url, params, headers) {
    var queryString = '';
    if (params instanceof Map) {
        queryString = [...params].map(([k, v]) => {
            return `${k}=${v}`;
        }).join('&');
    } else if (params instanceof Object) {
        queryString = Object.keys(params).map((k) => {
            return `${k}=${params[k]}`;
        }).join('&');
    }
    var fullUrl = url;
    if (queryString !== '') {
        fullUrl += '?' + queryString;
    }
    return new Promise((resolve, reject) => {
        ajax(fullUrl, 'GET', null, headers, (response) => {
            try {
                resolve(JSON.parse(response));
            } catch (e) {
                reject(e);
            }
        });
    });
}

function post(url, payload, headers) {
    return new Promise((resolve, reject) => {
        ajax(url, 'POST', payload, headers,  (response) => {
            try {
                resolve(JSON.parse(response));
            } catch (e) {
                reject(e);
            }
        });
    });
}


export {get, post};

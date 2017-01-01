function ajax(url, method, headers, callback) {
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
    xhr.send('');
}

function get(url, params, headers) {
    return new Promise((resolve, reject) => {
        ajax(url, 'GET', headers, (response) => {
            resolve(JSON.parse(response));
        });
    });
}

export {get};

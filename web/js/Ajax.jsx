function ajax(url, method, callback) {
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
    xhr.send('');
}

function get(url, params) {
    return new Promise((resolve, reject) => {
        ajax(url, 'GET', (response) => {
            resolve(JSON.parse(response));
        });
    });
}

export {get};

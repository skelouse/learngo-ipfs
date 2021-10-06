// file-upload.service.js

import axios from 'axios';

const BASE_URL = 'http://localhost:1323';

function upload(formData) {
    const url = `${BASE_URL}/add/`;
    return axios.post(url, {
        data:formData,
        header: {
            'Content-Type': 'application/json'
        }
    }).then((response) =>{
        this.resp = response.data
    }).catch((error) => {
        this.resp = error
    });
    
}

export { upload }
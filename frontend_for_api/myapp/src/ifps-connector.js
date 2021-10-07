// file-upload.service.js

import axios from 'axios';

const BASE_URL = 'http://localhost:1323';


function upload(formData) {
    const url = `${BASE_URL}/add`;
    var post_data = {
        method:"post",
        url: url,
        data: formData,
        headers: {'Content-Type': 'multipart/form-data'},
    }
    return axios(post_data)
}

function download(cid, password){
    const url = `${BASE_URL}/get/` + cid
    var get_data = {
        method:"get",
        url: url,
        params: {"password": password},
        headers: {'Content-Type': 'multipart/form-data'},
    }
    return axios(get_data)
}

export { upload, download }
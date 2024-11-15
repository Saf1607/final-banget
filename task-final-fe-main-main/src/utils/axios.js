// src/utils/axios.js
import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:8080/api';

// Menambahkan Authorization header dengan token
axios.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers['Authorization'] = token;
  }
  return config;
}, error => {
  return Promise.reject(error);
});

// Export axios untuk digunakan di file lain
export default axios;

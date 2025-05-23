import axios from 'axios';

const axiosInstance = axios.create({
  baseURL: 'http://localhost:8080', // бекенд
  withCredentials: true,             // чтобы передавать сессионные куки
});

export default axiosInstance;

import axios from 'axios';
import qs from 'qs';

// 创建axios实例
let service: any = {};
if (process.env.NODE_ENV === 'development') {
  service = axios.create({
    baseURL: '/apis', // api的base_url
    timeout: 50000, // 请求超时时间
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    withCredentials: true, // 允许携带cookie
  });
} else {
  // 生产环境下
  service = axios.create({
    baseURL: '/apis', // api的base_url
    timeout: 50000, // 请求超时时间
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    withCredentials: true, // 允许携带cookie
  });
}

// console.log('process.env.BASE_URL',process.env.BASE_URL)
// request拦截器 axios的一些配置
service.interceptors.request.use(
  (config: any) => {
    config.data = qs.stringify(config.data);
    return config;
  },
  (error: any) => {
    // Do something with request error
    console.error('error:', error); // for debug
    Promise.reject(error);
  },
);



export default service;

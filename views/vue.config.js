const path = require("path");
const sourceMap = process.env.NODE_ENV === "development";

module.exports = {
    productionSourceMap: false,
    //这里配置代理服务器，因为开发环境中会发生跨域问题
    devServer: {
        proxy: {
            '/apis': {
                target: 'http://127.0.0.1:9527',
                changeOrigin: true,
                ws: true,
                pathRewrite: {
                  '^/apis': ''
                }
            },
            '/github': {
                target: 'https://api.github.com',
                changeOrigin: true,
                ws: true,
                pathRewrite: {
                  '^/github': ''
                }
            }
        }
    }
};

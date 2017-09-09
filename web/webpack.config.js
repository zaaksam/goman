const path = require('path');
const webpack = require('webpack');

module.exports = {
    entry: {
        vendor: ['vue', 'vue-router', 'lodash', 'axios', 'moment', 'iview-style', 'iview', 'highlight.js-style', 'highlight.js'],
        app: ['./ts/main.ts']
    },
    output: {
        filename: '[name].min.js',
        path: path.resolve(__dirname, 'static', 'js'),
        publicPath: '/static/js/'
    },
    resolve: {
        extensions: ['.ts', '.js', '.vue'],
        alias: {
            'iview-style': 'iview/dist/styles/iview.css',
            'highlight.js-style': 'highlight.js/styles/default.css'
        }
    },
    watch: true,
    module: {
        rules: [
            {
                test: /\.css$/,
                loader: 'style-loader!css-loader'
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader',
                options: {
                    esModule: true
                }
            },
            {
                test: /\.ts$/,
                loader: 'ts-loader',
                exclude: /node_modules/,
                options: {
                    appendTsSuffixTo: [/\.vue$/]
                }
            },
            {
                test: /\.(png|jpg|gif|woff|woff2|svg|eot|ttf)$/,
                loader: 'url-loader',
                options: {
                    name: './static/img/[hash].[ext]'
                }
            }
        ]
    },
    plugins: [
        new webpack.DefinePlugin({
            'process.env': {
                NODE_ENV: '"production"'
                // NODE_ENV: '"development"'
            }
        }),
        new webpack.optimize.CommonsChunkPlugin({
            name: 'vendor',
            minChunks: Infinity
            // filename: 'vendor.min1.js',
            // children: true
            // minChunks: 2
        }),
        new webpack.optimize.UglifyJsPlugin({
            sourceMap: false,
            compress: {
                warnings: false
            },
            // 混淆
            mangle: true
        })
    ],
    devtool: 'source-map'
}
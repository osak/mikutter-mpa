const webpack = require('webpack');

module.exports = {
    entry: ['babel-polyfill', './web/js/Main.jsx'],
    output: { path: './web-build', filename: 'static/bundle.js', sourceMapFilename: 'bundle.map' },
    devtool: '#source-map',
    module: {
        loaders: [{
            test: /.jsx?$/,
            loader: 'babel-loader',
            exclude: /node_modules/,
            query: {
                presets: ['react', 'es2015', 'stage-3']
            }
        }]
    },
    plugins: [
        new webpack.optimize.DedupePlugin(),
        new webpack.optimize.AggressiveMergingPlugin(),
        //new webpack.optimize.UglifyJsPlugin(),
        new webpack.EnvironmentPlugin(['NODE_ENV']),
    ]
}

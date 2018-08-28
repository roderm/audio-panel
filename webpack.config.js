var webpack = require("webpack");
var path = require("path");
var {execSync} = require("child_process");

var config = {
    cache: false,
    entry: "./react/index.js",
    output: {
        path: __dirname + "/views",
        filename: "bundle.js",
        publicPath: "/views",
    },
    devtool: "sourceMap",
    module: {
        rules: [
            {
                test: /\.jsx?$/,
                exclude: /node_modules/,
                use: "babel-loader",
            },
          {
            test: /\.css$/,
            use: [ 'style-loader', 'css-loader' ]
          }
        ]
      }
    
};

module.exports = config;

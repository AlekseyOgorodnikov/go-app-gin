var http = require('follow-redirects').http;
var fs = require('fs');

var options = {
  method: 'GET',
  hostname: 'localhost',
  port: 8080,
  path: '/videos',
  headers: {
    Authorization: 'Basic YWxla3NleTpoZWxsbw==',
  },
  maxRedirects: 20,
};

var req = http.request(options, function (res) {
  var chunks = [];

  res.on('data', function (chunk) {
    chunks.push(chunk);
  });

  res.on('end', function (chunk) {
    var body = Buffer.concat(chunks);
    console.log(body.toString());
  });

  res.on('error', function (error) {
    console.error(error);
  });
});

req.end();

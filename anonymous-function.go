
var test = require('tape')
var spec = require('stream-spec')
var through = require('../')

/*
  I'm using these two functions, and not streams and pipe
  so there is less to break. if this test fails it must be
  the implementation of _through_
*/

function write(array, stream) {
  array = array.slice()
  function next() {
    while(array.length)
      if(stream.write(array.shift()) === false)
        return stream.once('drain', next)
    
    str
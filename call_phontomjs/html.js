"use strict";
var system = require('system');
if (system.args.length === 1) {
    console.log('Try to pass some args when invoking this script!');
} 

var page = require('webpage').create();
//console.log('The default user agent is ' + page.settings.userAgent);
//page.settings.userAgent = 'SpecialAgent';
//page.address = system.args[1];
//page.open("http://www.hnkcsj.com/sdnews.jsp?classid=2&parentid=5", function(status) {
page.open(system.args[1], function(status) {
  if (status !== 'success') {
    console.log('Unable to access network');
  } else {
    var ua = page.evaluate(function() {
		//return document.getElementById('myagent').textContent;
    	return document.documentElement.outerHTML;
    });
	console.log("rst=========")
    console.log(ua);
	console.log("=========")
  }
  phantom.exit();
});
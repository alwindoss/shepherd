(window.webpackJsonp=window.webpackJsonp||[]).push([[3],{396:function(t,e,r){r(45),r(7),r(42),r(43),r(18),r(19),r(20),r(44);var n=r(208);r(207),r(143);var o=r(397);r(93);var c=r(398),f=r(399),h=r(400);function l(t,e){var r;if("undefined"==typeof Symbol||null==t[Symbol.iterator]){if(Array.isArray(t)||(r=function(t,e){if(!t)return;if("string"==typeof t)return y(t,e);var r=Object.prototype.toString.call(t).slice(8,-1);"Object"===r&&t.constructor&&(r=t.constructor.name);if("Map"===r||"Set"===r)return Array.from(t);if("Arguments"===r||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))return y(t,e)}(t))||e&&t&&"number"==typeof t.length){r&&(t=r);var i=0,n=function(){};return{s:n,n:function(){return i>=t.length?{done:!0}:{done:!1,value:t[i++]}},e:function(t){throw t},f:n}}throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}var o,c=!0,f=!1;return{s:function(){r=t[Symbol.iterator]()},n:function(){var t=r.next();return c=t.done,t},e:function(t){f=!0,o=t},f:function(){try{c||null==r.return||r.return()}finally{if(f)throw o}}}}function y(t,e){(null==e||e>t.length)&&(e=t.length);for(var i=0,r=new Array(e);i<e;i++)r[i]=t[i];return r}r(35),r(39),r(5),r(4),r(9);var d=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:[];return Object.keys(t).filter((function(t){return e.includes(t)})).reduce((function(e,r){return Object.assign(e,h({},r,t[r]))}),{})},v=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:[];return Object.keys(t).filter((function(t){return!e.includes(t)})).reduce((function(e,r){return Object.assign(e,h({},r,t[r]))}),{})},m=function(){"use strict";function t(e,r){var n=e.query,path=e.path,o=e.init,text=e.text,f=e.postprocess,h=void 0===f?[]:f;c(this,t),this.query=n,this.path=path,this.init=o,this.postprocess=h,this.options=r||{},this.onlyKeys=null,this.withoutKeys=null,text||this.postprocess.unshift((function(data){return data.map((function(t){return v(t,["text"])}))}))}var e;return f(t,[{key:"only",value:function(t){return this.onlyKeys=Array.isArray(t)?t:[t],this}},{key:"without",value:function(t){return this.withoutKeys=Array.isArray(t)?t:[t],this}},{key:"sortBy",value:function(t,e){return this.query=this.query.simplesort(t,{desc:"desc"===e}),this}},{key:"where",value:function(t){return this.query=this.query.find(t),this}},{key:"search",value:function(t,e){var r;return r="object"===o(t)?t:e?{query:{type:"match",field:t,value:e,prefix_length:1,fuzziness:1,extended:!0,minimum_should_match:1}}:{query:{type:"bool",should:this.options.fullTextSearchFields.map((function(e){return{type:"match",field:e,value:t,prefix_length:1,operator:"and",minimum_should_match:1,fuzziness:1,extended:!0}}))}},this.query=this.query.find({$fts:r}),this}},{key:"surround",value:function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},r=e.before,n=void 0===r?1:r,o=e.after,c=void 0===o?1:o;this.onlyKeys&&this.onlyKeys.push("slug"),this.withoutKeys&&(this.withoutKeys=this.withoutKeys.filter((function(t){return"slug"!==t})));var f=function(data){var e=data.findIndex((function(e){return e.slug===t})),r=new Array(n+c).fill(null,0);if(-1===e)return r;for(var o=data.slice(e-n,e),f=data.slice(e+1,e+1+c),h=0,i=n-1;i>=0;i--)r[i]=o[h]||null,h++;for(var l=0,y=n;y<=c;y++)r[y]=f[l]||null,l++;return r};return this.postprocess.push(f),this}},{key:"limit",value:function(t){return"string"==typeof t&&(t=parseInt(t)),this.query=this.query.limit(t),this}},{key:"skip",value:function(t){return"string"==typeof t&&(t=parseInt(t)),this.query=this.query.offset(t),this}},{key:"fetch",value:(e=n(regeneratorRuntime.mark((function t(){var data,e,r,n,o,c,f=this;return regeneratorRuntime.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:data=this.query.data({removeMeta:!0}),this.onlyKeys&&(e=function(data){return data.map((function(t){return d(t,f.onlyKeys)}))},this.postprocess.unshift(e)),this.withoutKeys&&(r=function(data){return data.map((function(t){return v(t,f.withoutKeys)}))},this.postprocess.unshift(r)),n=l(this.postprocess);try{for(n.s();!(o=n.n()).done;)c=o.value,data=c(data)}catch(t){n.e(t)}finally{n.f()}if(data){t.next=7;break}throw new Error("".concat(this.path," not found"));case 7:return t.abrupt("return",JSON.parse(JSON.stringify(data)));case 8:case"end":return t.stop()}}),t,this)}))),function(){return e.apply(this,arguments)})}]),t}();t.exports=m},411:function(t,e,r){"use strict";r.r(e);r(93),r(40);var n=r(17),o=(r(42),r(43),r(396)),c=r.n(o),f=r(401),h=r.n(f),l=r(402);r.n(l).a.register();var y,d,v=["/"],m=function(){var t={},e=[];Array.from(arguments).forEach((function(r){"string"==typeof r?e.push(r):"object"===Object(n.a)(r)&&(t=r)}));var path="/".concat(e.join("/").replace(/\/+/g,"/")),r=!path||!!v.find((function(t){return t===path})),o=r?{dir:t.deep?{$regex:path}:path}:{path:path},f=r?[]:[function(data){return data[0]}];return new c.a({query:d.chain().find(o,!r),postprocess:f},{fullTextSearchFields:["title","description","slug","text"]})};e.default=function(t){return(y=new h.a("content.db")).loadJSONObject(t),d=y.getCollection("items"),m}}}]);
(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{363:function(t,e,r){"use strict";r(11),r(7),r(9);var n=r(2),o=(r(93),r(5),r(4),r(121),r(35),r(39),r(10)),c=r(74),l=r(98);function d(object,t){var e=Object.keys(object);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(object);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(object,t).enumerable}))),e.push.apply(e,r)}return e}function f(t){for(var i=1;i<arguments.length;i++){var source=null!=arguments[i]?arguments[i]:{};i%2?d(Object(source),!0).forEach((function(e){Object(n.a)(t,e,source[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(source)):d(Object(source)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(source,e))}))}return t}e.a=Object(o.a)(c.a,Object(l.b)("form")).extend({name:"v-form",provide:function(){return{form:this}},inheritAttrs:!1,props:{disabled:Boolean,lazyValidation:Boolean,readonly:Boolean,value:Boolean},data:function(){return{inputs:[],watchers:[],errorBag:{}}},watch:{errorBag:{handler:function(t){var e=Object.values(t).includes(!0);this.$emit("input",!e)},deep:!0,immediate:!0}},methods:{watchInput:function(input){var t=this,e=function(input){return input.$watch("hasError",(function(e){t.$set(t.errorBag,input._uid,e)}),{immediate:!0})},r={_uid:input._uid,valid:function(){},shouldValidate:function(){}};return this.lazyValidation?r.shouldValidate=input.$watch("shouldValidate",(function(n){n&&(t.errorBag.hasOwnProperty(input._uid)||(r.valid=e(input)))})):r.valid=e(input),r},validate:function(){return 0===this.inputs.filter((function(input){return!input.validate(!0)})).length},reset:function(){this.inputs.forEach((function(input){return input.reset()})),this.resetErrorBag()},resetErrorBag:function(){var t=this;this.lazyValidation&&setTimeout((function(){t.errorBag={}}),0)},resetValidation:function(){this.inputs.forEach((function(input){return input.resetValidation()})),this.resetErrorBag()},register:function(input){this.inputs.push(input),this.watchers.push(this.watchInput(input))},unregister:function(input){var t=this.inputs.find((function(i){return i._uid===input._uid}));if(t){var e=this.watchers.find((function(i){return i._uid===t._uid}));e&&(e.valid(),e.shouldValidate()),this.watchers=this.watchers.filter((function(i){return i._uid!==t._uid})),this.inputs=this.inputs.filter((function(i){return i._uid!==t._uid})),this.$delete(this.errorBag,t._uid)}}},render:function(t){var e=this;return t("form",{staticClass:"v-form",attrs:f({novalidate:!0},this.attrs$),on:{submit:function(t){return e.$emit("submit",t)}}},this.$slots.default)}})},416:function(t,e,r){"use strict";r.r(e);r(18),r(44);var n=r(6),o={layout:"unauthz",auth:!1,data:function(){return{name:"",email:"",userName:"",password:""}},methods:{register:function(){var t=this;return Object(n.a)(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,t.$axios.post("/user",{name:t.name,userName:t.userName,email:t.email,password:t.password});case 2:t.$auth.loginWith("local",{data:{email:t.email,password:t.password}}),t.$router.push("/app");case 4:case"end":return e.stop()}}),e)})))()}}},c=r(49),l=r(61),d=r.n(l),f=r(151),m=r(342),h=r(328),v=r(407),w=r(363),y=r(408),O=r(324),_=r(403),V=r(38),x=r(94),component=Object(c.a)(o,(function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-row",{attrs:{align:"center",justify:"center"}},[r("v-col",{attrs:{cols:"12",sm:"8",md:"4"}},[r("v-card",{staticClass:"elevation-12"},[r("v-toolbar",{attrs:{color:"primary",dark:"",flat:""}},[r("v-toolbar-title",[t._v("Register")]),t._v(" "),r("v-spacer")],1),t._v(" "),r("v-card-text",[r("v-form",[r("v-text-field",{attrs:{label:"Name",name:"name","prepend-icon":"mdi-face",type:"text"},model:{value:t.name,callback:function(e){t.name=e},expression:"name"}}),t._v(" "),r("v-text-field",{attrs:{label:"Email",name:"email","prepend-icon":"mdi-email",type:"text"},model:{value:t.email,callback:function(e){t.email=e},expression:"email"}}),t._v(" "),r("v-text-field",{attrs:{label:"User Name",name:"userName","prepend-icon":"mdi-account",type:"text"},model:{value:t.userName,callback:function(e){t.userName=e},expression:"userName"}}),t._v(" "),r("v-text-field",{attrs:{id:"password",label:"Password",name:"password","prepend-icon":"mdi-lock",type:"password"},model:{value:t.password,callback:function(e){t.password=e},expression:"password"}})],1)],1),t._v(" "),r("v-card-actions",[r("v-spacer"),t._v(" "),r("v-btn",{attrs:{color:"primary"},on:{click:t.register}},[t._v("Register")])],1)],1)],1)],1)}),[],!1,null,null,null);e.default=component.exports;d()(component,{VBtn:f.a,VCard:m.a,VCardActions:h.a,VCardText:h.b,VCol:v.a,VForm:w.a,VRow:y.a,VSpacer:O.a,VTextField:_.a,VToolbar:V.a,VToolbarTitle:x.a})}}]);
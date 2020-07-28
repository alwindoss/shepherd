(window.webpackJsonp=window.webpackJsonp||[]).push([[5],{417:function(e,t,r){"use strict";r.r(t);var n={data:function(){return{dialog:!1,headers:[{text:"Dessert (100g serving)",align:"start",sortable:!1,value:"name"},{text:"Calories",value:"calories"},{text:"Fat (g)",value:"fat"},{text:"Carbs (g)",value:"carbs"},{text:"Protein (g)",value:"protein"},{text:"Actions",value:"actions",sortable:!1}],desserts:[],editedIndex:-1,editedItem:{name:"",calories:0,fat:0,carbs:0,protein:0},defaultItem:{name:"",calories:0,fat:0,carbs:0,protein:0}}},computed:{formTitle:function(){return-1===this.editedIndex?"New Item":"Edit Item"}},watch:{dialog:function(e){e||this.close()}},created:function(){this.initialize()},methods:{initialize:function(){this.desserts=[{name:"Frozen Yogurt",calories:159,fat:6,carbs:24,protein:4},{name:"Ice cream sandwich",calories:237,fat:9,carbs:37,protein:4.3},{name:"Eclair",calories:262,fat:16,carbs:23,protein:6},{name:"Cupcake",calories:305,fat:3.7,carbs:67,protein:4.3},{name:"Gingerbread",calories:356,fat:16,carbs:49,protein:3.9},{name:"Jelly bean",calories:375,fat:0,carbs:94,protein:0},{name:"Lollipop",calories:392,fat:.2,carbs:98,protein:0},{name:"Honeycomb",calories:408,fat:3.2,carbs:87,protein:6.5},{name:"Donut",calories:452,fat:25,carbs:51,protein:4.9},{name:"KitKat",calories:518,fat:26,carbs:65,protein:7}]},editItem:function(e){this.editedIndex=this.desserts.indexOf(e),this.editedItem=Object.assign({},e),this.dialog=!0},deleteItem:function(e){var t=this.desserts.indexOf(e);confirm("Are you sure you want to delete this item?")&&this.desserts.splice(t,1)},close:function(){var e=this;this.dialog=!1,this.$nextTick((function(){e.editedItem=Object.assign({},e.defaultItem),e.editedIndex=-1}))},save:function(){this.editedIndex>-1?Object.assign(this.desserts[this.editedIndex],this.editedItem):this.desserts.push(this.editedItem),this.close()}}},o=r(49),l=r(61),c=r.n(l),d=r(151),m=r(342),v=r(328),f=r(407),x=r(326),I=r(412),h=r(410),_=r(347),k=r(133),C=r(408),V=r(324),w=r(403),y=r(38),T=r(94),component=Object(o.a)(n,(function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("v-data-table",{staticClass:"elevation-1",attrs:{headers:e.headers,items:e.desserts,"sort-by":"calories"},scopedSlots:e._u([{key:"top",fn:function(){return[r("v-toolbar",{attrs:{flat:"",color:"white"}},[r("v-toolbar-title",[e._v("My CRUD")]),e._v(" "),r("v-divider",{staticClass:"mx-4",attrs:{inset:"",vertical:""}}),e._v(" "),r("v-spacer"),e._v(" "),r("v-dialog",{attrs:{"max-width":"500px"},scopedSlots:e._u([{key:"activator",fn:function(t){var n=t.on,o=t.attrs;return[r("v-btn",e._g(e._b({staticClass:"mb-2",attrs:{color:"primary",dark:""}},"v-btn",o,!1),n),[e._v("New Item")])]}}]),model:{value:e.dialog,callback:function(t){e.dialog=t},expression:"dialog"}},[e._v(" "),r("v-card",[r("v-card-title",[r("span",{staticClass:"headline"},[e._v(e._s(e.formTitle))])]),e._v(" "),r("v-card-text",[r("v-container",[r("v-row",[r("v-col",{attrs:{cols:"12",sm:"6",md:"4"}},[r("v-text-field",{attrs:{label:"Dessert name"},model:{value:e.editedItem.name,callback:function(t){e.$set(e.editedItem,"name",t)},expression:"editedItem.name"}})],1),e._v(" "),r("v-col",{attrs:{cols:"12",sm:"6",md:"4"}},[r("v-text-field",{attrs:{label:"Calories"},model:{value:e.editedItem.calories,callback:function(t){e.$set(e.editedItem,"calories",t)},expression:"editedItem.calories"}})],1),e._v(" "),r("v-col",{attrs:{cols:"12",sm:"6",md:"4"}},[r("v-text-field",{attrs:{label:"Fat (g)"},model:{value:e.editedItem.fat,callback:function(t){e.$set(e.editedItem,"fat",t)},expression:"editedItem.fat"}})],1),e._v(" "),r("v-col",{attrs:{cols:"12",sm:"6",md:"4"}},[r("v-text-field",{attrs:{label:"Carbs (g)"},model:{value:e.editedItem.carbs,callback:function(t){e.$set(e.editedItem,"carbs",t)},expression:"editedItem.carbs"}})],1),e._v(" "),r("v-col",{attrs:{cols:"12",sm:"6",md:"4"}},[r("v-text-field",{attrs:{label:"Protein (g)"},model:{value:e.editedItem.protein,callback:function(t){e.$set(e.editedItem,"protein",t)},expression:"editedItem.protein"}})],1)],1)],1)],1),e._v(" "),r("v-card-actions",[r("v-spacer"),e._v(" "),r("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:e.close}},[e._v("Cancel")]),e._v(" "),r("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:e.save}},[e._v("Save")])],1)],1)],1)],1)]},proxy:!0},{key:"item.actions",fn:function(t){var n=t.item;return[r("v-icon",{staticClass:"mr-2",attrs:{small:""},on:{click:function(t){return e.editItem(n)}}},[e._v("mdi-pencil")]),e._v(" "),r("v-icon",{attrs:{small:""},on:{click:function(t){return e.deleteItem(n)}}},[e._v("mdi-delete")])]}},{key:"no-data",fn:function(){return[r("v-btn",{attrs:{color:"primary"},on:{click:e.initialize}},[e._v("Reset")])]},proxy:!0}])})}),[],!1,null,null,null);t.default=component.exports;c()(component,{VBtn:d.a,VCard:m.a,VCardActions:v.a,VCardText:v.b,VCardTitle:v.c,VCol:f.a,VContainer:x.a,VDataTable:I.a,VDialog:h.a,VDivider:_.a,VIcon:k.a,VRow:C.a,VSpacer:V.a,VTextField:w.a,VToolbar:y.a,VToolbarTitle:T.a})}}]);
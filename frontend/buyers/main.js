new Vue({
  el: '#app',
  vuetify: new Vuetify(),
})

//replace callback to return a json
var buyers = "http://localhost:8080/buyers"
function load(url, callback){
  var xhr = new XMLHttpRequest();
  
  xhr.onreadystatechange = function(){
      if(xhr.readyState == 4){
          callback(xhr.response);
      }
  }
  xhr.open('GET', url, true);
  xhr.send('');
}
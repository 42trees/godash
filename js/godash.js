document.addEventListener("DOMContentLoaded", function(e) {
    main();

    var es = new EventSource("/woo");

    es.onmessage = function(e) {
      console.log(e.data);
      console.log(e);
    
      d = JSON.parse(e.data);
      console.log(d); 
    }


});

function main() {
  console.log("Main"); 
}




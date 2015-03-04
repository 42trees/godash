document.addEventListener("DOMContentLoaded", function(e) {
    main();

    var es = new EventSource("/woo");

    var content = document.getElementById("content");


    es.onmessage = function(e) {
      console.log(e.data);
      console.log(e);
    
      d = JSON.parse(e.data);
      console.log(d);

      //Loop over the objects received
      for(var i = 0; i < d.length; i++) {
        console.log(d[i].Title);
        var w = Handlebars.templates.text(d[i]);
        
        var existing = document.getElementById("widget-"+d[i].Id);
        if (typeof(existing) != 'undefined' && existing != null) {
          existing.outerHTML = w;
        } else {
          content.insertAdjacentHTML("beforeEnd", w);
        }


      
      }



    }


});

function main() {
  console.log("Main"); 
}




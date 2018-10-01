

window.onload= function(){
  $("#suggestionButton").click(function(){
      var src= $("#suggestionButton").attr("src");
      if (src == "./static/down.png"){
          $("#suggestionButton").attr('src', './static/uparrow.png');

      }
      else {
          $("#suggestionButton").attr('src', './static/down.png')
          var suggestionsCount= $.get("/userHASH/sugestionsCount", data,
              function (data, textStatus, jqXHR) {
                  alert(data)
              }
          );
          
      }
  })  
}

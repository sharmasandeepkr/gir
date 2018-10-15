
 function toggleSuggestions () { 
    if ($('.suggestionRowClass').hasClass("d-none")) {
        $("#txtSuggest").html('<mark style="color:`red`;">Click To Hide Suggested Services</mark>');
            $(".suggestionRowClass").removeClass("d-none");
            $(".suggestionRowClass").removeClass("d-lg-none");
            $(".suggestionRowClass").removeClass("d-md-none");
            $(".suggestionRowClass").removeClass("d-sm-none");
        
    }
    else {
            $("#txtSuggest").html('<mark style="color:`#31a2b8`;">Click For Services Suggestions</mark>');
            $(".suggestionRowClass").addClass("d-none");
            $(".suggestionRowClass").addClass("d-lg-none");
            $(".suggestionRowClass").addClass("d-md-none");
            $(".suggestionRowClass").addClass("d-sm-none");            

    }
 }

window.onload= function(){
   
    
  $("#suggestionButton").click(function(){
      var src= $("#suggestionButton").attr("src");
      if (src == "./static/down.png"){
          $("#suggestionButton").attr('src', './static/uparrow.png');

      }
      else {
          $("#suggestionButton").attr('src', './static/down.png')
                    
      }
  })  
}


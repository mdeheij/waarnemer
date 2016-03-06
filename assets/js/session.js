
         $( window ).load(function() {
            $.ajax({
              url: "/login/session",
              context: document.body
            }).done(function(context) {
              $( this ).addClass( "done" );
                if (context.username != "") {
                    $("#loginbtn").html('<i class="fa fa-user"></i> '+context.username);
                    //$("#loginbtn").attr('href', "/login/logout");
                    $("#loginbtn").attr('href', "#");
                    $("#loginbtn").attr('onclick', "UIkit.modal.confirm('Are you sure you want to log out?', function(){ window.location.href = '/login/logout'; });");
                } else {
                    $("#loginbtn").html("Login");
                }
                 $("#canvasmenu").append('<li class="uk-nav-header">Options</li>');
                 $("#canvasmenu").append("<li>"+$("#loginbtn").parent().html()+"</li>");
              console.log(context.username);
            });
            //$("#canvasmenu").html('<li class="uk-nav-header">Navigation</li>'+$("#mainmenu").html());

         });

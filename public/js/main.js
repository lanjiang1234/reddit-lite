
$(document).ready(function () {
    $('.downvote').css('background', 'red');

    $('.upvote').on('click', function() {
        var id = $(this).data("id");
        $.post( "/posts/vote", {id: id, voteType: 1} )
            .done(function(data) {
                $( ".result" ).html( data );
            });
    });

    $('.downvote').on('click', function() {
        var id = $(this).data("id");
        $.post( "/posts/vote", {id: id, voteType: 0} )
            .done(function(data) {
                $( ".result" ).html( data );
            });
    });
})
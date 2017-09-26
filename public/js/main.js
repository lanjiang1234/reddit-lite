
$(document).ready(function () {

    $('.up-vote').on('click', function() {
        vote($(this), 1);
    });

    $('.down-vote').on('click', function() {
        vote($(this), 0);
    });

    function vote($elem, voteType) {
        var id = $elem.data("id");
        $.post( "/posts/vote", {id: id, voteType: voteType} )
            .done(function(data) {
                var prop = voteType == 0 ? "downvote" : "upvote";
                $elem.text(data[prop]);
            });
    }
})
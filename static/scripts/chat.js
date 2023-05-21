$(document).ready(function() {
    $("#chatForm").submit(function(e) {
        e.preventDefault();
        var message = $("#messageInput").val();
        if (message.trim() !== "") {
            $("#responseContainer").append("<div class='user-message'><p class='message-text'>You: " + message + "</p></div>");
            $.get("/chat", { message: message }, function(response) {
                $("#responseContainer").append("<div class='bot-message'><p class='message-text'>BullyShield: " + response + "</p></div>");
                $("#messageInput").val("");
                $("#responseContainer").scrollTop($("#responseContainer")[0].scrollHeight);
            });
        }
    });
});

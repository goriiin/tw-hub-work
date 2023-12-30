function checkFollow() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            var response = JSON.parse(this.responseText);
            if (response.ans === "yes") {
                document.getElementById("followButton").style.backgroundColor = "#3e8e41";
            }
        }
    };
    xhttp.open("GET", "/ru/user/{id}/isfollow", true);
    xhttp.send();
}
function follow() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            var response = JSON.parse(this.responseText);
            if (response.ans === "yes") {
                document.getElementById("followButton").style.backgroundColor = "#3e8e41";
            } else {
                document.getElementById("followButton").style.backgroundColor = "#4CAF50";
            }
        }
    };
    xhttp.open("GET", "/user/{id}/follow", true);
    xhttp.send();
    document.getElementById("followButton").style.backgroundColor = "#3e8e41";
}




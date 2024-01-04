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




const newsList = document.querySelector('.news-list');

function addNews(authorId, author, description) {
    const newsItem = document.createElement('div');
    newsItem.className = 'news-item';

    const newsDescription = document.createElement('p');
    newsDescription.textContent = description;

    const newsAuthor = document.createElement('a'); /* Добавили элемент для автора */
    newsAuthor.textContent = author;
    newsAuthor.id = authorId
    newsAuthor.href = '/ru/user/' + authorId;

    // const likeButton = document.createElement('button'); /* Добавили кнопку "Нравится" */
    // likeButton.textContent = '❤️';
    // likeButton.className = 'heart-btn';
    // likeButton.onclick = 'changeColor()';

    // const commentButton = document.createElement('button'); /* Добавили кнопку "Комментарий" */
    // commentButton.textContent = 'Комментарий';
    newsItem.appendChild(newsAuthor);
    newsItem.appendChild(newsDescription);
    // newsItem.appendChild(likeButton);

    // newsItem.appendChild(commentButton);

    newsList.appendChild(newsItem);
}

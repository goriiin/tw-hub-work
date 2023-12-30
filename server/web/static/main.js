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

// Добавляем случайные новости для тестирования
// for (let i = 0; i < 20; i++) {
//     const author = `Автор новости ${i + 1}`; /* Добавили автора новости */
//     const description = `Краткое описание новости ${i + 1}`;
//     addNews(author, description, i);
// }

function changeColor() {
    var btn = document.querySelector('.heart-btn');
    btn.classList.toggle('active');
}

function addOwnNews() {
    const text = document.getElementById('own-news').value;
    fetch('/ru/news', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ text: text })
    })
        .then(response => response.json())
        .then(data => console.log(data))
        .catch(error => console.error('Ошибка:', error));
}

function renderNews() {
    fetch('/news/render')
        .then(response => response.json())
        .then(data => {
            for (let i = 0; i < data.length; i++) {
                const news = data[i];
                addNews(news.userId, news.username, news.text);
                console.log(news.userId, news.username, news.text)
            }
        })
        .catch(error => console.error(error));
}


document.querySelector('form').addEventListener('submit', async function(event) {
    event.preventDefault();
    // Получаем значение поля ввода
    const search = document.querySelector('#search').value;
    // Отправляем запрос на сервер
    const response = await fetch('/search/' + search)
        .then(function(response) {
            return response.json();
        })
        .then(function(data) {
            console.log(data)
            // Очищаем результаты предыдущего поиска
            document.getElementById('results').innerHTML = '';
            // Выводим результаты поиска
            data.forEach(function(user) {
                var result = document.createElement('div');
                result.className = 'result';

                var img = document.createElement("img");
                img.src = "/web/static/isITQX_Q7OM.jpg";
                result.appendChild(img);

                var a = document.createElement("a");
                a.href = '/user/' + user.id;
                a.textContent = user.nick;
                result.appendChild(a);
                document.getElementById('results').appendChild(result);
            });
        });
});


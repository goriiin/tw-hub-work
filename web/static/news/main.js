const form = document.getElementById('add-news-form');
const newsList = document.querySelector('.news-list');

form.addEventListener('submit', function(event) {
    event.preventDefault();
    const titleInput = document.getElementById('title');
    const descriptionInput = document.getElementById('description');
    const authorInput = document.getElementById('author');

    addNews(titleInput.value, descriptionInput.value, authorInput.value);

    titleInput.value = '';
    descriptionInput.value = '';
    authorInput.value = '';
});

function addNews(title, description, author) {
    const newsItem = document.createElement('div');
    newsItem.className = 'news-item';

    const newsTitle = document.createElement('h2');
    newsTitle.textContent = title;

    const newsDescription = document.createElement('p');
    newsDescription.textContent = description;

    const newsAuthor = document.createElement('p'); /* Добавили элемент для автора */
    newsAuthor.textContent = 'Автор: ' + author;

    const likeButton = document.createElement('button'); /* Добавили кнопку "Нравится" */
    likeButton.textContent = 'Нравится';

    // const commentButton = document.createElement('button'); /* Добавили кнопку "Комментарий" */
    // commentButton.textContent = 'Комментарий';

    newsItem.appendChild(newsTitle);
    newsItem.appendChild(newsDescription);
    newsItem.appendChild(newsAuthor);
    newsItem.appendChild(likeButton);
    // newsItem.appendChild(commentButton);

    newsList.appendChild(newsItem);
}

// Добавляем случайные новости для тестирования
for (let i = 0; i < 100; i++) {
    const title = `Заголовок новости ${i + 1}`; /* Исправили ошибку со строками */
    const description = `Краткое описание новости ${i + 1}`;
    const author = `Автор новости ${i + 1}`; /* Добавили автора новости */
    addNews(title, description, author);
}
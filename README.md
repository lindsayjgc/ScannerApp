# Grocery Scanner

An application that enables users to search common groceries, quickly view ingredients they may be personally allergic to in each product, and add desired items to a convenient shopping list.

## Members

Front-end: Dylan Tallon, Sarah Patel<br>
Back-end: Lindsay Goldberg-Custer, Jordan Sheehan<br>

<br/>

### Running the Frontend Locally

---

**NOTE:** Always run the frontend using <code>$ npm start</code> instead of alternatives like <code>$ ng serve</code>. <code>$ npm start</code> is configured to include the proxy configuration that allows for API requests to be made to the same port that the frontend runs on.

### Running the Backend Locally

---

1.  Navigate into back-end/api
2.  <details>
    <summary>Create/update your local .env file - be sure to include all listed variables:</summary>

    > | name                  | value (do not wrap these in quotes)                                                                        |
    > | --------------------- | ---------------------------------------------------------------------------------------------------------- |
    > | `SECRET_KEY`          | use key generator to create your own                                                                       |
    > | `PORT`                | 9000                                                                                                       |
    > | `MAIL`                | cen3031groceryapp@gmail.com                                                                                |
    > | `PW`                  | hyvowpezafvisvws                                                                                           |
    > | `API_KEY`             | 3ZUwh4W1oWTjCsqkbe9Del7axRUyKG1XR4Y6KMUN                                                                   |
    > | `RECIPE_API_BASE_URL` | https://api.edamam.com/api/recipes/v2?type=public&app_id=fba7cec1&app_key=c25a3700d1ed1dd6026d19bc0939be30 |

    </details>

3.  Run <code>$ go build</code> to create an executable (you must build this locally because the file is large and all .exe are included in .gitignore)
4.  Run <code>$ ./ScannerApp</code> to start up the back-end

**NOTE:** Running this will output a message that the back-end is listening on port 9000. However, the proxy configuration of our frontend means that all requests to the API made from the Angular client should be made to the same URL the frontend is running on (e.g. <code>http://localhost:4200/api/signup</code>). Requests should only be made to port 9000 when making requests from Postman or similar applications.

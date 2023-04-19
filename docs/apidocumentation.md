# Scanner App API Documentation

### Running the Frontend Locally

---

**NOTE:** Always run the frontend using <code>$ npm start</code> instead of alternatives like <code>$ ng serve</code>. <code>$ npm start</code> is configured to include the proxy configuration that allows for API requests to be made to the same port that the frontend runs on.

<br/>

### Running the Backend Locally

---

1.  Navigate into back-end/api
2.  <details>
    <summary>Create/update your local .env file - be sure to include all listed variables:</summary>

    > | name         | value (do not wrap these in quotes)      |
    > | ------------ | ---------------------------------------- |
    > | `SECRET_KEY` | use key generator to create your own     |
    > | `PORT`       | 9000                                     |
    > | `MAIL`       | cen3031groceryapp@gmail.com              |
    > | `PW`         | hyvowpezafvisvws                         |
    > | `API_KEY`    | 3ZUwh4W1oWTjCsqkbe9Del7axRUyKG1XR4Y6KMUN |

    </details>

3.  Run <code>$ go build</code> to create an executable (you must build this locally because the file is large and all .exe are included in .gitignore)
4.  Run <code>$ ./ScannerApp</code> to start up the back-end

**NOTE:** Running this will output a message that the back-end is listening on port 9000. However, the proxy configuration of our frontend means that all requests to the API made from the Angular client should be made to the same URL the frontend is running on (e.g. <code>http://localhost:4200/api/signup</code>). Requests should only be made to port 9000 when making requests from Postman or similar applications.

<br/>

### User Auth/Creation/Deletion/Info

---

<details>
    <summary><code>POST</code> <code><b>/api/signup</b></code> <code>Adds user info and credentials to database</code></summary>

##### Parameters

> | name        | type     | data type | description |
> | ----------- | -------- | --------- | ----------- |
> | `firstname` | required | string    | N/A         |
> | `lastname`  | required | string    | N/A         |
> | `email`     | required | string    | N/A         |
> | `password`  | required | string    | N/A         |

##### Responses

> | http code | content-type       | response                                                  |
> | --------- | ------------------ | --------------------------------------------------------- |
> | `201`     | `application/json` | `{"message":"User successfully created"}`                 |
> | `400`     | `application/json` | `{"message":"All fields are required"}`                   |
> | `409`     | `application/json` | `{"message":"Email is already registered to an account"}` |
> | `500`     | `application/json` | `{"message":"Could not generate password hash"}`          |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`                  |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/login</b></code> <code>Authenticates user and saves cookie to be used by frontend</code></summary>

##### Parameters

> | name       | type     | data type | description |
> | ---------- | -------- | --------- | ----------- |
> | `email`    | required | string    | N/A         |
> | `password` | required | string    | N/A         |

##### Responses

> | http code | content-type       | response                                            |
> | --------- | ------------------ | --------------------------------------------------- |
> | `202`     | `application/json` | `{"message":"User successfully logged in"}`         |
> | `400`     | `application/json` | `{"message":"Email not registered to any account"}` |
> | `401`     | `application/json` | `{"message":"Incorrect password"}`                  |
> | `500`     | `application/json` | `{"message":"Error creating JWT"}`                  |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`            |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/logout</b></code> <code>Deletes cookie storing the logged in user's email, effectively logging user out</code></summary>

##### Parameters

> `none (the user currently logged in will be logged out)`

##### Responses

> | http code | content-type       | response                                                                             |
> | --------- | ------------------ | ------------------------------------------------------------------------------------ |
> | `200`     | `application/json` | `{"email": "*email that was logged out*", "message":"User successfully logged out"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                                    |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                           |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                    |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`                                              |

</details>

<details>
    <summary><code>GET</code> <code><b>/api/logged-in</b></code> <code>Checks whether any user is logged in and returns email if so</code></summary>

##### Parameters

> `none`

##### Responses

> | http code | content-type       | response                                                                         |
> | --------- | ------------------ | -------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"email":"*current email logged in*", "message":"User is currently logged in"}` |
> | `401`     | `application/json` | `{"message":"No user logged in"}`                                                |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`                                         |

</details>

<details>
    <summary><code>DELETE</code> <code><b>/api/delete-user</b></code> <code>Deletes all records of the logged in user</code></summary>

##### Parameters

> `none (the user currently logged in will be logged out)`

##### Responses

> | http code | content-type       | response                                                                               |
> | --------- | ------------------ | -------------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"email": "*email of user that was deleted*", "message":"User successfully deleted"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                                      |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                             |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                      |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`                                                |

</details>

<details>
    <summary><code>GET</code> <code><b>/api/user-info</b></code> <code>Retrieves specific user's information from database</code></summary>

##### Parameters

> `none`

##### Responses

> | http code | content-type       | response                                                                                                                                                                                     |
> | --------- | ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"firstname":"*user's first name*"`<br>`"lastname":"*user's last name*"`<br> `"email":"*user's email*"`<br> `"password":"*user's password*"`<br> `"allergies":"*comma delimited or NONE*"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                                                                                                                                            |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                                                                                                                                   |
> | `404`     | `application/json` | `{"message":"User Not Found"}`                                                                                                                                                               |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                                                                                                                            |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`                                                                                                                                                      |

</details>

### Email Verification for Signup/Password Reset

---

<details>
    <summary><code>POST</code> <code><b>/api/verify/signup</b></code> <code>Sends signup verification email with randomly generated code</code></summary>

##### Parameters

> | name    | type     | data type | description                                  |
> | ------- | -------- | --------- | -------------------------------------------- |
> | `email` | required | string    | email the user is attempting to sign up with |

##### Responses

> | http code | content-type       | response                                             |
> | --------- | ------------------ | ---------------------------------------------------- |
> | `200`     | `application/json` | `{"message":"Verification email sent successfully"}` |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`             |
> | `500`     | `application/json` | `{"message":"*email sending-related error*"}`        |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/verify/reset</b></code> <code>Sends password reset email with randomly generated code</code></summary>

##### Parameters

> | name    | type     | data type | description                                                 |
> | ------- | -------- | --------- | ----------------------------------------------------------- |
> | `email` | required | string    | email to the account the user is attempting to reset pw for |

##### Responses

> | http code | content-type       | response                                             |
> | --------- | ------------------ | ---------------------------------------------------- |
> | `200`     | `application/json` | `{"message":"Verification email sent successfully"}` |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`             |
> | `500`     | `application/json` | `{"message":"*email sending-related error*"}`        |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/check-code</b></code> <code>Checks code submitted by user against the code issued in verification email</code></summary>

##### Parameters

> | name    | type     | data type | description                                            |
> | ------- | -------- | --------- | ------------------------------------------------------ |
> | `code`  | required | string    | code submitted by user                                 |
> | `email` | required | string    | email to the account signing up or having its pw reset |

##### Responses

> | http code | content-type       | response                                                            |
> | --------- | ------------------ | ------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"isVerified": "true","message": "Email successfully verified"}`   |
> | `400`     | `application/json` | `{"message":"Email has not been issued a verification code"}`       |
> | `401`     | `application/json` | `{"isVerified": "false","message": "*wrong code or expired code*"}` |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`                            |
> | `500`     | `application/json` | `{"message":"*email sending-related error*"}`                       |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/reset-password</b></code> <code>Changes the user's password after being verified by the previous routes</code></summary>

##### Parameters

> | name       | type     | data type | description                    |
> | ---------- | -------- | --------- | ------------------------------ |
> | `email`    | required | string    | email of the existing account  |
> | `password` | required | string    | new password submitted by user |

##### Responses

> | http code | content-type       | response                                         |
> | --------- | ------------------ | ------------------------------------------------ |
> | `200`     | `application/json` | `{"message":"Password reset successfully"}`      |
> | `400`     | `application/json` | `{"message":"Email not found"}`                  |
> | `400`     | `application/json` | `{"message":"All fields are required"}`          |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`         |
> | `500`     | `application/json` | `{"message":"Could not generate password hash"}` |

</details>

### Allergies

---

<details>
    <summary><code>PUT</code> <code><b>/api/add-allergies</b></code> <code>Adds allergies to user's database information</code></summary>

##### Parameters

> | name        | type     | data type | description                                    |
> | ----------- | -------- | --------- | ---------------------------------------------- |
> | `allergies` | required | string    | allergies that are to be added to the database |

##### Responses

> | http code | content-type       | response                                                                              |
> | --------- | ------------------ | ------------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"addedAllergies":"*new allergies*", "existingAllergies":"*preexisting allergies*"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                                     |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                            |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                     |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`                                               |

</details>

<details>
    <summary><code>DELETE</code> <code><b>/api/delete-allergies</b></code> <code>Deletes allergies from user's database information</code></summary>

##### Parameters

> | name        | type     | data type | description                                        |
> | ----------- | -------- | --------- | -------------------------------------------------- |
> | `allergies` | required | string    | allergies that are to be deleted from the database |

##### Responses

> | http code | content-type       | response                                                                                                 |
> | --------- | ------------------ | -------------------------------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"deletedAllergies":"*allergies that existed*", "notDeletedAllergies":"*allergies that didn't exist*"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                                                        |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                                               |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                                        |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`                                                                  |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/check-allergies</b></code> <code>Check for user allergies present in product ingredients</code></summary>

##### Parameters

> | name          | type     | data type | description                                                                |
> | ------------- | -------- | --------- | -------------------------------------------------------------------------- |
> | `ingredients` | required | string    | product ingredients, comma-delimited without spaces (e.g. "milk,eggs,soy") |

##### Responses

> | http code | content-type       | response                                                                    |
> | --------- | ------------------ | --------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"allergies":"*allergies found in ingredients","allergiesPresent":"true"}` |
> | `200`     | `application/json` | `{"allergiesPresent":"false"}`                                              |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                           |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                  |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                           |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`                                     |
> | `500`     | `application/json` | `{"message":"Error searching for user allergies"}`                          |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`                                    |

</details>

### Grocery Lists

---

<details>
    <summary><code>PUT</code> <code><b>/api/create-list</b></code> <code>Creates an empty grocery list in database</code></summary>

##### Parameters

> | name    | type     | data type | description            |
> | ------- | -------- | --------- | ---------------------- |
> | `title` | required | string    | new grocery list title |

##### Responses

> | http code | content-type       | response                                   |
> | --------- | ------------------ | ------------------------------------------ |
> | `200`     | `application/json` | `{"list successfully created"}`            |
> | `400`     | `application/json` | `{"message":"No user logged in"}`          |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}` |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`          |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `   |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/add-list-items</b></code> <code>Adds items to an existing grocery list</code></summary>

##### Parameters

> | name    | type     | data type | description                           |
> | ------- | -------- | --------- | ------------------------------------- |
> | `title` | required | string    | grocery list title                    |
> | `items` | required | string    | new items to be added to grocery list |

##### Responses

> | http code | content-type       | response                                                                      |
> | --------- | ------------------ | ----------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"addedItems":"*new allergies*", "existingAllergies":"*preexisting items*"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                             |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                    |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                             |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `                                      |

</details>

<details>
    <summary><code>DELETE</code> <code><b>/api/delete-lists</b></code> <code>Deletes entire grocery lists</code></summary>

##### Parameters

> | name     | type     | data type | description                       |
> | -------- | -------- | --------- | --------------------------------- |
> | `titles` | required | string    | grocery list titles to be deleted |

##### Responses

> | http code | content-type       | response                                                                                 |
> | --------- | ------------------ | ---------------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"deletedLists":"*lists that existed*", "notDeletedLists":"*lists that didn't exist*"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                                        |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                               |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                        |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `                                                 |

</details>

<details>
    <summary><code>DELETE</code> <code><b>/api/delete-list-items</b></code> <code>Deletes specified grocery list items</code></summary>

##### Parameters

> | name    | type     | data type | description                                         |
> | ------- | -------- | --------- | --------------------------------------------------- |
> | `title` | required | string    | grocery list title from which items will be deleted |
> | `items` | required | string    | items to be deleted from grocery list               |

##### Responses

> | http code | content-type       | response                                                                                 |
> | --------- | ------------------ | ---------------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"deletedItems":"*lists that existed*", "notDeletedItems":"*lists that didn't exist*"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                                        |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                               |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                        |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `                                                 |

</details>

<details>
    <summary><code>GET</code> <code><b>/api/get-lists</b></code> <code>Provides a list of all grocery list titles</code></summary>

##### Parameters

> `none`

##### Responses

> | http code | content-type       | response                                   |
> | --------- | ------------------ | ------------------------------------------ |
> | `200`     | `application/json` | `{"titles":*all titles or NONE*"}`         |
> | `400`     | `application/json` | `{"message":"No user logged in"}`          |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}` |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`          |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`    |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/get-list</b></code> <code>Provides the contents of requested list</code></summary>

##### Parameters

> | name    | type     | data type | description                                 |
> | ------- | -------- | --------- | ------------------------------------------- |
> | `title` | required | string    | grocery list title for requested list items |

##### Responses

> | http code | content-type       | response                                                         |
> | --------- | ------------------ | ---------------------------------------------------------------- |
> | `200`     | `application/json` | `{"title":"*title*", "items":"*comma delimited items or NONE*"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                       |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `                         |

</details>

### Favorites

---

<details>
    <summary><code>GET</code> <code><b>/api/favorite</b></code> <code>Returns all favorited products of the logged-in user</code></summary>

##### Parameters

> `none`

##### Responses

> | http code | content-type       | response                                                                 |
> | --------- | ------------------ | ------------------------------------------------------------------------ |
> | `200`     | `application/json` | `{"favorite":"*product name*","code": "*code*","image": "*image link*"}` |
> | `204`     | `application/json` | `{"message":"No favorites found"}`                                       |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                        |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                               |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                        |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `                                 |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/favorite</b></code> <code>Adds favorite to the logged-in user's account</code></summary>

##### Parameters

> | name       | type     | data type | description                                          |
> | ---------- | -------- | --------- | ---------------------------------------------------- |
> | `favorite` | required | string    | name of the favorite product                         |
> | `code`     | required | string    | barcode of product from OpenFood API                 |
> | `image`    | required | string    | link to image thumbnail of product from OpenFood API |

##### Responses

> | http code | content-type       | response                                       |
> | --------- | ------------------ | ---------------------------------------------- |
> | `201`     | `application/json` | `{"message":"Product successfully favorited"}` |
> | `400`     | `application/json` | `{"message":"Product is already favorited"}`   |
> | `400`     | `application/json` | `{"message":"No user logged in"}`              |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`     |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`       |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`              |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `       |

</details>

<details>
    <summary><code>DELETE</code> <code><b>/api/favorite</b></code> <code>Deletes favorite from logged-in user's account</code></summary>

##### Parameters

> | name   | type     | data type | description                          |
> | ------ | -------- | --------- | ------------------------------------ |
> | `code` | required | string    | barcode of product from OpenFood API |

##### Responses

> | http code | content-type       | response                                            |
> | --------- | ------------------ | --------------------------------------------------- |
> | `200`     | `application/json` | `{"message":"Favorite successfully deleted"}`       |
> | `400`     | `application/json` | `{"message":"Product not found in user favorites"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                   |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`          |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`            |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                   |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `            |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/check-favorite</b></code> <code>Returns whether the product is a favorite of the logged-in user</code></summary>

##### Parameters

> | name   | type     | data type | description                          |
> | ------ | -------- | --------- | ------------------------------------ |
> | `code` | required | string    | barcode of product from OpenFood API |

##### Responses

> | http code | content-type       | response                                            |
> | --------- | ------------------ | --------------------------------------------------- |
> | `200`     | `application/json` | `{"code":"*code*","isFavorite": "*true or false*"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                   |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`          |
> | `500`     | `application/json` | `{"message":"Error decoding JSON body"}`            |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                   |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `            |

</details>

### Search/Similar

---

<details>
    <summary><code>POST</code> <code><b>/api/search</b></code> <code>Adds a query to the database for recommendation purposes</code></summary>

##### Parameters

> | name    | type     | data type | description        |
> | ------- | -------- | --------- | ------------------ |
> | `query` | required | string    | the search string  |

##### Responses

> | http code | content-type       | response                                   |
> | --------- | ------------------ | ------------------------------------------ |
> | `200`     | `application/json` | `{"message":"Query count updated"}`        |
> | `400`     | `application/json` | `{"message":"No user logged in"}`          |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}` |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`          |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `   |

</details>

<details>
    <summary><code>DELETE</code> <code><b>/api/search</b></code> <code>Deletes query from user's database information</code></summary>

##### Parameters

> | name        | type     | data type | description                   |
> | ----------- | -------- | --------- | ----------------------------- |
> | `query` | required | string    | query to delete from the database |

##### Responses

> | http code | content-type       | response                                                                                                 |
> | --------- | ------------------ | -------------------------------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"message": "Query successfully deleted"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                                                        |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                                               |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                                        |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`                                                                  |

</details>

<details>
    <summary><code>GET</code> <code><b>/api/search</b></code> <code>Provides user search history and how many times each search has been made</code></summary>

##### Parameters

> `none`

##### Responses

> | http code | content-type       | response                                   |
> | --------- | ------------------ | ------------------------------------------ |
> | `200`     | `application/json` | `[{"*query*":*count*}]`                    |
> | `400`     | `application/json` | `{"message":"No user logged in"}`          |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}` |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`          |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`    |

</details>

<details>
    <summary><code>POST</code> <code><b>/api/similar</b></code> <code>Gets top five similar foods</code></summary>

##### Parameters

> | name    | type     | data type | description        |
> | ------- | -------- | --------- | ------------------ |
> | `query` | required | string    | the search food    |

##### Responses

> | http code | content-type       | response                                   |
> | --------- | ------------------ | ------------------------------------------ |
> | `200`     | `application/json` | `{"*similarfood*":"*FDCID*"}`              |
> | `400`     | `application/json` | `{"message":"No user logged in"}`          |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}` |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`          |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"} `   |

</details>

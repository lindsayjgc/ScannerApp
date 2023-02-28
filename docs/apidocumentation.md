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

    > | name         | value (do not wrap these in quotes)  |
    > | ------------ | ------------------------------------ |
    > | `SECRET_KEY` | use key generator to create your own |
    > | `PORT`       | 9000                                 |

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

> | http code | content-type       | response                                                                                  |
> | --------- | ------------------ | ----------------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"deletedAllergies":"*new allergies*", "notDeletedAllergies":"*preexisting allergies*"}` |
> | `400`     | `application/json` | `{"message":"No user logged in"}`                                                         |
> | `400`     | `application/json` | `{"message":"Other cookie-related error"}`                                                |
> | `500`     | `application/json` | `{"message":"Error parsing JWT"}`                                                         |
> | `500`     | `application/json` | `{"message":"Other JWT-related error"}`                                                   |

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

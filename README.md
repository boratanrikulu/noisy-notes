<p align="center">
	<img src="logo.png" alt="Noisy Notes">
</p>

#

Noisy Notes lets you **keep your notes with noise.**

Each sent noise is converted to text by using google's speech-to-text api.  
In this way, you'll be able to search in noises.

For example, you have 100 noisy notes.  
And, you want to **find the noisy notes that includes speech** "Rabbit is cute".  
Noisy Notes gives you this ability.

> **Note:** This project only keeps the API.  
 [**Click**](https://github.com/batin/NoisyNotes) to see the frontend repo.

## TODOs

> **These project do not have any release yet.**  
**It is havily under development.**

- [x] Noise recognition.
- [x] Session create and delete endpoints.
- [x] CRUD for the noise.
- [x] Make a queue structure for recognition.
- [ ] Search endpoint.

## API Endpoints

- [**Signup**](#signup)
- [**Login**](#login)
- User
	- [**Me**](#me)
	- [**Create Noise**](#create-noise)
	- [**Get Noises**](#get-noises)
	- [**Get Noise**](#get-noise)
	- [**Get Noise's File**](#get-noises-file)
	- [**Logout**](#logout)

## Signup

> Creates an user account.

**URL:** /signup

**Request:**

- Type: **POST**
- Body: 
    - Form Data :
		- username (must) (at least 2 characters)
		- password (must) (at least 8 characters)
		- name (must)
		- surname (must)

**Response:**

- Type: **403**
	- That means the inputs are not valid.
- Type: **202**
	- That means the user is created.
	- Includes the created user.

**Example:**

Request:

```bash
curl --location --request POST 'localhost:3000/signup' \
	--form 'username=testing-user' \
	--form 'password=testing-pass' \
	--form 'name=Test' \
	--form 'surname=Account'
```

Response:

```json
{
    "Message": "Account is created.",
    "User": {
        "ID": 15,
        "CreatedAt": "2020-06-02T16:39:21.182619391+03:00",
        "UpdatedAt": "2020-06-02T16:39:21.182619391+03:00",
        "DeletedAt": null,
        "Name": "Test",
        "Surname": "Account",
        "Username": "testing-user",
        "Noises": null,
        "Tags": null
    }
}
```

## Login

> Creates a session for the user.

**URL:** /login

**Request:**

- Type: **POST**
- Body: 
    - Form Data :
		- username (must)
		- password (must)

**Response:**

- Type: **401**
	- That means the login info is wrong.
- Type: **403**
	- That means the inputs are not valid.
- Type: **202**
	- That means the session is created.
	- Includes the token value.

**Example:**

Request:

```bash
curl --location --request POST 'localhost:3000/login' \
	--form 'username=testing-user' \
	--form 'password=testing-pass'
```

Response:

```json
{
    "Token": "2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1",
    "TokenType": "Bearer",
    "ExpiresIn": 3600
}
```

## Me

> Returns the current user.

**URL:** /user/me

**Request:**

- Type: **GET**
- Header: 
    - Authorization: **Bearer** Token (must)

**Response:**

- Type: **401**
	- That means the token is not valid.
- Type: **200**
	- That means the login is okay.
	- Includes the current user.

**Example:**

Request:

```bash
curl --location --request GET 'localhost:3000/user/me' \
	--header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1'
```

Response:

```json
{
    "ID": 15,
    "CreatedAt": "2020-06-02T16:39:21.182619+03:00",
    "UpdatedAt": "2020-06-02T16:39:21.182619+03:00",
    "DeletedAt": null,
    "Name": "Test",
    "Surname": "Account",
    "Username": "testing-user",
    "Noises": null,
    "Tags": null
}
```

## Create Noise

> Creates a noise.  
Recognition is run in a queue. 

>So,  
It will be inactive firstly.  
When recognition is done, the noise will be activated.

**URL:** /user/noises

**Request:**

- Type: **POST**
- Header: 
    - Authorization: **Bearer** Token (must)
- Body:
	- Form Data :
		- title (must)
		- file (must) (audio/mpeg)

**Response:**

- Type: **401**
	- That means the token is not valid.
- Type: **403**
	- That means input are not valid.
- Type: **202**
	- That means the noisy note was create.
	- Includes the noisy notes.

**Example:**

Request:

```bash
curl --location --request POST 'localhost:3000/user/noises' \
	--header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1' \
	--form 'title=First noise note' \
	--form 'file=@/path/to/audio_test.mp3'
```

Response:

```json
{
    "ID": 51,
    "CreatedAt": "2020-06-02T16:53:00.500487332+03:00",
    "UpdatedAt": "2020-06-02T16:53:00.500487332+03:00",
    "DeletedAt": null,
    "Title": "First noise note",
    "Tags": null,
    "Text": "",
    "IsActive": false
}
```

## Get Noises

> Returns all noises as noise array.

**URL:** /user/noises

**Request:**

- Type: **GET**
- Header: 
    - Authorization: **Bearer** Token (must)

**Response:**

- Type: **401**
	- That means the token is not valid.
- Type: **200**
	- That means the noises were listed.
	- Includes a noise array.

**Example:**

Request:

```bash
curl --location --request GET 'localhost:3000/user/noises' \
	--header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1'
```

Response:

```json
[
    {
        "ID": 51,
        "CreatedAt": "2020-06-02T16:53:00.500487+03:00",
        "UpdatedAt": "2020-06-02T16:53:02.703716+03:00",
        "DeletedAt": null,
        "Title": "First noise note",
        "Tags": null,
        "Text": "Tavşan ile kuşun macerası",
        "IsActive": true
    }
]
```

## Get Noise

> Return a specific noise. 

**URL:** /user/noises/{id}

**Request:**

- Type: **GET**
- Header: 
    - Authorization: **Bearer** Token (must)
- Path:
	- ID

**Response:**

- Type: **401**
	- That means the token is not valid.
- Type: **403**
	- That means record was not found.
- Type: **200**
	- That means the noise were listed.
	- Includes a noise array.

**Example:**

Request:

```bash
curl --location --request GET 'localhost:3000/user/noises/51' \
	--header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1'
```

Response:

```json
{
    "ID": 51,
    "CreatedAt": "2020-06-02T16:53:00.500487+03:00",
    "UpdatedAt": "2020-06-02T16:53:02.703716+03:00",
    "DeletedAt": null,
    "Title": "First noise note",
    "Tags": null,
    "Text": "Tavşan ile kuşun macerası",
    "IsActive": true
}
```

## Get Noise's File

> Returns noise's file as audio/mpeg format.

**URL:** /user/noises/{id}/file

**Request:**

- Type: **GET**
- Header: 
    - Authorization: **Bearer** Token (must)
- Path:
	- ID

**Response:**

- Type: **401**
	- That means the token is not valid.
- Type: **403**
	- That means record was not found.
- Type: **200**
	- That means the noise were listed.
	- Includes a noise array.

**Example:**

Request:

```bash
curl --location --request GET 'localhost:3000/user/noises/51/file' \
	--header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1'
```

Response:

```json
file
```

## Logout

> Removes the current user's session.

**URL:** /user/logout

**Request:**

- Type: **POST**
- Header: 
    - Authorization: **Bearer** Token (must)

**Response:**

- Type: **401**
	- That means the token is not valid.
- Type: **202**
	- That means the session were removed.

**Example:**

Request:

```sh
curl --location --request POST 'localhost:3000/user/logout' \
--header 'Authorization: Bearer ce356970-ea02-4cb8-8291-112d61cef3aa'
```

Response:

```json
{
    "Message": "Sessions is removed."
}
```

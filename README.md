<p align="center">
	<img src="logo.png" alt="Noisy Notes">
</p>

#

[![Go Report Card](https://goreportcard.com/badge/github.com/boratanrikulu/noisy-notes)](https://goreportcard.com/report/github.com/boratanrikulu/noisy-notes)

Noisy Notes lets you **keep your notes with noise.**

Each sent noise is converted to text by using google's speech-to-text api.  
In this way, you'll be able to search in noises.

For example, you have 100 noisy notes.  
And, you want to **find the noisy notes that includes speech** "Rabbit is cute".  
Noisy Notes gives you this ability.

> **Note:** This project only keeps the API.  
 [**Click**](https://github.com/batin/NoisyNotes) to see the frontend repo.

## TODOs

> **This project do not have any release yet.**  
**It is havily under development.**

- [x] Noise recognition.
- [x] Session create and delete endpoints.
- [x] CRUD for the noise.
- [x] Make a queue structure for recognition.
- [x] Add tag option for noises.
- [x] Search endpoint.
- [ ] Tags endpoint that is returned noises for user's tag.

## API Endpoints

- [**Signup**](#signup)
- [**Login**](#login)
- User
	- [**Me**](#me)
	- [**Create Noise**](#create-noise)
	- [**Get Noises**](#get-noises)
	- [**Get Noise**](#get-noise)
	- [**Get Noise's File**](#get-noises-file)
    - [**Update Noise**](#update-noise)
    - [**Delete Noise**](#delete-noise)
	- [**Logout**](#logout)

## Signup

> Creates an user account.

**URL:** /signup

**Request:**

- Type: **POST**
- Body: 
	- Form Data :
		- username  
        > `must`  
        `at least 2 characters`
		- password  
        > `must`  
        `at least 8 characters`
		- name  
        > `must`  
		- surname  
        > `must`  

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
		- username  
        > `must`  
		- password  
        > `must`  

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
	- Authorization: **Bearer** Token  
    > `must`  

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
	- Authorization: **Bearer** Token  
    > `must` 
- Body:
	- Form Data :
		- title  
        > `must` 
		- file  
        > `must`  
        > Audio format may be one of them;  
        `audio/mpeg`, `audio/mp3`, `audio/ogg`, `audio/wav`, `audio/flac`, `audio/aac`
		- tags
        > `not must`  
        example: `Tag 1, Tag 2, Tag3`

**Response:**

- Type: **401**
	- That means the token is not valid.
- Type: **403**
	- That means input are not valid.
- Type: **202**
	- That means the noisy note was created.
	- Includes the noisy notes.

**Example:**

Request:

```bash
curl --location --request POST 'localhost:3000/user/noises' \
	--header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1' \
	--form 'title=First noise note' \
	--form 'tags=Tag 1, Tag 2, Tag 3' \
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
    "Tags": [
        {
            "ID": 1,
            "CreatedAt": "2020-06-04T01:55:19.545495012+03:00",
            "UpdatedAt": "2020-06-04T01:55:19.561174344+03:00",
            "DeletedAt": null,
            "Title": "Tag 1"
        },
        {
            "ID": 2,
            "CreatedAt": "2020-06-04T01:55:19.55087435+03:00",
            "UpdatedAt": "2020-06-04T01:55:19.562559461+03:00",
            "DeletedAt": null,
            "Title": "Tag 2"
        },
        {
            "ID": 3,
            "CreatedAt": "2020-06-04T01:55:19.556297971+03:00",
            "UpdatedAt": "2020-06-04T01:55:19.563764474+03:00",
            "DeletedAt": null,
            "Title": "Tag 3"
        }
    ],
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
	- Authorization: **Bearer** Token  
    > `must` 
- Params:
    - q
    > `not must`  
    > to search,  
    it checks `noises' titles`, `noises' texts` and `tags' titles`
    - take
    > `not must`  
    > to limit size to take,  
    default: `-1`
    - sort
    > `not must`  
    > to sort by updated_at, only allowed: `asc`, `desc`.  
    default: `desc`

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
    --header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1' \
    --form 'q=tavşan' \
    --form 'sort=asc' \
    --form 'take=2'
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
        "Tags": [
            {
                "ID": 1,
                "CreatedAt": "2020-06-04T01:55:19.545495012+03:00",
                "UpdatedAt": "2020-06-04T01:55:19.561174344+03:00",
                "DeletedAt": null,
                "Title": "Tag 1"
            },
            {
                "ID": 2,
                "CreatedAt": "2020-06-04T01:55:19.55087435+03:00",
                "UpdatedAt": "2020-06-04T01:55:19.562559461+03:00",
                "DeletedAt": null,
                "Title": "Tag 2"
            },
            {
                "ID": 3,
                "CreatedAt": "2020-06-04T01:55:19.556297971+03:00",
                "UpdatedAt": "2020-06-04T01:55:19.563764474+03:00",
                "DeletedAt": null,
                "Title": "Tag 3"
            }
        ],
        "Text": "Tavşan ile kuşun macerası",
        "IsActive": true
    },
    {
        "ID": 52,
        "CreatedAt": "2020-06-02T16:59:00.500487+03:00",
        "UpdatedAt": "2020-06-02T16:59:02.703716+03:00",
        "DeletedAt": null,
        "Title": "Second noise note",
        "Tags": [
            {
                "ID": 1,
                "CreatedAt": "2020-06-04T01:55:19.545495012+03:00",
                "UpdatedAt": "2020-06-04T01:55:19.561174344+03:00",
                "DeletedAt": null,
                "Title": "Tag 1"
            },
        ],
        "Text": "Tavşan tatlıdır",
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
	- Authorization: **Bearer** Token  
    > `must` 
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
    "Tags": [
        {
            "ID": 1,
            "CreatedAt": "2020-06-04T01:55:19.545495012+03:00",
            "UpdatedAt": "2020-06-04T01:55:19.561174344+03:00",
            "DeletedAt": null,
            "Title": "Tag 1"
        },
        {
            "ID": 2,
            "CreatedAt": "2020-06-04T01:55:19.55087435+03:00",
            "UpdatedAt": "2020-06-04T01:55:19.562559461+03:00",
            "DeletedAt": null,
            "Title": "Tag 2"
        },
        {
            "ID": 3,
            "CreatedAt": "2020-06-04T01:55:19.556297971+03:00",
            "UpdatedAt": "2020-06-04T01:55:19.563764474+03:00",
            "DeletedAt": null,
            "Title": "Tag 3"
        }
    ],
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
	- Authorization: **Bearer** Token  
    > `must` 
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

## Update Noise

> Updates the noise.
> It works just like [creating](#create-noise)

> Recognition with the file will be work on a queue,  
So,  
It will be inactive firstly.  
When recognition is done, the noise will be activated.

**URL:** /user/noises/{id}

**Request:**

- Type: **PUT**
- Header: 
    - Authorization: **Bearer** Token  
    > `must` 
- Path:
    - ID
- Body:
    - Form Data :
        - title  
        > `must` 
        - file  
        > `must`  
        > Audio format may be one of them;  
        `audio/mpeg`, `audio/mp3`, `audio/ogg`, `audio/wav`, `audio/flac`, `audio/aac`
        - tags
        > `not must`  
        example: `Tag 1, Tag 2, Tag3`

**Response:**

- Type: **401**
    - That means the token is not valid.
- Type: **403**
    - That means input are not valid.
- Type: **202**
    - That means the noisy note was updated.
    - Includes the new noisy notes.

**Example:**

Request:

```bash
curl --location --request PUT 'localhost:3000/user/noises' \
    --header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1' \
    --form 'title=A new title' \
    --form 'tags=Tag 4' \
    --form 'file=@/path/to/audio_test.mp3'
```

Response:

```json
{
    "ID": 51,
    "CreatedAt": "2020-06-02T16:53:00.500487332+03:00",
    "UpdatedAt": "2020-06-03T20:53:00.500487332+03:00",
    "DeletedAt": null,
    "Title": "A new title",
    "Tags": [
        {
            "ID": 4,
            "CreatedAt": "2020-06-03T20:53:0.556297971+03:00",
            "UpdatedAt": "2020-06-03T20:53:0.563764474+03:00",
            "DeletedAt": null,
            "Title": "Tag 4"
        }
    ],
    "Text": "Tavşan ile kuşun macerası",
    "IsActive": false
}
```

## Delete Noise

> Deletes the noise. 

**URL:** /user/noises/{id}

**Request:**

- Type: **DELETE**
- Header: 
    - Authorization: **Bearer** Token  
    > `must` 
- Path:
    - ID

**Response:**

- Type: **401**
    - That means the token is not valid.
- Type: **403**
    - That means record was not found.
- Type: **200**
    - That means the noise were deleted.

**Example:**

Request:

```bash
curl --location --request DELETE 'localhost:3000/user/noises/51' \
    --header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1'
```

Response:

```json
{
    "Message": "The noise is deleted."
}
```

## Logout

> Removes the current user's session.

**URL:** /user/logout

**Request:**

- Type: **POST**
- Header: 
	- Authorization: **Bearer** Token  
    > `must` 

**Response:**

- Type: **401**
	- That means the token is not valid.
- Type: **202**
	- That means the session were removed.

**Example:**

Request:

```sh
curl --location --request POST 'localhost:3000/user/logout' \
    --header 'Authorization: Bearer 2ec69760-ef38-4a47-af5c-0ef1a7e6ecf1'
```

Response:

```json
{
    "Message": "Sessions is removed."
}
```

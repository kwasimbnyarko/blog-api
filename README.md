# blog-api
An API for a blog service that allows users to create, view, update and delete posts without authorization.

No sign-ups required.
Create a post right away with a username and some text.

This api could be used to for a web app like [Telegraph](https://telegra.ph). (👈🏾 Click on the link for context if necessary)

This Project was built as a submission for a [Slightly Techie Network](https://slightlytechie.com/) Challenge
<br>


## API Documentation
Main URL: https://blog-api-kf6i.onrender.com/api/
<br>

### Create Post
End Point: /post <br>
Method: POST <br>
Request Body:<br>

```json 
{
    "username" : "mj",
    "title":"on beefs",
    "text" : "i took that personally"
}
```

Response:201 Created<br>
```json 
{
    "success" : "post created"
}
```
<br>

### View Post
End Point: /post/:postId <br>
Method: GET <br>
Example Request: /post/649c20f10f9ddfa97a24ef48<br>
Response:200 OK<br>
```json 
[
    {
        "Id": "649c20f10f9ddfa97a24ef48",
        "username": "mj",
        "title": "on beefs",
        "text": "i took that personally",
        "created_At": "2023-06-28T12:00:49Z",
        "updated_At": "2023-06-28T12:00:49Z",
        "post_id": "649c20f10f9ddfa97a24ef48"
    }
]
```
<br>

### View All Posts From A User
End Point: /post/user/:username <br>
Method: GET <br>
Example Request: /post/user/mj<br>
Response:200 OK<br>
```json 
[
    {
        "Id": "649c20f10f9ddfa97a24ef48",
        "username": "mj",
        "title": "on beefs",
        "text": "i took that personally",
        "created_At": "2023-06-28T12:00:49Z",
        "updated_At": "2023-06-28T12:00:49Z",
        "post_id": "649c20f10f9ddfa97a24ef48"
    },
    {
        "Id": "649c24ad0f9ddfa97a24ef4a",
        "username": "mj",
        "title": "on kids",
        "text": "f*ck 'em kids",
        "created_At": "2023-06-28T12:16:45Z",
        "updated_At": "2023-06-28T12:16:45Z",
        "post_id": "649c24ad0f9ddfa97a24ef4a"
    }
]
```
<br>

### View All Posts
End Point: /post <br>
Method: GET<br>
Response:200 OK<br>
```json 
[
    {
        "Id": "649c20f10f9ddfa97a24ef48",
        "username": "mj",
        "title": "on beefs",
        "text": "i took that personally",
        "created_At": "2023-06-28T12:00:49Z",
        "updated_At": "2023-06-28T12:00:49Z",
        "post_id": "649c20f10f9ddfa97a24ef48"
    },
    {
        "Id": "649b8651ac77266faa2fd853",
        "username": "ja",
        "title": "that interview",
        "text": "im fine in the west",
        "created_At": "2023-06-28T01:01:05Z",
        "updated_At": "2023-06-28T01:01:05Z",
        "post_id": "649b8651ac77266faa2fd853"
    },
    {
        "Id": "649c24ad0f9ddfa97a24ef4a",
        "username": "mj",
        "title": "on kids",
        "text": "f*ck 'em kids",
        "created_At": "2023-06-28T12:16:45Z",
        "updated_At": "2023-06-28T12:16:45Z",
        "post_id": "649c24ad0f9ddfa97a24ef4a"
    }
]
```
<br>

### Update Post
End Point: /post/:postId <br>
Method: PUT <br>
Example Request:/post/649b8632ac77266faa2fd851
```json 
{
    "username":"tidylearner",
    "title":"alomo gyata",
    "text":"me na me de me ho aba nea mo p3 biaa mo nfa ny3 me"
}
```
Set unrequired fields to ""<br>
An example for updating just title and text<br>
```json 
{
    "username":"",
    "title" : "RW",
    "text" : "real wai"
}
```
Response:200 OK<br>
```json 
{
    "success" : "post updated"
}
```
<br>



### Delete Post
End Point: /post/:postId <br>
Method: DELETE <br>
Example Request:/post/649b8632ac77266faa2fd851<br>
Response:200 OK<br>
```json 
{
    "success" : "post deleted"
}
```

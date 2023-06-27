# blog-api
An API for a blog service that allows users to create, view, update and delete posts without authorization.

Users are not required to create accounts first.
Create a post right away with a username and some text

## API Documentation
Main URL: https://blog-api-kf6i.onrender.com

### Create Post
End Point: /api/post <br>
Method: POST <br>
Request Body:<br>

```json 
{
    "username" : " ",
    "text" : " "
}
```

### View Post
End Point: /api/post/:postId <br>
Method: GET <br>

### View All Posts From A User
End Point: /api/post/user/:username <br>
Method: GET <br>

### View All Posts
End Point: /api/post <br>
Method: GET<br>

### Update Post
End Point: /api/post/:postId <br>
Method: PUT <br>
```json 
{
    "text" : " "
}
```
### Delete Post
End Point: /api/post/:postId <br>
Method: DELETE <br>

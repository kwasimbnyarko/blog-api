# blog-api
An API for a blog service that allows users to create, view, update and delete posts without authorization.

No sign ups required.
Create a post right away with a username and some text.

This api could be used to for a web app like [Telegraph](https://telegra.ph).👈🏾 Click on the link for context if neccessary

This Project was built as a submission for a [Slightly Techie Network](https://slightlytechie.com/) Challenge


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

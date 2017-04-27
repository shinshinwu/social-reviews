<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="//fonts.googleapis.com/css?family=Raleway:400,300,600" rel="stylesheet" type="text/css">
  <link rel="stylesheet" href="/static/css/normalize.css">
  <link rel="stylesheet" href="/static/css/skeleton.css">
</head>

<body>
  <header>
    <h1 class="logo">Welcome to Beego</h1>
    <div class="description">
      We are currently on the page {{.PageTitle}}
      <br />
      User Name: {{.UserName}}
      <br />
      Email: {{.Email}}
    </div>
  </header>
  <div class="backdrop"></div>

  <form action={{urlfor "MainController.CreateReview"}} accept-charset="UTF-8" method="post">
    <div class="row">
      <div class="six columns">
        <label for="exampleEmailInput">Your email</label>
        <input class="u-full-width" type="email" placeholder="youremail@host.com" id="user_email" name="user[email]">
      </div>
      <div class="six columns">
        <label for="exampleEmailInput">Username</label>
        <input class="u-full-width" type="text" id="user_username" name="user[username]">
      </div>
    </div>
    <div class="row">
      <div class="twelve columns">
      <label for="exampleMessage">Review Title</label>
      <input class="u-full-width" type="text" placeholder="Title" id="review_title" name="review[title]">
    </div>
    <div class="row">
      <div class="twelve columns">
        <label for="exampleMessage">Comment</label>
        <textarea class="u-full-width" placeholder="My review …" id="review_comment" name="review[comment]"></textarea>
      </div>
    </div>
    <input class="button-primary" type="submit" value="Submit">
  </form>

</body>
</html>

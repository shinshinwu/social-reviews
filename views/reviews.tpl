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

  <form>
    <div class="row">
      <div class="six columns">
        <label for="exampleEmailInput">Your email</label>
        <input class="u-full-width" type="email" placeholder="test@mailbox.com" id="exampleEmailInput">
      </div>
      <div class="six columns">
        <label for="exampleRecipientInput">Reason for contacting</label>
        <select class="u-full-width" id="exampleRecipientInput">
          <option value="Option 1">Questions</option>
          <option value="Option 2">Admiration</option>
          <option value="Option 3">Can I get your number?</option>
        </select>
      </div>
    </div>
    <label for="exampleMessage">Message</label>
    <textarea class="u-full-width" placeholder="Hi Dave …" id="exampleMessage"></textarea>
    <label class="example-send-yourself-copy">
      <input type="checkbox">
      <span class="label-body">Send a copy to yourself</span>
    </label>
    <input class="button-primary" type="submit" value="Submit">
  </form>

</body>
</html>

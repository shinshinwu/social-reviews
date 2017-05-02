<div class="docs-section">
  <h3>Log In</h3>
  <form action={{urlfor "SessionController.Create"}} method="post" id="login-form">
    <div class="row">
      <div class="six columns">
        <label>Email</label>
        <input class="u-full-width" type="email" placeholder="youremail@host.com" id="email" name="email">
      </div>
      <div class="six columns">
        <label>Password</label>
        <input class="u-full-width" type="password" id="password" name="password">
      </div>
    </div>
    <input class="button-primary" type="submit" value="Login">
  </form>
</div>

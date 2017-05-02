<div class="docs-section">
  <h3>Log In</h3>
  <form action={{urlfor "UsersController.Create"}} method="post" id="login-form">
    <div class="row">
      <div class="twelve columns">
        <label>Email</label>
        <input class="u-full-width" type="email" placeholder="youremail@host.com" id="email" name="email">
      </div>
      <div class="twelve columns">
        <label>Username</label>
        <input class="u-full-width" type="text" id="username" name="username">
      </div>
      <div class="twelve columns">
        <label>Password</label>
        <input class="u-full-width" type="password" id="password" name="password">
      </div>
    </div>
    <input class="button-primary" type="submit" value="Sign Up">
  </form>
</div>

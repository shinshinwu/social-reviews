<div class="docs-section">
  <h3>Post A Review</h3>
  <form action={{urlfor "MainController.CreateReview"}} method="post" id="review-form">
    <div class="row">
      <div class="six columns">
        <label for="exampleEmailInput">Your email</label>
        <input class="u-full-width" type="email" placeholder="youremail@host.com" id="email" name="user[email]">
      </div>
      <div class="six columns">
        <label for="exampleEmailInput">Username</label>
        <input class="u-full-width" type="text" id="username" name="user[username]">
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
</div>

<div class="docs-section">
  <h3>Reviews:</h3>
  <div id="reviews-panel">
    {{range $index, $review := .Reviews}}
      <div class="row">
        <div class="one column"></div>
        <div class="ten columns">
          <div class="review">
            <h5>{{$review.Title}}</h5>
            <p>{{$review.Comment}}</p>
            <small>- Posted by {{$review.User.UserName}}</small>
          </div>
        </div>
      </div>
    {{end}}
  </div>
</div>

<script type="text/javascript">
  $(document).ready(function() {
    var form = $('#review-form');
    form.submit(function(e) {
      e.preventDefault();
      $.ajax({
        type: "POST",
        url: {{urlfor "MainController.CreateReview"}},
        data: {
          username: $("#username").val(),
          email:    $("#email").val(),
          title:    $("#review_title").val(),
          comment:  $("#review_comment").val()
        } // The data passed to / by POST method
      }).done(function(response) {
        // reset all form fields
        form.find("input[type=text], input[type=email], textarea").val("");
        $("#reviews-panel").prepend(
          `<div class="row" style="padding-bottom:25px;">
            <div class="one column"></div>
            <div class="ten columns">
              <div class="review">
                <h5>${response.review.Title}</h5>
                <p>${response.review.Comment}</p>
                <small>- Posted by ${response.user.UserName}</small>
              </div>
            </div>
          </div>`
        )
      }).fail(function(error) {
        // do error handling here
      });
    });
  });
</script>

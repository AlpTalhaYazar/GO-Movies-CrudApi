<h1>This is my Firs Go CRUD API project</h1>

<br>
<p>In this project I've created an CRUD API with GoLang</p>
<p>Here is a Movie Struct and Director Struct</p>
<p>We define a global array of Movie struct as named movies</p>
<p>Inside our Endpoints we get and modify our movies array</p>
<br>
<p>With Get request to /movies endpoint we get all movies inside our array.</p>
<p>With Get request to /movies/{id} endpoint we get only movie with defined id in the querystring.</p>
<p>With Post request to /movies endpoint we create a movie which we defined as JSON in request body, and add it to our movies array.</p>
<p>With Put request to /movies/{id} endpoint we update the movie that have {id} with the movie we defined in request body.</p>
<p>With Delete request to /movies/{id} endpoint we delete the movie with ID equal to {id}.</p>
<br>
<p>I've also shared the postman collection inside my repo.</p>
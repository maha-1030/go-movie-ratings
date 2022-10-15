# go-movie-ratings
Rest API for movie details and ratings
1. Run the mysql server
   ```
   docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql
   ```
2. Enter into mysql container once it is UP
    ```
    docker exec -it mysql mysql -u root -p
    ```
    When asked for password enter the password as `password` 
3. After entering into mysql server, create movies database
   ```
   create database movies;
   ```
4. Clone this repo into your local system
   ```
   git clone git@github.com:maha-1030/go-movie-ratings.git
   ```
5. Change the working directory to cloned git repository directory
   ```
   cd absolute-path-to-go-movie-ratings-folder 
   ```
6. Run the api
   ```
   go run main.go
   ```
   or run the api executable directly by following command
   ```
   ./go-movie-ratings
   ```
   It will automatically creates the tables in database if they not exist
7. Use `populateData` flag to start api with fresh data existed in csv file in data folder
   ```
   go run main.go -populateData
   ```
   or
   ```
   ./go-movie-ratings -populateData
   ```
   It deletes all the existing data in the tables and inserts all the records in csv files in data folder
8. Test the api by importing the postman collection of movies. You can find the following apis in collection to test:
   * new-movie: http://localhost:8080/api/v1/new-movie
   ```
    curl --location --request POST 'localhost:8080/api/v1/new-movie' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "Movie": {
        "Tconst": "tt0000021",
        "TitleType": "short",
        "PrimaryTitle": "funny gang",
        "RuntimeMinutes": 15,
        "Genres": "comedy"
      },
      "Rating": {
        "Tconst": "tt0000021",
        "AverageRating": 8,
        "NumVotes": 99
      }
    }'
   ```
   * longest-duration-movies: http://localhost:8080/api/v1/longest-duration-movies
   ```
    curl --location --request GET 'localhost:8080/api/v1/longest-duration-movies' \
    --header 'Content-Type: application/json'
   ```

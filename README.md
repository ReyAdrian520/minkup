# **MinkUp**

MinkUp is an app that lets people meet up and casually converse over a drink. The app matches people based on their top 10 interests and drinks of choice. Users must be 21 and older to use the app, and profiles must be verified by a picture. Users can see each other's first name, interests, 500 character bio, and one picture.

## **Technologies Used**

The MinkUp app is built using the following technologies:

- Go (Golang) programming language
- PostgreSQL database
- Gorilla mux router
- GORM ORM
- HTML/CSS

## **Installation**

To install the app, please follow the steps below:

1. Install Go and PostgreSQL on your machine.
2. Clone the repository to your local machine using **`git clone https://github.com/ReyAdrian520/minkup.git`**.
3. Navigate to the project directory using the command line.
4. Create a new PostgreSQL database called "minkup".
5. Update the database connection details in the main.go file.
6. Seed the database by running **`go run main.go -seed`**.
7. Start the app by running **`go run main.go`**.

## **Usage**

Once the app is running, you can access the home page at **`http://localhost:8080`**. The home page displays a welcome message.

To create a new user, send a POST request to **`http://localhost:8080/users`** with the following JSON payload:

```
jsonCopy code
{
  "firstName": "John",
  "age": 25,
  "interests": ["Hiking", "Cooking", "Photography"],
  "bio": "I'm a software developer who loves the outdoors and trying new recipes.",
  "drink": "Beer",
  "picture": "https://example.com/john.jpg"
}

```

To get a list of all users, send a GET request to **`http://localhost:8080/users`**.

To get information about a specific user, send a GET request to **`http://localhost:8080/users/{id}`**, where **`{id}`** is the ID of the user.

To update a user's information, send a PUT request to **`http://localhost:8080/users/{id}`**, where **`{id}`** is the ID of the user, with the following JSON payload:

```
jsonCopy code
{
  "firstName": "John",
  "age": 27,
  "interests": ["Hiking", "Cooking", "Photography", "Travel"],
  "bio": "I'm a software developer who loves the outdoors, trying new recipes, and exploring new places.",
  "drink": "Wine",
  "picture": "https://example.com/john.jpg"
}

```

To delete a user, send a DELETE request to **`http://localhost:8080/users/{id}`**, where **`{id}`** is the ID of the user.

## **Credits**

This app was created by Adrian Reyes.
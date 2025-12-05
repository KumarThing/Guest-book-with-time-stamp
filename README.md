# Guest book with Timestamp - Go web app
This is a simple web application written in Go that lets users add guest names to a list.
Each guest entry is automatically saved with a timestamp, showing when it was added.

- The app demonstrates:
- Go’s built-in net/http web server
- HTML templates using html/template
- Serving static CSS files
- Form handling (POST requests)
- Dynamic rendering of lists
- Basic in-memory data storage

---

## Features

- Add a guest name using a simple form.
- Each added guest automatically receives a timestamp.
- Guests are displayed in a clean HTML list.
- CSS styling using /static/style.css.
- Logging on each page visit or POST action.
- In-memory storage (resets on server restart).

___

## Project Structure

```
project-folder/
│
├── main.go
├── go.mod
│
├── template/
│   └── index.html
│
└── static/
    └── style.css
```

___
##  How to Run the Project 
1. Clone or download the project files.
2. Open terminal inside the project folder.
3. Run the Go server: 

``` sh
go run .
```
4. Open your browser and visit: 

```
http://localhost:8080
```


___
 
 ## Example Output 

 ```
 Guest Book with Time Stamp

John Doe — 2025-02-23 14:52:11
Sarah Kim — 2025-02-23 14:55:02

```

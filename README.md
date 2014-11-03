GoCityGame
==============

GoCityGame is a city simulation game made by Patric Vormstein.

## Setup

 * Install **go** language
 * Set your **$GOPATH**
 * Create the folders **$GOPATH/src** and **$GOPATH/src/vormstein.eu**
 * Clone repository insinde **$GOPATH/src/vormstein.eu**
 * On Windows: Run *install_pkg.bat* for code completion
 * Run "*[go get gopkg.in/mgo.v2](https://labix.org/mgo)*" and "*[go get github.com/gorilla/mux](https://labix.org/mgo)*"
 
 
## Coding practices

 * Install [Sublime Text](http://www.sublimetext.com/) 2+ or the editor of your choice
 * For Sublime: Install [package manager](https://sublime.wbond.net/installation)
 * For Sublime: Press Ctrl + Shift + P (Win) and select "Install Package" and then **GoSublime**
 * For Sublime: If not yet done set **$GOROOT** and **$GOPATH**
 * Run the project in terminal with the command "*go run webapp.go*"
 * On Windows: Every time, when code completion needs to get updated run *install_pkg.bat*

## Structure

```
project/
├── app/
│   |  ^ "Contains config and database control"
│   ├── controller/
|   |      ^ "Contains all the controller logic"
│   ├── models/
|   |      ^ "Contains all data models"
|   └── views/
|          ^ "Contains all views and templating functions"
├── dev/
│   └── migrations/
|          ^ "Contains migrations for populating database and tests"
├── static/
│      ^ "Static content like js, img, css"
|
└── webapp.go "<< Main Entrance"
```
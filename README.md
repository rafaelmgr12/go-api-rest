<h1 align="center">
  <img alt="Logo" src="./images/download.png" alt="Go">
</h1>

<h1 align="center">Alura Personalities</h1>
<p align = "center"> An API to add Names and Stories from street names</p>


<p align="center">
  <a href="#-technology">Technology</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
    <a href="#-project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-run">How to Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>
</p>

<p align="center">
  <img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=8257E5&labelColor=000000">
</p>

## ✨ Technology

The Project was develop as using the following techs:
- [Go](https://go.dev/)
- [Postgres](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Gorilla Handlers](https://github.com/gorilla/handlers) 
- [Gorm](https://github.com/go-gorm/gorm)


## 💻 Project
Repository developing during the course *Go: desenvolvendo uma API Rest*. The points covered by the course is

- Create a web application from scratch with Go, following the main of conventions
- Third party Libraries
- Connect the application in a database
- Docker


###  📓 Requirements 
As educational project, and the purpose is use the concepts of CRUD, then we requirements of the our applications is to : 
1. List all the personalities
2. Search
3. Add a personality
4. Delete personality
5. Update a personality

As can be seen above the list of requirements is , in summary, a simple CRUD.

### Products Attributes

| Personality | Type   |
| ----------- | ------ |
| ID          | int    |
| Name        | string |
| Story       | string |



### End Points
The API has the following end points:

| Type         | url                       | Functionality        |
| ------------ | ------------------------- | -------------------- |
| ```get```    | "/api/personalities"      | List Personalities   |
| ```post```   | "/api/personalities"      | Create a personality |
| ```delete``` | "/api/personalities/{id}" | Delete a personality |
| ```post```   | "/api/personalities/{id}" | Update a personality |



## 🚀 How to Run

To run the this project 

- Clone the repo and access the directory;
- Edit and run the docker using `docker-compose -d`
- Init the instance in [`localhost: 8000`](http://localhost:8000) with `go run main.go`;



## 📄 License
The projects is under the MIT license. See the file [LICENSE](LICENSE) fore more details

---
## Author

Made with ♥ by Rafael 👋🏻


[![Linkedin Badge](https://img.shields.io/badge/-Rafael-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/tgmarinho/)](https://www.linkedin.com/in/rafael-mgr/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-red?style=flat-square&link=mailto:nelsonsantosaraujo@hotmail.com)](mailto:ribeirorafaelmatehus@gmail.com)

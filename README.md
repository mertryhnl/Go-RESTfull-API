# Go RESTfull API
 The backend is written in "Go lang" programming language. Frontend is written with "React". Basic CRUD and user listing operations were performed. Restful API was used.
Go Restful API 


In Go, RESTful API (Representational State Transferful Application Programming Interface) defines an API that represents resources by communicating over the HTTP protocol. RESTful API provides a structure where resources can be accessed through URIs with a unique identifier and operations can be performed on these resources using standard HTTP methods (GET, POST, PUT, DELETE, etc.).

The concept of "Standard Layout" or "Project Layout", a widely used project layout and file structure convention in Go, is often used to facilitate the organization and maintenance of projects. This layout was developed in Go to ensure consistency in large projects and to promote easy understanding within the team.
	Below is an overview of a typical project layout with Go's Standard Layout:
 
    myproject/
    |-- cmd/
    |   |-- myapp/
    |       |-- main.go
    |-- internal/
    |   |-- pkgutil/
    |       |-- util.go
    |-- pkg/
    |   |-- mypackage/
    |       |-- mypackage.go
    |-- api/
    |   |-- openapi.yaml
    |-- web/
    |   |-- static/
    |   |-- templates/
    |-- scripts/
    |-- build/
    |-- deployments/
    |-- configs/
    |-- test/
    |-- .gitignore
    |-- README.md
    |-- go.mod
    |-- go.sum

MY Project
	My project is a website that performs basic CRUD operations that will list all users. Go RESTful API and React & Nextjs are used in this project. Before moving on to the project codes, we should design the project according to the "Standard Layout" I mentioned above. This design will provide us convenience.

 You can see this design below:
 
     restApi/
     |-- cmd/
     |   |-- main.go
     |   |-- user.db
     |-- internal\storage
     |       |-- sqlite.go
     |   |-- pkg
     |       |-- util.go
     |-- pkg/
     |       |-- api
     	|--api.go
     |       |-- handler
     	|--handler.go
     |       |-- model
     	|--model.go
     |       |-- repository
     	|--repository.go
     |       |-- service
     	|--service.go
     |-- go.mod
     |-- go.sum

If we talk about this design, we should first mention the "cmd" folder. The "cmd" folder is a directory in Go projects that usually contains the main entry point of the application. This folder is where the main application of the project is located and usually contains the files containing the main function from which the application is started.
	main.go: This file contains the main entry point of the application. The main function is where the application is initialized and runs. This file usually contains the basic configuration and initialization of the application using other packages and modules.	
	The "Internal" folder is a directory in Go projects that usually contains code that belongs to the internal part of the project. This folder contains modules that are used only inside the project and not exposed to the outside world. The internal folder ensures that packages are only used by other packages in the same project.
	The "pkg" folder is a directory in Go projects that contains libraries and packages that are usually open to the outside world and can be used by other projects. The packages in this folder can be imported and used by other Go projects outside the project.
	If we examine the project files in general:
 
     restApi/: Home directory of the project.
     cmd/: The directory containing the main entry points of the application.
     main.go: The main entry point of the application, the main file that starts the program.
     user.db: SQLite or other database file that will be used to store user data.
     internal/: The directory that houses the internals of the project.
     storage/: Directory containing packages related to database operations and storage.
     sqlite.go: File that interacts with the SQLite database.
     pkg/: Directory containing the various packages in the project.
     util.go: Package containing general helper functions.
     api/: File that defines the basic structure of the API.
     api.go: The main configuration and routes of the API can be found here.
     handler/: File that handles HTTP requests.
     handler.go: HTTP handlers that handle API requests are defined here.
     model/: File used for data modeling.
     model.go: Defines the data models to be used in the application.
     repository/: The file that interacts with the database.
     repository.go: Access operations to the database are performed here.
     service/: File used for business logic services.
     service.go: Business logic services such as CRUD operations are defined here.
     go.mod and go.sum: Go module files that manage project dependencies.

This organization is a common structure for a basic Go project. Different responsibilities, such as the business logic of the API, HTTP handlers, database access and auxiliary functions, are kept in separate packages, creating a modular structure. This can improve project maintainability and make the code more manageable.

REPOSITORY

The file "repository.go" contains a repository structure in Golang language that serves as a database operations package. Let's analyze the code in this file.

    package repository
    
    import (
        "fillabs_intern_project/pkg/model"
        "fmt"
    
        "gorm.io/gorm"
    )
    
    // Repository struct represents the repository that interacts with the database.
    type Repository struct {
        db *gorm.DB
    }
    
    // NewRepository creates a new Repository instance with the given database connection.
    func NewRepository(db *gorm.DB) *Repository {
        return &Repository{
            db: db,
        }
    }

First, the repository package is defined and the necessary packages (gorm and the model package of your own project) are imported. The "gorm" library is a library that facilitates the use of "sql" and allows you to do a lot with little code. A structure (struct) called repository is defined and this structure contains a gorm.DB object to interact with the database. Also, the NewRepository function to create this repository structure takes a gorm.DB connection and returns a new Repository instance. After these operations are done, we can proceed to the user listing and CRUD operations that I need.

    // AddUser adds a new user to the database.
    func (repo *Repository) AddUser(user *model.User) (*model.User, error) {
        if err := repo.db.Create(&user).Error; err != nil {
            fmt.Println("Error occurred when saving user: " + err.Error())
        }
        return user, nil
    }

The "AddUser" function is used to add a new user to the database. This function adds the user via the "Create" method of the "gorm" library. If an error occurs, it prints the error to the console and returns a "nil" user and error status.

    // UpdateUser updates an existing user in the database.
    func (repo *Repository) UpdateUser(userID uint, newUser *model.UserUpdateRequest) (*model.User, error) {
        // Check if the user with the given ID exists in the database.
        err := repo.db.Where("id = ?", userID).Find(&model.User{}).Error
        if err != nil {
            return nil, err
        }
    
        // Update the user with the new data.
        if err := repo.db.Model(&model.User{}).Where("id = ?", userID).Updates(&newUser).Error; err != nil {
            fmt.Println("Error occurred when saving user: " + err.Error())
            return nil, err
        }
    
        // Retrieve and return the updated user.
        var user *model.User
        if err := repo.db.Where("id = ?", userID).Find(&user).Error; err != nil {
            return nil, err
        }
    
        return user, nil
    }


“The "UpdateUser" function is used to update an existing user. First, it checks if the user exists in the database. Then, the user is updated via gorm's Model and Updates methods. If an error occurs, it prints the error to the console and returns a "nil" user and error status.









     // DeleteUser deletes a user from the database.
     func (repo *Repository) DeleteUser(userID uint) error {
         // Soft delete (using Unscoped to include soft-deleted records) the user with the given ID.
         if err := repo.db.Unscoped().Where("id = ?", userID).Delete(&model.User{}).Error; err != nil {
             fmt.Println("Error occurred when deleting user: " + err.Error())
         }
         return nil
     }

The "DeleteUser" function is used to delete a specific user from the database. Using soft delete (soft-delete is a way of marking data instead of deleting it completely), it checks if the user exists in the database and then deletes the user using gorm's Unscoped and Delete methods. If an error occurs, it prints the error to the console.

        // GetUsers retrieves all users from the database.
        func (repo *Repository) GetUsers() (*[]model.User, error) {
            var users *[]model.User
            if err := repo.db.Find(&users).Error; err != nil {
                return nil, err
            }
            return users, nil
        }

The "GetUsers" function is used to retrieve all users in the database. All users are retrieved using gorm's Find method and if an error occurs, it prints the error to the console.

     // GetUserByID retrieves a user from the database based on the provided user ID.
     func (repo *Repository) GetUserByID(userID uint) (*model.User, error) {
         var user *model.User
         if err := repo.db.Where("id = ?", userID).First(&user).Error; err != nil {
             return nil, err
         }
         return user, nil
     }


The "GetUserByID" function is used to retrieve a specific user from the database by user ID. If an error occurs, it prints the error to the console and returns a "nil" user and error status.

The "repository.go" file defines a package of database operations that performs CRUD (Create, Read, Update, Delete) and user listing operations.


     
     MODEL
     package model
     
     import "gorm.io/gorm"
     
     type User struct {
         gorm.Model
         Name    string `json:"name"`
         Surname string `json:"surname"`
         Age     int    `json:"age"`
     }
     
     type UserUpdateRequest struct {
         Name    string `json:"name"`
         Surname string `json:"surname"`
         Age     int    `json:"age"`
     }

The file "model.go" contains a package of data structures and special requests for these data structures in the Golang language. First, the model package is defined and includes the Gorm library, which provides support for ORM (Object-Relational Mapping) operations.

A data structure (struct) called "User" is defined. This structure is extended with "gorm.Model" so that it has the basic model properties provided by Gorm (ID, CreatedAt, UpdatedAt, DeletedAt). In addition, the fields Name, Surname, and Age are defined, which are other attributes of the user. These fields are marked with tags (json:"...") corresponding to data in JSON format.

Another data structure called "UserUpdateRequest" is defined. The purpose of this structure is to represent the data structure to be used when making an HTTP request to update a user. The Name, Surname, and Age fields are marked with JSON tags corresponding to the updated values

These two data structures are typically used to represent and update users in the database. The first data structure (User) contains the basic model properties provided by Gorm, while the second data structure (UserUpdateRequest) represents the parameters to be used for update operations.




     	
     SERVİCE
     package service
     
     import (
         "errors"
         "fillabs_intern_project/pkg/model"
         "fillabs_intern_project/pkg/repository"
     )
     
     type Service struct {
         repository *repository.Repository
     }
     
     func NewService(repo *repository.Repository) *Service {
         return &Service{
             repository: repo,
         }
     }

First, the service package is defined and the required packages (model and repository packages of your own project) are imported. A structure (struct) called "Service" is defined. This structure contains a repository element that will be used to perform database operations. The "NewService" function creates a new instance of Service by taking a repository element and returns it.

     func (service Service) AddUser(user *model.User) (*model.User, error) {
         if user.Name == "" || user.Surname == "" || user.Age <= 0 {
             return nil, errors.New("user's field can't be empty")
         }
         user, err := service.repository.AddUser(user)
         if err != nil {
             return nil, err
         }
         return user, nil
     }


The "AddUser" function is used to add a user. The function checks the required fields of the incoming user (Name, Surname, and Age). If one of these fields is empty or the age is less than 0, an error is returned. Otherwise, it calls the repository's "AddUser" method to add the user.



     
     func (service Service) UpdateUser(userID uint, newUser *model.UserUpdateRequest) (*model.User, error) {
         if newUser.Name == "" || newUser.Surname == "" || newUser.Age <= 0 {
             return nil, errors.New("user's field can't be empty")
         }
         updatedUser, err := service.repository.UpdateUser(userID, newUser)
         if err != nil {
             return nil, err
         }
         return updatedUser, nil
     }

The "UpdateUser" function is used to update a user's information. It receives a UserUpdateRequest structure containing the ID of the user to be updated and the new data. Again, the required fields are checked and then the update is performed by calling the repository's "UpdateUser" method.

    func (service Service) DeleteUser(userID uint) error {
        if userID <= 0 {
            return errors.New("user id can't be equal or less then zero")
        }
        err := service.repository.DeleteUser(userID)
        if err != nil {
            return err
        }
        return nil
    }

The "DeleteUser" function is used to delete a user. The ID of the user to be deleted is checked and an error is returned if the ID is less than 0. Otherwise, the user is deleted by calling the "DeleteUser" method of the repository

    
    func (service Service) GetUsers() (*[]model.User, error) {
        users, err := service.repository.GetUsers()
        if err != nil {
            return nil, err
        }
        return users, nil
    }

The "GetUsers" function is used to retrieve all users. It retrieves all users by calling the "GetUsers" method of the repository and returns the error if an error occurs.
    
    func (service Service) GetUserByID(userID uint) (*model.User, error) {
        if userID <= 0 {
            return nil, errors.New("user id cant be equal or less then zero")
        }
        user, err := service.repository.GetUserByID(userID)
        if err != nil {
            return nil, err
        }
        return user, nil
    }

The "GetUserByID" function is used to retrieve a specific user. The user ID is checked and an error is returned if the ID is less than 0. Otherwise, it retrieves a specific user by calling the repository's "GetUserByID" method and returns the error if an error occurs.

    API
    package api
    
    import (
        "encoding/json"
        "fillabs_intern_project/pkg/model"
        "fillabs_intern_project/pkg/service"
        "net/http"
        "strconv"
    
        "github.com/gorilla/mux"
    )
    
    type Api struct {
        service *service.Service
    }
    
    func NewApi(service *service.Service) *Api {
        return &Api{
            service: service,
        }
    }

Define the "Api" package and import the required packages (json, model, service, net/http, strconv, github.com/gorilla/mux). A structure (struct) called "Api" is defined. This structure contains a service element that will be used to handle HTTP requests. The "NewApi" function takes a service element, creates a new instance of Api and returns it.



     func (api Api) AddUser(w http.ResponseWriter, r *http.Request) {
         var user *model.User
         err := json.NewDecoder(r.Body).Decode(&user)
         if err != nil {
             ReturnError(w, "error occured when decoding body "+err.Error(), http.StatusInternalServerError)
         }
         addUser, err := api.service.AddUser(user)
         if err != nil {
             ReturnError(w, "error occured when adding user "+err.Error(), http.StatusInternalServerError)
         }
         w.Header().Set("Content-Type", "application/json")
         err = json.NewEncoder(w).Encode(&addUser)
         if err != nil {
             ReturnError(w, "error occured when encoding user "+err.Error(), http.StatusInternalServerError)
         }
     }

The "AddUser" function processes the HTTP POST request to add a new user. It decodes the incoming JSON data, performs the necessary error checking, then adds the user by calling the service's "AddUser" method and returns the result in JSON format.


















    func (api Api) UpdateUser(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        user_id, exists := params["user_id"]
        if !exists {
            ReturnError(w, "User ID not found ", http.StatusNotFound)
            return
        }
        userID, err := strconv.Atoi(user_id)
        if err != nil {
            ReturnError(w, "User ID cant converting integer "+err.Error(), http.StatusBadRequest)
            return
        }
        if userID <= 0 {
            ReturnError(w, "User ID can't be equal or less then zero "+err.Error(), http.StatusBadRequest)
            return
        }
        var userUpdateRequest model.UserUpdateRequest
        err = json.NewDecoder(r.Body).Decode(&userUpdateRequest)
        if err != nil {
            ReturnError(w, "Update Request Can't Decoding "+err.Error(), http.StatusBadRequest)
            return
        }
        updatedUser, err := api.service.UpdateUser(uint(userID), &userUpdateRequest)
        if err != nil {
            ReturnError(w, "User Can't Be Updating "+err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(updatedUser)
    }

The "UpdateUser" function processes an HTTP PUT request to update the information of a specific user. It checks the required parameters, decodes the user update data, then performs the update operation by calling the service's "UpdateUser" method and returns the result in JSON format.








    
    func (api Api) DeleteUser(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        user_id, exists := params["user_id"]
        if !exists {
            ReturnError(w, "User ID not found ", http.StatusNotFound)
            return
        }
        userID, err := strconv.Atoi(user_id)
        if err != nil {
            ReturnError(w, "User ID cant converting integer "+err.Error(), http.StatusBadRequest)
            return
        }
        if userID <= 0 {
            ReturnError(w, "User ID can't be equal or less then zero "+err.Error(), http.StatusBadRequest)
            return
        }
        err = api.service.DeleteUser(uint(userID))
        if err != nil {
            ReturnError(w, "User Can't Be Deleted "+err.Error(), http.StatusBadRequest)
            return
        }
        successfullyDeleted := map[string]string{
            "message": "User Successfully Deleted",
        }
        byteMessage, err := json.Marshal(successfullyDeleted)
        if err != nil {
            ReturnError(w, "User Can't Be Deleted "+err.Error(), http.StatusBadRequest)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(byteMessage))
    }

The "DeleteUser" function processes an HTTP DELETE request to delete a specific user. It checks the required parameters, then deletes the user by calling the service's "DeleteUser" method and returns the result in JSON format.








    func (api Api) GetUsers(w http.ResponseWriter, r *http.Request) {
        users, err := api.service.GetUsers()
        if err != nil {
            ReturnError(w, "Users Can't Listing "+err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&users)
    }

The "GetUsers" function processes the HTTP GET request to list all users. It retrieves all users by calling service's "GetUsers" method and returns the result in JSON format.

    func (api Api) GetUserByID(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        user_id, exists := params["user_id"]
        if !exists {
            ReturnError(w, "User ID not found ", http.StatusNotFound)
            return
        }
        userID, err := strconv.Atoi(user_id)
        if err != nil {
            ReturnError(w, "User ID cant converting integer "+err.Error(), http.StatusBadRequest)
            return
        }
        if userID <= 0 {
            ReturnError(w, "User ID can't be equal or less then zero "+err.Error(), http.StatusBadRequest)
            return
        }
        user, err := api.service.GetUserByID(uint(userID))
        if err != nil {
            ReturnError(w, "Users Can't Get By ID "+err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&user)
    }


The "GetUserByID" function processes an HTTP GET request to fetch a specific user by ID. It checks the required parameters, then retrieves the user by calling the service's "GetUserByID" method and returns the result in JSON format.



    
    func ReturnError(w http.ResponseWriter, message string, statusCode int) {
        http.Error(w, message, statusCode)
    }

The "ReturnError" function is used to add an error message to the HTTP response. It sends an HTTP response containing a specific error message and status code to the client via the http.Error function. This function is used to reduce repetitive error checking codes.

HANDLER

    package handler
    
    import (
        "fillabs_intern_project/internal/storage"
        "fillabs_intern_project/pkg/api"
        "fillabs_intern_project/pkg/repository"
        "fillabs_intern_project/pkg/service"
    
        "github.com/gorilla/mux"
    )
    
    func InitHandler() (*mux.Router, error) {
        db, err := storage.OpenDB()
        if err != nil {
            return nil, err
        }
        repo := repository.NewRepository(db)
        service := service.NewService(repo)
        api := api.NewApi(service)
        router := mux.NewRouter()
        router.HandleFunc("/users", api.GetUsers).Methods("GET")
        router.HandleFunc("/user/{user_id:[0-9]+}", api.GetUserByID).Methods("GET")
        router.HandleFunc("/users", api.AddUser).Methods("POST")
        router.HandleFunc("/users/{user_id:[0-9]+}", api.UpdateUser).Methods("PUT")
        router.HandleFunc("/users/{user_id:[0-9]+}", api.DeleteUser).Methods("DELETE")
        return router, nil
    }

The "handler" package is defined and the required packages (internal/storage, pkg/api, pkg/repository, pkg/service, github.com/gorilla/mux) are imported. 

The "InitHandler" function is responsible for initializing all services and redirects within the application. First, the database connection is opened (using the storage.OpenDB() function).

Then the repository, service, and api objects are created. The repository object is created using the database connection with the NewRepository function. Then, the service object is created using the repository with the NewService function. Finally, the api object is created using the service with the NewApi function.

Finally, a mux.Router instance is created. On this router, the routing (handling) of the functions provided by the api object for specific URLs is defined. For example, a GET request to the URL "/users" is handled by the api.GetUsers function. Specific HTTP methods are allowed for each URL. The InitHandler function is completed by returning the created router and nil in case of an error. This router contains the main routing and service configuration that will be used when the application starts.

The "model.go" file contains the data structures used in the Go language. The User struct represents a basic user data model, while the UserUpdateRequest struct defines the parameters used during an HTTP request to update a user.

The file "repository.go" contains a repository struct that performs database operations. The repository struct uses the Gorm library to execute database operations. Basic CRUD operations such as adding, updating, deleting, listing and fetching users by ID are defined in this file.

The "service.go" file contains the services in the business logic layer. This file contains a Service struct that interacts with the repository to perform user operations. Basic CRUD operations are defined along with the necessary error checking and business logic rules.

The "api.go" file contains an API struct that handles HTTP requests. The api struct contains a service element for calling service functions. HTTP handler functions for basic CRUD operations are defined in this file.

The "handler.go" file contains a mux.Router that routes and processes HTTP requests. The InitHandler function initializes the router to be used at the start of the application and sets up the related services, repositories and API. Requests to the defined URLs are redirected to the functions provided by the api object.

When these files come together, a RESTful API project is created. model.go defines the data structures, repository.go handles database operations, service.go implements the business logic, api.go handles HTTP requests, and handler.go routes and handles all system operations.




REACT

	React is a JavaScript library developed by Facebook and used for building user interfaces. React is widely used in the development of modern web applications, especially Single Page Applications (SPA). In the picture below, you can see the project that I created with React and wrote the backend side with the Go language I mentioned above.

 


In this project, basic CRUD operations and user listing operations were done with React and Go. If we talk about React, the most important place to talk about for this project is the codes I wrote in my "App.js" file.








    “App.js”
    // src/App.js
    import React, { useState } from "react";
    import axios from "axios";
    
    import Form from "./Components/Form";
    import CustomButton from "./Components/Button";
    
    import Box from "@mui/material/Box";
    import Container from "@mui/material/Container";
    import Stack from "@mui/material/Stack";
    import Table from "@mui/material/Table";
    import TableBody from "@mui/material/TableBody";
    import TableCell from "@mui/material/TableCell";
    import TableContainer from "@mui/material/TableContainer";
    import TableHead from "@mui/material/TableHead";
    import TableRow from "@mui/material/TableRow";
    
    -We import React and useState from React.
    
    We import to make HTTP requests using the -axios library.
    
    -Import Form, CustomButton, Box, Container, Stack, Table, TableBody, TableCell, TableContainer, TableHead, and TableRow components
    
    State İşlemleri
    
    const [formData, setFormData] = useState({
        id: "",
        name: "",
        surname: "",
        age: "",
      });
    
      const [readData, setReadData] = useState([]);

-The state named -formData holds the form data entered by the user.
-The state named -readData holds the data received from the server.



    
Form Operations

    const handleInputChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
        console.log("Form Data:", formData);
      };

-When the user enters values in any input box on the form, it updates these values to the formData state.
-Prints the updated form data to the console with -console.log.
   
    const handleCreate = async () => {
       	********************
     };
   
     const handleRead = async () => {
   	********************
     };
   
     const handleUpdate = async () => {
       	********************
     };
   
     const handleDelete = async () => {
   	********************
     };


-The "handleCreate", "handleRead", "handleUpdate", and "handleDelete" functions add new data to the server, read data, update data, and delete data respectively.









User Interface

    return (
       <div className="App">
         <Container maxWidth="sm">
           ********************
           <TableContainer>
             *******************
               <TableBody>
                 {readData.map((item, index) => (
                   <TableRow key={index}>
                     <TableCell>{item.ID}</TableCell>
                     <TableCell>{item.name}</TableCell>
                     <TableCell>{item.surname}</TableCell>
                     <TableCell>{item.age}</TableCell>
                   </TableRow>
                 ))}
               </TableBody>
             </Table>
           </TableContainer>
         </Container>
       </div>
     );

-Container component renders a user interface with a form, buttons and a table.
-Table and related components render a table with data retrieved from the server.

This React application provides a user interface and performs CRUD (Create, Read, Update, Delete) operations by communicating with a server using the Axios library. The form component stores the data entered by the user in the formData state and the handleCreate function associated with the "Create" button is used to post this data to the server. Also, to retrieve data from the server, the "Read" button and the associated handleRead function display the available data in a table. The "Update" and "Delete" buttons are associated with the handleUpdate and handleDelete functions respectively, and allow the user to update or delete the selected data. The user interface is organized using Material-UI components and the data is displayed in a tabular format.

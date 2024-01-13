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

function App() {
  // State for form data
  const [formData, setFormData] = useState({
    id: "",
    name: "",
    surname: "",
    age: "",
  });

  // State to store data fetched from the server
  const [readData, setReadData] = useState([]);

  // Function to handle input changes in the form
  const handleInputChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
    console.log("Form Data:", formData);
  };

  // Function to handle creating a new user
  const handleCreate = async () => {
    try {
      const { name, surname, age } = formData;
      const ageInt = parseInt(age, 10);
      if (isNaN(ageInt)) {
        console.error("Invalid age. Please enter a valid integer for age.");
        return;
      }
      const requestData = { name, surname, age: ageInt };
      await axios.post("http://localhost:8080/users", requestData, {
        headers: {
          "Content-Type": "application/json",
        },
      });

      console.log("Form Data:", requestData);
      handleRead(); // Refresh the data after creating a new user
    } catch (error) {
      console.error("Create error:", error);
    }
  };

  // Function to handle reading data from the server
  const handleRead = async () => {
    try {
      const response = await axios.get("http://localhost:8080/users");
      console.log("Form Data:", formData);
      setReadData(response.data);
    } catch (error) {
      console.error("Read error:", error);
    }
  };

  // Function to handle updating user data
  const handleUpdate = async () => {
    try {
      const { id, name, surname, age } = formData;

      const ageInt = parseInt(age, 10);
      const idInt = parseInt(id, 10);

      if (isNaN(ageInt)) {
        console.error("Invalid age. Please enter a valid integer for age.");
        return;
      }

      if (isNaN(idInt) || idInt <= 0) {
        console.error(
          "Invalid ID. Please enter a valid positive integer for ID."
        );
        return;
      }

      const requestData = { id: idInt, name, surname, age: ageInt };
      await axios.put(`http://localhost:8080/users/${idInt}`, requestData, {
        headers: {
          "Content-Type": "application/json",
        },
      });
      handleRead(); // Refresh the data after updating
    } catch (error) {
      console.error("Update error:", error);
    }
  };

  // Function to handle deleting a user
  const handleDelete = async () => {
    try {
      await axios.delete(`http://localhost:8080/users/${formData.id}`);
      handleRead(); // Refresh the data after deleting
    } catch (error) {
      console.error("Delete error:", error);
    }
  };

  return (
    <div className="App">
      <Container maxWidth="sm">
        <Box>
          <Stack spacing={2}>
            {/* Form component for user input */}
            <Form formData={formData} onInputChange={handleInputChange} />

            {/* Buttons for CRUD operations */}
            <CustomButton
              onClick={handleCreate}
              label="Create"
              color="primary"
            />
            <CustomButton onClick={handleRead} label="Read" color="primary" />
            <CustomButton
              onClick={handleUpdate}
              label="Update"
              color="primary"
            />
            <CustomButton
              onClick={handleDelete}
              label="Delete"
              color="secondary"
            />
          </Stack>
        </Box>
        {/* Table to display user data */}
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                {/* Table headers */}
                <TableCell
                  style={{ backgroundColor: "#f5f5f5", fontWeight: "bold" }}
                >
                  ID
                </TableCell>
                <TableCell
                  style={{ backgroundColor: "#f5f5f5", fontWeight: "bold" }}
                >
                  Name
                </TableCell>
                <TableCell
                  style={{ backgroundColor: "#f5f5f5", fontWeight: "bold" }}
                >
                  Surname
                </TableCell>
                <TableCell
                  style={{ backgroundColor: "#f5f5f5", fontWeight: "bold" }}
                >
                  Age
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {/* Mapping through the data to display in the table */}
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
}

export default App;

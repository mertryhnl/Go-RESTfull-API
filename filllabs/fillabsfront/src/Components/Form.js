// src/components/Form.js
import React from "react";
import { Box, Stack, TextField } from "@mui/material/";

const Form = ({ onInputChange, formData }) => {
  const handleInputChange = (e) => {
    onInputChange(e);
  };

  return (
    <Box pt={5}>
      <form>
        <Stack spacing={2}>
          <TextField
            label="ID"
            name="id"
            value={formData.id}
            onChange={handleInputChange}
            fullWidth
          />
          <TextField
            label="Name"
            name="name"
            value={formData.name}
            onChange={handleInputChange}
            fullWidth
          />
          <TextField
            label="Surname"
            name="surname"
            value={formData.surname}
            onChange={handleInputChange}
            fullWidth
          />
          <TextField
            label="Age"
            name="age"
            value={formData.age}
            onChange={handleInputChange}
            fullWidth
          />
        </Stack>
      </form>
    </Box>
  );
};

export default Form;

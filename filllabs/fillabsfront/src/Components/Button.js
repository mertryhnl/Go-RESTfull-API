// src/components/Button.js
import React from "react";
import Button from "@mui/material/Button";

const CustomButton = ({ onClick, label, color }) => {
  return (
    <Button onClick={onClick} variant="contained" color={color || "default"}>
      {label}
    </Button>
  );
};

export default CustomButton;

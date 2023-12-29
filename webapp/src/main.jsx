import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";
import Auth from "./components/Auth/Auth";
import { ChakraProvider, extendTheme } from "@chakra-ui/react";

const colors = {
  brand: {
    900: "#1a365d",
    800: "#153e75",
    700: "#2a69ac",
    100: "#26394B",
  },
  sec: {
    100: "#FFC926",
  },
  bg: {
    100: "#f5f5f7",
  },
  footer: {
    100: "#e9cdb3",
  },
};
const Button = {
  variants: {
    solid: {
      _active: { outline: "none !important" },
      _focus: { boxShadow: "none" },
    },
  },
};
const theme = extendTheme({ colors, components: { Button } });

ReactDOM.createRoot(document.getElementById("root")).render(
  <BrowserRouter>
    <ChakraProvider theme={theme}>
      <Auth>
        <App />
      </Auth>
    </ChakraProvider>
  </BrowserRouter>,
);

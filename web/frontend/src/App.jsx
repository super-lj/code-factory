import React from "react";

import { ThemeProvider } from "@material-ui/core/styles";
import { CssBaseline } from "@material-ui/core";

import { ApolloProvider } from "@apollo/client";
import { client } from "./ApolloClient";

import appTheme from "./theme";
import Entry from "./Entry";

export default function App() {
  return (
    <ApolloProvider client={client}>
      <CssBaseline />
      <ThemeProvider theme={appTheme}>
        <Entry />
      </ThemeProvider>
    </ApolloProvider>
  );
}

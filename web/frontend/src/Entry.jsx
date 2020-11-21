import React from "react";

import { makeStyles } from "@material-ui/core/styles";
import Toolbar from "@material-ui/core/Toolbar";

import RepoDrawer from "./RepoDrawer";
import CIAppBar from "./CIAppBar";
import RepoDetail from "./RepoDetail";

const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(2),
  },
}));

export default function Entry() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <CIAppBar />
      <RepoDrawer />
      <main className={classes.content}>
        <Toolbar />
        <RepoDetail />
      </main>
    </div>
  );
}

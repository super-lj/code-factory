import React, { useState } from "react";

import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogContent from "@material-ui/core/DialogContent";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContentText from "@material-ui/core/DialogContentText";

const useStyles = makeStyles((theme) => ({
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  aboutButton: {
    marginLeft: theme.spacing(2),
    fontSize: "large",
  },
}));

export default function CIAppBar() {
  const classes = useStyles();
  const [aboutDialogOpen, setAboutDialogOpen] = useState(false);

  return (
    <div>
      <AppBar position="fixed" className={classes.appBar}>
        <Toolbar>
          <Typography variant="h6">CodeFactoryCI</Typography>
          <Button
            variant="contained"
            color="primary"
            disableElevation
            onClick={() => setAboutDialogOpen(true)}
            className={classes.aboutButton}
          >
            About
          </Button>
        </Toolbar>
      </AppBar>
      <Dialog open={aboutDialogOpen} onClose={() => setAboutDialogOpen(false)}>
        <DialogTitle>About</DialogTitle>
        <DialogContent>
          <DialogContentText>
            CodeFactoryCI is a light-weighted CI/CD system. It is a course
            project in the early stage of development, and subject to change
            over time.
          </DialogContentText>
          <DialogContentText>
            <b>Authors: Peixuan Li, Xingyou Ji</b>
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button
            size="large"
            onClick={() => setAboutDialogOpen(false)}
            color="primary"
          >
            OK
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

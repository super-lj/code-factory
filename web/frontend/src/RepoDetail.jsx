import React, { useEffect, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import Grid from "@material-ui/core/Grid";
import BookIcon from "@material-ui/icons/Book";
import Paper from "@material-ui/core/Paper";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";

import { useReactiveVar } from "@apollo/client";

import { repoNameVar } from "./ApolloClient";
import CurrentRun from "./Current/CurrentRun";
import Branches from "./Branches/Branches";
import BuildHistory from "./BuildHistory/BuildHistory";

const useStyles = makeStyles((theme) => ({
  cardContent: {
    display: "flex",
  },
  bookIcon: {
    paddingLeft: theme.spacing(1),
    paddingRight: theme.spacing(1),
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  repoTitle: {
    paddingLeft: theme.spacing(1),
    paddingRight: theme.spacing(1),
  },
}));

export default function RepoDetail() {
  const classes = useStyles();
  const repoName = useReactiveVar(repoNameVar);
  const [selectedTab, setSelectedTab] = useState(0);

  useEffect(() => {
    setSelectedTab(0);
  }, [repoName]);

  if (repoName === "") {
    return <></>;
  }

  const renderTab = (tab) => {
    switch (tab) {
      case 0:
        return <CurrentRun repoName={repoName} />;
      case 1:
        return <Branches repoName={repoName} />;
      case 2:
        return <BuildHistory repoName={repoName} />;
      default:
        return <></>;
    }
  };

  return (
    <Grid container direction="column" spacing={2}>
      <Grid item container xs>
        <Card>
          <CardContent className={classes.cardContent}>
            <div className={classes.bookIcon}>
              <BookIcon fontSize="large" />
            </div>
            <Typography variant="h4" className={classes.repoTitle}>
              {repoName}
            </Typography>
          </CardContent>
        </Card>
      </Grid>
      <Grid item container xs>
        <Paper>
          <Tabs
            value={selectedTab}
            indicatorColor="primary"
            textColor="primary"
            onChange={(_, newValue) => setSelectedTab(newValue)}
          >
            <Tab label="Current" />
            <Tab label="Branches" />
            <Tab label="Build History" />
          </Tabs>
        </Paper>
      </Grid>
      {renderTab(selectedTab)}
    </Grid>
  );
}

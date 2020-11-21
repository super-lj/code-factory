import React, { useEffect } from "react";

import Drawer from "@material-ui/core/Drawer";
import Toolbar from "@material-ui/core/Toolbar";
import { Grid, makeStyles } from "@material-ui/core";

import { useQuery, gql, useReactiveVar } from "@apollo/client";

import RepoInfoCard from "./RepoInfoCard";
import { repoNameVar } from "./ApolloClient";

const drawerWidth = 400;

const useStyle = makeStyles((theme) => ({
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  paper: {
    backgroundColor: theme.palette.background.default,
  },
  grid: {
    width: drawerWidth,
    padding: theme.spacing(2),
  },
}));

const GET_REPO_BASIC_INFO = gql`
  query GetRepoBasicInfo {
    repos {
      name
      branchesConnection(first: 1) {
        edges {
          node {
            runsConnection(first: 1) {
              edges {
                node {
                  num
                  startTimestamp
                  duration
                  status
                }
              }
            }
          }
        }
      }
    }
  }
`;

const renderRepoCards = ({ loading, error, data, selectedRepoName }) => {
  // check if data is valid and get all needed fields
  if (loading || error) {
    return <></>;
  }
  return data.repos.map((repo) => {
    const {
      name,
      branchesConnection: {
        edges: [
          {
            node: {
              runsConnection: {
                edges: [
                  {
                    node: { num, startTimestamp, duration, status },
                  },
                ],
              },
            },
          },
        ],
      },
    } = repo;
    return (
      <Grid item xs key={name}>
        <RepoInfoCard
          repoName={name}
          selected={selectedRepoName === name}
          status={status}
          runNum={num}
          duration={duration}
          startTimestamp={startTimestamp}
        />
      </Grid>
    );
  });
};

export default function RepoDrawer() {
  const classes = useStyle();
  const { loading, error, data } = useQuery(GET_REPO_BASIC_INFO, {
    pollInterval: 500,
  });
  const selectedRepoName = useReactiveVar(repoNameVar);

  useEffect(() => {
    if (repoNameVar() === "" && !loading && !error && data.repos.length > 0) {
      repoNameVar(data.repos[0].name);
    }
  });

  return (
    <Drawer
      classes={{ paper: classes.paper }}
      className={classes.drawer}
      variant="permanent"
    >
      <Toolbar />
      <Grid container className={classes.grid} spacing={2} direction="column">
        {renderRepoCards({ loading, error, data, selectedRepoName })}
      </Grid>
    </Drawer>
  );
}

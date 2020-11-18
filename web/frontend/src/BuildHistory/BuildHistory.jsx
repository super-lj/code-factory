import React from "react";
import Grid from "@material-ui/core/Grid";

import { useQuery, gql } from "@apollo/client";

import PropTypes from "prop-types";
import RunCard from "./RunCard";

const GET_REPO_RUNS = gql`
  query GetRepoRunsInfo($repoName: String) {
    repos(name: $repoName) {
      name
      runsConnection {
        edges {
          node {
            num
            startTimestamp
            duration
            status
            branch {
              name
            }
            commit {
              hash
              msg
              author
            }
          }
        }
      }
    }
  }
`;

export default function BuildHistory({ repoName }) {
  const { loading, error, data } = useQuery(GET_REPO_RUNS, {
    variables: { repoName },
    pollInterval: 500,
  });

  // check if there is valid data
  if (loading || error) {
    return <></>;
  }

  return (
    <Grid item container xs direction="column" spacing={2}>
      {data.repos[0].runsConnection.edges
        .slice()
        .sort((e1, e2) => e2.node.num - e1.node.num)
        .map((edge) => {
          const {
            node: {
              num: runNum,
              startTimestamp,
              duration,
              status,
              branch: { name: brName },
              commit: { hash: commitHash, msg: commitMsg, author },
            },
          } = edge;
          return (
            <Grid item container xs key={runNum}>
              <RunCard
                brName={brName}
                commitHash={commitHash}
                commitMsg={commitMsg}
                runNum={runNum}
                duration={duration}
                startTimestamp={startTimestamp}
                author={author}
                status={status}
              />
            </Grid>
          );
        })}
    </Grid>
  );
}

BuildHistory.propTypes = {
  repoName: PropTypes.string.isRequired,
};

import React from "react";
import Grid from "@material-ui/core/Grid";

import { useQuery, gql } from "@apollo/client";

import PropTypes from "prop-types";

import BranchCard from "./BranchCard";

const GET_REPO_BRANCHES_INFO = gql`
  query GetRepoBranchesInfo($repoName: String) {
    repos(name: $repoName) {
      name
      branchesConnection {
        edges {
          node {
            name
            commit {
              hash
              msg
              author
              runsConnection(first: 1) {
                edges {
                  node {
                    num
                    startTimestamp
                    status
                  }
                }
              }
            }
          }
        }
      }
    }
  }
`;

export default function Branches({ repoName }) {
  const { loading, error, data } = useQuery(GET_REPO_BRANCHES_INFO, {
    variables: { repoName },
    pollInterval: 500,
  });

  // check if there is valid data
  if (loading || error) {
    return <></>;
  }

  return (
    <Grid item container xs direction="column" spacing={2}>
      {data.repos[0].branchesConnection.edges.map((branchInfo) => {
        const {
          node: {
            name,
            commit: {
              hash,
              msg,
              author,
              runsConnection: {
                edges: [
                  {
                    node: { num, startTimestamp, status },
                  },
                ],
              },
            },
          },
        } = branchInfo;
        return (
          <Grid item container xs key={name}>
            <BranchCard
              brName={name}
              commitHash={hash}
              commitMsg={msg}
              runNum={num}
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

Branches.propTypes = {
  repoName: PropTypes.string.isRequired,
};

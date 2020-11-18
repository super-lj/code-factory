import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import Paper from "@material-ui/core/Paper";
import Typography from "@material-ui/core/Typography";
import Grid from "@material-ui/core/Grid";
import Divider from "@material-ui/core/Divider";
import CheckIcon from "@material-ui/icons/Check";
import CloseIcon from "@material-ui/icons/Close";
import UpdateIcon from "@material-ui/icons/Update";
import TodayIcon from "@material-ui/icons/Today";
import PeopleIcon from "@material-ui/icons/People";
import ScheduleIcon from "@material-ui/icons/Schedule";

import { SourceCommit, Pound } from "mdi-material-ui";

import { formatDuration, formatDistanceToNow } from "date-fns";

import PropTypes from "prop-types";

const useStyles = makeStyles((theme) => ({
  card: {
    display: "flex",
    width: "100%",
  },
  cardContent: {
    width: "100%",
  },
  statusInd: {
    backgroundColor: ({ status }) => {
      switch (status) {
        case "IN_PROGRESS":
          return theme.palette.warning.light;
        case "FAILED":
          return theme.palette.error.light;
        case "SUCCEED":
        default:
          return theme.palette.success.light;
      }
    },
    width: "5px",
    height: "100%",
  },
  branchNameLine: {
    display: "flex",
    color: ({ status }) => {
      switch (status) {
        case "IN_PROGRESS":
          return theme.palette.warning.light;
        case "FAILED":
          return theme.palette.error.light;
        case "SUCCEED":
        default:
          return theme.palette.success.light;
      }
    },
  },
  statusIcon: {
    marginTop: "4px",
  },
  branchName: {
    marginLeft: theme.spacing(2),
  },
  commitMsg: {
    marginLeft: theme.spacing(5),
  },
  runNum: {
    color: ({ status }) => {
      switch (status) {
        case "IN_PROGRESS":
          return theme.palette.warning.light;
        case "FAILED":
          return theme.palette.error.light;
        case "SUCCEED":
        default:
          return theme.palette.success.light;
      }
    },
  },
}));

export default function RunCard({
  brName,
  commitHash,
  commitMsg,
  runNum,
  duration,
  startTimestamp,
  author,
  status,
}) {
  const classes = useStyles({ status });

  return (
    <Card className={classes.card}>
      <Paper elevation={0} className={classes.statusInd} />
      <CardContent className={classes.cardContent}>
        <Grid container item xs>
          <Grid container item xs={4} direction="column">
            <Grid container item xs alignItems="center">
              <Grid
                container
                item
                xs
                className={classes.branchNameLine}
                alignItems="center"
                wrap="nowrap"
              >
                {((s) => {
                  switch (s) {
                    case "IN_PROGRESS":
                      return <UpdateIcon className={classes.statusIcon} />;
                    case "FAILED":
                      return <CloseIcon className={classes.statusIcon} />;
                    case "SUCCEED":
                    default:
                      return <CheckIcon className={classes.statusIcon} />;
                  }
                })(status)}
                <Typography variant="h5" className={classes.branchName}>
                  <b>{brName}</b>
                </Typography>
              </Grid>
              <Grid container item xs alignItems="center" wrap="nowrap">
                <PeopleIcon />
                <Typography variant="h6" className={classes.branchName}>
                  {author}
                </Typography>
              </Grid>
            </Grid>
            <Grid container item xs alignItems="center">
              <Typography variant="h6" className={classes.commitMsg}>
                {commitMsg}
              </Typography>
            </Grid>
          </Grid>
          <Grid container item xs={1}>
            <Divider variant="middle" orientation="vertical" flexItem />
          </Grid>
          <Grid container item xs>
            <Grid
              container
              item
              xs={6}
              alignItems="center"
              className={classes.runNum}
              wrap="nowrap"
            >
              <Pound />
              <Typography variant="h6" className={classes.branchName}>
                {runNum}
                <span>&nbsp;&nbsp;</span>
                {((s) => {
                  switch (s) {
                    case "IN_PROGRESS":
                      return "running";
                    case "FAILED":
                      return "failed";
                    case "SUCCEED":
                    default:
                      return "passed";
                  }
                })(status)}
              </Typography>
            </Grid>
            <Grid container item xs={6} alignItems="center" wrap="nowrap">
              <ScheduleIcon />
              <Typography variant="h6" className={classes.branchName}>
                {formatDuration({
                  hours: Math.floor(duration / 3600),
                  minutes: Math.floor(duration / 60) % 60,
                  seconds: duration % 60,
                })}
              </Typography>
            </Grid>
            <Grid container item xs={6} alignItems="center" wrap="nowrap">
              <SourceCommit />
              <Typography variant="h6" className={classes.branchName}>
                {commitHash}
              </Typography>
            </Grid>
            <Grid container item xs={6} alignItems="center" wrap="nowrap">
              <TodayIcon />
              <Typography variant="h6" className={classes.branchName}>
                {formatDistanceToNow(new Date(startTimestamp) * 1000, {
                  addSuffix: true,
                })}
              </Typography>
            </Grid>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

RunCard.propTypes = {
  brName: PropTypes.string.isRequired,
  commitHash: PropTypes.string.isRequired,
  commitMsg: PropTypes.string.isRequired,
  runNum: PropTypes.number.isRequired,
  duration: PropTypes.number.isRequired,
  startTimestamp: PropTypes.number.isRequired,
  author: PropTypes.string.isRequired,
  status: PropTypes.string.isRequired,
};

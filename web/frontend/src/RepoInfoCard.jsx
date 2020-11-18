import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import CardActionArea from "@material-ui/core/CardActionArea";
import Paper from "@material-ui/core/Paper";
import CheckIcon from "@material-ui/icons/Check";
import CloseIcon from "@material-ui/icons/Close";
import UpdateIcon from "@material-ui/icons/Update";
import TodayIcon from "@material-ui/icons/Today";
import ScheduleIcon from "@material-ui/icons/Schedule";
import Grid from "@material-ui/core/Grid";

import PropTypes from "prop-types";

import { formatDistanceToNow, formatDuration } from "date-fns";

import { repoNameVar } from "./ApolloClient";

const useStyles = makeStyles((theme) => ({
  card: {
    backgroundColor: ({ selected }) =>
      selected ? theme.palette.primary.main : theme.palette.common.white,
    color: ({ selected }) =>
      selected
        ? theme.palette.primary.contrastText
        : theme.palette.text.primary,
    display: "flex",
    height: "100%",
  },
  title: {
    fontSize: 14,
    color: ({ selected }) =>
      selected
        ? theme.palette.primary.contrastText
        : theme.palette.text.secondary,
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
  repoTitleLine: {
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
  repoTitle: {
    marginLeft: theme.spacing(2),
  },
  durLine: {
    display: "flex",
  },
}));

export default function RepoInfoCard({
  selected,
  repoName,
  status,
  runNum,
  duration,
  startTimestamp,
}) {
  const classes = useStyles({ selected, status });

  return (
    <Card className={classes.card}>
      <Paper elevation={0} className={classes.statusInd} />
      <CardActionArea onClick={() => repoNameVar(repoName)}>
        <CardContent className={classes.durLine}>
          <Grid container direction="column" spacing={1}>
            <Grid container item xs className={classes.repoTitleLine}>
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
              <Typography variant="h5" className={classes.repoTitle}>
                {repoName}
                <span>&nbsp;&nbsp;</span>#{runNum}
              </Typography>
            </Grid>
            <Grid container item xs>
              <ScheduleIcon />
              <Typography className={classes.repoTitle}>
                {formatDuration({
                  hours: Math.floor(duration / 3600),
                  minutes: Math.floor(duration / 60) % 60,
                  seconds: duration % 60,
                })}
              </Typography>
            </Grid>
            <Grid container item xs>
              <TodayIcon />
              <Typography className={classes.repoTitle}>
                {formatDistanceToNow(new Date(startTimestamp) * 1000, {
                  addSuffix: true,
                })}
              </Typography>
            </Grid>
          </Grid>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}

RepoInfoCard.propTypes = {
  repoName: PropTypes.string.isRequired,
  selected: PropTypes.bool.isRequired,
  status: PropTypes.string.isRequired,
  runNum: PropTypes.number.isRequired,
  duration: PropTypes.number.isRequired,
  startTimestamp: PropTypes.number.isRequired,
};

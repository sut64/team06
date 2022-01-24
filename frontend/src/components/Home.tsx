import React from "react";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import { Theme, makeStyles, createStyles } from "@material-ui/core";

const useStyles = makeStyles((theme: Theme) => 
  createStyles({
    root: { flexGrow: 1 },
    container: { marginTop: theme.spacing(2) },
    paper: { padding: theme.spacing(3) },

  })
);

export default function Home() {
  const classes = useStyles();
  return (
    <div>
      <Container className={classes.container}>
        <Paper className={classes.paper} elevation={1}>

        </Paper>
      </Container>
    </div>
  );
}
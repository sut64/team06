import React, { useState, useEffect, ChangeEvent } from "react";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import ButtonGroup from "@material-ui/core/ButtonGroup";
import FormControl from "@material-ui/core/FormControl";
import TextField from "@material-ui/core/TextField";
import Typography from "@material-ui/core/Typography";
import Autocomplete from "@material-ui/lab/Autocomplete";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TablePagination from "@material-ui/core/TablePagination";
import TableRow from "@material-ui/core/TableRow";
import AddIcon from "@material-ui/icons/Add";
import RemoveIcon from "@material-ui/icons/Remove";
import DeleteIcon from "@material-ui/icons/Delete";
import { MuiPickersUtilsProvider, DateTimePicker } from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import { createStyles, Theme, makeStyles } from "@material-ui/core";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import { MembersInterface } from "../../models/IUser";
import { PremiumMemberPeriodInterface } from "../../models/IPremiumMemberPeriod";
import { MemberClassInterface } from "../../models/IMemberClass";
import { PremiumMemberInterface } from "../../models/IPremiumMember";

import { UsersInterface } from "../../models/ISignIn";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: { flexGrow: 1 },
    container: { marginTop: theme.spacing(2) },
    paper: {
      padding: theme.spacing(3),
    },
    tableContainer: { maxHeight: 480 },
    tableHead: {
      "& .MuiTableCell-head": {
        color: "white",
        backgroundColor: "navy",
      },
    },
    typoHeader: {
      fontWeight: "bold"
    }
  })
);

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

function CreatePremiumMember() {
  const classes = useStyles();

  const [user, setUser] = useState<UsersInterface>();

  const [member, setMember] = useState<MembersInterface[]>([]);
  const [premiumMemberPeriod, setPremiumMemberPeriod] = useState<PremiumMemberPeriodInterface[]>([]);
  const [memberClass, setMemberClass] = useState<MemberClassInterface[]>([]);
  const [premiumMember, setPremiumMember] = useState<Partial<PremiumMemberInterface>>({
    MemberClassID: 0,
    PremiumMemberPeriodID: 0,
  });
  //const [orderItem, setOrderItem] = useState<Partial<PurchaseOrderItemsInterface>[]>([]);
  const [createTime, setCreateTime] = useState<Date | null>(new Date());

  // TABLE CUSTOMIZE
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(5);

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  // QUERY DATA FROM DATABASE
  const getUser = async () => {
    const getToken = localStorage.getItem("token");
    if (getToken) {
      setUser(JSON.parse(localStorage.getItem("user") || ""));
    }
  }
  const getPremiumMember = async () => {
    const apiUrl = "http://localhost:8080/premium_members";
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    }

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setPremiumMember(res.data);
        } else {
          console.log("else");
        }
      })
  }

  const getPremiumMemberPeriod = async () => {
    const apiUrl = "http://localhost:8080/premium_member_period";
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      }
    }

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setPremiumMemberPeriod(res.data);
        } else {
          console.log("else");
        }
      })
  }

  const getMemberClasss = async () => {
    const apiUrl = "http://localhost:8080/member_class";
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    }

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setMemberClass(res.data);
        } else {
          console.log("else");
        }
      })
  }

  // SUTMIT AND MESSAGE
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errorMsg, setErrorMsg] = useState<string>("");

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };


  const submit = () => {
    let data = {
      MemberID: user?.ID,
      MemberClassID: premiumMember.MemberClassID,
      PremiumMemberPeriodID: convertType(premiumMember.PremiumMemberPeriodID),
      CrateAt: createTime,
      Point: convertType(premiumMember.Point),
      PremiumMemberID: premiumMember.PremiumMemberID,
    }

    const apiUrl = "http://localhost:8080/premium_members";
    const requestPostOptions = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    }

    fetch(apiUrl, requestPostOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log(res.data);
          setError(false);
          setSuccess(true);
        } else {
          console.log(res.error);
          setError(true);
          setErrorMsg(res.error);
        }
      });
  }

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
    setErrorMsg("");
  };

  // HANDLE ON-CHANGE
  const handleChange = (event: ChangeEvent<{ name?: string; value: unknown }>) => {
    const name = event.target.name as keyof typeof premiumMember;
    setPremiumMember({ ...premiumMember, [name]: event.target.value });
  };

  const handleInputChange = (event: ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof premiumMember;
    const { value } = event.target;
    setPremiumMember({ ...premiumMember, [id]: value });
  };

  const handleCreateTime = (date: Date | null) => {
    setCreateTime(date);
  }


  useEffect(() => {
    getUser();
    getPremiumMember();
    getPremiumMemberPeriod();
    getMemberClasss();

  }, []);

  // DEBUG CONSOLE MONITOR
  console.log("user: ", user);
  console.log("class: ", memberClass);
  console.log("period: ", premiumMemberPeriod);

  return (
    <div>
      <Container className={classes.container} maxWidth="lg">
        <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="success">
            สมัครสำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="error">
            {errorMsg}
          </Alert>
        </Snackbar>
        <Paper className={classes.paper} style={{ backgroundColor: "#81c784", }}>
          <Grid container spacing={1}>
            {/* Head */}
            <Grid item xs={12}>
              <Typography component="h4" variant="h5">
                สมัครสมาชิกพรีเมียม
              </Typography>
            </Grid>

            {/* Left Column */}
            <Grid item xs={8}>
              <Paper elevation={1} className={classes.paper} style={{ marginRight: "-9.5rem" }}>
                <Grid container spacing={2}>
                  <Grid item xs={12}>
                    <FormControl fullWidth variant="outlined">
                      <p>Member Class</p>
                      <Select
                        value={premiumMember.MemberClassID}
                        inputProps={{ name: "MemberClassID" }}
                        onChange={handleChange}
                      >
                        <MenuItem value={0} key={0}>Select Class</MenuItem>
                        {memberClass.map((item: MemberClassInterface) => (
                          <MenuItem value={item.ID} key={item.ID}>{item.Name}</MenuItem>
                        ))}
                      </Select>
                    </FormControl>
                  </Grid>
                  <Grid item xs={12}>
                    <FormControl fullWidth variant="outlined">
                      <p>Period</p>
                      <Select
                        value={premiumMember.PremiumMemberPeriodID}
                        inputProps={{ name: "PremiumMemberPeriodID" }}
                        onChange={handleChange}
                      >
                        <MenuItem value={0} key={0}>Select Period</MenuItem>
                        {premiumMemberPeriod.map((item: PremiumMemberPeriodInterface) => (
                          <MenuItem value={item.ID} key={item.ID}>{item.Period}</MenuItem>
                        ))}
                      </Select>
                    </FormControl>
                  </Grid>
                  <Grid item xs={12}>
                    <FormControl fullWidth variant="outlined">
                      <p>Premium Member ID</p>
                      <TextField
                        fullWidth
                        id="outlined-basic"
                        variant="outlined"
                        placeholder="ID"
                        value={premiumMember.PremiumMemberID}
                        onChange={handleInputChange}
                        inputProps={{ name: "Name" }}
                      />
                      <p>Point</p>
                      <TextField
                        fullWidth
                        id="outlined-basic"
                        variant="outlined"
                        type="number"
                        placeholder="Point"
                        value={premiumMember.Point}
                        onChange={handleInputChange}
                        inputProps={{ name: "Point" }}
                      />
                    </FormControl>
                  </Grid>
                </Grid>
              </Paper>
            </Grid>

            {/* Bottom */}
            <Grid item xs={12} style={{ marginTop: ".5rem" }}>
              <Button
                color="primary"
                variant="contained"
                onClick={submit}
              >
                สมัครสมาชิกพรีเมียม
              </Button>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    </div>
  );
}
export default CreatePremiumMember;
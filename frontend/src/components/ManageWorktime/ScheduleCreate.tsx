import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
    createStyles,
    makeStyles,
    Theme,
    // ThemeProvider,
    // createTheme
} from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import FormControl from '@material-ui/core/FormControl';
// import NativeSelect from '@material-ui/core/NativeSelect';
import {
    Grid,
    Paper,
    Select,
    Button,
    Snackbar,
    Box,
    // Typography
} from "@material-ui/core";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert"
import TextField from '@material-ui/core/TextField';
import {
    MuiPickersUtilsProvider,
    DatePicker
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

// import { green } from "@material-ui/core/colors";
import AccountCircle from '@material-ui/icons/AccountCircle';
import AccessAlarmIcon from '@material-ui/icons/AccessAlarm';

import {
    DayInterface,
    WeeklyInterface,
    MonthInterface,
    WorkingTimeInterface,
    ManageWorkTimeInterface
} from '../../models/IManageWorkTime'
import { EmployeesInterface } from "../../models/IUser";


const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            flexGrow: 1,
        },
        container: {
            marginTop: theme.spacing(2),
        },
        table: {
            minWidth: 650,
        },
        tableSpace: {
            marginTop: 20,
        },
        formControl: {
            margin: theme.spacing(1),
            minWidth: 200,
        },
        selectEmpty: {
            marginTop: theme.spacing(2),
        },
        margin: {
            margin: theme.spacing(1),
        },
    })
);

function Alert(props: AlertProps) {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
}

// const theme = createTheme({
//     palette: {
//         primary: green,
//     },
// });

export default function ScheduleCreate() {
    const classes = useStyles();
    const [employee, setEmployee] = useState<EmployeesInterface[]>([]);
    const [day, setDay] = useState<DayInterface[]>([]);
    const [weekly, setWeekly] = useState<WeeklyInterface[]>([]);
    const [month, setMonth] = useState<MonthInterface[]>([]);
    const [workingTime, setWorkingTime] = useState<WorkingTimeInterface[]>([]);
    const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
    const [manageWorkTime, setManageWorkTime] = useState<Partial<ManageWorkTimeInterface>>({});

    const [openAlertSucess, setOpenAlertSucess] = useState(false);
    const [openAlertError, setOpenAlertError] = useState(false);
    const handleClose = () => {
        setOpenAlertSucess(false);
        setOpenAlertError(false);
    };

    const BaseURL = "http://127.0.0.1:8080";
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };


    const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>) => {
        console.log(event.target.name, event.target.value)
        const name = event.target.name as keyof typeof manageWorkTime;
        setManageWorkTime({
            ...manageWorkTime,
            [name]: event.target.value,
        });
        // console.log(manageWorkTime)
    };

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof manageWorkTime;
        const { value } = event.target;
        setManageWorkTime({ ...manageWorkTime, [id]: value });
        // console.log(manageWorkTime)
    };

    const handleDateChange = (date: Date | null) => {
        console.log(date)
        setSelectedDate(date);
    };

    const getEmployee = async () => {
        fetch(`${BaseURL}/manage/employee/all`, requestOptions)
            .then(response => response.json())
            .then(res => {
                if (res.data) {
                    console.log("[GET] Employee => ", res.data);
                    setEmployee(res.data);
                }
            })
            .catch(error => console.log('error', error));
    };
    const getDay = async () => {
        fetch(`${BaseURL}/manage/day/all`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log("[GET] Day => ", res.data)
                    setDay(res.data)
                } else {
                    console.log("FAIL!")
                }
            })
    };
    const getWeek = async () => {
        fetch(`${BaseURL}/manage/week/all`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log("[GET] Week => ", res.data)
                    setWeekly(res.data)
                }
            })
    };
    const getMonth = async () => {
        fetch(`${BaseURL}/manage/month/all`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log("[GET] Month => ", res.data)
                    setMonth(res.data)
                }
            })
    };
    const getWorkTime = async () => {
        fetch(`${BaseURL}/manage/work/all`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log("[GET] WorkingTime => ", res.data)
                    setWorkingTime(res.data)
                }
            })
    };

    useEffect(() => {
        getEmployee();
        getDay();
        getWeek();
        getMonth();
        getWorkTime();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    const submit = async () => {
        let data = {
            NameSchedule: manageWorkTime.NameSchedule,
            WorkingDate: selectedDate,
            TimeTotal: convertType(manageWorkTime.TimeTotal),
            EmployeeID: convertType(manageWorkTime.EmployeeID),
            DayID: convertType(manageWorkTime.DayID),
            WeeklyID: convertType(manageWorkTime.WeeklyID),
            MonthID: convertType(manageWorkTime.MonthID),
            WorkingTimeID: convertType(manageWorkTime.WorkingTimeID)
        }
        console.log(data)
        const requestOptionsPost = {
            method: "POST",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        };
        fetch(`${BaseURL}/manage/create`, requestOptionsPost)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log("บันทึกได้")
                    setOpenAlertSucess(true);
                } else {
                    console.log("บันทึกไม่ได้")
                    setOpenAlertError(true);
                }
            });
        // console.log("here!")
        // setOpenAlertError(true);
        // setTimeout(() => {
        //     setOpenAlertError(false);
        // }, 2000)
    }

    return (
        <div>
            <Container className={classes.container} maxWidth="md">
                <Snackbar open={openAlertSucess} anchorOrigin={{ vertical: "top", horizontal: "right" }} autoHideDuration={6000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="success">
                        บันทึกข้อมูลสำเร็จ
                    </Alert>
                </Snackbar>
                <Snackbar open={openAlertError} anchorOrigin={{ vertical: "top", horizontal: "right" }} autoHideDuration={6000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="error">
                        บันทึกข้อมูลไม่สำเร็จ
                    </Alert>
                </Snackbar>
                <h1 style={{ textAlign: "center" }}>จัดตารางงาน</h1>
                <Grid container spacing={5}>
                    <Grid item xs={12} container direction="row" justifyContent="flex-end" alignItems="center">
                        <Box display="flex">
                            <Box>
                                <Button
                                    component={RouterLink}
                                    to="/manager/manage-schedule/table"
                                    variant="contained"
                                    color="primary"
                                >
                                    ข้อมูลทั้งหมด
                                </Button>
                            </Box>
                        </Box>
                    </Grid>
                </Grid>
                <Paper elevation={2} className={classes.root}>
                    <Grid container spacing={5}>
                        <Grid item xs={12} container direction="row" justifyContent="center" alignContent="center">
                            <FormControl className={classes.formControl}>
                                {/* <InputLabel htmlFor="age-native-helper">Age</InputLabel> */}
                                <Select
                                    native
                                    value={manageWorkTime.EmployeeID}
                                    onChange={handleChange}
                                    inputProps={{
                                        name: 'EmployeeID'
                                    }}
                                >
                                    <option aria-label="None" value="">กรุณาเลือกชื่อพนักงาน</option>
                                    {employee.map((item: EmployeesInterface) => (
                                        <option value={item.ID} key={item.ID}>
                                            {item.UserDetail.FirstName}
                                        </option>
                                    ))}
                                </Select>
                                {/* <FormHelperText>Some important helper text</FormHelperText> */}
                            </FormControl>
                        </Grid>
                        <Grid container spacing={0} justifyContent="center">
                            <Grid item>
                                <FormControl className={classes.formControl}>
                                    {/* <InputLabel htmlFor="age-native-helper">Age</InputLabel> */}
                                    <Select
                                        native
                                        value={manageWorkTime.WeeklyID}
                                        onChange={handleChange}
                                        inputProps={{
                                            name: 'WeeklyID'
                                        }}
                                    >
                                        <option aria-label="None" value="">กรุณาเลือกสัปดาห์</option>
                                        {weekly.map((item: WeeklyInterface) => (
                                            <option value={item.ID} key={item.ID}>
                                                {item.WeekAt}
                                            </option>
                                        ))}
                                    </Select>
                                    {/* <FormHelperText>Some important helper text</FormHelperText> */}
                                </FormControl>
                            </Grid>
                            <Grid item>
                                <FormControl className={classes.formControl}>
                                    {/* <InputLabel htmlFor="age-native-helper">Age</InputLabel> */}
                                    <Select
                                        native
                                        value={manageWorkTime.MonthID}
                                        onChange={handleChange}
                                        inputProps={{
                                            name: 'MonthID'
                                        }}
                                    >
                                        <option aria-label="None" value="">กรุณาเลือกเดือน</option>
                                        {month.map((item: MonthInterface) => (
                                            <option value={item.ID} key={item.ID}>
                                                {item.MonthOfYear}
                                            </option>
                                        ))}
                                    </Select>
                                    {/* <FormHelperText>Some important helper text</FormHelperText> */}
                                </FormControl>
                            </Grid>
                            <Grid item>
                                <FormControl className={classes.formControl}>
                                    {/* <InputLabel htmlFor="age-native-helper">Age</InputLabel> */}
                                    <Select
                                        native
                                        value={manageWorkTime.DayID}
                                        onChange={handleChange}
                                        inputProps={{
                                            name: 'DayID'
                                        }}
                                    >
                                        <option aria-label="None" value="">กรุณาเลือกวัน</option>
                                        {day.map((item: DayInterface) => (
                                            <option value={item.ID} key={item.ID}>
                                                {item.DayOfWeek}
                                            </option>
                                        ))}
                                    </Select>
                                    {/* <FormHelperText>Some important helper text</FormHelperText> */}
                                </FormControl>
                            </Grid>
                        </Grid>
                        <Grid item xs={12} container direction="row" justifyContent="center" alignContent="center">
                            <FormControl className={classes.formControl}>
                                {/* <InputLabel htmlFor="age-native-helper">Age</InputLabel> */}
                                <Select
                                    native
                                    value={manageWorkTime.WorkingTimeID}
                                    onChange={handleChange}
                                    inputProps={{
                                        name: 'WorkingTimeID'
                                    }}
                                >
                                    <option aria-label="None" value="">กรุณาเลือกเวลาทำงาน</option>
                                    {workingTime.map((item: WorkingTimeInterface) => (
                                        <option value={item.ID} key={item.ID}>
                                            {item.TimeToTime}
                                        </option>
                                    ))}
                                </Select>
                                {/* <FormHelperText>Some important helper text</FormHelperText> */}
                            </FormControl>
                        </Grid>
                        <Grid container spacing={1} alignItems="flex-end" direction="row" justifyContent="center" alignContent="center">
                            <Grid item>
                                <AccountCircle />
                            </Grid>
                            <Grid item>
                                <TextField
                                    id="NameSchedule"
                                    name="NameSchedule"
                                    type="text"
                                    label="กรอกชื่อตารางงาน"
                                    value={manageWorkTime.NameSchedule || ""}
                                    onChange={handleInputChange}
                                />
                            </Grid>
                        </Grid>
                        <Grid container spacing={1} alignItems="flex-end" direction="row" justifyContent="center" alignContent="center">
                            <Grid item>
                                <AccessAlarmIcon />
                            </Grid>
                            <Grid item>
                                <TextField
                                    id="TimeTotal"
                                    name="TimeTotal"
                                    type="number"
                                    label="กรอกเวลางาน"
                                    value={manageWorkTime.TimeTotal || ""}
                                    onChange={handleInputChange}
                                />
                            </Grid>
                        </Grid>
                        <Grid item xs={12} container direction="row" justifyContent="center" alignContent="center">
                            <FormControl variant="outlined">
                                {/* <p>วันที่และเวลา</p> */}
                                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                                    <DatePicker
                                        style={{ justifyContent: "center" }}
                                        name="WorkingDate"
                                        value={selectedDate}
                                        onChange={handleDateChange}
                                        label="กรุณาเลือกวันที่และเวลา"
                                        minDate={new Date("2018-01-01T00:00")}
                                        format="yyyy/MM/dd"
                                    />
                                </MuiPickersUtilsProvider>
                            </FormControl>
                        </Grid>
                        <Grid container spacing={5} justifyContent="center" style={{ marginTop: "10px", marginBottom: "5px" }}>
                            {/* <Grid item>
                                <Button
                                    variant="contained"
                                    size="large"
                                ><b>กลับ</b></Button>
                            </Grid> */}
                            <Grid item>
                                <Button
                                    style={{ float: "right" }}
                                    onClick={submit}
                                    variant="contained"
                                    color="primary"
                                    size="large"
                                ><b>บันทึก</b></Button>
                            </Grid>
                        </Grid>
                    </Grid>
                </Paper>
            </Container>
        </div >
    );
}
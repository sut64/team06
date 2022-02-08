import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
    createStyles,
    makeStyles,
    Theme,
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
    MenuItem,
    InputLabel,
} from "@material-ui/core";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert"
import TextField from '@material-ui/core/TextField';
import {
    MuiPickersUtilsProvider,
    DatePicker
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

// import { green } from "@material-ui/core/colors";
import AccessAlarmIcon from '@material-ui/icons/AccessAlarm';
import TableChartIcon from '@material-ui/icons/TableChart';
import DehazeIcon from '@material-ui/icons/Dehaze';

import {
    DayInterface,
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

export default function ScheduleCreate() {
    const classes = useStyles();
    const [employee, setEmployee] = useState<EmployeesInterface[]>([]);
    const [day, setDay] = useState<DayInterface[]>([]);
    const [stepDay, setStepDay] = useState<number[]>([]);
    const [month, setMonth] = useState<MonthInterface[]>([]);
    const [workingTime, setWorkingTime] = useState<WorkingTimeInterface[]>([]);
    const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
    const [manageWorkTime, setManageWorkTime] = useState<Partial<ManageWorkTimeInterface>>({});
    const [allManageWorkTime, setAllManageWorkTime] = useState<ManageWorkTimeInterface[]>([]);
    const [allDayOfMonth, setAllDayOfMonth] = useState<number>(0);
    const [createdDayWorking, setCreatedDayWorking] = useState<number[]>([]);

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
    const totalWorkTime = (eId: any) => {
        const tempManage = allManageWorkTime
        let total = 0;
        for (const data of tempManage) {
            if (data.EmployeeID === eId) {
                console.log(data.TimeTotal)
                total = total + data.TimeTotal;
            }
        }
        manageWorkTime.TimeTotal = total
        // return total;
    }

    const addTimeTotal = () => {
        // manageWorkTime.TimeTotal = manageWorkTime.TimeTotal ? manageWorkTime.TimeTotal + 8 : 8
        manageWorkTime.TimeTotal = 8;
    }
    const getCreatedDay = (eId: any, mId: any) => {
        const saveDay: any = [];
        allManageWorkTime.map((item: ManageWorkTimeInterface) => {
            // console.log(manageWorkTime.MonthID)
            if (item.EmployeeID === eId && item.MonthID === mId) {
                saveDay.push(item.Day.DayNumber)
            }
        });
        setCreatedDayWorking(saveDay);
    }

    const getWeeksInMonth = (month: any) => {
        let date = new Date(),
            year = date.getFullYear(),
            tMonth = month - 1
        const weeks = [],
            firstDate = new Date(year, tMonth, 1),
            lastDate = new Date(year, tMonth + 1, 0),
            numDays = lastDate.getDate();
        setAllDayOfMonth(numDays);

        let offset = 0;
        if (firstDate.getDay() === 0) offset = 1;
        // else if (firstDate.getDay() === 6) offset = 2;

        let start = 1 + offset;
        let end: number;

        if (offset === 0) end = 6 - firstDate.getDay() + 1;
        else end = 6 - start + 1 + (offset * 2);

        while (start <= numDays) {
            let temp = start;
            for (let i = start; i <= end; i++) {
                weeks.push(temp);
                temp += 1;
            }
            start = end + 2;
            end = end + 7;
            end = start === 1 && end === 8 ? 1 : end;
            if (end > numDays) end = numDays;
        }

        // return weeks;
        // console.log(weeks)
        setStepDay(weeks);
    }

    const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>) => {
        // console.log(event.target.name, event.target.value)
        const name = event.target.name as keyof typeof manageWorkTime;
        if (name === "EmployeeID") {
            totalWorkTime(event.target.value)
            getCreatedDay(event.target.value, manageWorkTime.MonthID);
        } else if (name === "WorkingTimeID") {
            addTimeTotal();
        } else if (name === "MonthID") {
            manageWorkTime.DayID = manageWorkTime.MonthID !== event.target.value ? 0 : manageWorkTime.DayID;
            getWeeksInMonth(event.target.value);
            getCreatedDay(manageWorkTime.EmployeeID, event.target.value);
        }
        setManageWorkTime({
            ...manageWorkTime,
            [name]: event.target.value,
        });
    };

    const handleInputChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
        const id = event.target.id as keyof typeof manageWorkTime;
        const { value } = event.target;
        setManageWorkTime({ ...manageWorkTime, [id]: value });
    };

    const handleDateChange = (date: Date | null) => {
        // console.log(date)
        setSelectedDate(date);
    };


    const getEmployee = async () => {
        fetch(`${BaseURL}/manage/employee/all`, requestOptions)
            .then(response => response.json())
            .then(res => {
                if (res.data) {
                    // console.log("[GET] Employee => ", res.data);
                    setEmployee(res.data);
                } else {
                    console.log("FAIL!")
                }
            })
            .catch(error => console.log('error', error));
    };
    const getDay = async () => {
        fetch(`${BaseURL}/manage/day/all`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    // console.log("[GET] Day => ", res.data)
                    setDay(res.data)
                } else {
                    console.log("FAIL!")
                }
            })
    };
    const getMonth = async () => {
        fetch(`${BaseURL}/manage/month/all`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    // console.log("[GET] Month => ", res.data)
                    setMonth(res.data)
                } else {
                    console.log("FAIL!")
                }
            })
    };
    const getWorkTime = async () => {
        fetch(`${BaseURL}/manage/work/all`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    // console.log("[GET] WorkingTime => ", res.data)
                    setWorkingTime(res.data)
                } else {
                    console.log("FAIL!")
                }
            })
    };

    const getAllManageWorkTime = async () => {
        fetch(`${BaseURL}/manage/all`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    // console.log("[GET] All ManageWorkingTime => ", res.data)
                    setAllManageWorkTime(res.data)
                    getCreatedDay(manageWorkTime.EmployeeID, manageWorkTime.MonthID);
                } else {
                    console.log("FAIL!")
                }
            })
    }

    useEffect(() => {
        getEmployee();
        getDay();
        getMonth();
        getWorkTime();
        getAllManageWorkTime();
    }, []);

    const minDateSelect = () => {
        let date = new Date();
        return date.setDate(date.getDate() - 1);
    }
    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    const submit = async () => {
        let data = {
            Comment: manageWorkTime.Comment,
            WorkingDate: selectedDate,
            TimeTotal: convertType(manageWorkTime.TimeTotal),
            EmployeeID: convertType(manageWorkTime.EmployeeID),
            DayID: convertType(manageWorkTime.DayID),
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
                    getAllManageWorkTime();
                    manageWorkTime.DayID = 0;
                    manageWorkTime.MonthID = 0;
                    setOpenAlertSucess(true);
                } else {
                    console.log("บันทึกไม่ได้")
                    setOpenAlertError(true);
                }
            });
    }

    return (
        <div>
            <Container className={classes.container} maxWidth="md">
                <Snackbar open={openAlertSucess} anchorOrigin={{ vertical: "top", horizontal: "right" }} autoHideDuration={5000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="success">
                        บันทึกข้อมูลสำเร็จ
                    </Alert>
                </Snackbar>
                <Snackbar open={openAlertError} anchorOrigin={{ vertical: "top", horizontal: "right" }} autoHideDuration={5000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="error">
                        บันทึกข้อมูลไม่สำเร็จ
                    </Alert>
                </Snackbar>
                <h1 style={{ textAlign: "center" }}>จัดตารางงาน</h1>
                <Paper elevation={2} className={classes.root}>
                    <Grid container spacing={2}>
                        <Grid item xs={12} style={{ marginTop: "5px", marginRight: "20px" }} container direction="row" justifyContent="flex-end" alignItems="center">
                            <Button
                                component={RouterLink}
                                to="/manager/manage-schedule/table"
                                variant="contained"
                                color="primary"
                                startIcon={<TableChartIcon />}
                            >
                                ข้อมูลทั้งหมด
                            </Button>
                        </Grid>
                    </Grid>
                    <Grid container spacing={5}>
                        <Grid item xs={12} container direction="row" justifyContent="center" alignContent="center">
                            <FormControl required className={classes.formControl}>
                                <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                                    ชื่อพนักงาน
                                </InputLabel>
                                <Select
                                    // native
                                    labelId="demo-simple-select-placeholder-label-label"
                                    value={manageWorkTime.EmployeeID || ""}
                                    onChange={handleChange}
                                    displayEmpty
                                    className={classes.selectEmpty}
                                    inputProps={{
                                        name: 'EmployeeID',
                                    }}
                                >
                                    <MenuItem value="" disabled><em>กรุณาเลือกชื่อพนักงาน</em></MenuItem>
                                    {employee.map((item: EmployeesInterface) => (
                                        <MenuItem value={item.ID} key={item.ID}>
                                            {item.UserDetail.FirstName + " " + item.UserDetail.LastName}
                                        </MenuItem>
                                    ))}
                                </Select>
                                {/* <FormHelperText>Some important helper text</FormHelperText> */}
                            </FormControl>
                        </Grid>
                        <Grid container spacing={0} justifyContent="center">
                            <Grid item>
                                <FormControl className={classes.formControl}>
                                    <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                                        เดือน
                                    </InputLabel>
                                    <Select
                                        // native
                                        value={manageWorkTime.MonthID || ""}
                                        onChange={handleChange}
                                        displayEmpty
                                        className={classes.selectEmpty}
                                        inputProps={{
                                            name: 'MonthID'
                                        }}
                                    >
                                        <MenuItem value="" disabled><em>กรุณาเลือกเดือน</em></MenuItem>
                                        {month.map((item: MonthInterface) => (
                                            <MenuItem value={item.ID} key={item.ID}>
                                                {item.MonthOfYear}
                                            </MenuItem>
                                        ))}
                                    </Select>
                                    {/* <FormHelperText>Some important helper text</FormHelperText> */}
                                </FormControl>
                            </Grid>
                            <Grid item>
                                <FormControl className={classes.formControl} disabled={manageWorkTime.MonthID ? false : true}>
                                    <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                                        วัน
                                    </InputLabel>
                                    <Select
                                        // native
                                        value={manageWorkTime.DayID || ""}
                                        onChange={handleChange}
                                        displayEmpty
                                        className={classes.selectEmpty}
                                        inputProps={{
                                            name: 'DayID'
                                        }}
                                    >
                                        <MenuItem value="" disabled><em>กรุณาเลือกวัน</em></MenuItem>
                                        {(() => {
                                            let items: any = [];
                                            let newDay = day.slice(0, allDayOfMonth)
                                            newDay.map((item: DayInterface) => (
                                                items.push(createdDayWorking.indexOf(parseInt(item.DayNumber)) < 0 ? (
                                                    <MenuItem value={item.ID} key={item.ID}>
                                                        {item.DayNumber}
                                                    </MenuItem>
                                                ) : (
                                                <MenuItem value={item.ID} key={item.ID} disabled>
                                                    {item.DayNumber} - สร้างตารางแล้ว
                                                </MenuItem>
                                                ))
                                            ))
                                            return items;
                                        })()}
                                    </Select>
                                    {/* <FormHelperText>Some important helper text</FormHelperText> */}
                                </FormControl>
                            </Grid>
                        </Grid>
                        <Grid item xs={12} container direction="row" justifyContent="center" alignContent="center">
                            <FormControl className={classes.formControl}>
                                <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                                    เวลาทำงาน
                                </InputLabel>
                                <Select
                                    // native
                                    value={manageWorkTime.WorkingTimeID || ""}
                                    onChange={handleChange}
                                    displayEmpty
                                    className={classes.selectEmpty}
                                    inputProps={{
                                        name: 'WorkingTimeID'
                                    }}
                                >
                                    <MenuItem value="" disabled><em>กรุณาเลือกเวลาทำงาน</em></MenuItem>
                                    {workingTime.map((item: WorkingTimeInterface) => (
                                        <MenuItem value={item.ID} key={item.ID}>
                                            {item.TimeToTime}
                                        </MenuItem>
                                    ))}
                                </Select>
                                {/* <FormHelperText>Some important helper text</FormHelperText> */}
                            </FormControl>
                        </Grid>
                        <Grid container spacing={3} alignItems="flex-end" direction="row" justifyContent="center" alignContent="center">
                            <Grid item>
                                <AccessAlarmIcon />
                            </Grid>
                            <Grid item>
                                <TextField
                                    id="TimeTotal"
                                    name="TimeTotal"
                                    // type="number"
                                    label="กรอกเวลางาน"
                                    value={manageWorkTime.TimeTotal || 0}
                                    onChange={handleInputChange}
                                    InputProps={{
                                        readOnly: true,
                                    }}
                                />
                            </Grid>
                        </Grid>
                        <Grid item xs={12} container direction="row" justifyContent="center" alignContent="center">
                            <FormControl variant="outlined">
                                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                                    <DatePicker
                                        style={{ justifyContent: "center" }}
                                        name="WorkingDate"
                                        value={selectedDate}
                                        onChange={handleDateChange}
                                        label="กรุณาเลือกวันที่และเวลา"
                                        minDate={minDateSelect}
                                        format="yyyy/MM/dd"
                                    />
                                </MuiPickersUtilsProvider>
                            </FormControl>
                        </Grid>
                        <Grid container spacing={3} alignItems="flex-end" direction="row" justifyContent="center" alignContent="center">
                            <Grid item>
                                <DehazeIcon />
                            </Grid>
                            <Grid item>
                                <TextField
                                    id="Comment"
                                    name="Comment"
                                    type="text"
                                    label="กรอกชื่อตารางงาน"
                                    value={manageWorkTime.Comment || ""}
                                    onChange={handleInputChange}
                                />
                            </Grid>
                        </Grid>
                        <Grid container spacing={5} justifyContent="center" style={{ marginTop: "10px", marginBottom: "5px" }}>
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
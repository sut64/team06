import React from "react";
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { ManageWorkTimeInterface } from "../../models/IManageWorkTime";
import { format } from 'date-fns';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        container: {
            marginTop: theme.spacing(2),
        },
        table: {
            minWidth: 650,
        },
        tableSpace: {
            marginTop: 20,
        },
    })
);
export default function ScheduleTable() {
    const classes = useStyles();
    const [manageWorkTime, setManageWorkTime] = useState<ManageWorkTimeInterface[]>([]);
    const BaseURL = "http://127.0.0.1:8080";
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    const getWatchVideos = async () => {
        fetch(`${BaseURL}/manage/all`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    setManageWorkTime(res.data);
                } else {
                    console.log("else");
                }
            })
            .catch(error => console.log('error', error));
    };

    useEffect(() => {
        getWatchVideos();
    }, []);

    return (
        <div>
            <Container className={classes.container} maxWidth="md">
                <Box display="flex">
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            ข้อมูลการจัดตารางงาน
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/manager/manage-schedule/create"
                            variant="contained"
                            color="primary"
                        >
                            สร้างข้อมูล
                        </Button>
                    </Box>
                </Box>
                <TableContainer component={Paper} className={classes.tableSpace}>
                    <Table stickyHeader className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center" width="5%">
                                    ลำดับ
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    ชื่อพนักงาน
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    ชื่อตาราง
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    สัปดาห์ที่
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    วัน
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    เดือน
                                </TableCell>
                                <TableCell align="center" width="60%">
                                    ช่วงเวลาทำงาน
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    เวลารวม
                                </TableCell>
                                <TableCell align="center" width="40%">
                                    วันทำงาน
                                </TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {manageWorkTime.map((item: ManageWorkTimeInterface) => (
                                <TableRow key={item.ID}>
                                    <TableCell align="center">{item.ID}</TableCell>
                                    <TableCell align="center">{item.Employee.UserDetail.FirstName}</TableCell>
                                    <TableCell align="center">{item.NameSchedule}</TableCell>
                                    <TableCell align="center">{item.Weekly.WeekAt}</TableCell>
                                    <TableCell align="center">{item.Day.DayOfWeek}</TableCell>
                                    <TableCell align="center">{item.Month.MonthOfYear}</TableCell>
                                    <TableCell align="center">{item.WorkingTime.TimeToTime}</TableCell>
                                    <TableCell align="center">{item.TimeTotal}</TableCell>
                                    <TableCell align="center">{format((new Date(item.WorkingDate)), 'dd-MMMM-yyyy')}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Container>
        </div>
    );
}
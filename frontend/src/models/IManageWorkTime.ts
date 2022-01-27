import { EmployeesInterface } from "./IUser";

export interface DayInterface {
    ID: number,
    DayOfWeek: string
}

export interface WeeklyInterface {
    ID: number,
    WeekAt: string
}

export interface MonthInterface {
    ID: number,
    MonthOfYear: string
}

export interface WorkingTimeInterface {
    ID: number,
    TimeToTime: string
}

export interface ManageWorkTimeInterface {
    ID: number,
    NameSchedule: string,
    WorkingDate: string,
    TimeTotal: number,

    ManagerID: number,
    Manager: EmployeesInterface,

    EmployeeID: number,
    Employee: EmployeesInterface,

    DayID: number,
    Day: DayInterface,

    WeeklyID: number,
    Weekly: WeeklyInterface,

    MonthID: number,
    Month: MonthInterface,
    
    WorkingTimeID: number,
    WorkingTime: WorkingTimeInterface
}
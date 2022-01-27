import { EmployeesInterface } from "./IUser";
export interface Productstockinterface {
    ID : number,
    Amount_remain:number,
    Update_datetime : Date,
    Detail:string,
    ProductID :number,
    Product :Productinterface,
    ShelfstoreID:number,
    Shelfstore:Shelfstoreinterface,
    EmployeeID:number,
    Employee:EmployeesInterface
}

export interface Productinterface {
    ID : number,
    Name : string,
    Price : number,
    TypeproductID : number,
    Typeproduct: Typeproductinterface,
}
export interface Typeproductinterface {
    ID : number,
    Name : string,
}
export interface Shelfstoreinterface {
    ID: number,
    Zone:string,
}
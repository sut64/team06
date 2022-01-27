import { ProductstocksInterface } from "./IProductstock";
import { PurchaseOrdersInterface } from "./IPurchaseOrder";

export interface PurchaseOrderItemsInterface {
    ID: number;

    OrderID: number;
    Order: PurchaseOrdersInterface;

    ProductstockID: number;
    Productstock: ProductstocksInterface;
    OrderAmount: number;
    ItemPrice: number;

    AmountPrice: number;
}
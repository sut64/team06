import { MembersInterface } from "./IUser";
import { ManagePromotionsInterface } from "./IManagePromotion";
import { PaymentMethodsInterface } from "./IPaymentMethod";
import { PurchaseOrderItemsInterface } from "./IPurchaseOrderItem";

export interface PurchaseOrdersInterface {
    ID: number;

    MemberID: number;
    Member: MembersInterface;

    PromotionID: number;
    Promotion: ManagePromotionsInterface;

    PaymentMethodID: number;
    PaymentMethod: PaymentMethodsInterface;

    OrderTime: Date;
    DeliveryAddress: string;
    OrderDiscount: number;
    OrderTotalPrice: number;

    OrderItems: PurchaseOrderItemsInterface[];
}
import { NamePromotionInterface } from "./INamePromotion";
import { PromotionPeriodInterface } from "./IPromotionPeriod";
import { PromotionTypeInterface } from "./IPromotionType";

export interface ManagePromotionInterface {
    ID: number;

   NamePromotionID: number;
   NamePromotion: NamePromotionInterface;

   PromotionPeriodID: number;
   PromotionPeriod: PromotionPeriodInterface;

   PromotionTypeID: number;
   PromotionType: PromotionTypeInterface;

   MinPrice: number;
   Quantity: number;
   Discount: number; 
   
   PromotionCode: string;

   Createdatetime: Date;


}
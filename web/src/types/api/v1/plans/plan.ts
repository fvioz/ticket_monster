import { BasePlan } from "./basePlan";
import { Zone } from "./zone";

export interface Plan {
  id: string;
  planStartDate: string;
  planEndDate: string;
  sellTo: string;
  sollOut: boolean;
  basePlan: BasePlan;
  zones: Zone[];
}

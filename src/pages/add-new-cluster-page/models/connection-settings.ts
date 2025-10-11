import {Prod} from './prod';
import {Mon} from './mon';

export interface ConnectionSettings {
    prod: Prod;
    mon: Mon
}

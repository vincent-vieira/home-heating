import { Manageable } from './manageable';

export interface ItemWriter<Output, Setup, Teardown> extends Manageable<Setup, Teardown> {
  write(value: Output);
}
import { App } from './app';
import { SocketIOItemReader } from './data/socket-io.item-reader';
import { RethinkDBItemWriter } from './data/rethinkdb.item-writer';

let measurementsObservable = new App(new SocketIOItemReader(), new RethinkDBItemWriter()).setup();
this.measurementsSubscription = measurementsObservable.subscribe((measurement) => this.writer.write(measurement));

// Loop and wait for interrupt.
['SIGINT', 'SIGTERM', 'SIGQUIT', 'SIGKILL'].forEach(signal => process.on(signal, () => {
  this.measurementsSubscription.unsubscribe();
}));
process.stdin.resume();
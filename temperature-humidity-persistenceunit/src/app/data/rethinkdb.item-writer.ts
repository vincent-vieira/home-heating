import { ItemWriter } from './item-writer';
import { Measurement } from './measurement';
import { Connection, connect, table, dbList, dbCreate, db } from 'rethinkdb';

export class RethinkDBItemWriter implements ItemWriter<Measurement, void, void> {

  private rethinkClient: Connection;
  private tableName: string;
  private databaseName: string;

  setup(environment): Promise<void> {
    this.tableName = environment.RETHINKDB_TABLE_NAME;
    this.databaseName = environment.RETHINKDB_DATABASE_NAME;

    // TODO : some logs ?
    return connect(<string>environment.RETHINKDB_REMOTE)
      .then((connection: Connection) => {
        this.rethinkClient = connection;
        return dbList().run(this.rethinkClient);
      })
      .then(databaseList => {
        if(databaseList.indexOf(this.databaseName) === -1) {
          return dbCreate(this.databaseName).run(this.rethinkClient);
        }
        return Promise.resolve();
      })
      .then(() => {
        this.rethinkClient.use(this.databaseName);
        return db(this.databaseName).tableList().run(this.rethinkClient);
      })
      .then(tableList => {
        if(tableList.indexOf(this.tableName) === -1) {
          return db(this.databaseName).tableCreate(this.tableName).run(this.rethinkClient);
        }
        return Promise.resolve();
      })
      .then(() => {});
  }

  write(value: Measurement) {
    table(this.tableName).insert(value).run(this.rethinkClient).then(() => {}, (error) => {
      // TODO : handle error
    });
  }

  teardown(): Promise<void> {
    return this.rethinkClient.close();
  }
}
export interface Manageable<Setup, Teardown> {
  teardown(): Promise<Teardown>;
  setup(environment: any): Promise<Setup>;
}
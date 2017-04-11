module.exports = (config) => {
  config.set({
    frameworks: ['jasmine', 'karma-typescript', 'es6-shim'],
    reporters: ['spec', 'karma-typescript'],
    browsers: ['PhantomJS'],
    files: [
      'src/test/*.spec.ts',
      'src/app/**/*.ts'
    ],
    exclude: [
      'src/app/**/*.item-writer.ts',
      'src/app/**/*.item-reader.ts',
      'src/app/launcher.ts'
    ],
    singleRun: true,
    preprocessors: {
      'src/**/*.ts': ['karma-typescript']
    },
    karmaTypescriptConfig: {
      compilerOptions: {
        module: 'commonjs',
        target: 'es5',
        typeRoots: ['node_modules/@types']
      },
      bundlerOptions: {
        transforms: [
          require('karma-typescript-es6-transform')()
        ],
        validateSyntax: true
      }
    }
  });
};
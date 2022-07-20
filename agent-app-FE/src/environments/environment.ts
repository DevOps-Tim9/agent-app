// This file can be replaced during build by using the `fileReplacements` array.
// `ng build` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
  production: false,
  api_url: 'http://localhost:9094/',
  auth: {
    clientID: 'ch0E73GSl25SIsXeftx4f2ByO5lJqlTT',
    domain: 'dev-4l1tkzmy.eu.auth0.com',
    audience: 'http://localhost:9094',
    redirect: 'http://localhost:4201/homepage',
    logout: 'http://localhost:4201/homepage',
    scope: 'openid profile'
  }
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/plugins/zone-error';  // Included with Angular CLI.

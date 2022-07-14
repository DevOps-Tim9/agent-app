import { TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { AuthenticationService } from 'src/app/core/services/authentication.service';

import { NoAuthGuard } from './no-auth.guard';

describe('NoAuthGuard', () => {
  let guard: NoAuthGuard;
  let router: Router;
  let authenticationService: AuthenticationService;

  beforeEach(() => {
    const authenticationServiceMocked = {
      isAuthenticated: jasmine.createSpy('isAuthenticated').and.returnValue(true)
    };
    const routerMocked = jasmine.createSpyObj('router', ['navigate']);
    TestBed.configureTestingModule({
      providers: [
        NoAuthGuard,
        {provide: AuthenticationService, useValue: authenticationServiceMocked},
        {provide: Router, useValue: routerMocked }
      ]
    });
  });

  it('should be created', () => {
    guard = TestBed.inject(NoAuthGuard);
    router = TestBed.inject(Router);
    authenticationService = TestBed.inject(AuthenticationService);
    expect(guard).toBeTruthy();
  });

  it('canActivate() should return false', () => {
    guard = TestBed.inject(NoAuthGuard);
    router = TestBed.inject(Router);
    authenticationService = TestBed.inject(AuthenticationService);
    const val = guard.canActivate();

    expect(val).toBeFalsy();
    expect(authenticationService.isAuthenticated).toHaveBeenCalled();
    expect(router.navigate).toHaveBeenCalledWith(['/']);
  });

  it('canActivate() should return true', () => {
    const authenticationServiceMocked2 = {
      isAuthenticated: jasmine.createSpy('isAuthenticated').and.returnValue(false)
    };

    TestBed.overrideProvider(AuthenticationService, {useValue: authenticationServiceMocked2});
    guard = TestBed.inject(NoAuthGuard);
    router = TestBed.inject(Router);
    authenticationService = TestBed.inject(AuthenticationService);
    const val = guard.canActivate();

    expect(val).toBeTruthy();
    expect(authenticationService.isAuthenticated).toHaveBeenCalled();
    expect(router.navigate).toHaveBeenCalledTimes(0);
  });

});

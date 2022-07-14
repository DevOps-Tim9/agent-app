import { TestBed } from '@angular/core/testing';
import { ActivatedRouteSnapshot, Router } from '@angular/router';
import { JwtService } from 'src/app/core/services/jwt.service';
import { of } from 'rxjs';
import { RoleGuard } from './role.guard';

describe('RoleGuard', () => {
  let guard: RoleGuard;
  let router: Router;
  let jwtService: JwtService;

  beforeEach(() => {
    const jwtServiceMocked = {
      getToken: jasmine.createSpy('getToken').and.returnValue({
        accessToken: 'eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJzcHJpbmctc2VjdXJpdHktZXhhbXBsZSIsInN1YiI6ImVtYWlsX2FkcmVzYTFAZ21haWwuY29tIiwiYXVkIjoid2ViIiwiaWF0IjoxNjExMDE0MDk0LCJleHAiOjE2MTEwMTU4OTQsInJvbGUiOiJST0xFX0FETUlOIn0.bl85WgHa1NStDvrlxrWa3P6WO7TUMqcrDDd1WPtxVkOV0gh8Mb0kAcaPU3UzlTR0xsNAl8b08N8rLMTGoJD4HQ',
        expiresIn: 1800000
      })
    };
    const routerMocked = jasmine.createSpyObj('router', ['navigate']);
    TestBed.configureTestingModule({
      providers: [RoleGuard,  {provide: JwtService, useValue: jwtServiceMocked}, {provide: Router, useValue: routerMocked }]
    });
  });

  it('should be created', () => {
    router = TestBed.inject(Router);
    jwtService = TestBed.inject(JwtService);
    guard = TestBed.inject(RoleGuard);
    expect(guard).toBeTruthy();
  });

  it('canActivate() should return true', () => {
    guard = TestBed.inject(RoleGuard);
    router = TestBed.inject(Router);
    jwtService = TestBed.inject(JwtService);
    const val = guard.canActivate(({data: { expectedRoles: 'ROLE_ADMIN'}} as any) as ActivatedRouteSnapshot);

    expect(val).toBeTruthy();
    expect(jwtService.getToken).toHaveBeenCalled();
    expect(router.navigate).toHaveBeenCalledTimes(0);
  });

  it('canActivate() should return false', () => {
    const jwtServiceMocked2 = {
      getToken: jasmine.createSpy('getToken').and.returnValue(null)
    };

    TestBed.overrideProvider(JwtService, {useValue: jwtServiceMocked2});
    guard = TestBed.inject(RoleGuard);
    router = TestBed.inject(Router);
    jwtService = TestBed.inject(JwtService);
    const val = guard.canActivate(({data: { expectedRoles: 'ROLE_ADMIN'}} as any) as ActivatedRouteSnapshot);

    expect(val).toBeFalsy();
    expect(jwtService.getToken).toHaveBeenCalled();
    expect(router.navigate).toHaveBeenCalledWith(['/auth/sign-in']);
  });

});

# Access Token and Refresh Token Guide

This document explains what access tokens and refresh tokens are, how they are used, and why refresh tokens should be stored in Secure + HttpOnly cookies.

---

## Access Token

- Short lifetime (minutes)
- Sent with each API request using the Authorization header
- Safe to expose to frontend JavaScript
- If stolen, it becomes useless shortly after

---

## Refresh Token

- Long lifetime (days or weeks)
- Used only to request new access tokens
- Must never be readable by frontend JavaScript
- If stolen, an attacker can generate new access tokens

---

## Why Store Refresh Tokens in Secure HttpOnly Cookies

A refresh token acts as a long-term session key.  
Storing it in JavaScript-accessible storage is unsafe.

Use a cookie with these attributes:

```
HttpOnly; Secure; SameSite=Lax; Path=/auth/refresh; Max-Age=2592000
```


Explanation:

- HttpOnly: JavaScript cannot read the cookie
- Secure: only sent over HTTPS
- SameSite: reduces CSRF attacks
- Path: limits when the cookie is sent
- Max-Age: controls expiration time

---

## Who Can See the Refresh Token

Cannot see:
- Frontend JavaScript
- Injected XSS scripts
- Network attackers when HTTPS is used

Can see:
- The user through browser DevTools
- Browser extensions with cookie permissions
- Malware or a compromised device

HttpOnly protects against JavaScript access, not against the user inspecting their own cookies.

---

## Token Refresh Flow

1. User logs in  
   - Backend issues an access token (JSON)  
   - Backend sends refresh token as a Secure + HttpOnly cookie

2. Frontend stores the access token in memory.

3. When the access token expires, frontend calls:  
   `fetch("/auth/refresh", { method: "POST", credentials: "include" })`

4. Browser automatically includes the refresh_token cookie.

5. Backend validates the refresh token and returns a new access token.

---

## Logout

Backend should:
1. Delete refresh token from server storage
2. Clear the cookie:
```
Max-Age=0
```

---

## Summary

- Access token: short-lived, safe for frontend use
- Refresh token: long-lived, must be protected
- Store refresh tokens only in Secure + HttpOnly cookies
- Browser handles refresh cookies automatically
- JavaScript never reads the refresh token

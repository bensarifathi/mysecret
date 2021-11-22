# mysecret

---
this is a CLI app that store and fetch user account information in a secure way.

## How to use it:

---

**Register Account**:

```Bash
$ mysecret register 
Name: github
Email: myemail@gmail.com
Passowrd: ****************
Confirm password: *****************

$ [+] new account successfully saved
```

**Fetch Account**:

```bash
$ mysecret fetch
Name: github
=============================================================
Name      Email                                   Password  
github    <myemail>                              <mypassword>  

```

## How it work in backend:

---

- The user must create an account with the command: **mysecret register**.

- The user must provide a secret key that will be used to encrypt the stored account password with the command: **mysecret cipher**.

- To store and fetch an account it's simple just run the command above.

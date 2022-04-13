# TansSactional

So how does things work?
* Import Postman Requests [Curl](https://www.getpostman.com/collections/7b0cb4ae55f7e9fe3020)
(Import using JSON link)
## Things I used:-
* Go(lang)
* SQL
* JWT

## Explaining APIs:-
* Get Single User & all Users
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/user/getuser (remover userID from header and get all users)
* Create User
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/user/createuser (to create user we need company id,so create company first)
* Patch User Data
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/user/updateuser
* Create Product
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/user/createproduct
* Get Single Company & all Companies
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/company/getcompany (remover companyId from header and get all users)
* Create Company
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/company/createcompany
* Patch Company Data
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/company/updatecompany
* Get Invoices(sort by total,filter by greater,less or equal to amount,
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/invoices/getinvoice
* Create Invoice
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/invoices/createinvoice
* Patch Company Invoice
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/invoices/updateinvoices
* Send Company Invoice
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/invoices/sendinvoice
* Acknowledge Company Invoice
  - Sample URL:- https://tranzactional.herokuapp.com/api/v0/invoices/acknowledgeinvoice
 
## Overall Schematic 
![Group Schema (8)](https://user-images.githubusercontent.com/60891544/163135586-5e9e20c3-8180-4813-94c5-4ddcb2d80419.png)
## User Schematic
![System Flow (13)](https://user-images.githubusercontent.com/60891544/163108333-97ee3311-355e-4ff6-98d0-c0738263b209.png)
## Company Schematic
![Source Diagram (2)](https://user-images.githubusercontent.com/60891544/163108338-882c1501-8bce-406e-b9b3-009e45b98ab7.png)
## Invoices Schematic
![Source Diagram (3)](https://user-images.githubusercontent.com/60891544/163131509-ba6bd736-a211-4e22-9922-a5ac5e2505a6.png)
